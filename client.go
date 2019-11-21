package ilclient

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
	"unsafe"
)

/*
#cgo CFLAGS: -Wno-unused-variable -Wall -Wno-deprecated -g -DRASPBERRY_PI -DSTANDALONE -D__STDC_CONSTANT_MACROS  -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -I/opt/vc/include/IL -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux/ -I./include
#cgo LDFLAGS: -L/opt/vc/lib -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt

#include <OMX_Core.h>
#include <OMX_Component.h>

#include <stdlib.h>
#include <stdio.h>
#include <bcm_host.h>
#include <ilclient.h>

extern COMPONENT_T* ilclient_create_component_wrapper(ILCLIENT_T *handle, int * ret, char * name, ILCLIENT_CREATE_FLAGS_T flags);
extern void enable_trace_logging();

extern void setup_callbacks(ILCLIENT_T * handle);

extern int get_component_state(COMPONENT_T * comp, OMX_STATETYPE * state);

extern void initialize_struct(void * p, size_t size);
extern OMX_ERRORTYPE set_parameter(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param);
extern OMX_ERRORTYPE get_parameter(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param);
extern OMX_ERRORTYPE set_config(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param);
extern OMX_ERRORTYPE get_config(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param);

extern int calc_stride(unsigned int width, unsigned int alignment);
*/
import "C"

const default_timeout = 0

type Client struct {
	client     *C.ILCLIENT_T
	Timeout    time.Duration
	components map[*C.COMPONENT_T]*Component
	tunnels    map[*C.TUNNEL_T]*Tunnel
	clientID   C.int
	lock       sync.Locker
}
type Event struct{}
type Component struct {
	component *C.COMPONENT_T
}
type Tunnel struct {
	tunnel *C.TUNNEL_T
}
type Buffer struct {
	buffer *C.OMX_BUFFERHEADERTYPE
}
type ComponentPort struct {
	component *Component
	port      PortIndex
}

var (
	client *Client
)

func init() {
	C.enable_trace_logging()

	fmt.Fprintf(os.Stderr, "bcm_host_init\n")
	C.bcm_host_init()

	fmt.Fprintf(os.Stderr, "ilclient_init ")
	c := C.ilclient_init()
	fmt.Fprintf(os.Stderr, "response: %v\n", client)

	client = &Client{
		client:  c,
		Timeout: default_timeout,

		components: make(map[*C.COMPONENT_T]*Component),
		tunnels:    make(map[*C.TUNNEL_T]*Tunnel),
		lock:       &sync.Mutex{},
	}

	C.setup_callbacks(client.client)

	fmt.Fprintf(os.Stderr, "OMX_Init ")
	err := C.OMX_Init()
	fmt.Fprintf(os.Stderr, "ret: %v\n", Error(err))

}

func CalculateStride(width uint, alignment uint) int {
	return int(C.calc_stride(C.uint(width), C.uint(alignment)))
}

//export goErrorHandler
func goErrorHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T, data C.OMX_U32) {
	client.handleError(comp, data)
}

//export goPortSettingsChangedHandler
func goPortSettingsChangedHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T, data C.OMX_U32) {
	client.portSettingsChanged(comp, data)
}

//export goEOSHandler
func goEOSHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T, data C.OMX_U32) {
	client.handleEOS(comp, data)
}

//export goConfigChangedHandler
func goConfigChangedHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T, data C.OMX_U32) {
	client.handleConfigChanged(comp, data)
}

//export goFillBufferHandler
func goFillBufferHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T) {
	client.handleFillBuffer(comp)
}

//export goEmptyBufferHandler
func goEmptyBufferHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T) {
	client.handleEmptyBuffer(comp)
}

func Get() *Client {
	return client
}

func (c *Client) NewComponent(name string, flags ...CreateFlag) (*Component, error) {
	ret := &Component{}
	var e C.int

	var fin int
	for _, f := range flags {
		fin = fin | int(f)
	}

	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))
	ret.component = C.ilclient_create_component_wrapper(c.client, &e, C.CString(name), C.ILCLIENT_CREATE_FLAGS_T(fin))

	if e != 0 {
		return nil, fmt.Errorf("ilclient: could not create component: %v", Error(e))
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.components[ret.component] = ret

	return ret, nil
}

func (c *Client) NewTunnel(source, sink ComponentPort) (*Tunnel, error) {
	t := &C.TUNNEL_T{}
	t.source = source.component.component
	t.source_port = C.int(source.port)
	t.sink = sink.component.component
	t.sink_port = C.int(sink.port)

	if e := TunnelError(C.ilclient_setup_tunnel(t, 0, C.int(1000*c.Timeout.Seconds()))); e != TunnelErrorNone {
		return nil, e
	}

	ret := &Tunnel{t}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.tunnels[t] = ret

	return ret, nil
}

func (c *Client) Close() {
	// cleanup tunnels
	// list of null terminated tunnels

	c.lock.Lock()
	defer c.lock.Unlock()

	tuns := make([]C.TUNNEL_T, len(c.tunnels)+1)
	i := 0
	for t, _ := range c.tunnels {
		tuns[i].source = t.source
		tuns[i].source_port = t.source_port
		tuns[i].sink = t.sink
		tuns[i].sink_port = t.sink_port

		delete(c.tunnels, t)
		i++
	}
	tuns[i].source = nil
	tuns[i].source_port = 0
	tuns[i].sink = nil
	tuns[i].sink_port = 0

	C.ilclient_teardown_tunnels(&tuns[0])

	// cleanup components
	// list of components is null terminated
	comps := make([]*C.COMPONENT_T, len(c.components)+1)

	i = 0
	for k, _ := range c.components {
		comps[i] = k
		delete(c.components, k)
		i++
	}
	// this shouldn't be necessary, but just being explicit
	comps[i] = nil
	C.ilclient_cleanup_components(&comps[0])

	C.ilclient_destroy(c.client)
}

func (c *Client) handleError(comp *C.COMPONENT_T, data C.OMX_U32) {
	log.Printf("error: %v", Error(data))
}

func (c *Client) portSettingsChanged(comp *C.COMPONENT_T, data C.OMX_U32) {
	log.Printf("port settings changed")
}

func (c *Client) handleEOS(comp *C.COMPONENT_T, data C.OMX_U32) {
	log.Printf("end of stream")
}

func (c *Client) handleConfigChanged(comp *C.COMPONENT_T, data C.OMX_U32) {
	log.Printf("config changed")
}

func (c *Client) handleFillBuffer(comp *C.COMPONENT_T) {
	log.Printf("buffer filled")
}
func (c *Client) handleEmptyBuffer(comp *C.COMPONENT_T) {
	log.Printf("buffer emptied")
}

func (c *Component) SuggestBufferSize(size int) error {
	if e := C.ilclient_suggest_bufsize(c.component, C.uint(size)); e != 0 {
		return fmt.Errorf("suggest buffer size failed")
	}

	return nil
}

