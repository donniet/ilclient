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

void setup_callbacks(ILCLIENT_T * handle) {
	ilclient_set_error_callback(handle, goErrorHandler, NULL);
	ilclient_set_port_settings_callback(handle, goPortSettingsChangedHandler, NULL);
	ilclient_set_eos_callback(handle, goEOSHandler, NULL);
	ilclient_set_configchanged_callback(handle, goConfigChangedHandler, NULL);
	ilclient_set_fill_buffer_done_callback(handle, goFillBufferHandler, NULL);
	ilclient_set_empty_buffer_done_callback(handle, goEmptyBufferHandler, NULL);
}

// important because OMX_GetState is a macro
int get_component_state(COMPONENT_T * comp, OMX_STATETYPE * state) {
	fprintf(stderr, "OMX_GetState\n");
	return OMX_GetState(ilclient_get_handle(comp), state);
}