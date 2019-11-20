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
	VideoCodingUnused     VideoCoding = C.OMX_VIDEO_CodingUnused
	VideoCodingAutoDetect VideoCoding = C.OMX_VIDEO_CodingAutoDetect
	VideoCodingMPEG2      VideoCoding = C.OMX_VIDEO_CodingMPEG2
	VideoCodingH263       VideoCoding = C.OMX_VIDEO_CodingH263
	VideoCodingMPEG4      VideoCoding = C.OMX_VIDEO_CodingMPEG4
	VideoCodingWMV        VideoCoding = C.OMX_VIDEO_CodingWMV
	VideoCodingRV         VideoCoding = C.OMX_VIDEO_CodingRV
	VideoCodingAVC        VideoCoding = C.OMX_VIDEO_CodingAVC
	VideoCodingMJPEG      VideoCoding = C.OMX_VIDEO_CodingMJPEG
)

func (v VideoCoding) String() string {
	switch v {
	case VideoCodingUnused:
		return "VideoCodingUnused"
	case VideoCodingAutoDetect:
		return "VideoCodingAutoDetect"
	case VideoCodingMPEG2:
		return "VideoCodingMPEG2"
	case VideoCodingH263:
		return "VideoCodingH263"
	case VideoCodingMPEG4:
		return "VideoCodingMPEG4"
	case VideoCodingWMV:
		return "VideoCodingWMV"
	case VideoCodingRV:
		return "VideoCodingRV"
	case VideoCodingAVC:
		return "VideoCodingAVC"
	case VideoCodingMJPEG:
		return "VideoCodingMJPEG"
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
	ControlRateDisable            VideoControlRate = C.OMX_Video_ControlRateDisable
	ControlRateVariable           VideoControlRate = C.OMX_Video_ControlRateVariable
	ControlRateConstant           VideoControlRate = C.OMX_Video_ControlRateConstant
	ControlRateVariableSkipFrames VideoControlRate = C.OMX_Video_ControlRateVariableSkipFrames
	ControlRateConstantSkipFrames VideoControlRate = C.OMX_Video_ControlRateConstantSkipFrames
	ControlRateKhronosExtensions  VideoControlRate = C.OMX_Video_ControlRateKhronosExtensions
	ControlRateVendorStartUnused  VideoControlRate = C.OMX_Video_ControlRateVendorStartUnused
	ControlRateMax                VideoControlRate = C.OMX_Video_ControlRateMax
)

func (c VideoControlRate) String() string {
	switch c {
	case ControlRateDisable:
		return "ControlRateDisable"
	case ControlRateVariable:
		return "ControlRateVariable"
	case ControlRateConstant:
		return "ControlRateConstant"
	case ControlRateVariableSkipFrames:
		return "ControlRateVariableSkipFrames"
	case ControlRateConstantSkipFrames:
		return "ControlRateConstantSkipFrames"
	case ControlRateKhronosExtensions:
		return "ControlRateKhronosExtensions"
	case ControlRateVendorStartUnused:
		return "ControlRateVendorStartUnused"
	case ControlRateMax:
		return "ControlRateMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(c))
}

type VideoMotionVectorAccuracy C.OMX_VIDEO_MOTIONVECTORTYPE

const (
	MotionVectorPixel             VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorPixel
	MotionVectorHalfPel           VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorHalfPel
	MotionVectorQuarterPel        VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorQuarterPel
	MotionVectorEighthPel         VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorEighthPel
	MotionVectorKhronosExtensions VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorKhronosExtensions
	MotionVectorVendorStartUnused VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorVendorStartUnused
	MotionVectorMax               VideoMotionVectorAccuracy = C.OMX_Video_MotionVectorMax
)

func (c VideoMotionVectorAccuracy) String() string {
	switch c {
	case MotionVectorPixel:
		return "MotionVectorPixel"
	case MotionVectorHalfPel:
		return "MotionVectorHalfPel"
	case MotionVectorQuarterPel:
		return "MotionVectorQuarterPel"
	case MotionVectorEighthPel:
		return "MotionVectorEighthPel"
	case MotionVectorKhronosExtensions:
		return "MotionVectorKhronosExtensions"
	case MotionVectorVendorStartUnused:
		return "MotionVectorVendorStartUnused"
	case MotionVectorMax:
		return "MotionVectorMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(c))
}

type VideoIntraRefreshMode C.OMX_VIDEO_INTRAREFRESHTYPE

const (
	VideoIntraRefreshCyclic            VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshCyclic
	VideoIntraRefreshAdaptive          VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshAdaptive
	VideoIntraRefreshBoth              VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshBoth
	VideoIntraRefreshKhronosExtensions VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshKhronosExtensions
	VideoIntraRefreshVendorStartUnused VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshVendorStartUnused
	VideoIntraRefreshCyclicMrows       VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshCyclicMrows
	VideoIntraRefreshPseudoRand        VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshPseudoRand
	VideoIntraRefreshMax               VideoIntraRefreshMode = C.OMX_VIDEO_IntraRefreshMax
)

func (v VideoIntraRefreshMode) String() string {
	switch v {
	case VideoIntraRefreshCyclic:
		return "VideoIntraRefreshCyclic"
	case VideoIntraRefreshAdaptive:
		return "VideoIntraRefreshAdaptive"
	case VideoIntraRefreshBoth:
		return "VideoIntraRefreshBoth"
	case VideoIntraRefreshKhronosExtensions:
		return "VideoIntraRefreshKhronosExtensions"
	case VideoIntraRefreshVendorStartUnused:
		return "VideoIntraRefreshVendorStartUnused"
	case VideoIntraRefreshCyclicMrows:
		return "VideoIntraRefreshCyclicMrows"
	case VideoIntraRefreshPseudoRand:
		return "VideoIntraRefreshPseudoRand"
	case VideoIntraRefreshMax:
		return "VideoIntraRefreshMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoH263Profile C.OMX_VIDEO_H263PROFILETYPE

const (
	VideoH263ProfileBaseline           VideoH263Profile = C.OMX_VIDEO_H263ProfileBaseline
	VideoH263ProfileH320Coding         VideoH263Profile = C.OMX_VIDEO_H263ProfileH320Coding
	VideoH263ProfileBackwardCompatible VideoH263Profile = C.OMX_VIDEO_H263ProfileBackwardCompatible
	VideoH263ProfileISWV2              VideoH263Profile = C.OMX_VIDEO_H263ProfileISWV2
	VideoH263ProfileISWV3              VideoH263Profile = C.OMX_VIDEO_H263ProfileISWV3
	VideoH263ProfileHighCompression    VideoH263Profile = C.OMX_VIDEO_H263ProfileHighCompression
	VideoH263ProfileInternet           VideoH263Profile = C.OMX_VIDEO_H263ProfileInternet
	VideoH263ProfileInterlace          VideoH263Profile = C.OMX_VIDEO_H263ProfileInterlace
	VideoH263ProfileHighLatency        VideoH263Profile = C.OMX_VIDEO_H263ProfileHighLatency
	VideoH263ProfileKhronosExtensions  VideoH263Profile = C.OMX_VIDEO_H263ProfileKhronosExtensions
	VideoH263ProfileVendorStartUnused  VideoH263Profile = C.OMX_VIDEO_H263ProfileVendorStartUnused
	VideoH263ProfileMax                VideoH263Profile = C.OMX_VIDEO_H263ProfileMax
)

func (v VideoH263Profile) String() string {
	switch v {
	case VideoH263ProfileBaseline:
		return "VideoH263ProfileBaseline"
	case VideoH263ProfileH320Coding:
		return "VideoH263ProfileH320Coding"
	case VideoH263ProfileBackwardCompatible:
		return "VideoH263ProfileBackwardCompatible"
	case VideoH263ProfileISWV2:
		return "VideoH263ProfileISWV2"
	case VideoH263ProfileISWV3:
		return "VideoH263ProfileISWV3"
	case VideoH263ProfileHighCompression:
		return "VideoH263ProfileHighCompression"
	case VideoH263ProfileInternet:
		return "VideoH263ProfileInternet"
	case VideoH263ProfileInterlace:
		return "VideoH263ProfileInterlace"
	case VideoH263ProfileHighLatency:
		return "VideoH263ProfileHighLatency"
	case VideoH263ProfileKhronosExtensions:
		return "VideoH263ProfileKhronosExtensions"
	case VideoH263ProfileVendorStartUnused:
		return "VideoH263ProfileVendorStartUnused"
	case VideoH263ProfileMax:
		return "VideoH263ProfileMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoH263Level C.OMX_VIDEO_H263LEVELTYPE

const (
	VideoH263Level10                VideoH263Level = C.OMX_VIDEO_H263Level10
	VideoH263Level20                VideoH263Level = C.OMX_VIDEO_H263Level20
	VideoH263Level30                VideoH263Level = C.OMX_VIDEO_H263Level30
	VideoH263Level40                VideoH263Level = C.OMX_VIDEO_H263Level40
	VideoH263Level45                VideoH263Level = C.OMX_VIDEO_H263Level45
	VideoH263Level50                VideoH263Level = C.OMX_VIDEO_H263Level50
	VideoH263Level60                VideoH263Level = C.OMX_VIDEO_H263Level60
	VideoH263Level70                VideoH263Level = C.OMX_VIDEO_H263Level70
	VideoH263LevelKhronosExtensions VideoH263Level = C.OMX_VIDEO_H263LevelKhronosExtensions
	VideoH263LevelVendorStartUnused VideoH263Level = C.OMX_VIDEO_H263LevelVendorStartUnused
	VideoH263LevelMax               VideoH263Level = C.OMX_VIDEO_H263LevelMax
)

func (v VideoH263Level) String() string {
	switch v {
	case VideoH263Level10:
		return "VideoH263Level10"
	case VideoH263Level20:
		return "VideoH263Level20"
	case VideoH263Level30:
		return "VideoH263Level30"
	case VideoH263Level40:
		return "VideoH263Level40"
	case VideoH263Level45:
		return "VideoH263Level45"
	case VideoH263Level50:
		return "VideoH263Level50"
	case VideoH263Level60:
		return "VideoH263Level60"
	case VideoH263Level70:
		return "VideoH263Level70"
	case VideoH263LevelKhronosExtensions:
		return "VideoH263LevelKhronosExtensions"
	case VideoH263LevelVendorStartUnused:
		return "VideoH263LevelVendorStartUnused"
	case VideoH263LevelMax:
		return "VideoH263LevelMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoPicture C.OMX_VIDEO_PICTURETYPE

const (
	VideoPictureI                 VideoPicture = C.OMX_VIDEO_PictureTypeI
	VideoPictureP                 VideoPicture = C.OMX_VIDEO_PictureTypeP
	VideoPictureB                 VideoPicture = C.OMX_VIDEO_PictureTypeB
	VideoPictureSI                VideoPicture = C.OMX_VIDEO_PictureTypeSI
	VideoPictureSP                VideoPicture = C.OMX_VIDEO_PictureTypeSP
	VideoPictureEI                VideoPicture = C.OMX_VIDEO_PictureTypeEI
	VideoPictureEP                VideoPicture = C.OMX_VIDEO_PictureTypeEP
	VideoPictureS                 VideoPicture = C.OMX_VIDEO_PictureTypeS
	VideoPictureKhronosExtensions VideoPicture = C.OMX_VIDEO_PictureTypeKhronosExtensions
	VideoPictureVendorStartUnused VideoPicture = C.OMX_VIDEO_PictureTypeVendorStartUnused
	VideoPictureMax               VideoPicture = C.OMX_VIDEO_PictureTypeMax
)

func (v VideoPicture) String() string {
	switch v {
	case VideoPictureI:
		return "VideoPictureI"
	case VideoPictureP:
		return "VideoPictureP"
	case VideoPictureB:
		return "VideoPictureB"
	case VideoPictureSI:
		return "VideoPictureSI"
	case VideoPictureSP:
		return "VideoPictureSP"
	case VideoPictureEI:
		return "VideoPictureEI"
	case VideoPictureEP:
		return "VideoPictureEP"
	case VideoPictureS:
		return "VideoPictureS"
	case VideoPictureKhronosExtensions:
		return "VideoPictureKhronosExtensions"
	case VideoPictureVendorStartUnused:
		return "VideoPictureVendorStartUnused"
	case VideoPictureMax:
		return "VideoPictureMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoMPEG2Profile C.OMX_VIDEO_MPEG2PROFILETYPE

const (
	VideoMPEG2ProfileSimple            VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileSimple
	VideoMPEG2ProfileMain              VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileMain
	VideoMPEG2Profile422               VideoMPEG2Profile = C.OMX_VIDEO_MPEG2Profile422
	VideoMPEG2ProfileSNR               VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileSNR
	VideoMPEG2ProfileSpatial           VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileSpatial
	VideoMPEG2ProfileHigh              VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileHigh
	VideoMPEG2ProfileKhronosExtensions VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileKhronosExtensions
	VideoMPEG2ProfileVendorStartUnused VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileVendorStartUnused
	VideoMPEG2ProfileMax               VideoMPEG2Profile = C.OMX_VIDEO_MPEG2ProfileMax
)

func (v VideoMPEG2Profile) String() string {
	switch v {
	case VideoMPEG2ProfileSimple:
		return "VideoMPEG2ProfileSimple"
	case VideoMPEG2ProfileMain:
		return "VideoMPEG2ProfileMain"
	case VideoMPEG2Profile422:
		return "VideoMPEG2Profile422"
	case VideoMPEG2ProfileSNR:
		return "VideoMPEG2ProfileSNR"
	case VideoMPEG2ProfileSpatial:
		return "VideoMPEG2ProfileSpatial"
	case VideoMPEG2ProfileHigh:
		return "VideoMPEG2ProfileHigh"
	case VideoMPEG2ProfileKhronosExtensions:
		return "VideoMPEG2ProfileKhronosExtensions"
	case VideoMPEG2ProfileVendorStartUnused:
		return "VideoMPEG2ProfileVendorStartUnused"
	case VideoMPEG2ProfileMax:
		return "VideoMPEG2ProfileMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoMPEG2Level C.OMX_VIDEO_MPEG2LEVELTYPE

const (
	VideoMPEG2LevelLL                VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelLL
	VideoMPEG2LevelML                VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelML
	VideoMPEG2LevelH14               VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelH14
	VideoMPEG2LevelHL                VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelHL
	VideoMPEG2LevelKhronosExtensions VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelKhronosExtensions
	VideoMPEG2LevelVendorStartUnused VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelVendorStartUnused
	VideoMPEG2LevelMax               VideoMPEG2Level = C.OMX_VIDEO_MPEG2LevelMax
)

func (v VideoMPEG2Level) String() string {
	switch v {
	case VideoMPEG2LevelLL:
		return "VideoMPEG2LevelLL"
	case VideoMPEG2LevelML:
		return "VideoMPEG2LevelML"
	case VideoMPEG2LevelH14:
		return "VideoMPEG2LevelH14"
	case VideoMPEG2LevelHL:
		return "VideoMPEG2LevelHL"
	case VideoMPEG2LevelKhronosExtensions:
		return "VideoMPEG2LevelKhronosExtensions"
	case VideoMPEG2LevelVendorStartUnused:
		return "VideoMPEG2LevelVendorStartUnused"
	case VideoMPEG2LevelMax:
		return "VideoMPEG2LevelMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
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

type VideoMPEG4Profile C.OMX_VIDEO_MPEG4PROFILETYPE

const (
	VideoMPEG4ProfileSimple            VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileSimple
	VideoMPEG4ProfileSimpleScalable    VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileSimpleScalable
	VideoMPEG4ProfileCore              VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileCore
	VideoMPEG4ProfileMain              VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileMain
	VideoMPEG4ProfileNbit              VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileNbit
	VideoMPEG4ProfileScalableTexture   VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileScalableTexture
	VideoMPEG4ProfileSimpleFace        VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileSimpleFace
	VideoMPEG4ProfileSimpleFBA         VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileSimpleFBA
	VideoMPEG4ProfileBasicAnimated     VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileBasicAnimated
	VideoMPEG4ProfileHybrid            VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileHybrid
	VideoMPEG4ProfileAdvancedRealTime  VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileAdvancedRealTime
	VideoMPEG4ProfileCoreScalable      VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileCoreScalable
	VideoMPEG4ProfileAdvancedCoding    VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileAdvancedCoding
	VideoMPEG4ProfileAdvancedCore      VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileAdvancedCore
	VideoMPEG4ProfileAdvancedScalable  VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileAdvancedScalable
	VideoMPEG4ProfileAdvancedSimple    VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileAdvancedSimple
	VideoMPEG4ProfileKhronosExtensions VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileKhronosExtensions
	VideoMPEG4ProfileVendorStartUnused VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileVendorStartUnused
	VideoMPEG4ProfileMax               VideoMPEG4Profile = C.OMX_VIDEO_MPEG4ProfileMax
)

func (v VideoMPEG4Profile) String() string {
	switch v {
	case VideoMPEG4ProfileSimple:
		return "VideoMPEG4ProfileSimple"
	case VideoMPEG4ProfileSimpleScalable:
		return "VideoMPEG4ProfileSimpleScalable"
	case VideoMPEG4ProfileCore:
		return "VideoMPEG4ProfileCore"
	case VideoMPEG4ProfileMain:
		return "VideoMPEG4ProfileMain"
	case VideoMPEG4ProfileNbit:
		return "VideoMPEG4ProfileNbit"
	case VideoMPEG4ProfileScalableTexture:
		return "VideoMPEG4ProfileScalableTexture"
	case VideoMPEG4ProfileSimpleFace:
		return "VideoMPEG4ProfileSimpleFace"
	case VideoMPEG4ProfileSimpleFBA:
		return "VideoMPEG4ProfileSimpleFBA"
	case VideoMPEG4ProfileBasicAnimated:
		return "VideoMPEG4ProfileBasicAnimated"
	case VideoMPEG4ProfileHybrid:
		return "VideoMPEG4ProfileHybrid"
	case VideoMPEG4ProfileAdvancedRealTime:
		return "VideoMPEG4ProfileAdvancedRealTime"
	case VideoMPEG4ProfileCoreScalable:
		return "VideoMPEG4ProfileCoreScalable"
	case VideoMPEG4ProfileAdvancedCoding:
		return "VideoMPEG4ProfileAdvancedCoding"
	case VideoMPEG4ProfileAdvancedCore:
		return "VideoMPEG4ProfileAdvancedCore"
	case VideoMPEG4ProfileAdvancedScalable:
		return "VideoMPEG4ProfileAdvancedScalable"
	case VideoMPEG4ProfileAdvancedSimple:
		return "VideoMPEG4ProfileAdvancedSimple"
	case VideoMPEG4ProfileKhronosExtensions:
		return "VideoMPEG4ProfileKhronosExtensions"
	case VideoMPEG4ProfileVendorStartUnused:
		return "VideoMPEG4ProfileVendorStartUnused"
	case VideoMPEG4ProfileMax:
		return "VideoMPEG4ProfileMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoMPEG4Level C.OMX_VIDEO_MPEG4LEVELTYPE

const (
	VideoMPEG4Level0                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level0
	VideoMPEG4Level0b                VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level0b
	VideoMPEG4Level1                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level1
	VideoMPEG4Level2                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level2
	VideoMPEG4Level3                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level3
	VideoMPEG4Level4                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level4
	VideoMPEG4Level4a                VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level4a
	VideoMPEG4Level5                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level5
	VideoMPEG4Level6                 VideoMPEG4Level = C.OMX_VIDEO_MPEG4Level6
	VideoMPEG4LevelKhronosExtensions VideoMPEG4Level = C.OMX_VIDEO_MPEG4LevelKhronosExtensions
	VideoMPEG4LevelVendorStartUnused VideoMPEG4Level = C.OMX_VIDEO_MPEG4LevelVendorStartUnused
	VideoMPEG4LevelMax               VideoMPEG4Level = C.OMX_VIDEO_MPEG4LevelMax
)

func (v VideoMPEG4Level) String() string {
	switch v {
	case VideoMPEG4Level0:
		return "VideoMPEG4Level0"
	case VideoMPEG4Level0b:
		return "VideoMPEG4Level0b"
	case VideoMPEG4Level1:
		return "VideoMPEG4Level1"
	case VideoMPEG4Level2:
		return "VideoMPEG4Level2"
	case VideoMPEG4Level3:
		return "VideoMPEG4Level3"
	case VideoMPEG4Level4:
		return "VideoMPEG4Level4"
	case VideoMPEG4Level4a:
		return "VideoMPEG4Level4a"
	case VideoMPEG4Level5:
		return "VideoMPEG4Level5"
	case VideoMPEG4Level6:
		return "VideoMPEG4Level6"
	case VideoMPEG4LevelKhronosExtensions:
		return "VideoMPEG4LevelKhronosExtensions"
	case VideoMPEG4LevelVendorStartUnused:
		return "VideoMPEG4LevelVendorStartUnused"
	case VideoMPEG4LevelMax:
		return "VideoMPEG4LevelMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoWMVFormat C.OMX_VIDEO_WMVFORMATTYPE

const (
	VideoWMVFormatUnused            VideoWMVFormat = C.OMX_VIDEO_WMVFormatUnused
	VideoWMVFormat7                 VideoWMVFormat = C.OMX_VIDEO_WMVFormat7
	VideoWMVFormat8                 VideoWMVFormat = C.OMX_VIDEO_WMVFormat8
	VideoWMVFormat9                 VideoWMVFormat = C.OMX_VIDEO_WMVFormat9
	VideoWMFFormatKhronosExtensions VideoWMVFormat = C.OMX_VIDEO_WMFFormatKhronosExtensions
	VideoWMFFormatVendorStartUnused VideoWMVFormat = C.OMX_VIDEO_WMFFormatVendorStartUnused
	VideoWMVFormatMax               VideoWMVFormat = C.OMX_VIDEO_WMVFormatMax
)

func (v VideoWMVFormat) String() string {
	switch v {
	case VideoWMVFormatUnused:
		return "VideoWMVFormatUnused"
	case VideoWMVFormat7:
		return "VideoWMVFormat7"
	case VideoWMVFormat8:
		return "VideoWMVFormat8"
	case VideoWMVFormat9:
		return "VideoWMVFormat9"
	case VideoWMFFormatKhronosExtensions:
		return "VideoWMFFormatKhronosExtensions"
	case VideoWMFFormatVendorStartUnused:
		return "VideoWMFFormatVendorStartUnused"
	case VideoWMVFormatMax:
		return "VideoWMVFormatMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoRVFormat C.OMX_VIDEO_RVFORMATTYPE

const (
	VideoRVFormatUnused            VideoRVFormat = C.OMX_VIDEO_RVFormatUnused
	VideoRVFormat8                 VideoRVFormat = C.OMX_VIDEO_RVFormat8
	VideoRVFormat9                 VideoRVFormat = C.OMX_VIDEO_RVFormat9
	VideoRVFormatG2                VideoRVFormat = C.OMX_VIDEO_RVFormatG2
	VideoRVFormatKhronosExtensions VideoRVFormat = C.OMX_VIDEO_RVFormatKhronosExtensions
	VideoRVFormatVendorStartUnused VideoRVFormat = C.OMX_VIDEO_RVFormatVendorStartUnused
	VideoRVFormatMax               VideoRVFormat = C.OMX_VIDEO_RVFormatMax
)

func (v VideoRVFormat) String() string {
	switch v {
	case VideoRVFormatUnused:
		return "VideoRVFormatUnused"
	case VideoRVFormat8:
		return "VideoRVFormat8"
	case VideoRVFormat9:
		return "VideoRVFormat9"
	case VideoRVFormatG2:
		return "VideoRVFormatG2"
	case VideoRVFormatKhronosExtensions:
		return "VideoRVFormatKhronosExtensions"
	case VideoRVFormatVendorStartUnused:
		return "VideoRVFormatVendorStartUnused"
	case VideoRVFormatMax:
		return "VideoRVFormatMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}
