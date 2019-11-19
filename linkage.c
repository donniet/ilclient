#include <OMX_Core.h>
#include <OMX_Component.h>

#include <stdlib.h>
#include <stdio.h>
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
	fprintf(stderr, "ilclient_create_component\n");
	*ret = ilclient_create_component(handle, &comp, name, flags);
	return comp;
}

void enable_trace_logging() {
	fprintf(stderr, "VC_LOGLEVEL=%s\n", getenv("VC_LOGLEVEL"));

	putenv("VC_LOGLEVEL=ilclient:trace");

	fprintf(stderr, "VC_LOGLEVEL=%s\n", getenv("VC_LOGLEVEL"));
}

int ilclient_enable_port_buffers_wrapper(COMPONENT_T * comp, int port_index) {
	fprintf(stderr, "ilclient_enable_port_buffers\n");
	return ilclient_enable_port_buffers(comp, port_index, NULL, NULL, NULL);
}
void ilclient_disable_port_buffers_wrapper(COMPONENT_T * comp, int port_index) {
	fprintf(stderr, "ilclient_disable_port_buffers\n");
	ilclient_disable_port_buffers(comp, port_index, NULL, NULL, NULL);
}
void ilclient_set_error_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_error_callback\n");
	ilclient_set_error_callback(handle, goErrorHandler, (void*)userdata);
}
void ilclient_set_port_settings_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_port_settings_callback\n");
	ilclient_set_port_settings_callback(handle, goPortSettingsChangedHandler, (void*)userdata);
}
void ilclient_set_eos_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_eos_callback\n");
	ilclient_set_eos_callback(handle, goEOSHandler, (void*)userdata);
}
void ilclient_set_configchanged_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_configchanged_callback\n");
	ilclient_set_configchanged_callback(handle, goConfigChangedHandler, (void*)userdata);
}
void ilclient_set_fill_buffer_done_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_fill_buffer_done_callback\n");
	ilclient_set_fill_buffer_done_callback(handle, goFillBufferHandler, (void*)userdata);
}
void ilclient_set_empty_buffer_done_callback_wrapper(ILCLIENT_T * handle, int * userdata) {
	fprintf(stderr, "ilclient_set_empty_buffer_done_callback\n");
	ilclient_set_empty_buffer_done_callback(handle, goEmptyBufferHandler, (void*)userdata);
}
int get_component_state(COMPONENT_T * comp, OMX_STATETYPE * state) {
	fprintf(stderr, "OMX_GetState\n");
	return OMX_GetState(ilclient_get_handle(comp), state);
}