package ilclient

/*
#cgo CFLAGS: -Wall -Wno-deprecated -g -DRASPBERRY_PI -DSTANDALONE -D__STDC_CONSTANT_MACROS  -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -I/opt/vc/include/IL -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux/ -I/opt/vc/src/hello_pi/libs/ilclient
#cgo LDFLAGS: -L /opt/vc/lib -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt -L/opt/vc/src/hello_pi/libs/ilclient -lilclient -rdynamic -Wl,-rpath-link,/opt/vc/lib

#include <OMX_Core.h>
#include <OMX_Component.h>

#include <bcm_host.h>
#include <ilclient.h>

extern int goErrorHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);

void error_handler(void * userdata, COMPONENT_T *comp, OMX_U32 data) {
	goErrorHandler(userdata, comp, data);
}

COMPONENT_T* ilclient_create_component_wrapper(ILCLIENT_T *handle, int * ret, char * name, ILCLIENT_CREATE_FLAGS_T flags) {
	COMPONENT_T * comp = NULL;
	*ret = ilclient_create_component(handle, &comp, name, flags);
	return comp;
}
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"
)

type Client struct {
	client *C.ILCLIENT_T
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

func init() {
	C.bcm_host_init()

	C.OMX_Init()
}

func goErrorHandler(userdata unsafe.Pointer, comp *C.COMPONENT_T, data C.OMX_U32) C.int {
	client := (*Client)(userdata)

	return C.int(client.handleError(comp, data))
}

func New() *Client {
	c := &Client{
		client: C.ilclient_init(),
	}
	return c
}

func (c *Client) NewComponent(name string, flags C.ILCLIENT_CREATE_FLAGS_T) (*Component, error) {
	ret := &Component{}
	var e C.int
	ret.component = C.ilclient_create_component_wrapper(c.client, &e, C.CString(name), flags)

	if e != 0 {
		return nil, fmt.Errorf("ilclient: could not create component")
	}

	return ret, nil
}

func (c *Client) Close() {
	C.ilclient_destroy(c.client)
}

func (c *Client) handleError(comp *C.COMPONENT_T, data C.OMX_U32) int {
	log.Printf("error!")
	return 0
}

func (c *Component) SetState(state C.OMX_STATETYPE) error {
	if ret := C.ilclient_change_component_state(c.component, state); ret != 0 {
		return fmt.Errorf("error changing component state")
	}
	return nil
}

func (c *Component) EnablePort(port_index int) {
	C.ilclient_enable_port(c.component, C.int(port_index))
}

func (c *Component) DisablePort(port_index int) {
	C.ilclient_disable_port(c.component, C.int(port_index))
}

func (c *Component) EnablePortBuffers(port_index int) {
	C.ilclient_enable_port_buffers(c.component, port_index, unsafe.Pointer(nil), unsafe.Pointer(nil), unsafe.Pointer(nil))
}

func main() {
	c := New()
	defer c.Close()

	fmt.Printf("%v\n", c)
}