func (c *Component) queryPorts(dir C.OMX_DIRTYPE, domain C.OMX_PORTDOMAINTYPE) []ComponentPort {
	i := 0
	var ret []ComponentPort

	for {
		p := C.ilclient_get_port_index(c.component, dir, domain, C.int(i))
		if p < 0 {
			break
		}

		ret = append(ret, ComponentPort{c, PortIndex(p)})
		i++
	}

	return ret
}

func (c *Component) InputPorts() []ComponentPort {
	return c.queryPorts(C.OMX_DirInput, C.OMX_PORTDOMAINTYPE(0XFFFFFFFF))
}

func (c *Component) OutputPorts() []ComponentPort {
	return c.queryPorts(C.OMX_DirOutput, C.OMX_PORTDOMAINTYPE(0XFFFFFFFF))
}

func (c *Component) OutputBuffer(port_index PortIndex) (*Buffer, error) {
	fmt.Fprintf(os.Stderr, "ilclient_get_output_buffer\n")
	buf := C.ilclient_get_output_buffer(c.component, C.int(port_index), 0)
	if buf == nil {
		return nil, fmt.Errorf("output buffer not available for port %d", port_index)
	}
	return &Buffer{buf}, nil
}

func (c *Component) InputBuffer(port_index PortIndex) (*Buffer, error) {
	fmt.Fprintf(os.Stderr, "ilclient_get_input_buffer\n")
	buf := C.ilclient_get_input_buffer(c.component, C.int(port_index), 0)
	if buf == nil {
		return nil, fmt.Errorf("input buffer not available for port %d", port_index)
	}
	return &Buffer{buf}, nil
}

func (c *Component) State() (State, error) {
	var s C.OMX_STATETYPE
	e := C.get_component_state(c.component, &s)

	if e != C.OMX_ErrorNone {
		return State(s), Error(e)
	}

	return State(s), nil
}

func (c *Component) SetState(state State) error {
	if ret := C.ilclient_change_component_state(c.component, C.OMX_STATETYPE(state)); ret != 0 {
		return fmt.Errorf("ilclient: could not change component state: %v", Error(ret))
	}
	return nil
}

func (c *Component) Port(port_index PortIndex) ComponentPort {
	return ComponentPort{c, port_index}
}

func (c ComponentPort) Enable() {
	C.ilclient_enable_port(c.component.component, C.int(c.port))
}

func (c ComponentPort) Disable() {
	C.ilclient_disable_port(c.component.component, C.int(c.port))
}

func (c ComponentPort) EnableBuffers() error {
	if e := C.ilclient_enable_port_buffers(c.component.component, C.int(c.port), nil, nil, nil); e != 0 {
		return fmt.Errorf("error: EnablePortBuffers: %v", Error(e))
	}
	return nil
}

func (c ComponentPort) DisableBuffers() {
	C.ilclient_disable_port_buffers(c.component.component, C.int(c.port), nil, nil, nil)
}

func (t *Tunnel) Close() {
	C.ilclient_disable_tunnel(t.tunnel)
}

func (t *Tunnel) Open() error {
	if ret := C.ilclient_enable_tunnel(t.tunnel); ret != 0 {
		return TunnelErrorNoEnable
	}
	return nil
}

func (t *Tunnel) Flush() {
	C.ilclient_flush_tunnels(t.tunnel, 0)
}

type ImageFormat struct {
	Compression ImageCoding
	Color       ColorFormat
}

func (c ComponentPort) SetImagePortFormat(formats []ImageFormat) error {
	for i, f := range formats {
		var p C.OMX_IMAGE_PARAM_PORTFORMATTYPE
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nIndex = C.OMX_U32(i)
		p.eCompressionFormat = C.OMX_IMAGE_CODINGTYPE(f.Compression)
		p.eColorFormat = C.OMX_COLOR_FORMATTYPE(f.Color)

		e := C.set_parameter(c.component.component, C.OMX_IndexParamImagePortFormat, unsafe.Pointer(&p))

		if e != C.OMX_ErrorNone {
			return Error(e)
		}
	}

	return nil
}

func (c ComponentPort) GetImagePortFormat() ([]ImageFormat, error) {
	var ret []ImageFormat
	var e C.OMX_ERRORTYPE

	for i := uint(0); ; i++ {
		var p C.OMX_IMAGE_PARAM_PORTFORMATTYPE
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamImagePortFormat, unsafe.Pointer(&p))

		if e == C.OMX_ErrorNone {
			ret = append(ret, ImageFormat{
				ImageCoding(p.eCompressionFormat),
				ColorFormat(p.eColorFormat),
			})
		} else {
			break
		}
	}

	if e == C.OMX_ErrorNoMore {
		return ret, nil
	}
	return ret, Error(e)
}

func toQ16(x float64) C.OMX_U32 {
	return C.OMX_U32(math.Floor(x * 65536.))
}
func fromQ16(y C.OMX_U32) float64 {
	return float64(y) / 65536.
}

type VideoFormat struct {
	Compression VideoCoding
	Color       ColorFormat
	Framerate   float64
}

func (c ComponentPort) SetVideoPortFormat(formats []VideoFormat) error {
	for i, f := range formats {
		var p C.OMX_VIDEO_PARAM_PORTFORMATTYPE
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nIndex = C.OMX_U32(i)
		p.eCompressionFormat = C.OMX_VIDEO_CODINGTYPE(f.Compression)
		p.eColorFormat = C.OMX_COLOR_FORMATTYPE(f.Color)
		p.xFramerate = toQ16(f.Framerate)

		e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoPortFormat, unsafe.Pointer(&p))
		if e != C.OMX_ErrorNone {
			return Error(e)
		}
	}

	return nil
}

func (c ComponentPort) GetVideoPortFormat() ([]VideoFormat, error) {
	var ret []VideoFormat
	var e C.OMX_ERRORTYPE

	for i := uint(0); true; i++ {
		var p C.OMX_VIDEO_PARAM_PORTFORMATTYPE
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamVideoPortFormat, unsafe.Pointer(&p))
		if e == C.OMX_ErrorNone {
			ret = append(ret, VideoFormat{
				VideoCoding(p.eCompressionFormat),
				ColorFormat(p.eColorFormat),
				fromQ16(p.xFramerate),
			})
		} else {
			break
		}
	}

	if e == C.OMX_ErrorNoMore {
		return ret, nil
	}
	return ret, Error(e)
}

type VideoQuantization struct {
	QpI uint
	QpP uint
	QpB uint
}

func (c ComponentPort) SetVideoQuantization(q VideoQuantization) error {
	var p C.OMX_VIDEO_PARAM_QUANTIZATIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nQpI = C.OMX_U32(q.QpI)
	p.nQpP = C.OMX_U32(q.QpP)
	p.nQpB = C.OMX_U32(q.QpB)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoQuantization, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoQuantization() (ret VideoQuantization, err error) {
	var p C.OMX_VIDEO_PARAM_QUANTIZATIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoQuantization, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return ret, Error(e)
	}

	ret.QpI = uint(p.nQpI)
	ret.QpP = uint(p.nQpP)
	ret.QpB = uint(p.nQpB)

	return ret, nil
}

