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
*/
import "C"

import (
	"fmt"
	"log"
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
	port      int
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

	var fin C.ILCLIENT_CREATE_FLAGS_T
	for f := range flags {
		fin = fin | C.ILCLIENT_CREATE_FLAGS_T(f)
	}

	str := C.CString(name)
	defer C.free(unsafe.Pointer(str))
	ret.component = C.ilclient_create_component_wrapper(c.client, &e, C.CString(name), fin)

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

		ret = append(ret, ComponentPort{c, int(p)})
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

func (c *Component) OutputBuffer(port_index int) (*Buffer, error) {
	fmt.Fprintf(os.Stderr, "ilclient_get_output_buffer\n")
	buf := C.ilclient_get_output_buffer(c.component, C.int(port_index), 0)
	if buf == nil {
		return nil, fmt.Errorf("output buffer not available for port %d", port_index)
	}
	return &Buffer{buf}, nil
}

func (c *Component) InputBuffer(port_index int) (*Buffer, error) {
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

func (c *Component) Port(port_index int) ComponentPort {
	return ComponentPort{c, port_index}
}

func (c *Component) EnablePort(port_index int) {
	C.ilclient_enable_port(c.component, C.int(port_index))
}

func (c *Component) DisablePort(port_index int) {
	C.ilclient_disable_port(c.component, C.int(port_index))
}

func (c *Component) EnablePortBuffers(port_index int) error {
	if e := C.ilclient_enable_port_buffers(c.component, C.int(port_index), nil, nil, nil); e != 0 {
		return fmt.Errorf("error: EnablePortBuffers: %v", Error(e))
	}
	return nil
}

func (c *Component) DisablePortBuffers(port_index int) {
	C.ilclient_disable_port_buffers(c.component, C.int(port_index), nil, nil, nil)
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

type State C.OMX_STATETYPE

func (s State) String() string {
	switch C.OMX_STATETYPE(s) {
	case C.OMX_StateInvalid:
		return "OMX_StateInvalid"
	case C.OMX_StateLoaded:
		return "OMX_StateLoaded"
	case C.OMX_StateIdle:
		return "OMX_StateIdle"
	case C.OMX_StateExecuting:
		return "OMX_StateExecuting"
	case C.OMX_StatePause:
		return "OMX_StatePause"
	case C.OMX_StateWaitForResources:
		return "OMX_StateWaitForResources"
	}
	return fmt.Sprintf("UNKONWN %v", int(s))
}

const (
	StateIdle             State = C.OMX_StateIdle
	StateLoaded           State = C.OMX_StateLoaded
	StateInvalid          State = C.OMX_StateInvalid
	StateExecuting        State = C.OMX_StateExecuting
	StatePause            State = C.OMX_StatePause
	StateWaitForResources State = C.OMX_StateWaitForResources
)

type TunnelError int

const (
	TunnelErrorNone              TunnelError = 0
	TunnelErrorTimeout           TunnelError = -1
	TunnelErrorParameter         TunnelError = -2
	TunnelErrorNoStreams         TunnelError = -3
	TunnelErrorStreamUnavailable TunnelError = -4
	TunnelErrorDataFormat        TunnelError = -5
	TunnelErrorNoEnable          TunnelError = -0x7fff
)

func (e TunnelError) String() string {
	switch e {
	case TunnelErrorNone:
		return "TunnelErrorNone"
	case TunnelErrorTimeout:
		return "TunnelErrorTimeout"
	case TunnelErrorParameter:
		return "TunnelErrorParameter"
	case TunnelErrorNoStreams:
		return "TunnelErrorNoStreams"
	case TunnelErrorStreamUnavailable:
		return "TunnelErrorStreamUnavailable"
	case TunnelErrorDataFormat:
		return "TunnelErrorDataFormat"
	case TunnelErrorNoEnable:
		return "TunnelErrorNoEnable"
	}
	return fmt.Sprintf("UNKONWN %x", int(e))
}
func (e TunnelError) Error() string {
	return e.String()
}

type Error C.OMX_ERRORTYPE

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	switch C.OMX_ERRORTYPE(e) {
	case C.OMX_ErrorNone:
		return "OMX_ErrorNone"
	case C.OMX_ErrorInsufficientResources:
		return "OMX_ErrorInsufficientResources"
	case C.OMX_ErrorUndefined:
		return "OMX_ErrorUndefined"
	case C.OMX_ErrorInvalidComponentName:
		return "OMX_ErrorInvalidComponentName"
	case C.OMX_ErrorComponentNotFound:
		return "OMX_ErrorComponentNotFound"
	case C.OMX_ErrorInvalidComponent:
		return "OMX_ErrorInvalidComponent"
	case C.OMX_ErrorBadParameter:
		return "OMX_ErrorBadParameter"
	case C.OMX_ErrorNotImplemented:
		return "OMX_ErrorNotImplemented"
	case C.OMX_ErrorUnderflow:
		return "OMX_ErrorUnderflow"
	case C.OMX_ErrorOverflow:
		return "OMX_ErrorOverflow"
	case C.OMX_ErrorHardware:
		return "OMX_ErrorHardware"
	case C.OMX_ErrorInvalidState:
		return "OMX_ErrorInvalidState"
	case C.OMX_ErrorStreamCorrupt:
		return "OMX_ErrorStreamCorrupt"
	case C.OMX_ErrorPortsNotCompatible:
		return "OMX_ErrorPortsNotCompatible"
	case C.OMX_ErrorResourcesLost:
		return "OMX_ErrorResourcesLost"
	case C.OMX_ErrorNoMore:
		return "OMX_ErrorNoMore"
	case C.OMX_ErrorVersionMismatch:
		return "OMX_ErrorVersionMismatch"
	case C.OMX_ErrorNotReady:
		return "OMX_ErrorNotReady"
	case C.OMX_ErrorTimeout:
		return "OMX_ErrorTimeout"
	case C.OMX_ErrorSameState:
		return "OMX_ErrorSameState"
	case C.OMX_ErrorResourcesPreempted:
		return "OMX_ErrorResourcesPreempted"
	case C.OMX_ErrorPortUnresponsiveDuringAllocation:
		return "OMX_ErrorPortUnresponsiveDuringAllocation"
	case C.OMX_ErrorPortUnresponsiveDuringDeallocation:
		return "OMX_ErrorPortUnresponsiveDuringDeallocation"
	case C.OMX_ErrorPortUnresponsiveDuringStop:
		return "OMX_ErrorPortUnresponsiveDuringStop"
	case C.OMX_ErrorIncorrectStateTransition:
		return "OMX_ErrorIncorrectStateTransition"
	case C.OMX_ErrorIncorrectStateOperation:
		return "OMX_ErrorIncorrectStateOperation"
	case C.OMX_ErrorUnsupportedSetting:
		return "OMX_ErrorUnsupportedSetting"
	case C.OMX_ErrorUnsupportedIndex:
		return "OMX_ErrorUnsupportedIndex"
	case C.OMX_ErrorBadPortIndex:
		return "OMX_ErrorBadPortIndex"
	case C.OMX_ErrorPortUnpopulated:
		return "OMX_ErrorPortUnpopulated"
	case C.OMX_ErrorComponentSuspended:
		return "OMX_ErrorComponentSuspended"
	case C.OMX_ErrorDynamicResourcesUnavailable:
		return "OMX_ErrorDynamicResourcesUnavailable"
	case C.OMX_ErrorMbErrorsInFrame:
		return "OMX_ErrorMbErrorsInFrame"
	case C.OMX_ErrorFormatNotDetected:
		return "OMX_ErrorFormatNotDetected"
	case C.OMX_ErrorContentPipeOpenFailed:
		return "OMX_ErrorContentPipeOpenFailed"
	case C.OMX_ErrorContentPipeCreationFailed:
		return "OMX_ErrorContentPipeCreationFailed"
	case C.OMX_ErrorSeperateTablesUsed:
		return "OMX_ErrorSeperateTablesUsed"
	case C.OMX_ErrorTunnelingUnsupported:
		return "OMX_ErrorTunnelingUnsupported"
	case C.OMX_ErrorDiskFull:
		return "OMX_ErrorDiskFull"
	case C.OMX_ErrorMaxFileSize:
		return "OMX_ErrorMaxFileSize"
	case C.OMX_ErrorDrmUnauthorised:
		return "OMX_ErrorDrmUnauthorised"
	case C.OMX_ErrorDrmExpired:
		return "OMX_ErrorDrmExpired"
	case C.OMX_ErrorDrmGeneral:
		return "OMX_ErrorDrmGeneral"
	}
	return fmt.Sprintf("UNKNOWN %x", int(e))
}

type CreateFlag C.ILCLIENT_CREATE_FLAGS_T

const (
	CreateFlagNone                CreateFlag = 0x0
	CreateFlagEnableInputBuffers  CreateFlag = 0x1
	CreateFlagEnableOutputBuffers CreateFlag = 0x2
	CreateFlagDisableAllPorts     CreateFlag = 0x4
	CreateFlagHostComponent       CreateFlag = 0x8
	CreateFlagOutputZeroBuffers   CreateFlag = 0x10
)

func (f CreateFlag) String() string {
	switch f {
	case CreateFlagNone:
		return "CreateFlagNone"
	case CreateFlagEnableInputBuffers:
		return "CreateFlagEnableInputBuffers"
	case CreateFlagEnableOutputBuffers:
		return "CreateFlagEnableOutputBuffers"
	case CreateFlagDisableAllPorts:
		return "CreateFlagDisableAllPorts"
	case CreateFlagHostComponent:
		return "CreateFlagHostComponent"
	case CreateFlagOutputZeroBuffers:
		return "CreateFlagOutputZeroBuffers"
	}
	return fmt.Sprintf("UNKNOWN: %x", int(f))
}
