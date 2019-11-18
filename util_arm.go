package ilclient

/*
#cgo CFLAGS: -Wno-unused-variable -Wall -Wno-deprecated -g -DRASPBERRY_PI -DSTANDALONE -D__STDC_CONSTANT_MACROS  -D__STDC_LIMIT_MACROS -DTARGET_POSIX -D_LINUX -fPIC -DPIC -D_REENTRANT -D_LARGEFILE64_SOURCE -D_FILE_OFFSET_BITS=64 -g -DHAVE_LIBOPENMAX=2 -DOMX -DOMX_SKIP64BIT -pipe -DUSE_EXTERNAL_OMX -DHAVE_LIBBCM_HOST -DUSE_EXTERNAL_LIBBCM_HOST -DUSE_VCHIQ_ARM -I/opt/vc/include/IL -I/opt/vc/include -I/opt/vc/include/interface/vcos/pthreads -I/opt/vc/include/interface/vmcs_host/linux/ -I/opt/vc/src/hello_pi/libs/ilclient
#cgo LDFLAGS: -L /opt/vc/lib -lopenmaxil -lbcm_host -lvcos -lvchiq_arm -lpthread -lrt -L/opt/vc/src/hello_pi/libs/ilclient -lilclient

#include <OMX_Core.h>
#include <OMX_Component.h>

#include <bcm_host.h>
#include <ilclient.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

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
	StateIdle State = C.OMX_StateIdle
	StateLoaded State = C.OMX_StateLoaded
	StateInvalid State = C.OMX_StateInvalid
	StateExecuting State = C.OMX_StateExecuting
	StatePause State = C.OMX_StatePause
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