type VideoFastUpdate struct {
	Enabled  bool
	FirstGOB uint
	FirstMB  uint
	NumMB    uint
}

func toOMXBool(x bool) C.OMX_BOOL {
	if x {
		return C.OMX_TRUE
	}
	return C.OMX_FALSE
}

func (c ComponentPort) SetVideoFastUpdate(v VideoFastUpdate) error {
	var p C.OMX_VIDEO_PARAM_VIDEOFASTUPDATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.bEnableVFU = toOMXBool(v.Enabled)
	p.nFirstGOB = C.OMX_U32(v.FirstGOB)
	p.nFirstMB = C.OMX_U32(v.FirstMB)
	p.nNumMBs = C.OMX_U32(v.NumMB)

	if e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoFastUpdate, unsafe.Pointer(&p)); e != C.OMX_ErrorNone {
		return Error(e)
	}

	return nil
}

func (c ComponentPort) GetVideoFastUpdate() (ret VideoFastUpdate, err error) {
	var p C.OMX_VIDEO_PARAM_VIDEOFASTUPDATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoFastUpdate, unsafe.Pointer(&p))

	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.Enabled = (p.bEnableVFU != C.OMX_FALSE)
	ret.FirstGOB = uint(p.nFirstGOB)
	ret.FirstMB = uint(p.nFirstMB)
	ret.NumMB = uint(p.nNumMBs)
	return
}

type VideoBitrate struct {
	ControlRate   VideoControlRate
	TargetBitrate uint
}

func (c ComponentPort) SetVideoBitrate(v VideoBitrate) error {
	var p C.OMX_VIDEO_PARAM_BITRATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eControlRate = C.OMX_VIDEO_CONTROLRATETYPE(v.ControlRate)
	p.nTargetBitrate = C.OMX_U32(v.TargetBitrate)

	if e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoBitrate, unsafe.Pointer(&p)); e != C.OMX_ErrorNone {
		return Error(e)
	}

	return nil
}

func (c ComponentPort) GetVideoBitrate() (ret VideoBitrate, err error) {
	var p C.OMX_VIDEO_PARAM_BITRATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoBitrate, unsafe.Pointer(&p))

	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.ControlRate = VideoControlRate(p.eControlRate)
	ret.TargetBitrate = uint(p.nTargetBitrate)
	return
}

type VideoMotionVector struct {
	Accuracy        VideoMotionVectorAccuracy
	UnrestrictedMVs bool
	FourMV          bool
	XSearchRange    int
	YSearchRange    int
}

func (c ComponentPort) SetVideoMotionVector(v VideoMotionVector) error {
	var p C.OMX_VIDEO_PARAM_MOTIONVECTORTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eAccuracy = C.OMX_VIDEO_MOTIONVECTORTYPE(v.Accuracy)
	p.bUnrestrictedMVs = toOMXBool(v.UnrestrictedMVs)
	p.bFourMV = toOMXBool(v.FourMV)
	p.sXSearchRange = C.OMX_S32(v.XSearchRange)
	p.sYSearchRange = C.OMX_S32(v.YSearchRange)

	if e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoMotionVector, unsafe.Pointer(&p)); e != C.OMX_ErrorNone {
		return Error(e)
	}

	return nil
}

func (c ComponentPort) GetVideoMotionVector() (ret VideoMotionVector, err error) {
	var p C.OMX_VIDEO_PARAM_MOTIONVECTORTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoMotionVector, unsafe.Pointer(&p))

	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.Accuracy = VideoMotionVectorAccuracy(p.eAccuracy)
	ret.UnrestrictedMVs = (p.bUnrestrictedMVs != C.OMX_FALSE)
	ret.FourMV = (p.bFourMV != C.OMX_FALSE)
	ret.XSearchRange = int(p.sXSearchRange)
	ret.YSearchRange = int(p.sYSearchRange)
	return
}

type VideoIntraRefresh struct {
	Mode   VideoIntraRefreshMode
	AirMBs uint
	AirRef uint
	CirMBs uint
	PirMBs uint
}

