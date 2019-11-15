package main

/*
#cgo CFLAGS: -Wall -Wno-deprecated -g -DRASPBERRY_PI -DSTANDALONE -D__STDC_CONSTANT_MACROS  -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -I/opt/vc/include/IL -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux/ -I/opt/vc/src/hello_pi/libs/ilclient
#cgo LDFLAGS: -L/opt/vc/lib/ -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt -L/opt/vc/src/hello_pi/libs/ilclient -lilclient

#include <OMX_Core.h>
#include <OMX_Component.h>

#include <bcm_host.h>
#include <ilclient.h>

extern int goErrorHandler(void * userdata, COMPONENT_T * comp, OMX_U32 data);

void error_handler(void * userdata, COMPONENT_T *comp, OMX_U32 data) {
	goErrorHandler(userdata, comp, data);
}
*/
import "C"


import (
	"fmt"
	"unsafe"
	"log"
)

type Client struct {
	client *C.ILCLIENT_T
}
type Event struct{}
type Component struct{}
type Tunnel struct{
	tunnel *C.TUNNEL_T
}

func goErrorHandler(userdata unsafe.Pointer, comp * C.COMPONENT_T, data C.OMX_U32) C.int {
	client := (*Client)(userdata)

	client.handleError(comp, data);
}

func New() *Client {
	c := Client{
		client: C.ilclient_init(),
	}
	C.ilclient_set_error_callback(c.client, )
} 

func (c *Client) Close() {
	C.ilclient_destroy(c.client)
}

func (c *Client) handleError(comp * C.COMPONENT_T, data C.OMX_U32) {
	log.Printf("error!")
}

func main() {
	c := New()
	defer c.Close()

	fmt.Printf("%v\n", c)
}
