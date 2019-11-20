package ilclient

/*
#include <OMX_Core.h>
#include <OMX_Component.h>

#include <stdlib.h>
#include <stdio.h>
#include <bcm_host.h>
#include <ilclient.h>
*/
import "C"

import (
	"fmt"
)

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
	CreateFlagNone                CreateFlag = C.ILCLIENT_FLAGS_NONE
	CreateFlagEnableInputBuffers  CreateFlag = C.ILCLIENT_ENABLE_INPUT_BUFFERS
	CreateFlagEnableOutputBuffers CreateFlag = C.ILCLIENT_ENABLE_OUTPUT_BUFFERS
	CreateFlagDisableAllPorts     CreateFlag = C.ILCLIENT_DISABLE_ALL_PORTS
	CreateFlagHostComponent       CreateFlag = C.ILCLIENT_HOST_COMPONENT
	CreateFlagOutputZeroBuffers   CreateFlag = C.ILCLIENT_OUTPUT_ZERO_BUFFERS
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
	return fmt.Sprintf("UNKNOWN[%x]", int(f))
}

type PortIndex int

const (
	CameraPreviewOut         PortIndex = 70
	CameraCaptureOut         PortIndex = 71
	CameraStillCaptureOut    PortIndex = 72
	CameraClockIn            PortIndex = 73
	VideoSplitterInputIn     PortIndex = 250
	VideoSplitterOutput1Out  PortIndex = 251
	VideoSplitterOutput2Out  PortIndex = 252
	VideoSplitterOutput3Out  PortIndex = 253
	VideoSplitterOutput4Out  PortIndex = 254
	ImageEncodeRawPixelsIn   PortIndex = 340
	ImageEncodeCompressedOut PortIndex = 341
	VideoEncodeRawVideoIn    PortIndex = 200
	VideoEncodeCompressedOut PortIndex = 201
)

func (p PortIndex) String() string {
	switch p {
	case CameraPreviewOut:
		return "CameraPreviewOut"
	case CameraCaptureOut:
		return "CameraCaptureOut"
	case CameraStillCaptureOut:
		return "CameraStillCaptureOut"
	case CameraClockIn:
		return "CameraClockIn"
	case VideoSplitterInputIn:
		return "VideoSplitterInputIn"
	case VideoSplitterOutput1Out:
		return "VideoSplitterOutput1Out"
	case VideoSplitterOutput2Out:
		return "VideoSplitterOutput2Out"
	case VideoSplitterOutput3Out:
		return "VideoSplitterOutput3Out"
	case VideoSplitterOutput4Out:
		return "VideoSplitterOutput4Out"
	case ImageEncodeRawPixelsIn:
		return "ImageEncodeRawPixelsIn"
	case ImageEncodeCompressedOut:
		return "ImageEncodeCompressedOut"
	case VideoEncodeRawVideoIn:
		return "VideoEncodeRawVideoIn"
	case VideoEncodeCompressedOut:
		return "VideoEncodeCompressedOut"
	}
	return fmt.Sprintf("UNKNOWN %d", int(p))
}