func (c ComponentPort) SetVideoIntraRefresh(v VideoIntraRefresh) error {
	var p C.OMX_VIDEO_PARAM_INTRAREFRESHTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eRefreshMode = C.OMX_VIDEO_INTRAREFRESHTYPE(v.Mode)
	p.nAirMBs = C.OMX_U32(v.AirMBs)
	p.nAirRef = C.OMX_U32(v.AirRef)
	p.nCirMBs = C.OMX_U32(v.CirMBs)
	p.nPirMBs = C.OMX_U32(v.PirMBs)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoIntraRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoIntraRefresh() (ret VideoIntraRefresh, err error) {
	var p C.OMX_VIDEO_PARAM_INTRAREFRESHTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoIntraRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.Mode = VideoIntraRefreshMode(p.eRefreshMode)
	ret.AirMBs = uint(p.nAirMBs)
	ret.AirRef = uint(p.nAirRef)
	ret.CirMBs = uint(p.nCirMBs)
	ret.PirMBs = uint(p.nPirMBs)
	return
}

type VideoErrorCorrection struct {
	EnableHEC              bool
	EnableResync           bool
	ResynchMarkerSpacing   uint
	EnableDataPartitioning bool
	EnableRVLC             bool
}

func (c ComponentPort) SetVideoErrorCorrection(v VideoErrorCorrection) error {
	var p C.OMX_VIDEO_PARAM_ERRORCORRECTIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.bEnableHEC = toOMXBool(v.EnableHEC)
	p.bEnableResync = toOMXBool(v.EnableResync)
	p.nResynchMarkerSpacing = C.OMX_U32(v.ResynchMarkerSpacing)
	p.bEnableDataPartitioning = toOMXBool(v.EnableDataPartitioning)
	p.bEnableRVLC = toOMXBool(v.EnableRVLC)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoErrorCorrection, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoErrorCorrection() (ret VideoErrorCorrection, err error) {
	var p C.OMX_VIDEO_PARAM_ERRORCORRECTIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoErrorCorrection, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.EnableHEC = (p.bEnableHEC != C.OMX_FALSE)
	ret.EnableResync = (p.bEnableResync != C.OMX_FALSE)
	ret.ResynchMarkerSpacing = uint(p.nResynchMarkerSpacing)
	ret.EnableDataPartitioning = (p.bEnableDataPartitioning != C.OMX_FALSE)
	ret.EnableRVLC = (p.bEnableRVLC != C.OMX_FALSE)
	return
}

type VideoVariableBlockSizeMotionCompensation struct {
	B16x16 bool
	B16x8  bool
	B8x16  bool
	B8x8   bool
	B8x4   bool
	B4x8   bool
	B4x4   bool
}

func (c ComponentPort) SetVideoVBSMC(v VideoVariableBlockSizeMotionCompensation) error {
	var p C.OMX_VIDEO_PARAM_VBSMCTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.b16x16 = toOMXBool(v.B16x16)
	p.b16x8 = toOMXBool(v.B16x8)
	p.b8x16 = toOMXBool(v.B8x16)
	p.b8x8 = toOMXBool(v.B8x8)
	p.b8x4 = toOMXBool(v.B8x4)
	p.b4x8 = toOMXBool(v.B16x16)
	p.b4x4 = toOMXBool(v.B16x16)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoVBSMC, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoVBSMC() (ret VideoVariableBlockSizeMotionCompensation, err error) {
	var p C.OMX_VIDEO_PARAM_VBSMCTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoVBSMC, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.B16x16 = (p.b16x16 != C.OMX_FALSE)
	ret.B16x8 = (p.b16x8 != C.OMX_FALSE)
	ret.B8x16 = (p.b8x16 != C.OMX_FALSE)
	ret.B8x8 = (p.b8x8 != C.OMX_FALSE)
	ret.B8x4 = (p.b8x4 != C.OMX_FALSE)
	ret.B4x8 = (p.b4x8 != C.OMX_FALSE)
	ret.B4x4 = (p.b4x4 != C.OMX_FALSE)
	return
}

type VideoH263 struct {
	PFrames                 uint
	BFrames                 uint
	Profile                 VideoH263Profile
	Level                   VideoH263Level
	PLUSPTYPEAllowed        bool
	AllowedPictureTypes     uint
	ForceRoundingTypeToZero bool
	PictureHeaderRepetition uint
	GOBHeaderInterval       uint
}

func (c ComponentPort) SetVideoH263(v VideoH263) error {
	var p C.OMX_VIDEO_PARAM_H263TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nPFrames = C.OMX_U32(v.PFrames)
	p.nBFrames = C.OMX_U32(v.BFrames)
	p.eProfile = C.OMX_VIDEO_H263PROFILETYPE(v.Profile)
	p.eLevel = C.OMX_VIDEO_H263LEVELTYPE(v.Level)
	p.bPLUSPTYPEAllowed = toOMXBool(v.PLUSPTYPEAllowed)
	p.nAllowedPictureTypes = C.OMX_U32(v.AllowedPictureTypes)
	p.bForceRoundingTypeToZero = toOMXBool(v.ForceRoundingTypeToZero)
	p.nPictureHeaderRepetition = C.OMX_U32(v.PictureHeaderRepetition)
	p.nGOBHeaderInterval = C.OMX_U32(v.GOBHeaderInterval)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoH263, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoH263() (ret VideoH263, err error) {
	var p C.OMX_VIDEO_PARAM_H263TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoH263, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.PFrames = uint(p.nPFrames)
	ret.BFrames = uint(p.nBFrames)
	ret.Profile = VideoH263Profile(p.eProfile)
	ret.Level = VideoH263Level(p.eLevel)
	ret.PLUSPTYPEAllowed = (p.bPLUSPTYPEAllowed != C.OMX_FALSE)
	ret.AllowedPictureTypes = uint(p.nAllowedPictureTypes)
	ret.ForceRoundingTypeToZero = (p.bForceRoundingTypeToZero != C.OMX_FALSE)
	ret.PictureHeaderRepetition = uint(p.nPictureHeaderRepetition)
	ret.GOBHeaderInterval = uint(p.nGOBHeaderInterval)
	return
}

type H263ProfileLevel struct {
	Profile VideoH263Profile
	Level   VideoH263Level
}

func (c ComponentPort) CurrentH263ProfileLevel() (ret H263ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelCurrent, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
	} else {
		ret.Profile = VideoH263Profile(p.eProfile)
		ret.Level = VideoH263Level(p.eLevel)
	}
	return
}

func (c ComponentPort) SupportedH263ProfileLevels() (ret []H263ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	var e C.OMX_ERRORTYPE

	for i := uint(0); ; i++ {
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nProfileIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelQuerySupported, unsafe.Pointer(&p))
		if e != C.OMX_ErrorNone {
			break
		}
		ret = append(ret, H263ProfileLevel{
			Profile: VideoH263Profile(p.eProfile),
			Level:   VideoH263Level(p.eLevel),
		})
	}

	if e != C.OMX_ErrorNoMore {
		err = Error(e)
	}
	return
}

type VideoMPEG2 struct {
	PFrames uint
	BFrames uint
	Profile VideoMPEG2Profile
	Level   VideoMPEG2Level
}

func (c ComponentPort) SetVideoMPEG2(v VideoMPEG2) error {
	var p C.OMX_VIDEO_PARAM_MPEG2TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nPFrames = C.OMX_U32(v.PFrames)
	p.nBFrames = C.OMX_U32(v.BFrames)
	p.eProfile = C.OMX_VIDEO_MPEG2PROFILETYPE(v.Profile)
	p.eLevel = C.OMX_VIDEO_MPEG2LEVELTYPE(v.Level)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoMpeg2, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoMPEG2() (ret VideoMPEG2, err error) {
	var p C.OMX_VIDEO_PARAM_MPEG2TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoMpeg2, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.PFrames = uint(p.nPFrames)
	ret.BFrames = uint(p.nBFrames)
	ret.Profile = VideoMPEG2Profile(p.eProfile)
	ret.Level = VideoMPEG2Level(p.eLevel)
	return
}

type MPEG2ProfileLevel struct {
	Profile VideoMPEG2Profile
	Level   VideoMPEG2Level
}

func (c ComponentPort) CurrentMPEG2ProfileLevel() (ret MPEG2ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelCurrent, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
	} else {
		ret.Profile = VideoMPEG2Profile(p.eProfile)
		ret.Level = VideoMPEG2Level(p.eLevel)
	}
	return
}

func (c ComponentPort) SupportedMPEG2ProfileLevels() (ret []MPEG2ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	var e C.OMX_ERRORTYPE

	for i := uint(0); ; i++ {
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nProfileIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelQuerySupported, unsafe.Pointer(&p))
		if e != C.OMX_ErrorNone {
			break
		}
		ret = append(ret, MPEG2ProfileLevel{
			Profile: VideoMPEG2Profile(p.eProfile),
			Level:   VideoMPEG2Level(p.eLevel),
		})
	}

	if e != C.OMX_ErrorNoMore {
		err = Error(e)
	}
	return
}

type VideoMPEG4 struct {
	SliceHeaderSpacing  uint
	SVH                 bool
	Gov                 bool
	PFrames             uint
	BFrames             uint
	IDCVLCThreshold     uint
	ACPred              bool
	MaxPacketSize       uint
	TimeIncRes          uint
	Profile             VideoMPEG4Profile
	Level               VideoMPEG4Level
	AllowedPictureTypes uint
	HeaderExtension     uint
	ReversibleVLC       bool
}

func (c ComponentPort) SetVideoMPEG4(v VideoMPEG4) error {
	var p C.OMX_VIDEO_PARAM_MPEG4TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nSliceHeaderSpacing = C.OMX_U32(v.SliceHeaderSpacing)
	p.bSVH = toOMXBool(v.SVH)
	p.bGov = toOMXBool(v.Gov)
	p.nPFrames = C.OMX_U32(v.PFrames)
	p.nBFrames = C.OMX_U32(v.BFrames)
	p.nIDCVLCThreshold = C.OMX_U32(v.IDCVLCThreshold)
	p.nTimeIncRes = C.OMX_U32(v.TimeIncRes)
	p.eProfile = C.OMX_VIDEO_MPEG4PROFILETYPE(v.Profile)
	p.eLevel = C.OMX_VIDEO_MPEG4LEVELTYPE(v.Level)
	p.nAllowedPictureTypes = C.OMX_U32(v.AllowedPictureTypes)
	p.bReversibleVLC = toOMXBool(v.ReversibleVLC)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoMpeg4, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoMPEG4() (ret VideoMPEG4, err error) {
	var p C.OMX_VIDEO_PARAM_MPEG4TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoMpeg4, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.SliceHeaderSpacing = uint(p.nSliceHeaderSpacing)
	ret.SVH = (p.bSVH != C.OMX_FALSE)
	ret.Gov = (p.bGov != C.OMX_FALSE)
	ret.PFrames = uint(p.nPFrames)
	ret.BFrames = uint(p.nBFrames)
	ret.IDCVLCThreshold = uint(p.nIDCVLCThreshold)
	ret.TimeIncRes = uint(p.nTimeIncRes)
	ret.Profile = VideoMPEG4Profile(p.eProfile)
	ret.Level = VideoMPEG4Level(p.eLevel)
	ret.AllowedPictureTypes = uint(p.nAllowedPictureTypes)
	ret.ReversibleVLC = (p.bReversibleVLC != C.OMX_FALSE)
	return
}

type MPEG4ProfileLevel struct {
	Profile VideoMPEG4Profile
	Level   VideoMPEG4Level
}

func (c ComponentPort) CurrentMPEG4ProfileLevel() (ret MPEG4ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelCurrent, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
	} else {
		ret.Profile = VideoMPEG4Profile(p.eProfile)
		ret.Level = VideoMPEG4Level(p.eLevel)
	}
	return
}

func (c ComponentPort) SupportedMPEG4ProfileLevels() (ret []MPEG4ProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	var e C.OMX_ERRORTYPE

	for i := uint(0); ; i++ {
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nProfileIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelQuerySupported, unsafe.Pointer(&p))
		if e != C.OMX_ErrorNone {
			break
		}
		ret = append(ret, MPEG4ProfileLevel{
			Profile: VideoMPEG4Profile(p.eProfile),
			Level:   VideoMPEG4Level(p.eLevel),
		})
	}

	if e != C.OMX_ErrorNoMore {
		err = Error(e)
	}
	return
}

