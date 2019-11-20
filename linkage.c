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

#define OMX_INIT_STRUCTURE(a) \
    memset(&(a), 0, sizeof(a)); \
    (a).nSize = sizeof(a); \
    (a).nVersion.nVersion = OMX_VERSION; \
    (a).nVersion.s.nVersionMajor = OMX_VERSION_MAJOR; \
    (a).nVersion.s.nVersionMinor = OMX_VERSION_MINOR; \
    (a).nVersion.s.nRevision = OMX_VERSION_REVISION; \
    (a).nVersion.s.nStep = OMX_VERSION_STEP

typedef struct header_t {
    OMX_U32 nSize;
    OMX_VERSIONTYPE nVersion;
} header;

void initialize_struct(void * p, size_t size) {
    memset(p, 0, size);
    header * h = (header*)p;
    h->nSize = size;
    h->nVersion.s.nVersionMajor = OMX_VERSION_MAJOR;
    h->nVersion.s.nVersionMinor = OMX_VERSION_MINOR;
    h->nVersion.s.nRevision = OMX_VERSION_REVISION;
    h->nVersion.s.nStep = OMX_VERSION_STEP;
}

OMX_ERRORTYPE set_parameter(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param) {
    return OMX_SetParameter(ilclient_get_handle(comp), index, param);
}
OMX_ERRORTYPE get_parameter(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param) {
    return OMX_GetParameter(ilclient_get_handle(comp), index, param);
}
OMX_ERRORTYPE set_config(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param) {
    return OMX_SetConfig(ilclient_get_handle(comp), index, param);
}
OMX_ERRORTYPE get_config(COMPONENT_T * comp, OMX_INDEXTYPE index, void * param) {
    return OMX_GetConfig(ilclient_get_handle(comp), index, param);
}


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
