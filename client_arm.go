package ilclient

/*
#cgo CFLAGS: -Wno-unused-variable -Wall -Wno-deprecated -g -DRASPBERRY_PI -DSTANDALONE -D__STDC_CONSTANT_MACROS  -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -I/opt/vc/include/IL -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux/ -I/opt/vc/src/hello_pi/libs/ilclient
#cgo LDFLAGS: -L /opt/vc/lib -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt -L/opt/vc/src/hello_pi/libs/ilclient -lilclient

#include <OMX_Core.h>
#include <OMX_Component.h>

#include <bcm_host.h>
#include <ilclient.h>

extern void goErrorHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);
extern void goPortSettingsChangedHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);
extern void goEOSHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);
extern void goConfigChangedHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);
extern void goFillBufferHandler(void * userdata, COMPONENT_T * comp);
extern void goEmptyBufferHandler(void * userdata, COMPONENT_T * comp);

COMPONENT_T* ilclient_create_component_wrapper(ILCLIENT_T *handle, int * ret, char * name, ILCLIENT_CREATE_FLAGS_T flags) {
	COMPONENT_T * comp = NULL;
	*ret = ilclient_create_component(handle, &comp, name, flags);
	return comp;
}

int ilclient_enable_port_buffers_wrapper(COMPONENT_T * comp, int port_index) {
	return ilclient_enable_port_buffers(comp, port_index, NULL, NULL, NULL);
}
void ilclient_disable_port_buffers_wrapper(COMPONENT_T * comp, int port_index) {
	ilclient_disable_port_buffers(comp, port_index, NULL, NULL, NULL);
}
void ilclient_set_error_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_error_callback(handle, goErrorHandler, (void*)userdata);
}
void ilclient_set_port_settings_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_port_settings_callback(handle, goPortSettingsChangedHandler, (void*)userdata);
}
void ilclient_set_eos_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_eos_callback(handle, goEOSHandler, (void*)userdata);
}
void ilclient_set_configchanged_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_configchanged_callback(handle, goConfigChangedHandler, (void*)userdata);
}
void ilclient_set_fill_buffer_done_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_fill_buffer_done_callback(handle, goFillBufferHandler, (void*)userdata);
}
void ilclient_set_empty_buffer_done_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	ilclient_set_empty_buffer_done_callback(handle, goEmptyBufferHandler, (void*)userdata);
}
*/
import "C"

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const default_timeout = 0

type Client struct {
	client     *C.ILCLIENT_T
	Timeout    time.Duration
	components map[*C.COMPONENT_T]*Component
	tunnels    map[*C.TUNNEL_T]*Tunnel
	clientID   C.int
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
	clients     map[C.int]*Client
	clientID    C.int
	clientsLock sync.Mutex
)

func init() {
	C.bcm_host_init()

	C.OMX_Init()

	clientID = 0
	clients = make(map[C.int]*Client)
}

func New() *Client {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	c := &Client{
		client:  C.ilclient_init(),
		Timeout: default_timeout,

		components: make(map[*C.COMPONENT_T]*Component),
		tunnels:    make(map[*C.TUNNEL_T]*Tunnel),
		clientID:   clientID,
	}
	clientID++
	clients[c.clientID] = c

	C.ilclient_set_error_callback_wrapper(c.client, &c.clientID)
	C.ilclient_set_port_settings_callback_wrapper(c.client, &c.clientID)
	C.ilclient_set_eos_callback_wrapper(c.client, &c.clientID)
	C.ilclient_set_configchanged_callback_wrapper(c.client, &c.clientID)
	C.ilclient_set_fill_buffer_done_callback_wrapper(c.client, &c.clientID)
	C.ilclient_set_empty_buffer_done_callback_wrapper(c.client, &c.clientID)

	return c
}

func (c *Client) NewComponent(name string, flags ...CreateFlag) (*Component, error) {
	ret := &Component{}
	var e C.int

	var fin C.ILCLIENT_CREATE_FLAGS_T
	for f := range flags {
		fin = fin | C.ILCLIENT_CREATE_FLAGS_T(f)
	}

	ret.component = C.ilclient_create_component_wrapper(c.client, &e, C.CString(name), fin)
	if e != 0 {
		return nil, fmt.Errorf("ilclient: could not create component: %v", Error(e))
	}
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
	c.tunnels[t] = ret

	return ret, nil
}

func (c *Client) Close() {
	// cleanup tunnels
	// list of null terminated tunnels
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

	clientsLock.Lock()
	defer clientsLock.Unlock()

	delete(clients, c.clientID)
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
	buf := C.ilclient_get_output_buffer(c.component, C.int(port_index), 0)
	if buf == nil {
		return nil, fmt.Errorf("output buffer not available for port %d", port_index)
	}
	return &Buffer{buf}, nil
}

func (c *Component) InputBuffer(port_index int) (*Buffer, error) {
	buf := C.ilclient_get_input_buffer(c.component, C.int(port_index), 0)
	if buf == nil {
		return nil, fmt.Errorf("input buffer not available for port %d", port_index)
	}
	return &Buffer{buf}, nil
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
	if e := C.ilclient_enable_port_buffers_wrapper(c.component, C.int(port_index)); e != 0 {
		return fmt.Errorf("error: EnablePortBuffers: %v", Error(e))
	}
	return nil
}

func (c *Component) DisablePortBuffers(port_index int) {
	C.ilclient_disable_port_buffers_wrapper(c.component, C.int(port_index))
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