type VideoCoding C.OMX_VIDEO_CODINGTYPE
const (
	VideoCodingUnused VideoCoding = C.OMX_VIDEO_CodingUnused
    VideoCodingAutoDetect VideoCoding = C.OMX_VIDEO_CodingAutoDetect
    VideoCodingMPEG2 VideoCoding = C.OMX_VIDEO_CodingMPEG2
    VideoCodingH263 VideoCoding = C.OMX_VIDEO_CodingH263
    VideoCodingMPEG4 VideoCoding = C.OMX_VIDEO_CodingMPEG4
    VideoCodingWMV VideoCoding = C.OMX_VIDEO_CodingWMV
    VideoCodingRV VideoCoding = C.OMX_VIDEO_CodingRV
    VideoCodingAVC VideoCoding = C.OMX_VIDEO_CodingAVC
    VideoCodingMJPEG VideoCoding = C.OMX_VIDEO_CodingMJPEG
)
func (v VideoCoding) String() string {
	switch v {
	case VideoCodingUnused: return "VideoCodingUnused"; 
    case VideoCodingAutoDetect: return "VideoCodingAutoDetect"; 
    case VideoCodingMPEG2: return "VideoCodingMPEG2"; 
    case VideoCodingH263: return "VideoCodingH263"; 
    case VideoCodingMPEG4: return "VideoCodingMPEG4"; 
    case VideoCodingWMV: return "VideoCodingWMV"; 
    case VideoCodingRV: return "VideoCodingRV"; 
    case VideoCodingAVC: return "VideoCodingAVC"; 
    case VideoCodingMJPEG: return "VideoCodingMJPEG"; 
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type ImagePortFormat C.OMX_IMAGE_CODINGTYPE

const (
	ImagePortFormatUnused     ImagePortFormat = C.OMX_IMAGE_CodingUnused
	ImagePortFormatAutodetect ImagePortFormat = C.OMX_IMAGE_CodingAutoDetect
	ImagePortFormatJPEG       ImagePortFormat = C.OMX_IMAGE_CodingJPEG
	ImagePortFormatJPEG2K     ImagePortFormat = C.OMX_IMAGE_CodingJPEG2K
	ImagePortFormatEXIF       ImagePortFormat = C.OMX_IMAGE_CodingEXIF
	ImagePortFormatTIFF       ImagePortFormat = C.OMX_IMAGE_CodingTIFF
	ImagePortFormatGIF        ImagePortFormat = C.OMX_IMAGE_CodingGIF
	ImagePortFormatPNG        ImagePortFormat = C.OMX_IMAGE_CodingPNG
	ImagePortFormatLZW        ImagePortFormat = C.OMX_IMAGE_CodingLZW
	ImagePortFormatBMP        ImagePortFormat = C.OMX_IMAGE_CodingBMP
	ImagePortFormatTGA        ImagePortFormat = C.OMX_IMAGE_CodingTGA
	ImagePortFormatPPM        ImagePortFormat = C.OMX_IMAGE_CodingPPM
)

func (f ImagePortFormat) String() string {
	switch f {
	case ImagePortFormatUnused:
		return "ImagePortFormatUnused"
	case ImagePortFormatAutodetect:
		return "ImagePortFormatAutodetect"
	case ImagePortFormatJPEG:
		return "ImagePortFormatJPEG"
	case ImagePortFormatJPEG2K:
		return "ImagePortFormatJPEG2K"
	case ImagePortFormatEXIF:
		return "ImagePortFormatEXIF"
	case ImagePortFormatTIFF:
		return "ImagePortFormatTIFF"
	case ImagePortFormatGIF:
		return "ImagePortFormatGIF"
	case ImagePortFormatPNG:
		return "ImagePortFormatPNG"
	case ImagePortFormatLZW:
		return "ImagePortFormatLZW"
	case ImagePortFormatBMP:
		return "ImagePortFormatBMP"
	case ImagePortFormatTGA:
		return "ImagePortFormatTGA"
	case ImagePortFormatPPM:
		return "ImagePortFormatPPM"
	}
	return fmt.Sprintf("UNKNOWN[%x]", int(f))
}

type ColorFormat C.OMX_COLOR_FORMATTYPE

const (
	ColorFormatUnused                 ColorFormat = C.OMX_COLOR_FormatUnused
	ColorFormatMonochrome             ColorFormat = C.OMX_COLOR_FormatMonochrome
	ColorFormat8bitRGB332             ColorFormat = C.OMX_COLOR_Format8bitRGB332
	ColorFormat12bitRGB444            ColorFormat = C.OMX_COLOR_Format12bitRGB444
	ColorFormat16bitARGB4444          ColorFormat = C.OMX_COLOR_Format16bitARGB4444
	ColorFormat16bitARGB1555          ColorFormat = C.OMX_COLOR_Format16bitARGB1555
	ColorFormat16bitRGB565            ColorFormat = C.OMX_COLOR_Format16bitRGB565
	ColorFormat16bitBGR565            ColorFormat = C.OMX_COLOR_Format16bitBGR565
	ColorFormat18bitRGB666            ColorFormat = C.OMX_COLOR_Format18bitRGB666
	ColorFormat18bitARGB1665          ColorFormat = C.OMX_COLOR_Format18bitARGB1665
	ColorFormat19bitARGB1666          ColorFormat = C.OMX_COLOR_Format19bitARGB1666
	ColorFormat24bitRGB888            ColorFormat = C.OMX_COLOR_Format24bitRGB888
	ColorFormat24bitBGR888            ColorFormat = C.OMX_COLOR_Format24bitBGR888
	ColorFormat24bitARGB1887          ColorFormat = C.OMX_COLOR_Format24bitARGB1887
	ColorFormat25bitARGB1888          ColorFormat = C.OMX_COLOR_Format25bitARGB1888
	ColorFormat32bitBGRA8888          ColorFormat = C.OMX_COLOR_Format32bitBGRA8888
	ColorFormat32bitARGB8888          ColorFormat = C.OMX_COLOR_Format32bitARGB8888
	ColorFormatYUV411Planar           ColorFormat = C.OMX_COLOR_FormatYUV411Planar
	ColorFormatYUV411PackedPlanar     ColorFormat = C.OMX_COLOR_FormatYUV411PackedPlanar
	ColorFormatYUV420Planar           ColorFormat = C.OMX_COLOR_FormatYUV420Planar
	ColorFormatYUV420PackedPlanar     ColorFormat = C.OMX_COLOR_FormatYUV420PackedPlanar
	ColorFormatYUV420SemiPlanar       ColorFormat = C.OMX_COLOR_FormatYUV420SemiPlanar
	ColorFormatYUV422Planar           ColorFormat = C.OMX_COLOR_FormatYUV422Planar
	ColorFormatYUV422PackedPlanar     ColorFormat = C.OMX_COLOR_FormatYUV422PackedPlanar
	ColorFormatYUV422SemiPlanar       ColorFormat = C.OMX_COLOR_FormatYUV422SemiPlanar
	ColorFormatYCbYCr                 ColorFormat = C.OMX_COLOR_FormatYCbYCr
	ColorFormatYCrYCb                 ColorFormat = C.OMX_COLOR_FormatYCrYCb
	ColorFormatCbYCrY                 ColorFormat = C.OMX_COLOR_FormatCbYCrY
	ColorFormatCrYCbY                 ColorFormat = C.OMX_COLOR_FormatCrYCbY
	ColorFormatYUV444Interleaved      ColorFormat = C.OMX_COLOR_FormatYUV444Interleaved
	ColorFormatRawBayer8bit           ColorFormat = C.OMX_COLOR_FormatRawBayer8bit
	ColorFormatRawBayer10bit          ColorFormat = C.OMX_COLOR_FormatRawBayer10bit
	ColorFormatRawBayer8bitcompressed ColorFormat = C.OMX_COLOR_FormatRawBayer8bitcompressed
	ColorFormatL2                     ColorFormat = C.OMX_COLOR_FormatL2
	ColorFormatL4                     ColorFormat = C.OMX_COLOR_FormatL4
	ColorFormatL8                     ColorFormat = C.OMX_COLOR_FormatL8
	ColorFormatL16                    ColorFormat = C.OMX_COLOR_FormatL16
	ColorFormatL24                    ColorFormat = C.OMX_COLOR_FormatL24
	ColorFormatL32                    ColorFormat = C.OMX_COLOR_FormatL32
	ColorFormatYUV420PackedSemiPlanar ColorFormat = C.OMX_COLOR_FormatYUV420PackedSemiPlanar
	ColorFormatYUV422PackedSemiPlanar ColorFormat = C.OMX_COLOR_FormatYUV422PackedSemiPlanar
	ColorFormat18BitBGR666            ColorFormat = C.OMX_COLOR_Format18BitBGR666
	ColorFormat24BitARGB6666          ColorFormat = C.OMX_COLOR_Format24BitARGB6666
	ColorFormat24BitABGR6666          ColorFormat = C.OMX_COLOR_Format24BitABGR6666
	ColorFormatKhronosExtensions      ColorFormat = C.OMX_COLOR_FormatKhronosExtensions
	ColorFormatVendorStartUnused      ColorFormat = C.OMX_COLOR_FormatVendorStartUnused
	ColorFormat32bitABGR8888          ColorFormat = C.OMX_COLOR_Format32bitABGR8888
	ColorFormat8bitPalette            ColorFormat = C.OMX_COLOR_Format8bitPalette
	ColorFormatYUVUV128               ColorFormat = C.OMX_COLOR_FormatYUVUV128
	ColorFormatRawBayer12bit          ColorFormat = C.OMX_COLOR_FormatRawBayer12bit
	ColorFormatBRCMEGL                ColorFormat = C.OMX_COLOR_FormatBRCMEGL
	ColorFormatBRCMOpaque             ColorFormat = C.OMX_COLOR_FormatBRCMOpaque
	ColorFormatYVU420PackedPlanar     ColorFormat = C.OMX_COLOR_FormatYVU420PackedPlanar
	ColorFormatYVU420PackedSemiPlanar ColorFormat = C.OMX_COLOR_FormatYVU420PackedSemiPlanar
	ColorFormatRawBayer16bit          ColorFormat = C.OMX_COLOR_FormatRawBayer16bit
	// ColorFormatYUV420_16PackedPlanar  ColorFormat = C.OMX_COLOR_FormatYUV420_16PackedPlanar
	// ColorFormatYUVUV64_16             ColorFormat = C.OMX_COLOR_FormatYUVUV64_16
	// ColorFormatYUV420_10PackedPlanar ColorFormat = C.OMX_COLOR_FormatYUV420_10PackedPlanar
	// ColorFormatYUVUV64_10             ColorFormat = C.OMX_COLOR_FormatYUVUV64_10
	// ColorFormatYUV420_UVSideBySide    ColorFormat = C.OMX_COLOR_FormatYUV420_UVSideBySide
	// ColorFormat32bitXRGB8888          ColorFormat = C.OMX_COLOR_Format32bitXRGB8888
	// ColorFormat32bitXBGR8888          ColorFormat = C.OMX_COLOR_Format32bitXBGR8888
	// ColorFormatYUV10bitColumn         ColorFormat = C.OMX_COLOR_FormatYUV10bitColumn
	ColorFormatMax ColorFormat = C.OMX_COLOR_FormatMax
)

func (f ColorFormat) String() string {
	switch f {
	case ColorFormatUnused:
		return "ColorFormatUnused"
	case ColorFormatMonochrome:
		return "ColorFormatMonochrome"
	case ColorFormat8bitRGB332:
		return "ColorFormat8bitRGB332"
	case ColorFormat12bitRGB444:
		return "ColorFormat12bitRGB444"
	case ColorFormat16bitARGB4444:
		return "ColorFormat16bitARGB4444"
	case ColorFormat16bitARGB1555:
		return "ColorFormat16bitARGB1555"
	case ColorFormat16bitRGB565:
		return "ColorFormat16bitRGB565"
	case ColorFormat16bitBGR565:
		return "ColorFormat16bitBGR565"
	case ColorFormat18bitRGB666:
		return "ColorFormat18bitRGB666"
	case ColorFormat18bitARGB1665:
		return "ColorFormat18bitARGB1665"
	case ColorFormat19bitARGB1666:
		return "ColorFormat19bitARGB1666"
	case ColorFormat24bitRGB888:
		return "ColorFormat24bitRGB888"
	case ColorFormat24bitBGR888:
		return "ColorFormat24bitBGR888"
	case ColorFormat24bitARGB1887:
		return "ColorFormat24bitARGB1887"
	case ColorFormat25bitARGB1888:
		return "ColorFormat25bitARGB1888"
	case ColorFormat32bitBGRA8888:
		return "ColorFormat32bitBGRA8888"
	case ColorFormat32bitARGB8888:
		return "ColorFormat32bitARGB8888"
	case ColorFormatYUV411Planar:
		return "ColorFormatYUV411Planar"
	case ColorFormatYUV411PackedPlanar:
		return "ColorFormatYUV411PackedPlanar"
	case ColorFormatYUV420Planar:
		return "ColorFormatYUV420Planar"
	case ColorFormatYUV420PackedPlanar:
		return "ColorFormatYUV420PackedPlanar"
	case ColorFormatYUV420SemiPlanar:
		return "ColorFormatYUV420SemiPlanar"
	case ColorFormatYUV422Planar:
		return "ColorFormatYUV422Planar"
	case ColorFormatYUV422PackedPlanar:
		return "ColorFormatYUV422PackedPlanar"
	case ColorFormatYUV422SemiPlanar:
		return "ColorFormatYUV422SemiPlanar"
	case ColorFormatYCbYCr:
		return "ColorFormatYCbYCr"
	case ColorFormatYCrYCb:
		return "ColorFormatYCrYCb"
	case ColorFormatCbYCrY:
		return "ColorFormatCbYCrY"
	case ColorFormatCrYCbY:
		return "ColorFormatCrYCbY"
	case ColorFormatYUV444Interleaved:
		return "ColorFormatYUV444Interleaved"
	case ColorFormatRawBayer8bit:
		return "ColorFormatRawBayer8bit"
	case ColorFormatRawBayer10bit:
		return "ColorFormatRawBayer10bit"
	case ColorFormatRawBayer8bitcompressed:
		return "ColorFormatRawBayer8bitcompressed"
	case ColorFormatL2:
		return "ColorFormatL2"
	case ColorFormatL4:
		return "ColorFormatL4"
	case ColorFormatL8:
		return "ColorFormatL8"
	case ColorFormatL16:
		return "ColorFormatL16"
	case ColorFormatL24:
		return "ColorFormatL24"
	case ColorFormatL32:
		return "ColorFormatL32"
	case ColorFormatYUV420PackedSemiPlanar:
		return "ColorFormatYUV420PackedSemiPlanar"
	case ColorFormatYUV422PackedSemiPlanar:
		return "ColorFormatYUV422PackedSemiPlanar"
	case ColorFormat18BitBGR666:
		return "ColorFormat18BitBGR666"
	case ColorFormat24BitARGB6666:
		return "ColorFormat24BitARGB6666"
	case ColorFormat24BitABGR6666:
		return "ColorFormat24BitABGR6666"
	// case ColorFormatKhronosExtensions: return "ColorFormatKhronosExtensions"
	// case ColorFormatVendorStartUnused: return "ColorFormatVendorStartUnused"
	case ColorFormat32bitABGR8888:
		return "ColorFormat32bitABGR8888"
	case ColorFormat8bitPalette:
		return "ColorFormat8bitPalette"
	case ColorFormatYUVUV128:
		return "ColorFormatYUVUV128"
	case ColorFormatRawBayer12bit:
		return "ColorFormatRawBayer12bit"
	case ColorFormatBRCMEGL:
		return "ColorFormatBRCMEGL"
	case ColorFormatBRCMOpaque:
		return "ColorFormatBRCMOpaque"
	case ColorFormatYVU420PackedPlanar:
		return "ColorFormatYVU420PackedPlanar"
	case ColorFormatYVU420PackedSemiPlanar:
		return "ColorFormatYVU420PackedSemiPlanar"
	case ColorFormatRawBayer16bit:
		return "ColorFormatRawBayer16bit"
		// case ColorFormatYUV420_16PackedPlanar:
		// 	return "ColorFormatYUV420_16PackedPlanar"
		// case ColorFormatYUVUV64_16:
		// 	return "ColorFormatYUVUV64_16"
		// case ColorFormatYUV420_10PackedPlanar:
		// 	return "ColorFormatYUV420_10PackedPlanar"
		// case ColorFormatYUVUV64_10:
		// 	return "ColorFormatYUVUV64_10"
		// case ColorFormatYUV420_UVSideBySide:
		// 	return "ColorFormatYUV420_UVSideBySide"
		// case ColorFormat32bitXRGB8888:
		// 	return "ColorFormat32bitXRGB8888"
		// case ColorFormat32bitXBGR8888:
		// 	return "ColorFormat32bitXBGR8888"
		// case ColorFormatYUV10bitColumn:
		// 	return "ColorFormatYUV10bitColumn"
		// case ColorFormatMax: return "ColorFormatMax"
	}
	return fmt.Sprintf("UNKNOWN[%x]", int(f))
}

type VideoControlRate C.OMX_VIDEO_CONTROLRATETYPE
const (
	ControlRateDisable VideoControlRate = C.OMX_Video_ControlRateDisable
    ControlRateVariable VideoControlRate = C.OMX_Video_ControlRateVariable
    ControlRateConstant VideoControlRate = C.OMX_Video_ControlRateConstant
    ControlRateVariableSkipFrames VideoControlRate = C.OMX_Video_ControlRateVariableSkipFrames
    ControlRateConstantSkipFrames VideoControlRate = C.OMX_Video_ControlRateConstantSkipFrames
    ControlRateKhronosExtensions VideoControlRate = C.OMX_Video_ControlRateKhronosExtensions
    ControlRateVendorStartUnused VideoControlRate = C.OMX_Video_ControlRateVendorStartUnused
    ControlRateMax VideoControlRate = C.OMX_Video_ControlRateMax
)
func (c VideoControlRate) String() string {
	switch c {
	case ControlRateDisable: return "ControlRateDisable"
	case ControlRateVariable: return "ControlRateVariable"
	case ControlRateConstant: return "ControlRateConstant"
	case ControlRateVariableSkipFrames: return "ControlRateVariableSkipFrames"
	case ControlRateConstantSkipFrames: return "ControlRateConstantSkipFrames"
	case ControlRateKhronosExtensions: return "ControlRateKhronosExtensions"
	case ControlRateVendorStartUnused: return "ControlRateVendorStartUnused"
	case ControlRateMax: return "ControlRateMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(c))
}



type Version struct {
	Major    uint8
	Minor    uint8
	Revision uint8
	Step     uint8
}

func (v Version) toBytes() (ret [4]byte) {
	ret[0] = byte(v.Major)
	ret[1] = byte(v.Minor)
	ret[2] = byte(v.Revision)
	ret[3] = byte(v.Step)
	return
}

func (v *Version) fromBytes(b [4]byte) {
	v.Major = uint8(b[0])
	v.Minor = uint8(b[1])
	v.Revision = uint8(b[2])
	v.Step = uint8(b[3])
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Revision, v.Step)
}
