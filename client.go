package ilclient

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

extern OMX_ERRORTYPE set_image_portformat(COMPONENT_T * comp, unsigned int port, unsigned int index,
	OMX_IMAGE_CODINGTYPE format, OMX_COLOR_FORMATTYPE color);
extern OMX_ERRORTYPE get_image_portformat(COMPONENT_T * comp, unsigned int port, unsigned int index,
	OMX_IMAGE_CODINGTYPE * format, OMX_COLOR_FORMATTYPE * color);
extern OMX_ERRORTYPE set_video_portformat(COMPONENT_T * comp, unsigned int port, unsigned int index,
	OMX_VIDEO_CODINGTYPE format, OMX_COLOR_FORMATTYPE color, OMX_U32 framerate);
extern OMX_ERRORTYPE get_video_portformat(COMPONENT_T * comp, unsigned int port, unsigned int index,
	OMX_VIDEO_CODINGTYPE * format, OMX_COLOR_FORMATTYPE * color, OMX_U32 * framerate);
extern OMX_ERRORTYPE set_video_quantization(COMPONENT_T * comp, unsigned int port, 
	OMX_U32 nQpI, OMX_U32 nQpP, OMX_U32 nQpB);
extern OMX_ERRORTYPE get_video_quantization(COMPONENT_T * comp, unsigned int port, 
    OMX_U32 * nQpI, OMX_U32 * nQpP, OMX_U32 * nQpB);
*/
import "C"

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"time"
	"unsafe"
)

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
	Compression ImagePortFormat
	Color       ColorFormat
}

func (c ComponentPort) SetImagePortFormat(formats []ImageFormat) error {
	for i, f := range formats {
		e := C.set_image_portformat(c.component.component, C.uint(c.port), C.uint(i),
			C.OMX_IMAGE_CODINGTYPE(f.Compression), C.OMX_COLOR_FORMATTYPE(f.Color))

		if e != C.OMX_ErrorNone {
			return Error(e)
		}
	}

	return nil
}

func (c ComponentPort) GetImagePortFormat() ([]ImageFormat, error) {
	var ret []ImageFormat
	var e C.OMX_ERRORTYPE

	for i := uint(0); true; i++ {
		var coding C.OMX_IMAGE_CODINGTYPE
		var color C.OMX_COLOR_FORMATTYPE

		fmt.Fprintf(os.Stderr, "getting image format: %s: %d\n", c.port, int(c.port))
		e = C.get_image_portformat(c.component.component, C.uint(c.port), C.uint(i),
			&coding, &color)

		if e == C.OMX_ErrorNone {
			ret = append(ret, ImageFormat{ImagePortFormat(coding), ColorFormat(color)})
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
		e := C.set_video_portformat(c.component.component, C.uint(c.port), C.uint(i),
			C.OMX_VIDEO_CODINGTYPE(f.Compression), C.OMX_COLOR_FORMATTYPE(f.Color), toQ16(f.Framerate))

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
		var coding C.OMX_VIDEO_CODINGTYPE
		var color C.OMX_COLOR_FORMATTYPE
		var framerate C.OMX_U32

		e = C.get_video_portformat(c.component.component, C.uint(c.port), C.uint(i),
			&coding, &color, &framerate)

		if e == C.OMX_ErrorNone {
			ret = append(ret, VideoFormat{
				VideoCoding(coding),
				ColorFormat(color),
				fromQ16(framerate),
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
	if e := C.set_video_quantization(c.component.component, C.uint(c.port),
		C.OMX_U32(q.QpI), C.OMX_U32(q.QpP), C.OMX_U32(q.QpB));
		e != C.OMX_ErrorNone {

		return Error(e)
	}
	return nil
} 

func (c ComponentPort) GetVideoQuantization() (VideoQuantization, error) {
	var nQpI, nQpP, nQpB C.OMX_U32

	ret := VideoQuantization{}

	if e := C.get_video_quantization(c.component.component, C.uint(c.port),
		&nQpI, &nQpP, &nQpB);
		e != C.OMX_ErrorNone {

		return ret, Error(e)
	}
	ret.QpI = uint(nQpI)
	ret.QpP = uint(nQpP)
	ret.QpB = uint(nQpB)

	return ret, nil
} 

type VideoFastUpdate struct {
	Enabled bool
	FirstGOB uint
	FirstMB uint
	NumMB uint
}

func (c ComponentPort) SetVideoFastUpdate(v VideoFastUpdate) error {
	return nil
}