func (c ComponentPort) SetVideoWMV(v VideoWMVFormat) error {
	var p C.OMX_VIDEO_PARAM_WMVTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eFormat = C.OMX_VIDEO_WMVFORMATTYPE(v)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoWmv, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoWMV() (ret VideoWMVFormat, err error) {
	var p C.OMX_VIDEO_PARAM_WMVTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoWmv, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret = VideoWMVFormat(p.eFormat)
	return
}

type VideoRV struct {
	Format                      VideoRVFormat
	BitsPerPixel                uint16
	PaddedWidth                 uint16
	PaddedHeight                uint16
	FrameRate                   uint
	BitstreamFlags              uint
	BitstreamVersion            uint
	MaxEncodeFrameSize          uint
	EnablePostFilter            bool
	EnableTemporalInterpolation bool
	EnableLatencyMode           bool
}

func (c ComponentPort) SetVideoRV(v VideoRV) error {
	var p C.OMX_VIDEO_PARAM_RVTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eFormat = C.OMX_VIDEO_RVFORMATTYPE(v.Format)
	p.nBitsPerPixel = C.OMX_U16(v.BitsPerPixel)
	p.nPaddedWidth = C.OMX_U16(v.PaddedWidth)
	p.nPaddedHeight = C.OMX_U16(v.PaddedHeight)
	p.nFrameRate = C.OMX_U32(v.FrameRate)
	p.nBitstreamFlags = C.OMX_U32(v.BitstreamFlags)
	p.nBitstreamVersion = C.OMX_U32(v.BitstreamVersion)
	p.nMaxEncodeFrameSize = C.OMX_U32(v.MaxEncodeFrameSize)
	p.bEnablePostFilter = toOMXBool(v.EnablePostFilter)
	p.bEnableTemporalInterpolation = toOMXBool(v.EnableTemporalInterpolation)
	p.bEnableLatencyMode = toOMXBool(v.EnableLatencyMode)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoRv, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoRV() (ret VideoRV, err error) {
	var p C.OMX_VIDEO_PARAM_RVTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoRv, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	ret.Format = VideoRVFormat(p.eFormat)
	ret.BitsPerPixel = uint16(p.nBitsPerPixel)
	ret.PaddedWidth = uint16(p.nPaddedWidth)
	ret.PaddedHeight = uint16(p.nPaddedHeight)
	ret.FrameRate = uint(p.nFrameRate)
	ret.BitstreamFlags = uint(p.nBitstreamFlags)
	ret.BitstreamVersion = uint(p.nBitstreamVersion)
	ret.MaxEncodeFrameSize = uint(p.nMaxEncodeFrameSize)
	ret.EnablePostFilter = (p.bEnablePostFilter != C.OMX_FALSE)
	ret.EnableTemporalInterpolation = (p.bEnableTemporalInterpolation != C.OMX_FALSE)
	ret.EnableLatencyMode = (p.bEnableLatencyMode != C.OMX_FALSE)

	return
}

type VideoAVC struct {
	SliceHeaderSpacing       uint
	PFrames                  uint
	BFrames                  uint
	UseHadamard              bool
	RefFrames                uint
	RefIdx10ActiveMinus1     uint
	RefIdx11ActiveMinus1     uint
	EnableUEP                bool
	EnableFMO                bool
	EnableASO                bool
	EnableRS                 bool
	Profile                  VideoAVCProfile
	Level                    VideoAVCLevel
	AllowedPictureTypes      uint
	FrameMBsOnly             bool
	MBAFF                    bool
	EntropyCodingCABAC       bool
	WeightedPPrediction      bool
	WeightedBipredicitonMode uint
	constIpred               bool
	Direct8x8Inference       bool
	DirectSpatialTemporal    bool
	CabacInitIdc             uint
	LoopFilterMode           VideoAVCLoopFilter
}

func (c ComponentPort) SetVideoAVC(v VideoAVC) error {
	var p C.OMX_VIDEO_PARAM_AVCTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nSliceHeaderSpacing = C.OMX_U32(v.SliceHeaderSpacing)
	p.nPFrames = C.OMX_U32(v.PFrames)
	p.nBFrames = C.OMX_U32(v.BFrames)
	p.bUseHadamard = toOMXBool(v.UseHadamard)
	p.nRefFrames = C.OMX_U32(v.RefFrames)
	p.nRefIdx10ActiveMinus1 = C.OMX_U32(v.RefIdx10ActiveMinus1)
	p.nRefIdx11ActiveMinus1 = C.OMX_U32(v.RefIdx11ActiveMinus1)
	p.bEnableUEP = toOMXBool(v.EnableUEP)
	p.bEnableFMO = toOMXBool(v.EnableFMO)
	p.bEnableASO = toOMXBool(v.EnableASO)
	p.bEnableRS = toOMXBool(v.EnableRS)
	p.eProfile = C.OMX_VIDEO_AVCPROFILETYPE(v.Profile)
	p.eLevel = C.OMX_VIDEO_AVCLEVELTYPE(v.Level)
	p.nAllowedPictureTypes = C.OMX_U32(v.AllowedPictureTypes)
	p.bFrameMBsOnly = toOMXBool(v.FrameMBsOnly)
	p.bMBAFF = toOMXBool(v.MBAFF)
	p.bEntropyCodingCABAC = toOMXBool(v.EntropyCodingCABAC)
	p.bWeightedPPrediction = toOMXBool(v.WeightedPPrediction)
	p.nWeightedBipredicitonMode = C.OMX_U32(v.WeightedBipredicitonMode)
	p.bconstIpred = toOMXBool(v.constIpred)
	p.bDirect8x8Inference = toOMXBool(v.Direct8x8Inference)
	p.bDirectSpatialTemporal = toOMXBool(v.DirectSpatialTemporal)
	p.nCabacInitIdc = C.OMX_U32(v.CabacInitIdc)
	p.eLoopFilterMode = C.OMX_VIDEO_AVCLOOPFILTERTYPE(v.LoopFilterMode)

	e := C.set_parameter(c.component.component, C.OMX_IndexParamVideoAvc, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) GetVideoAVC() (v VideoAVC, err error) {
	var p C.OMX_VIDEO_PARAM_AVCTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoAvc, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
		return
	}

	v.SliceHeaderSpacing = uint(p.nSliceHeaderSpacing)
	v.PFrames = uint(p.nPFrames)
	v.BFrames = uint(p.nBFrames)
	v.UseHadamard = (p.bUseHadamard != C.OMX_FALSE)
	v.RefFrames = uint(p.nRefFrames)
	v.RefIdx10ActiveMinus1 = uint(p.nRefIdx10ActiveMinus1)
	v.RefIdx11ActiveMinus1 = uint(p.nRefIdx11ActiveMinus1)
	v.EnableUEP = (p.bEnableUEP != C.OMX_FALSE)
	v.EnableFMO = (p.bEnableFMO != C.OMX_FALSE)
	v.EnableASO = (p.bEnableASO != C.OMX_FALSE)
	v.EnableRS = (p.bEnableRS != C.OMX_FALSE)
	v.Profile = VideoAVCProfile(p.eProfile)
	v.Level = VideoAVCLevel(p.eLevel)
	v.AllowedPictureTypes = uint(p.nAllowedPictureTypes)
	v.FrameMBsOnly = (p.bFrameMBsOnly != C.OMX_FALSE)
	v.MBAFF = (p.bMBAFF != C.OMX_FALSE)
	v.EntropyCodingCABAC = (p.bEntropyCodingCABAC != C.OMX_FALSE)
	v.WeightedPPrediction = (p.bWeightedPPrediction != C.OMX_FALSE)
	v.WeightedBipredicitonMode = uint(p.nWeightedBipredicitonMode)
	v.constIpred = (p.bconstIpred != C.OMX_FALSE)
	v.Direct8x8Inference = (p.bDirect8x8Inference != C.OMX_FALSE)
	v.DirectSpatialTemporal = (p.bDirectSpatialTemporal != C.OMX_FALSE)
	v.CabacInitIdc = uint(p.nCabacInitIdc)
	v.LoopFilterMode = VideoAVCLoopFilter(p.eLoopFilterMode)

	return
}

type AVCProfileLevel struct {
	Profile VideoAVCProfile
	Level   VideoAVCLevel
}

func (c ComponentPort) CurrentAVCProfileLevel() (ret AVCProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelCurrent, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		err = Error(e)
	} else {
		ret.Profile = VideoAVCProfile(p.eProfile)
		ret.Level = VideoAVCLevel(p.eLevel)
	}
	return
}

func (c ComponentPort) SupportedAVCProfileLevels() (ret []AVCProfileLevel, err error) {
	var p C.OMX_VIDEO_PARAM_PROFILELEVELTYPE
	var e C.OMX_ERRORTYPE

	for i := uint(0); ; i++ {
		C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
		p.nPortIndex = C.OMX_U32(c.port)
		p.nProfileIndex = C.OMX_U32(i)

		e = C.get_parameter(c.component.component, C.OMX_IndexParamVideoProfileLevelQuerySupported, unsafe.Pointer(&p))
		if e != C.OMX_ErrorNone {
			break
		}
		ret = append(ret, AVCProfileLevel{
			Profile: VideoAVCProfile(p.eProfile),
			Level:   VideoAVCLevel(p.eLevel),
		})
	}

	if e != C.OMX_ErrorNoMore {
		err = Error(e)
	}
	return
}

func (c ComponentPort) SetBitrate(bitrate uint) error {
	var p C.OMX_VIDEO_CONFIG_BITRATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nEncodeBitrate = C.OMX_U32(bitrate)

	e := C.set_config(c.component.component, C.OMX_IndexConfigVideoBitrate, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) Bitrate() (uint, error) {
	var p C.OMX_VIDEO_CONFIG_BITRATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_config(c.component.component, C.OMX_IndexConfigVideoBitrate, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return 0, Error(e)
	}
	return uint(p.nEncodeBitrate), nil
}

func (c ComponentPort) SetFramerate(framerate float64) error {
	var p C.OMX_CONFIG_FRAMERATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.xEncodeFramerate = toQ16(framerate)

	e := C.set_config(c.component.component, C.OMX_IndexConfigVideoFramerate, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) Framerate() (float64, error) {
	var p C.OMX_CONFIG_FRAMERATETYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_config(c.component.component, C.OMX_IndexConfigVideoFramerate, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return 0., Error(e)
	}
	return fromQ16(p.xEncodeFramerate), nil
}

func (c ComponentPort) SetIntraRefreshVOP(intraRefreshVOP bool) error {
	var p C.OMX_CONFIG_INTRAREFRESHVOPTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.IntraRefreshVOP = toOMXBool(intraRefreshVOP)

	e := C.set_config(c.component.component, C.OMX_IndexConfigVideoIntraVOPRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) IntraRefreshVOP() (bool, error) {
	var p C.OMX_CONFIG_INTRAREFRESHVOPTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_config(c.component.component, C.OMX_IndexConfigVideoIntraVOPRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return false, Error(e)
	}
	return (p.IntraRefreshVOP != C.OMX_FALSE), nil
}

type MacroBlockErrorMap struct {
	MapSize uint
	Hint    [1]uint8
}

func (c ComponentPort) SetMacroBlockErrorMap(m MacroBlockErrorMap) error {
	var p C.OMX_CONFIG_MACROBLOCKERRORMAPTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.nErrMapSize = C.OMX_U32(m.MapSize)
	p.ErrMap[0] = C.OMX_U8(m.Hint[0])

	e := C.set_config(c.component.component, C.OMX_IndexConfigVideoIntraMBRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) MacroBlockErrorMap() (m MacroBlockErrorMap, err error) {
	var p C.OMX_CONFIG_MACROBLOCKERRORMAPTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_config(c.component.component, C.OMX_IndexConfigVideoIntraMBRefresh, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return m, Error(e)
	}

	m.MapSize = uint(p.nErrMapSize)
	m.Hint[0] = uint8(p.ErrMap[0])
	return
}

func (c ComponentPort) SetMacroBlockErrorReporting(m bool) error {
	var p C.OMX_CONFIG_MBERRORREPORTINGTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.bEnabled = toOMXBool(m)

	e := C.set_config(c.component.component, C.OMX_IndexConfigVideoMBErrorReporting, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c ComponentPort) MacroBlockErrorReporting() (m bool, err error) {
	var p C.OMX_CONFIG_MBERRORREPORTINGTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_config(c.component.component, C.OMX_IndexConfigVideoMBErrorReporting, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return m, Error(e)
	}

	m = (p.bEnabled != C.OMX_FALSE)
	return
}

// type FaceDetectionControl struct {
// 	Mode       FaceDetectionControlMode
// 	Frames     uint
// 	MaxRegions uint
// 	Quality    uint
// }

// func (c ComponentPort) SetFaceDetectionControl(m FaceDetectionControl) error {
// 	var p C.OMX_CONFIG_FACEDETECTIONCONTROLTYPE
// 	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
// 	p.nPortIndex = C.OMX_U32(c.port)
// 	p.eMode = C.OMX_FACEDETECTIONCONTROLTYPE(m.Mode)
// 	p.nMaxRegions = C.OMX_U32(m.MaxRegions)
// 	p.nQuality = C.OMX_U32(m.Quality)

// 	e := C.set_config(c.component.component, C.OMX_IndexConfigCommonFaceDetectionControl, unsafe.Pointer(&p))
// 	if e != C.OMX_ErrorNone {
// 		return Error(e)
// 	}
// 	return nil
// }

// func (c ComponentPort) FaceDetectionControl() (m FaceDetectionControl, err error) {
// 	var p C.OMX_CONFIG_FACEDETECTIONCONTROLTYPE
// 	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
// 	p.nPortIndex = C.OMX_U32(c.port)

// 	e := C.get_config(c.component.component, C.OMX_IndexConfigCommonFaceDetectionControl, unsafe.Pointer(&p))
// 	if e != C.OMX_ErrorNone {
// 		return m, Error(e)
// 	}

// 	m.Mode = FaceDetectionControlMode(p.eMode)
// 	m.Frames = uint(p.nFrames)
// 	m.MaxRegions = uint(p.nMaxRegions)
// 	m.Quality = uint(p.nQuality)
// 	return
// }

// type FaceRegion struct {
// 	Left   int16
// 	Top    int16
// 	Width  uint16
// 	Height uint16
// 	Flags  FaceRegionFlags
// }

// func (c ComponentPort) SetFaceRegion(m FaceRegion) error {
// 	var p C.OMX_FACEREGIONTYPE
// 	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
// 	p.nPortIndex = C.OMX_U32(c.port)
// 	p.eMode = C.OMX_FaceRegionTYPE(m.Mode)
// 	p.nMaxRegions = C.OMX_U32(m.MaxRegions)
// 	p.nQuality = C.OMX_U32(m.Quality)

// 	e := C.set_config(c.component.component, C.OMX_IndexConfigCommonFaceRegion, unsafe.Pointer(&p))
// 	if e != C.OMX_ErrorNone {
// 		return Error(e)
// 	}
// 	return nil
// }

// func (c ComponentPort) FaceRegion() (m FaceRegion, err error) {
// 	var p C.OMX_FACEREGIONTYPE
// 	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
// 	p.nPortIndex = C.OMX_U32(c.port)

// 	e := C.get_config(c.component.component, C.OMX_IndexConfigCommonFaceRegion, unsafe.Pointer(&p))
// 	if e != C.OMX_ErrorNone {
// 		return m, Error(e)
// 	}

// 	m.Mode = FaceRegionMode(p.eMode)
// 	m.Frames = uint(p.nFrames)
// 	m.MaxRegions = uint(p.nMaxRegions)
// 	m.Quality = uint(p.nQuality)
// 	return
// }

func (c *Component) RequestCallback(index Index, enable bool) error {
	var p C.OMX_CONFIG_REQUESTCALLBACKTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_ALL
	p.nIndex = C.OMX_INDEXTYPE(index)
	p.bEnable = toOMXBool(enable)

	e := C.set_config(c.component, C.OMX_IndexConfigRequestCallback, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

func (c *Component) SetCameraDeviceNumber(n uint) error {
	var p C.OMX_PARAM_U32TYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_ALL
	p.nU32 = C.OMX_U32(n)

	e := C.set_parameter(c.component, C.OMX_IndexParamCameraDeviceNumber, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}
	return nil
}

type AudioPortDefinition struct {
	MIMEType string
	// NativeRender unsafe.Pointer
	ErrorConcealment bool
	Encoding         AudioCoding
}

type VideoPortDefinition struct {
	MIMEType string
	// NativeRender unsafe.Pointer
	Width            uint
	Height           uint
	Stride           int
	SliceHeight      uint
	Bitrate          uint
	Framerate        float64
	ErrorConcealment bool
	Compression      VideoCoding
	Color            ColorFormat
	// NativeWindow unsafe.Pointer
}

type ImagePortDefinition struct {
	MIMEType string
	// NativeRender unsafe.Pointer
	Width            uint
	Height           uint
	Stride           int
	SliceHeight      uint
	ErrorConcealment bool
	Compression      ImageCoding
	Color            ColorFormat
	// NativeWindow unsafe.Pointer
}

type OtherPortDefinition struct {
	Format OtherFormat
}

type PortDefinition struct {
	Direction         Direction
	BufferCountActual uint
	BufferCountMin    uint
	BufferSize        uint
	Enabled           bool
	Populated         bool
	Domain            PortDomain
	BuffersContiguous bool
	BufferAlignment   uint
	Audio             *AudioPortDefinition
	Video             *VideoPortDefinition
	Image             *ImagePortDefinition
	Other             *OtherPortDefinition
}

func (c ComponentPort) GetPortDefinition() (d PortDefinition, err error) {
	var p C.OMX_PARAM_PORTDEFINITIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)

	e := C.get_parameter(c.component.component, C.OMX_IndexParamPortDefinition, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return d, Error(e)
	}

	d.Direction = Direction(p.eDir)
	d.BufferCountActual = uint(p.nBufferCountActual)
	d.BufferCountMin = uint(p.nBufferCountMin)
	d.BufferSize = uint(p.nBufferSize)
	d.Enabled = (p.bEnabled != C.OMX_FALSE)
	d.Populated = (p.bPopulated != C.OMX_FALSE)
	d.Domain = PortDomain(p.eDomain)
	d.BuffersContiguous = (p.bBuffersContiguous != C.OMX_FALSE)
	d.BufferAlignment = uint(p.nBufferAlignment)

	switch d.Domain {
	case PortDomainAudio:
		// cast pointer to union
		d.Audio = getAudioPortDefinition(&p)
	case PortDomainImage:
		d.Image = getImagePortDefinition(&p)
	case PortDomainOther:
		d.Other = getOtherPortDefinition(&p)
	case PortDomainVideo:
		d.Video = getVideoPortDefinition(&p)
	}

	return
}

func getAudioPortDefinition(q *C.OMX_PARAM_PORTDEFINITIONTYPE) *AudioPortDefinition {
	p := (*C.OMX_AUDIO_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	return &AudioPortDefinition{
		MIMEType:         C.GoString(p.cMIMEType),
		ErrorConcealment: (p.bFlagErrorConcealment != C.OMX_FALSE),
		Encoding:         AudioCoding(p.eEncoding),
	}
}

func getImagePortDefinition(q *C.OMX_PARAM_PORTDEFINITIONTYPE) *ImagePortDefinition {
	p := (*C.OMX_IMAGE_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	return &ImagePortDefinition{
		MIMEType:         C.GoString(p.cMIMEType),
		Width:            uint(p.nFrameWidth),
		Height:           uint(p.nFrameHeight),
		Stride:           int(p.nStride),
		SliceHeight:      uint(p.nSliceHeight),
		ErrorConcealment: (p.bFlagErrorConcealment != C.OMX_FALSE),
		Compression:      ImageCoding(p.eCompressionFormat),
		Color:            ColorFormat(p.eColorFormat),
	}
}

func getOtherPortDefinition(q *C.OMX_PARAM_PORTDEFINITIONTYPE) *OtherPortDefinition {
	p := (*C.OMX_OTHER_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	return &OtherPortDefinition{
		Format: OtherFormat(p.eFormat),
	}
}

func getVideoPortDefinition(q *C.OMX_PARAM_PORTDEFINITIONTYPE) *VideoPortDefinition {
	p := (*C.OMX_VIDEO_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	return &VideoPortDefinition{
		MIMEType:         C.GoString(p.cMIMEType),
		Width:            uint(p.nFrameWidth),
		Height:           uint(p.nFrameHeight),
		Stride:           int(p.nStride),
		SliceHeight:      uint(p.nSliceHeight),
		Bitrate:          uint(p.nBitrate),
		Framerate:        fromQ16(p.xFramerate),
		ErrorConcealment: (p.bFlagErrorConcealment != C.OMX_FALSE),
		Compression:      VideoCoding(p.eCompressionFormat),
		Color:            ColorFormat(p.eColorFormat),
	}
}

func (c ComponentPort) SetPortDefinition(d PortDefinition) error {
	var p C.OMX_PARAM_PORTDEFINITIONTYPE
	C.initialize_struct(unsafe.Pointer(&p), C.uint(unsafe.Sizeof(p)))
	p.nPortIndex = C.OMX_U32(c.port)
	p.eDir = C.OMX_DIRTYPE(d.Direction)
	p.nBufferCountActual = C.OMX_U32(d.BufferCountActual)
	p.nBufferCountMin = C.OMX_U32(d.BufferCountMin)
	p.nBufferSize = C.OMX_U32(d.BufferSize)
	p.bEnabled = toOMXBool(d.Enabled)
	p.bPopulated = toOMXBool(d.Populated)
	p.eDomain = C.OMX_PORTDOMAINTYPE(d.Domain)
	p.bBuffersContiguous = toOMXBool(d.BuffersContiguous)
	p.nBufferAlignment = C.OMX_U32(d.BufferAlignment)

	var cleanup func()

	if d.Audio != nil {
		cleanup = setAudioPortDefinition(d.Audio, &p)
	} else if d.Video != nil {
		cleanup = setVideoPortDefinition(d.Video, &p)
	} else if d.Image != nil {
		cleanup = setImagePortDefinition(d.Image, &p)
	} else if d.Other != nil {
		cleanup = setOtherPortDefinition(d.Other, &p)
	}

	e := C.set_parameter(c.component.component, C.OMX_IndexParamPortDefinition, unsafe.Pointer(&p))
	if e != C.OMX_ErrorNone {
		return Error(e)
	}

	// free any allocated memory
	if cleanup != nil {
		cleanup()
	}
	return nil
}

func setAudioPortDefinition(d *AudioPortDefinition, q *C.OMX_PARAM_PORTDEFINITIONTYPE) (cleanup func()) {
	p := (*C.OMX_AUDIO_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	p.cMIMEType = C.CString(d.MIMEType)
	p.bFlagErrorConcealment = toOMXBool(d.ErrorConcealment)
	p.eEncoding = C.OMX_AUDIO_CODINGTYPE(d.Encoding)

	return func() {
		C.free(unsafe.Pointer(p.cMIMEType))
	}
}

func setVideoPortDefinition(d *VideoPortDefinition, q *C.OMX_PARAM_PORTDEFINITIONTYPE) (cleanup func()) {
	p := (*C.OMX_VIDEO_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	p.cMIMEType = C.CString(d.MIMEType)
	p.nFrameWidth = C.OMX_U32(d.Width)
	p.nFrameHeight = C.OMX_U32(d.Height)
	p.nStride = C.OMX_S32(d.Stride)
	p.nSliceHeight = C.OMX_U32(d.SliceHeight)
	p.nBitrate = C.OMX_U32(d.Bitrate)
	p.xFramerate = toQ16(d.Framerate)
	p.bFlagErrorConcealment = toOMXBool(d.ErrorConcealment)
	p.eCompressionFormat = C.OMX_VIDEO_CODINGTYPE(d.Compression)
	p.eColorFormat = C.OMX_COLOR_FORMATTYPE(d.Color)

	return func() {
		C.free(unsafe.Pointer(p.cMIMEType))
	}
}

func setImagePortDefinition(d *ImagePortDefinition, q *C.OMX_PARAM_PORTDEFINITIONTYPE) (cleanup func()) {
	p := (*C.OMX_IMAGE_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	p.cMIMEType = C.CString(d.MIMEType)
	p.nFrameWidth = C.OMX_U32(d.Width)
	p.nFrameHeight = C.OMX_U32(d.Height)
	p.nStride = C.OMX_S32(d.Stride)
	p.nSliceHeight = C.OMX_U32(d.SliceHeight)
	p.bFlagErrorConcealment = toOMXBool(d.ErrorConcealment)
	p.eCompressionFormat = C.OMX_IMAGE_CODINGTYPE(d.Compression)
	p.eColorFormat = C.OMX_COLOR_FORMATTYPE(d.Color)

	return func() {
		C.free(unsafe.Pointer(p.cMIMEType))
	}
}

func setOtherPortDefinition(d *OtherPortDefinition, q *C.OMX_PARAM_PORTDEFINITIONTYPE) (cleanup func()) {
	p := (*C.OMX_OTHER_PORTDEFINITIONTYPE)(unsafe.Pointer(&q.format[0]))

	p.eFormat = C.OMX_OTHER_FORMATTYPE(d.Format)

	return nil
}
