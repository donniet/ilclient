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

type OtherFormat C.OMX_OTHER_FORMATTYPE

const (
	OtherFormatTime              OtherFormat = C.OMX_OTHER_FormatTime
	OtherFormatPower             OtherFormat = C.OMX_OTHER_FormatPower
	OtherFormatStats             OtherFormat = C.OMX_OTHER_FormatStats
	OtherFormatBinary            OtherFormat = C.OMX_OTHER_FormatBinary
	OtherFormatVendorReserved    OtherFormat = C.OMX_OTHER_FormatVendorReserved
	OtherFormatKhronosExtensions OtherFormat = C.OMX_OTHER_FormatKhronosExtensions
	OtherFormatVendorStartUnused OtherFormat = C.OMX_OTHER_FormatVendorStartUnused
	OtherFormatText              OtherFormat = C.OMX_OTHER_FormatText
	OtherFormatTextSKM2          OtherFormat = C.OMX_OTHER_FormatTextSKM2
	OtherFormatText3GP5          OtherFormat = C.OMX_OTHER_FormatText3GP5
	OtherFormatMax               OtherFormat = C.OMX_OTHER_FormatMax
)

func (f OtherFormat) String() string {
	switch f {
	case OtherFormatTime:
		return "OtherFormatTime"
	case OtherFormatPower:
		return "OtherFormatPower"
	case OtherFormatStats:
		return "OtherFormatStats"
	case OtherFormatBinary:
		return "OtherFormatBinary"
	case OtherFormatVendorReserved:
		return "OtherFormatVendorReserved"
	case OtherFormatKhronosExtensions:
		return "OtherFormatKhronosExtensions"
	case OtherFormatVendorStartUnused:
		return "OtherFormatVendorStartUnused"
	case OtherFormatText:
		return "OtherFormatText"
	case OtherFormatTextSKM2:
		return "OtherFormatTextSKM2"
	case OtherFormatText3GP5:
		return "OtherFormatText3GP5"
	case OtherFormatMax:
		return "OtherFormatMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(f))
}

type ImageCoding C.OMX_IMAGE_CODINGTYPE

const (
	ImageCodingUnused     ImageCoding = C.OMX_IMAGE_CodingUnused
	ImageCodingAutodetect ImageCoding = C.OMX_IMAGE_CodingAutoDetect
	ImageCodingJPEG       ImageCoding = C.OMX_IMAGE_CodingJPEG
	ImageCodingJPEG2K     ImageCoding = C.OMX_IMAGE_CodingJPEG2K
	ImageCodingEXIF       ImageCoding = C.OMX_IMAGE_CodingEXIF
	ImageCodingTIFF       ImageCoding = C.OMX_IMAGE_CodingTIFF
	ImageCodingGIF        ImageCoding = C.OMX_IMAGE_CodingGIF
	ImageCodingPNG        ImageCoding = C.OMX_IMAGE_CodingPNG
	ImageCodingLZW        ImageCoding = C.OMX_IMAGE_CodingLZW
	ImageCodingBMP        ImageCoding = C.OMX_IMAGE_CodingBMP
	ImageCodingTGA        ImageCoding = C.OMX_IMAGE_CodingTGA
	ImageCodingPPM        ImageCoding = C.OMX_IMAGE_CodingPPM
)

func (f ImageCoding) String() string {
	switch f {
	case ImageCodingUnused:
		return "ImageCodingUnused"
	case ImageCodingAutodetect:
		return "ImageCodingAutodetect"
	case ImageCodingJPEG:
		return "ImageCodingJPEG"
	case ImageCodingJPEG2K:
		return "ImageCodingJPEG2K"
	case ImageCodingEXIF:
		return "ImageCodingEXIF"
	case ImageCodingTIFF:
		return "ImageCodingTIFF"
	case ImageCodingGIF:
		return "ImageCodingGIF"
	case ImageCodingPNG:
		return "ImageCodingPNG"
	case ImageCodingLZW:
		return "ImageCodingLZW"
	case ImageCodingBMP:
		return "ImageCodingBMP"
	case ImageCodingTGA:
		return "ImageCodingTGA"
	case ImageCodingPPM:
		return "ImageCodingPPM"
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

type VideoAVCProfile C.OMX_VIDEO_AVCPROFILETYPE

const (
	VideoAVCProfileBaseline            VideoAVCProfile = C.OMX_VIDEO_AVCProfileBaseline
	VideoAVCProfileMain                VideoAVCProfile = C.OMX_VIDEO_AVCProfileMain
	VideoAVCProfileExtended            VideoAVCProfile = C.OMX_VIDEO_AVCProfileExtended
	VideoAVCProfileHigh                VideoAVCProfile = C.OMX_VIDEO_AVCProfileHigh
	VideoAVCProfileHigh10              VideoAVCProfile = C.OMX_VIDEO_AVCProfileHigh10
	VideoAVCProfileHigh422             VideoAVCProfile = C.OMX_VIDEO_AVCProfileHigh422
	VideoAVCProfileHigh444             VideoAVCProfile = C.OMX_VIDEO_AVCProfileHigh444
	VideoAVCProfileConstrainedBaseline VideoAVCProfile = C.OMX_VIDEO_AVCProfileConstrainedBaseline
	VideoAVCProfileKhronosExtensions   VideoAVCProfile = C.OMX_VIDEO_AVCProfileKhronosExtensions
	VideoAVCProfileVendorStartUnused   VideoAVCProfile = C.OMX_VIDEO_AVCProfileVendorStartUnused
	VideoAVCProfileMax                 VideoAVCProfile = C.OMX_VIDEO_AVCProfileMax
)

func (v VideoAVCProfile) String() string {
	switch v {
	case VideoAVCProfileBaseline:
		return "VideoAVCProfileBaseline"
	case VideoAVCProfileMain:
		return "VideoAVCProfileMain"
	case VideoAVCProfileExtended:
		return "VideoAVCProfileExtended"
	case VideoAVCProfileHigh:
		return "VideoAVCProfileHigh"
	case VideoAVCProfileHigh10:
		return "VideoAVCProfileHigh10"
	case VideoAVCProfileHigh422:
		return "VideoAVCProfileHigh422"
	case VideoAVCProfileHigh444:
		return "VideoAVCProfileHigh444"
	case VideoAVCProfileConstrainedBaseline:
		return "VideoAVCProfileConstrainedBaseline"
	case VideoAVCProfileKhronosExtensions:
		return "VideoAVCProfileKhronosExtensions"
	case VideoAVCProfileVendorStartUnused:
		return "VideoAVCProfileVendorStartUnused"
	case VideoAVCProfileMax:
		return "VideoAVCProfileMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoAVCLevel C.OMX_VIDEO_AVCLEVELTYPE

const (
	VideoAVCLevel1                 VideoAVCLevel = C.OMX_VIDEO_AVCLevel1
	VideoAVCLevel1b                VideoAVCLevel = C.OMX_VIDEO_AVCLevel1b
	VideoAVCLevel11                VideoAVCLevel = C.OMX_VIDEO_AVCLevel11
	VideoAVCLevel12                VideoAVCLevel = C.OMX_VIDEO_AVCLevel12
	VideoAVCLevel13                VideoAVCLevel = C.OMX_VIDEO_AVCLevel13
	VideoAVCLevel2                 VideoAVCLevel = C.OMX_VIDEO_AVCLevel2
	VideoAVCLevel21                VideoAVCLevel = C.OMX_VIDEO_AVCLevel21
	VideoAVCLevel22                VideoAVCLevel = C.OMX_VIDEO_AVCLevel22
	VideoAVCLevel3                 VideoAVCLevel = C.OMX_VIDEO_AVCLevel3
	VideoAVCLevel31                VideoAVCLevel = C.OMX_VIDEO_AVCLevel31
	VideoAVCLevel32                VideoAVCLevel = C.OMX_VIDEO_AVCLevel32
	VideoAVCLevel4                 VideoAVCLevel = C.OMX_VIDEO_AVCLevel4
	VideoAVCLevel41                VideoAVCLevel = C.OMX_VIDEO_AVCLevel41
	VideoAVCLevel42                VideoAVCLevel = C.OMX_VIDEO_AVCLevel42
	VideoAVCLevel5                 VideoAVCLevel = C.OMX_VIDEO_AVCLevel5
	VideoAVCLevel51                VideoAVCLevel = C.OMX_VIDEO_AVCLevel51
	VideoAVCLevelKhronosExtensions VideoAVCLevel = C.OMX_VIDEO_AVCLevelKhronosExtensions
	VideoAVCLevelVendorStartUnused VideoAVCLevel = C.OMX_VIDEO_AVCLevelVendorStartUnused
	VideoAVCLevelMax               VideoAVCLevel = C.OMX_VIDEO_AVCLevelMax
)

func (v VideoAVCLevel) String() string {
	switch v {
	case VideoAVCLevel1:
		return "VideoAVCLevel1"
	case VideoAVCLevel1b:
		return "VideoAVCLevel1b"
	case VideoAVCLevel11:
		return "VideoAVCLevel11"
	case VideoAVCLevel12:
		return "VideoAVCLevel12"
	case VideoAVCLevel13:
		return "VideoAVCLevel13"
	case VideoAVCLevel2:
		return "VideoAVCLevel2"
	case VideoAVCLevel21:
		return "VideoAVCLevel21"
	case VideoAVCLevel22:
		return "VideoAVCLevel22"
	case VideoAVCLevel3:
		return "VideoAVCLevel3"
	case VideoAVCLevel31:
		return "VideoAVCLevel31"
	case VideoAVCLevel32:
		return "VideoAVCLevel32"
	case VideoAVCLevel4:
		return "VideoAVCLevel4"
	case VideoAVCLevel41:
		return "VideoAVCLevel41"
	case VideoAVCLevel42:
		return "VideoAVCLevel42"
	case VideoAVCLevel5:
		return "VideoAVCLevel5"
	case VideoAVCLevel51:
		return "VideoAVCLevel51"
	case VideoAVCLevelKhronosExtensions:
		return "VideoAVCLevelKhronosExtensions"
	case VideoAVCLevelVendorStartUnused:
		return "VideoAVCLevelVendorStartUnused"
	case VideoAVCLevelMax:
		return "VideoAVCLevelMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type VideoAVCLoopFilter C.OMX_VIDEO_AVCLOOPFILTERTYPE

const (
	VideoAVCLoopFilterEnable               VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterEnable
	VideoAVCLoopFilterDisable              VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterDisable
	VideoAVCLoopFilterDisableSliceBoundary VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterDisableSliceBoundary
	VideoAVCLoopFilterKhronosExtensions    VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterKhronosExtensions
	VideoAVCLoopFilterVendorStartUnused    VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterVendorStartUnused
	VideoAVCLoopFilterMax                  VideoAVCLoopFilter = C.OMX_VIDEO_AVCLoopFilterMax
)

func (v VideoAVCLoopFilter) String() string {
	switch v {
	case VideoAVCLoopFilterEnable:
		return "VideoAVCLoopFilterEnable"
	case VideoAVCLoopFilterDisable:
		return "VideoAVCLoopFilterDisable"
	case VideoAVCLoopFilterDisableSliceBoundary:
		return "VideoAVCLoopFilterDisableSliceBoundary"
	case VideoAVCLoopFilterKhronosExtensions:
		return "VideoAVCLoopFilterKhronosExtensions"
	case VideoAVCLoopFilterVendorStartUnused:
		return "VideoAVCLoopFilterVendorStartUnused"
	case VideoAVCLoopFilterMax:
		return "VideoAVCLoopFilterMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type FaceDetectionControlMode C.OMX_FACEDETECTIONCONTROLTYPE

const (
	FaceDetectionControlNone              FaceDetectionControlMode = C.OMX_FaceDetectionControlNone
	FaceDetectionControlOn                FaceDetectionControlMode = C.OMX_FaceDetectionControlOn
	FaceDetectionControlKhronosExtensions FaceDetectionControlMode = C.OMX_FaceDetectionControlKhronosExtensions
	FaceDetectionControlVendorStartUnused FaceDetectionControlMode = C.OMX_FaceDetectionControlVendorStartUnused
	FaceDetectionControlMax               FaceDetectionControlMode = C.OMX_FaceDetectionControlMax
)

func (v FaceDetectionControlMode) String() string {
	switch v {
	case FaceDetectionControlNone:
		return "FaceDetectionControlNone"
	case FaceDetectionControlOn:
		return "FaceDetectionControlOn"
	case FaceDetectionControlKhronosExtensions:
		return "FaceDetectionControlKhronosExtensions"
	case FaceDetectionControlVendorStartUnused:
		return "FaceDetectionControlVendorStartUnused"
	case FaceDetectionControlMax:
		return "FaceDetectionControlMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type FaceRegionFlags C.OMX_FACEREGIONFLAGSTYPE

const (
	FaceRegionFlagsNone              FaceRegionFlags = C.OMX_FaceRegionFlagsNone
	FaceRegionFlagsBlink             FaceRegionFlags = C.OMX_FaceRegionFlagsBlink
	FaceRegionFlagsSmile             FaceRegionFlags = C.OMX_FaceRegionFlagsSmile
	FaceRegionFlagsKhronosExtensions FaceRegionFlags = C.OMX_FaceRegionFlagsKhronosExtensions
	FaceRegionFlagsVendorStartUnused FaceRegionFlags = C.OMX_FaceRegionFlagsVendorStartUnused
	FaceRegionFlagsMax               FaceRegionFlags = C.OMX_FaceRegionFlagsMax
)

func (v FaceRegionFlags) String() string {
	switch v {
	case FaceRegionFlagsNone:
		return "FaceRegionFlagsNone"
	case FaceRegionFlagsBlink:
		return "FaceRegionFlagsBlink"
	case FaceRegionFlagsSmile:
		return "FaceRegionFlagsSmile"
	case FaceRegionFlagsKhronosExtensions:
		return "FaceRegionFlagsKhronosExtensions"
	case FaceRegionFlagsVendorStartUnused:
		return "FaceRegionFlagsVendorStartUnused"
	case FaceRegionFlagsMax:
		return "FaceRegionFlagsMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}

type PortDomain C.OMX_PORTDOMAINTYPE

const (
	PortDomainAudio             PortDomain = C.OMX_PortDomainAudio
	PortDomainVideo             PortDomain = C.OMX_PortDomainVideo
	PortDomainImage             PortDomain = C.OMX_PortDomainImage
	PortDomainOther             PortDomain = C.OMX_PortDomainOther
	PortDomainKhronosExtensions PortDomain = C.OMX_PortDomainKhronosExtensions
	PortDomainVendorStartUnused PortDomain = C.OMX_PortDomainVendorStartUnused
	PortDomainMax               PortDomain = C.OMX_PortDomainMax
)

func (d PortDomain) String() string {
	switch d {
	case PortDomainAudio:
		return "PortDomainAudio"
	case PortDomainVideo:
		return "PortDomainVideo"
	case PortDomainImage:
		return "PortDomainImage"
	case PortDomainOther:
		return "PortDomainOther"
	case PortDomainKhronosExtensions:
		return "PortDomainKhronosExtensions"
	case PortDomainVendorStartUnused:
		return "PortDomainVendorStartUnused"
	case PortDomainMax:
		return "PortDomainMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(d))
}

type Direction C.OMX_DIRTYPE

const (
	DirInput  Direction = C.OMX_DirInput
	DirOutput Direction = C.OMX_DirOutput
	DirMax    Direction = C.OMX_DirMax
)

func (d Direction) String() string {
	switch d {
	case DirInput:
		return "DirInput"
	case DirOutput:
		return "DirOutput"
	case DirMax:
		return "DirMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(d))
}

type AudioCoding C.OMX_AUDIO_CODINGTYPE

const (
	AudioCodingUnused            AudioCoding = C.OMX_AUDIO_CodingUnused
	AudioCodingAutoDetect        AudioCoding = C.OMX_AUDIO_CodingAutoDetect
	AudioCodingPCM               AudioCoding = C.OMX_AUDIO_CodingPCM
	AudioCodingADPCM             AudioCoding = C.OMX_AUDIO_CodingADPCM
	AudioCodingAMR               AudioCoding = C.OMX_AUDIO_CodingAMR
	AudioCodingGSMFR             AudioCoding = C.OMX_AUDIO_CodingGSMFR
	AudioCodingGSMEFR            AudioCoding = C.OMX_AUDIO_CodingGSMEFR
	AudioCodingGSMHR             AudioCoding = C.OMX_AUDIO_CodingGSMHR
	AudioCodingPDCFR             AudioCoding = C.OMX_AUDIO_CodingPDCFR
	AudioCodingPDCEFR            AudioCoding = C.OMX_AUDIO_CodingPDCEFR
	AudioCodingPDCHR             AudioCoding = C.OMX_AUDIO_CodingPDCHR
	AudioCodingTDMAFR            AudioCoding = C.OMX_AUDIO_CodingTDMAFR
	AudioCodingTDMAEFR           AudioCoding = C.OMX_AUDIO_CodingTDMAEFR
	AudioCodingQCELP8            AudioCoding = C.OMX_AUDIO_CodingQCELP8
	AudioCodingQCELP13           AudioCoding = C.OMX_AUDIO_CodingQCELP13
	AudioCodingEVRC              AudioCoding = C.OMX_AUDIO_CodingEVRC
	AudioCodingSMV               AudioCoding = C.OMX_AUDIO_CodingSMV
	AudioCodingG711              AudioCoding = C.OMX_AUDIO_CodingG711
	AudioCodingG723              AudioCoding = C.OMX_AUDIO_CodingG723
	AudioCodingG726              AudioCoding = C.OMX_AUDIO_CodingG726
	AudioCodingG729              AudioCoding = C.OMX_AUDIO_CodingG729
	AudioCodingAAC               AudioCoding = C.OMX_AUDIO_CodingAAC
	AudioCodingMP3               AudioCoding = C.OMX_AUDIO_CodingMP3
	AudioCodingSBC               AudioCoding = C.OMX_AUDIO_CodingSBC
	AudioCodingVORBIS            AudioCoding = C.OMX_AUDIO_CodingVORBIS
	AudioCodingWMA               AudioCoding = C.OMX_AUDIO_CodingWMA
	AudioCodingRA                AudioCoding = C.OMX_AUDIO_CodingRA
	AudioCodingMIDI              AudioCoding = C.OMX_AUDIO_CodingMIDI
	AudioCodingKhronosExtensions AudioCoding = C.OMX_AUDIO_CodingKhronosExtensions
	AudioCodingVendorStartUnused AudioCoding = C.OMX_AUDIO_CodingVendorStartUnused
)

func (c AudioCoding) String() string {
	switch c {
	case AudioCodingUnused:
		return "AudioCodingUnused"
	case AudioCodingAutoDetect:
		return "AudioCodingAutoDetect"
	case AudioCodingPCM:
		return "AudioCodingPCM"
	case AudioCodingADPCM:
		return "AudioCodingADPCM"
	case AudioCodingAMR:
		return "AudioCodingAMR"
	case AudioCodingGSMFR:
		return "AudioCodingGSMFR"
	case AudioCodingGSMEFR:
		return "AudioCodingGSMEFR"
	case AudioCodingGSMHR:
		return "AudioCodingGSMHR"
	case AudioCodingPDCFR:
		return "AudioCodingPDCFR"
	case AudioCodingPDCEFR:
		return "AudioCodingPDCEFR"
	case AudioCodingPDCHR:
		return "AudioCodingPDCHR"
	case AudioCodingTDMAFR:
		return "AudioCodingTDMAFR"
	case AudioCodingTDMAEFR:
		return "AudioCodingTDMAEFR"
	case AudioCodingQCELP8:
		return "AudioCodingQCELP8"
	case AudioCodingQCELP13:
		return "AudioCodingQCELP13"
	case AudioCodingEVRC:
		return "AudioCodingEVRC"
	case AudioCodingSMV:
		return "AudioCodingSMV"
	case AudioCodingG711:
		return "AudioCodingG711"
	case AudioCodingG723:
		return "AudioCodingG723"
	case AudioCodingG726:
		return "AudioCodingG726"
	case AudioCodingG729:
		return "AudioCodingG729"
	case AudioCodingAAC:
		return "AudioCodingAAC"
	case AudioCodingMP3:
		return "AudioCodingMP3"
	case AudioCodingSBC:
		return "AudioCodingSBC"
	case AudioCodingVORBIS:
		return "AudioCodingVORBIS"
	case AudioCodingWMA:
		return "AudioCodingWMA"
	case AudioCodingRA:
		return "AudioCodingRA"
	case AudioCodingMIDI:
		return "AudioCodingMIDI"
	case AudioCodingKhronosExtensions:
		return "AudioCodingKhronosExtensions"
	case AudioCodingVendorStartUnused:
		return "AudioCodingVendorStartUnused"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(c))
}

type Index C.OMX_INDEXTYPE

const (
	IndexComponentStartUnused            Index = C.OMX_IndexComponentStartUnused
	IndexParamPriorityMgmt               Index = C.OMX_IndexParamPriorityMgmt
	IndexParamAudioInit                  Index = C.OMX_IndexParamAudioInit
	IndexParamImageInit                  Index = C.OMX_IndexParamImageInit
	IndexParamVideoInit                  Index = C.OMX_IndexParamVideoInit
	IndexParamOtherInit                  Index = C.OMX_IndexParamOtherInit
	IndexParamNumAvailableStreams        Index = C.OMX_IndexParamNumAvailableStreams
	IndexParamActiveStream               Index = C.OMX_IndexParamActiveStream
	IndexParamSuspensionPolicy           Index = C.OMX_IndexParamSuspensionPolicy
	IndexParamComponentSuspended         Index = C.OMX_IndexParamComponentSuspended
	IndexConfigCapturing                 Index = C.OMX_IndexConfigCapturing
	IndexConfigCaptureMode               Index = C.OMX_IndexConfigCaptureMode
	IndexAutoPauseAfterCapture           Index = C.OMX_IndexAutoPauseAfterCapture
	IndexParamContentURI                 Index = C.OMX_IndexParamContentURI
	IndexParamCustomContentPipe          Index = C.OMX_IndexParamCustomContentPipe
	IndexParamDisableResourceConcealment Index = C.OMX_IndexParamDisableResourceConcealment
	IndexConfigMetadataItemCount         Index = C.OMX_IndexConfigMetadataItemCount
	IndexConfigContainerNodeCount        Index = C.OMX_IndexConfigContainerNodeCount
	IndexConfigMetadataItem              Index = C.OMX_IndexConfigMetadataItem
	IndexConfigCounterNodeID             Index = C.OMX_IndexConfigCounterNodeID
	IndexParamMetadataFilterType         Index = C.OMX_IndexParamMetadataFilterType
	IndexParamMetadataKeyFilter          Index = C.OMX_IndexParamMetadataKeyFilter
	IndexConfigPriorityMgmt              Index = C.OMX_IndexConfigPriorityMgmt
	IndexParamStandardComponentRole      Index = C.OMX_IndexParamStandardComponentRole

	IndexPortStartUnused         Index = C.OMX_IndexPortStartUnused
	IndexParamPortDefinition     Index = C.OMX_IndexParamPortDefinition
	IndexParamCompBufferSupplier Index = C.OMX_IndexParamCompBufferSupplier
	IndexReservedStartUnused     Index = C.OMX_IndexReservedStartUnused

	/* Audio parameters and configurations */
	IndexAudioStartUnused            Index = C.OMX_IndexAudioStartUnused
	IndexParamAudioPortFormat        Index = C.OMX_IndexParamAudioPortFormat
	IndexParamAudioPcm               Index = C.OMX_IndexParamAudioPcm
	IndexParamAudioAac               Index = C.OMX_IndexParamAudioAac
	IndexParamAudioRa                Index = C.OMX_IndexParamAudioRa
	IndexParamAudioMp3               Index = C.OMX_IndexParamAudioMp3
	IndexParamAudioAdpcm             Index = C.OMX_IndexParamAudioAdpcm
	IndexParamAudioG723              Index = C.OMX_IndexParamAudioG723
	IndexParamAudioG729              Index = C.OMX_IndexParamAudioG729
	IndexParamAudioAmr               Index = C.OMX_IndexParamAudioAmr
	IndexParamAudioWma               Index = C.OMX_IndexParamAudioWma
	IndexParamAudioSbc               Index = C.OMX_IndexParamAudioSbc
	IndexParamAudioMidi              Index = C.OMX_IndexParamAudioMidi
	IndexParamAudioGsm_FR            Index = C.OMX_IndexParamAudioGsm_FR
	IndexParamAudioMidiLoadUserSound Index = C.OMX_IndexParamAudioMidiLoadUserSound
	IndexParamAudioG726              Index = C.OMX_IndexParamAudioG726
	IndexParamAudioGsm_EFR           Index = C.OMX_IndexParamAudioGsm_EFR
	IndexParamAudioGsm_HR            Index = C.OMX_IndexParamAudioGsm_HR
	IndexParamAudioPdc_FR            Index = C.OMX_IndexParamAudioPdc_FR
	IndexParamAudioPdc_EFR           Index = C.OMX_IndexParamAudioPdc_EFR
	IndexParamAudioPdc_HR            Index = C.OMX_IndexParamAudioPdc_HR
	IndexParamAudioTdma_FR           Index = C.OMX_IndexParamAudioTdma_FR
	IndexParamAudioTdma_EFR          Index = C.OMX_IndexParamAudioTdma_EFR
	IndexParamAudioQcelp8            Index = C.OMX_IndexParamAudioQcelp8
	IndexParamAudioQcelp13           Index = C.OMX_IndexParamAudioQcelp13
	IndexParamAudioEvrc              Index = C.OMX_IndexParamAudioEvrc
	IndexParamAudioSmv               Index = C.OMX_IndexParamAudioSmv
	IndexParamAudioVorbis            Index = C.OMX_IndexParamAudioVorbis

	IndexConfigAudioMidiImmediateEvent   Index = C.OMX_IndexConfigAudioMidiImmediateEvent
	IndexConfigAudioMidiControl          Index = C.OMX_IndexConfigAudioMidiControl
	IndexConfigAudioMidiSoundBankProgram Index = C.OMX_IndexConfigAudioMidiSoundBankProgram
	IndexConfigAudioMidiStatus           Index = C.OMX_IndexConfigAudioMidiStatus
	IndexConfigAudioMidiMetaEvent        Index = C.OMX_IndexConfigAudioMidiMetaEvent
	IndexConfigAudioMidiMetaEventData    Index = C.OMX_IndexConfigAudioMidiMetaEventData
	IndexConfigAudioVolume               Index = C.OMX_IndexConfigAudioVolume
	IndexConfigAudioBalance              Index = C.OMX_IndexConfigAudioBalance
	IndexConfigAudioChannelMute          Index = C.OMX_IndexConfigAudioChannelMute
	IndexConfigAudioMute                 Index = C.OMX_IndexConfigAudioMute
	IndexConfigAudioLoudness             Index = C.OMX_IndexConfigAudioLoudness
	IndexConfigAudioEchoCancelation      Index = C.OMX_IndexConfigAudioEchoCancelation
	IndexConfigAudioNoiseReduction       Index = C.OMX_IndexConfigAudioNoiseReduction
	IndexConfigAudioBass                 Index = C.OMX_IndexConfigAudioBass
	IndexConfigAudioTreble               Index = C.OMX_IndexConfigAudioTreble
	IndexConfigAudioStereoWidening       Index = C.OMX_IndexConfigAudioStereoWidening
	IndexConfigAudioChorus               Index = C.OMX_IndexConfigAudioChorus
	IndexConfigAudioEqualizer            Index = C.OMX_IndexConfigAudioEqualizer
	IndexConfigAudioReverberation        Index = C.OMX_IndexConfigAudioReverberation
	IndexConfigAudioChannelVolume        Index = C.OMX_IndexConfigAudioChannelVolume

	/* Image specific parameters and configurations */
	IndexImageStartUnused       Index = C.OMX_IndexImageStartUnused
	IndexParamImagePortFormat   Index = C.OMX_IndexParamImagePortFormat
	IndexParamFlashControl      Index = C.OMX_IndexParamFlashControl
	IndexConfigFocusControl     Index = C.OMX_IndexConfigFocusControl
	IndexParamQFactor           Index = C.OMX_IndexParamQFactor
	IndexParamQuantizationTable Index = C.OMX_IndexParamQuantizationTable
	IndexParamHuffmanTable      Index = C.OMX_IndexParamHuffmanTable
	IndexConfigFlashControl     Index = C.OMX_IndexConfigFlashControl

	/* Video specific parameters and configurations */
	IndexVideoStartUnused                     Index = C.OMX_IndexVideoStartUnused
	IndexParamVideoPortFormat                 Index = C.OMX_IndexParamVideoPortFormat
	IndexParamVideoQuantization               Index = C.OMX_IndexParamVideoQuantization
	IndexParamVideoFastUpdate                 Index = C.OMX_IndexParamVideoFastUpdate
	IndexParamVideoBitrate                    Index = C.OMX_IndexParamVideoBitrate
	IndexParamVideoMotionVector               Index = C.OMX_IndexParamVideoMotionVector
	IndexParamVideoIntraRefresh               Index = C.OMX_IndexParamVideoIntraRefresh
	IndexParamVideoErrorCorrection            Index = C.OMX_IndexParamVideoErrorCorrection
	IndexParamVideoVBSMC                      Index = C.OMX_IndexParamVideoVBSMC
	IndexParamVideoMpeg2                      Index = C.OMX_IndexParamVideoMpeg2
	IndexParamVideoMpeg4                      Index = C.OMX_IndexParamVideoMpeg4
	IndexParamVideoWmv                        Index = C.OMX_IndexParamVideoWmv
	IndexParamVideoRv                         Index = C.OMX_IndexParamVideoRv
	IndexParamVideoAvc                        Index = C.OMX_IndexParamVideoAvc
	IndexParamVideoH263                       Index = C.OMX_IndexParamVideoH263
	IndexParamVideoProfileLevelQuerySupported Index = C.OMX_IndexParamVideoProfileLevelQuerySupported
	IndexParamVideoProfileLevelCurrent        Index = C.OMX_IndexParamVideoProfileLevelCurrent
	IndexConfigVideoBitrate                   Index = C.OMX_IndexConfigVideoBitrate
	IndexConfigVideoFramerate                 Index = C.OMX_IndexConfigVideoFramerate
	IndexConfigVideoIntraVOPRefresh           Index = C.OMX_IndexConfigVideoIntraVOPRefresh
	IndexConfigVideoIntraMBRefresh            Index = C.OMX_IndexConfigVideoIntraMBRefresh
	IndexConfigVideoMBErrorReporting          Index = C.OMX_IndexConfigVideoMBErrorReporting
	IndexParamVideoMacroblocksPerFrame        Index = C.OMX_IndexParamVideoMacroblocksPerFrame
	IndexConfigVideoMacroBlockErrorMap        Index = C.OMX_IndexConfigVideoMacroBlockErrorMap
	IndexParamVideoSliceFMO                   Index = C.OMX_IndexParamVideoSliceFMO
	IndexConfigVideoAVCIntraPeriod            Index = C.OMX_IndexConfigVideoAVCIntraPeriod
	IndexConfigVideoNalSize                   Index = C.OMX_IndexConfigVideoNalSize

	/* Image & Video common Configurations */
	IndexCommonStartUnused                 Index = C.OMX_IndexCommonStartUnused
	IndexParamCommonDeblocking             Index = C.OMX_IndexParamCommonDeblocking
	IndexParamCommonSensorMode             Index = C.OMX_IndexParamCommonSensorMode
	IndexParamCommonInterleave             Index = C.OMX_IndexParamCommonInterleave
	IndexConfigCommonColorFormatConversion Index = C.OMX_IndexConfigCommonColorFormatConversion
	IndexConfigCommonScale                 Index = C.OMX_IndexConfigCommonScale
	IndexConfigCommonImageFilter           Index = C.OMX_IndexConfigCommonImageFilter
	IndexConfigCommonColorEnhancement      Index = C.OMX_IndexConfigCommonColorEnhancement
	IndexConfigCommonColorKey              Index = C.OMX_IndexConfigCommonColorKey
	IndexConfigCommonColorBlend            Index = C.OMX_IndexConfigCommonColorBlend
	IndexConfigCommonFrameStabilisation    Index = C.OMX_IndexConfigCommonFrameStabilisation
	IndexConfigCommonRotate                Index = C.OMX_IndexConfigCommonRotate
	IndexConfigCommonMirror                Index = C.OMX_IndexConfigCommonMirror
	IndexConfigCommonOutputPosition        Index = C.OMX_IndexConfigCommonOutputPosition
	IndexConfigCommonInputCrop             Index = C.OMX_IndexConfigCommonInputCrop
	IndexConfigCommonOutputCrop            Index = C.OMX_IndexConfigCommonOutputCrop
	IndexConfigCommonDigitalZoom           Index = C.OMX_IndexConfigCommonDigitalZoom
	IndexConfigCommonOpticalZoom           Index = C.OMX_IndexConfigCommonOpticalZoom
	IndexConfigCommonWhiteBalance          Index = C.OMX_IndexConfigCommonWhiteBalance
	IndexConfigCommonExposure              Index = C.OMX_IndexConfigCommonExposure
	IndexConfigCommonContrast              Index = C.OMX_IndexConfigCommonContrast
	IndexConfigCommonBrightness            Index = C.OMX_IndexConfigCommonBrightness
	IndexConfigCommonBacklight             Index = C.OMX_IndexConfigCommonBacklight
	IndexConfigCommonGamma                 Index = C.OMX_IndexConfigCommonGamma
	IndexConfigCommonSaturation            Index = C.OMX_IndexConfigCommonSaturation
	IndexConfigCommonLightness             Index = C.OMX_IndexConfigCommonLightness
	IndexConfigCommonExclusionRect         Index = C.OMX_IndexConfigCommonExclusionRect
	IndexConfigCommonDithering             Index = C.OMX_IndexConfigCommonDithering
	IndexConfigCommonPlaneBlend            Index = C.OMX_IndexConfigCommonPlaneBlend
	IndexConfigCommonExposureValue         Index = C.OMX_IndexConfigCommonExposureValue
	IndexConfigCommonOutputSize            Index = C.OMX_IndexConfigCommonOutputSize
	IndexParamCommonExtraQuantData         Index = C.OMX_IndexParamCommonExtraQuantData
	IndexConfigCommonFocusRegion           Index = C.OMX_IndexConfigCommonFocusRegion
	IndexConfigCommonFocusStatus           Index = C.OMX_IndexConfigCommonFocusStatus
	IndexConfigCommonTransitionEffect      Index = C.OMX_IndexConfigCommonTransitionEffect

	/* Reserved Configuration range */
	IndexOtherStartUnused     Index = C.OMX_IndexOtherStartUnused
	IndexParamOtherPortFormat Index = C.OMX_IndexParamOtherPortFormat
	IndexConfigOtherPower     Index = C.OMX_IndexConfigOtherPower
	IndexConfigOtherStats     Index = C.OMX_IndexConfigOtherStats

	/* Reserved Time range */
	IndexTimeStartUnused                 Index = C.OMX_IndexTimeStartUnused
	IndexConfigTimeScale                 Index = C.OMX_IndexConfigTimeScale
	IndexConfigTimeClockState            Index = C.OMX_IndexConfigTimeClockState
	IndexConfigTimeActiveRefClock        Index = C.OMX_IndexConfigTimeActiveRefClock
	IndexConfigTimeCurrentMediaTime      Index = C.OMX_IndexConfigTimeCurrentMediaTime
	IndexConfigTimeCurrentWallTime       Index = C.OMX_IndexConfigTimeCurrentWallTime
	IndexConfigTimeCurrentAudioReference Index = C.OMX_IndexConfigTimeCurrentAudioReference
	IndexConfigTimeCurrentVideoReference Index = C.OMX_IndexConfigTimeCurrentVideoReference
	IndexConfigTimeMediaTimeRequest      Index = C.OMX_IndexConfigTimeMediaTimeRequest
	IndexConfigTimeClientStartTime       Index = C.OMX_IndexConfigTimeClientStartTime
	IndexConfigTimePosition              Index = C.OMX_IndexConfigTimePosition
	IndexConfigTimeSeekMode              Index = C.OMX_IndexConfigTimeSeekMode

	IndexKhronosExtensions Index = C.OMX_IndexKhronosExtensions
	/* Vendor specific area */
	IndexVendorStartUnused Index = C.OMX_IndexVendorStartUnused
	/* Vendor specific structures should be in the range of 0x7F000000
	   to 0x7FFFFFFE.  This range is not broken out by vendor, so
	   private indexes are not guaranteed unique and therefore should
	   only be sent to the appropriate component. */

	/* used for ilcs-top communication */
	IndexParamMarkComparison      Index = C.OMX_IndexParamMarkComparison
	IndexParamPortSummary         Index = C.OMX_IndexParamPortSummary
	IndexParamTunnelStatus        Index = C.OMX_IndexParamTunnelStatus
	IndexParamBrcmRecursionUnsafe Index = C.OMX_IndexParamBrcmRecursionUnsafe

	/* used for top-ril communication */
	IndexParamBufferAddress     Index = C.OMX_IndexParamBufferAddress
	IndexParamTunnelSetup       Index = C.OMX_IndexParamTunnelSetup
	IndexParamBrcmPortEGL       Index = C.OMX_IndexParamBrcmPortEGL
	IndexParamIdleResourceCount Index = C.OMX_IndexParamIdleResourceCount

	/* used for ril-ril communication */
	IndexParamImagePoolDisplayFunction    Index = C.OMX_IndexParamImagePoolDisplayFunction
	IndexParamBrcmDataUnit                Index = C.OMX_IndexParamBrcmDataUnit
	IndexParamCodecConfig                 Index = C.OMX_IndexParamCodecConfig
	IndexParamCameraPoolToEncoderFunction Index = C.OMX_IndexParamCameraPoolToEncoderFunction
	IndexParamCameraStripeFunction        Index = C.OMX_IndexParamCameraStripeFunction
	IndexParamCameraCaptureEventFunction  Index = C.OMX_IndexParamCameraCaptureEventFunction

	/* used for client-ril communication */
	IndexParamTestInterface Index = C.OMX_IndexParamTestInterface

	// 0x7f000010
	IndexConfigDisplayRegion               Index = C.OMX_IndexConfigDisplayRegion
	IndexParamSource                       Index = C.OMX_IndexParamSource
	IndexParamSourceSeed                   Index = C.OMX_IndexParamSourceSeed
	IndexParamResize                       Index = C.OMX_IndexParamResize
	IndexConfigVisualisation               Index = C.OMX_IndexConfigVisualisation
	IndexConfigSingleStep                  Index = C.OMX_IndexConfigSingleStep
	IndexConfigPlayMode                    Index = C.OMX_IndexConfigPlayMode
	IndexParamCameraCamplusId              Index = C.OMX_IndexParamCameraCamplusId
	IndexConfigCommonImageFilterParameters Index = C.OMX_IndexConfigCommonImageFilterParameters
	IndexConfigTransitionControl           Index = C.OMX_IndexConfigTransitionControl
	IndexConfigPresentationOffset          Index = C.OMX_IndexConfigPresentationOffset
	IndexParamSourceFunctions              Index = C.OMX_IndexParamSourceFunctions
	IndexConfigAudioMonoTrackControl       Index = C.OMX_IndexConfigAudioMonoTrackControl
	IndexParamCameraImagePool              Index = C.OMX_IndexParamCameraImagePool
	IndexConfigCameraISPOutputPoolHeight   Index = C.OMX_IndexConfigCameraISPOutputPoolHeight
	IndexParamImagePoolSize                Index = C.OMX_IndexParamImagePoolSize

	// 0x7f000020
	IndexParamImagePoolExternal              Index = C.OMX_IndexParamImagePoolExternal
	IndexParamRUTILFifoInfo                  Index = C.OMX_IndexParamRUTILFifoInfo
	IndexParamILFifoConfig                   Index = C.OMX_IndexParamILFifoConfig
	IndexConfigCameraSensorModes             Index = C.OMX_IndexConfigCameraSensorModes
	IndexConfigBrcmPortStats                 Index = C.OMX_IndexConfigBrcmPortStats
	IndexConfigBrcmPortBufferStats           Index = C.OMX_IndexConfigBrcmPortBufferStats
	IndexConfigBrcmCameraStats               Index = C.OMX_IndexConfigBrcmCameraStats
	IndexConfigBrcmIOPerfStats               Index = C.OMX_IndexConfigBrcmIOPerfStats
	IndexConfigCommonSharpness               Index = C.OMX_IndexConfigCommonSharpness
	IndexConfigCommonFlickerCancellation     Index = C.OMX_IndexConfigCommonFlickerCancellation
	IndexParamCameraSwapImagePools           Index = C.OMX_IndexParamCameraSwapImagePools
	IndexParamCameraSingleBufferCaptureInput Index = C.OMX_IndexParamCameraSingleBufferCaptureInput
	IndexConfigCommonRedEyeRemoval           Index = C.OMX_IndexConfigCommonRedEyeRemoval
	IndexConfigCommonFaceDetectionControl    Index = C.OMX_IndexConfigCommonFaceDetectionControl
	IndexConfigCommonFaceDetectionRegion     Index = C.OMX_IndexConfigCommonFaceDetectionRegion
	IndexConfigCommonInterlace               Index = C.OMX_IndexConfigCommonInterlace

	// 0x7f000030
	IndexParamISPTunerName               Index = C.OMX_IndexParamISPTunerName
	IndexParamCameraDeviceNumber         Index = C.OMX_IndexParamCameraDeviceNumber
	IndexParamCameraDevicesPresent       Index = C.OMX_IndexParamCameraDevicesPresent
	IndexConfigCameraInputFrame          Index = C.OMX_IndexConfigCameraInputFrame
	IndexConfigStillColourDenoiseEnable  Index = C.OMX_IndexConfigStillColourDenoiseEnable
	IndexConfigVideoColourDenoiseEnable  Index = C.OMX_IndexConfigVideoColourDenoiseEnable
	IndexConfigAFAssistLight             Index = C.OMX_IndexConfigAFAssistLight
	IndexConfigSmartShakeReductionEnable Index = C.OMX_IndexConfigSmartShakeReductionEnable
	IndexConfigInputCropPercentages      Index = C.OMX_IndexConfigInputCropPercentages
	IndexConfigStillsAntiShakeEnable     Index = C.OMX_IndexConfigStillsAntiShakeEnable
	IndexConfigWaitForFocusBeforeCapture Index = C.OMX_IndexConfigWaitForFocusBeforeCapture
	IndexConfigAudioRenderingLatency     Index = C.OMX_IndexConfigAudioRenderingLatency
	IndexConfigDrawBoxAroundFaces        Index = C.OMX_IndexConfigDrawBoxAroundFaces
	IndexParamCodecRequirements          Index = C.OMX_IndexParamCodecRequirements
	IndexConfigBrcmEGLImageMemHandle     Index = C.OMX_IndexConfigBrcmEGLImageMemHandle
	IndexConfigPrivacyIndicator          Index = C.OMX_IndexConfigPrivacyIndicator

	// 0x7f000040
	IndexParamCameraFlashType                   Index = C.OMX_IndexParamCameraFlashType
	IndexConfigCameraEnableStatsPass            Index = C.OMX_IndexConfigCameraEnableStatsPass
	IndexConfigCameraFlashConfig                Index = C.OMX_IndexConfigCameraFlashConfig
	IndexConfigCaptureRawImageURI               Index = C.OMX_IndexConfigCaptureRawImageURI
	IndexConfigCameraStripeFuncMinLines         Index = C.OMX_IndexConfigCameraStripeFuncMinLines
	IndexConfigCameraAlgorithmVersionDeprecated Index = C.OMX_IndexConfigCameraAlgorithmVersionDeprecated
	IndexConfigCameraIsoReferenceValue          Index = C.OMX_IndexConfigCameraIsoReferenceValue
	IndexConfigCameraCaptureAbortsAutoFocus     Index = C.OMX_IndexConfigCameraCaptureAbortsAutoFocus
	IndexConfigBrcmClockMissCount               Index = C.OMX_IndexConfigBrcmClockMissCount
	IndexConfigFlashChargeLevel                 Index = C.OMX_IndexConfigFlashChargeLevel
	IndexConfigBrcmVideoEncodedSliceSize        Index = C.OMX_IndexConfigBrcmVideoEncodedSliceSize
	IndexConfigBrcmAudioTrackGaplessPlayback    Index = C.OMX_IndexConfigBrcmAudioTrackGaplessPlayback
	IndexConfigBrcmAudioTrackChangeControl      Index = C.OMX_IndexConfigBrcmAudioTrackChangeControl
	IndexParamBrcmPixelAspectRatio              Index = C.OMX_IndexParamBrcmPixelAspectRatio
	IndexParamBrcmPixelValueRange               Index = C.OMX_IndexParamBrcmPixelValueRange
	IndexParamCameraDisableAlgorithm            Index = C.OMX_IndexParamCameraDisableAlgorithm

	// 0x7f000050
	IndexConfigBrcmVideoIntraPeriodTime                     Index = C.OMX_IndexConfigBrcmVideoIntraPeriodTime
	IndexConfigBrcmVideoIntraPeriod                         Index = C.OMX_IndexConfigBrcmVideoIntraPeriod
	IndexConfigBrcmAudioEffectControl                       Index = C.OMX_IndexConfigBrcmAudioEffectControl
	IndexConfigBrcmMinimumProcessingLatency                 Index = C.OMX_IndexConfigBrcmMinimumProcessingLatency
	IndexParamBrcmVideoAVCSEIEnable                         Index = C.OMX_IndexParamBrcmVideoAVCSEIEnable
	IndexParamBrcmAllowMemChange                            Index = C.OMX_IndexParamBrcmAllowMemChange
	IndexConfigBrcmVideoEncoderMBRowsPerSlice               Index = C.OMX_IndexConfigBrcmVideoEncoderMBRowsPerSlice
	IndexParamCameraAFAssistDeviceNumber_Deprecated         Index = C.OMX_IndexParamCameraAFAssistDeviceNumber_Deprecated
	IndexParamCameraPrivacyIndicatorDeviceNumber_Deprecated Index = C.OMX_IndexParamCameraPrivacyIndicatorDeviceNumber_Deprecated
	IndexConfigCameraUseCase                                Index = C.OMX_IndexConfigCameraUseCase
	IndexParamBrcmDisableProprietaryTunnels                 Index = C.OMX_IndexParamBrcmDisableProprietaryTunnels
	IndexParamBrcmOutputBufferSize                          Index = C.OMX_IndexParamBrcmOutputBufferSize
	IndexParamBrcmRetainMemory                              Index = C.OMX_IndexParamBrcmRetainMemory
	IndexConfigCanFocus_Deprecated                          Index = C.OMX_IndexConfigCanFocus_Deprecated
	IndexParamBrcmImmutableInput                            Index = C.OMX_IndexParamBrcmImmutableInput
	IndexParamDynamicParameterFile                          Index = C.OMX_IndexParamDynamicParameterFile

	// 0x7f000060
	IndexParamUseDynamicParameterFile     Index = C.OMX_IndexParamUseDynamicParameterFile
	IndexConfigCameraInfo                 Index = C.OMX_IndexConfigCameraInfo
	IndexConfigCameraFeatures             Index = C.OMX_IndexConfigCameraFeatures
	IndexConfigRequestCallback            Index = C.OMX_IndexConfigRequestCallback
	IndexConfigBrcmOutputBufferFullCount  Index = C.OMX_IndexConfigBrcmOutputBufferFullCount
	IndexConfigCommonFocusRegionXY        Index = C.OMX_IndexConfigCommonFocusRegionXY
	IndexParamBrcmDisableEXIF             Index = C.OMX_IndexParamBrcmDisableEXIF
	IndexConfigUserSettingsId             Index = C.OMX_IndexConfigUserSettingsId
	IndexConfigCameraSettings             Index = C.OMX_IndexConfigCameraSettings
	IndexConfigDrawBoxLineParams          Index = C.OMX_IndexConfigDrawBoxLineParams
	IndexParamCameraRmiControl_Deprecated Index = C.OMX_IndexParamCameraRmiControl_Deprecated
	IndexConfigBurstCapture               Index = C.OMX_IndexConfigBurstCapture
	IndexParamBrcmEnableIJGTableScaling   Index = C.OMX_IndexParamBrcmEnableIJGTableScaling
	IndexConfigPowerDown                  Index = C.OMX_IndexConfigPowerDown
	IndexConfigBrcmSyncOutput             Index = C.OMX_IndexConfigBrcmSyncOutput
	IndexParamBrcmFlushCallback           Index = C.OMX_IndexParamBrcmFlushCallback

	// 0x7f000070
	IndexConfigBrcmVideoRequestIFrame            Index = C.OMX_IndexConfigBrcmVideoRequestIFrame
	IndexParamBrcmNALSSeparate                   Index = C.OMX_IndexParamBrcmNALSSeparate
	IndexConfigConfirmView                       Index = C.OMX_IndexConfigConfirmView
	IndexConfigDrmView                           Index = C.OMX_IndexConfigDrmView
	IndexConfigBrcmVideoIntraRefresh             Index = C.OMX_IndexConfigBrcmVideoIntraRefresh
	IndexParamBrcmMaxFileSize                    Index = C.OMX_IndexParamBrcmMaxFileSize
	IndexParamBrcmCRCEnable                      Index = C.OMX_IndexParamBrcmCRCEnable
	IndexParamBrcmCRC                            Index = C.OMX_IndexParamBrcmCRC
	IndexConfigCameraRmiInUse_Deprecated         Index = C.OMX_IndexConfigCameraRmiInUse_Deprecated
	IndexConfigBrcmAudioSource                   Index = C.OMX_IndexConfigBrcmAudioSource
	IndexConfigBrcmAudioDestination              Index = C.OMX_IndexConfigBrcmAudioDestination
	IndexParamAudioDdp                           Index = C.OMX_IndexParamAudioDdp
	IndexParamBrcmThumbnail                      Index = C.OMX_IndexParamBrcmThumbnail
	IndexParamBrcmDisableLegacyBlocks_Deprecated Index = C.OMX_IndexParamBrcmDisableLegacyBlocks_Deprecated
	IndexParamBrcmCameraInputAspectRatio         Index = C.OMX_IndexParamBrcmCameraInputAspectRatio
	IndexParamDynamicParameterFileFailFatal      Index = C.OMX_IndexParamDynamicParameterFileFailFatal

	// 0x7f000080
	IndexParamBrcmVideoDecodeErrorConcealment   Index = C.OMX_IndexParamBrcmVideoDecodeErrorConcealment
	IndexParamBrcmInterpolateMissingTimestamps  Index = C.OMX_IndexParamBrcmInterpolateMissingTimestamps
	IndexParamBrcmSetCodecPerformanceMonitoring Index = C.OMX_IndexParamBrcmSetCodecPerformanceMonitoring
	IndexConfigFlashInfo                        Index = C.OMX_IndexConfigFlashInfo
	IndexParamBrcmMaxFrameSkips                 Index = C.OMX_IndexParamBrcmMaxFrameSkips
	IndexConfigDynamicRangeExpansion            Index = C.OMX_IndexConfigDynamicRangeExpansion
	IndexParamBrcmFlushCallbackId               Index = C.OMX_IndexParamBrcmFlushCallbackId
	IndexParamBrcmTransposeBufferCount          Index = C.OMX_IndexParamBrcmTransposeBufferCount
	IndexConfigFaceRecognitionControl           Index = C.OMX_IndexConfigFaceRecognitionControl
	IndexConfigFaceRecognitionSaveFace          Index = C.OMX_IndexConfigFaceRecognitionSaveFace
	IndexConfigFaceRecognitionDatabaseUri       Index = C.OMX_IndexConfigFaceRecognitionDatabaseUri
	IndexConfigClockAdjustment                  Index = C.OMX_IndexConfigClockAdjustment
	IndexParamBrcmThreadAffinity                Index = C.OMX_IndexParamBrcmThreadAffinity
	IndexParamAsynchronousOutput                Index = C.OMX_IndexParamAsynchronousOutput
	IndexConfigAsynchronousFailureURI           Index = C.OMX_IndexConfigAsynchronousFailureURI
	IndexConfigCommonFaceBeautification         Index = C.OMX_IndexConfigCommonFaceBeautification

	// 0x7f000090
	IndexConfigCommonSceneDetectionControl Index = C.OMX_IndexConfigCommonSceneDetectionControl
	IndexConfigCommonSceneDetected         Index = C.OMX_IndexConfigCommonSceneDetected
	IndexParamDisableVllPool               Index = C.OMX_IndexParamDisableVllPool
	IndexParamVideoMvc                     Index = C.OMX_IndexParamVideoMvc
	IndexConfigBrcmDrawStaticBox           Index = C.OMX_IndexConfigBrcmDrawStaticBox
	IndexConfigBrcmClockReferenceSource    Index = C.OMX_IndexConfigBrcmClockReferenceSource
	IndexParamPassBufferMarks              Index = C.OMX_IndexParamPassBufferMarks
	IndexConfigPortCapturing               Index = C.OMX_IndexConfigPortCapturing
	IndexConfigBrcmDecoderPassThrough      Index = C.OMX_IndexConfigBrcmDecoderPassThrough
	IndexParamBrcmDecoderPassThrough       Index = C.OMX_IndexParamBrcmDecoderPassThrough
	IndexParamBrcmMaxCorruptMBs            Index = C.OMX_IndexParamBrcmMaxCorruptMBs
	IndexConfigBrcmGlobalAudioMute         Index = C.OMX_IndexConfigBrcmGlobalAudioMute
	IndexParamCameraCaptureMode            Index = C.OMX_IndexParamCameraCaptureMode
	IndexParamBrcmDrmEncryption            Index = C.OMX_IndexParamBrcmDrmEncryption
	IndexConfigBrcmCameraRnDPreprocess     Index = C.OMX_IndexConfigBrcmCameraRnDPreprocess
	IndexConfigBrcmCameraRnDPostprocess    Index = C.OMX_IndexConfigBrcmCameraRnDPostprocess
	IndexConfigBrcmAudioTrackChangeCount   Index = C.OMX_IndexConfigBrcmAudioTrackChangeCount

	// 0x7f0000a0
	IndexParamCommonUseStcTimestamps      Index = C.OMX_IndexParamCommonUseStcTimestamps
	IndexConfigBufferStall                Index = C.OMX_IndexConfigBufferStall
	IndexConfigRefreshCodec               Index = C.OMX_IndexConfigRefreshCodec
	IndexParamCaptureStatus               Index = C.OMX_IndexParamCaptureStatus
	IndexConfigTimeInvalidStartTime       Index = C.OMX_IndexConfigTimeInvalidStartTime
	IndexConfigLatencyTarget              Index = C.OMX_IndexConfigLatencyTarget
	IndexConfigMinimiseFragmentation      Index = C.OMX_IndexConfigMinimiseFragmentation
	IndexConfigBrcmUseProprietaryCallback Index = C.OMX_IndexConfigBrcmUseProprietaryCallback
	IndexParamPortMaxFrameSize            Index = C.OMX_IndexParamPortMaxFrameSize
	IndexParamComponentName               Index = C.OMX_IndexParamComponentName
	IndexConfigEncLevelExtension          Index = C.OMX_IndexConfigEncLevelExtension
	IndexConfigTemporalDenoiseEnable      Index = C.OMX_IndexConfigTemporalDenoiseEnable
	IndexParamBrcmLazyImagePoolDestroy    Index = C.OMX_IndexParamBrcmLazyImagePoolDestroy
	IndexParamBrcmEEDEEnable              Index = C.OMX_IndexParamBrcmEEDEEnable
	IndexParamBrcmEEDELossRate            Index = C.OMX_IndexParamBrcmEEDELossRate
	IndexParamAudioDts                    Index = C.OMX_IndexParamAudioDts

	// 0x7f0000b0
	IndexParamNumOutputChannels       Index = C.OMX_IndexParamNumOutputChannels
	IndexConfigBrcmHighDynamicRange   Index = C.OMX_IndexConfigBrcmHighDynamicRange
	IndexConfigBrcmPoolMemAllocSize   Index = C.OMX_IndexConfigBrcmPoolMemAllocSize
	IndexConfigBrcmBufferFlagFilter   Index = C.OMX_IndexConfigBrcmBufferFlagFilter
	IndexParamBrcmVideoEncodeMinQuant Index = C.OMX_IndexParamBrcmVideoEncodeMinQuant
	IndexParamBrcmVideoEncodeMaxQuant Index = C.OMX_IndexParamBrcmVideoEncodeMaxQuant
	IndexParamRateControlModel        Index = C.OMX_IndexParamRateControlModel
	IndexParamBrcmExtraBuffers        Index = C.OMX_IndexParamBrcmExtraBuffers
	IndexConfigFieldOfView            Index = C.OMX_IndexConfigFieldOfView
	IndexParamBrcmAlignHoriz          Index = C.OMX_IndexParamBrcmAlignHoriz
	IndexParamBrcmAlignVert           Index = C.OMX_IndexParamBrcmAlignVert
	IndexParamColorSpace              Index = C.OMX_IndexParamColorSpace
	IndexParamBrcmDroppablePFrames    Index = C.OMX_IndexParamBrcmDroppablePFrames
	IndexParamBrcmVideoInitialQuant   Index = C.OMX_IndexParamBrcmVideoInitialQuant
	IndexParamBrcmVideoEncodeQpP      Index = C.OMX_IndexParamBrcmVideoEncodeQpP
	IndexParamBrcmVideoRCSliceDQuant  Index = C.OMX_IndexParamBrcmVideoRCSliceDQuant

	// 0x7f0000c0
	IndexParamBrcmVideoFrameLimitBits      Index = C.OMX_IndexParamBrcmVideoFrameLimitBits
	IndexParamBrcmVideoPeakRate            Index = C.OMX_IndexParamBrcmVideoPeakRate
	IndexConfigBrcmVideoH264DisableCABAC   Index = C.OMX_IndexConfigBrcmVideoH264DisableCABAC
	IndexConfigBrcmVideoH264LowLatency     Index = C.OMX_IndexConfigBrcmVideoH264LowLatency
	IndexConfigBrcmVideoH264AUDelimiters   Index = C.OMX_IndexConfigBrcmVideoH264AUDelimiters
	IndexConfigBrcmVideoH264DeblockIDC     Index = C.OMX_IndexConfigBrcmVideoH264DeblockIDC
	IndexConfigBrcmVideoH264IntraMBMode    Index = C.OMX_IndexConfigBrcmVideoH264IntraMBMode
	IndexConfigContrastEnhance             Index = C.OMX_IndexConfigContrastEnhance
	IndexParamCameraCustomSensorConfig     Index = C.OMX_IndexParamCameraCustomSensorConfig
	IndexParamBrcmHeaderOnOpen             Index = C.OMX_IndexParamBrcmHeaderOnOpen
	IndexConfigBrcmUseRegisterFile         Index = C.OMX_IndexConfigBrcmUseRegisterFile
	IndexConfigBrcmRegisterFileFailFatal   Index = C.OMX_IndexConfigBrcmRegisterFileFailFatal
	IndexParamBrcmConfigFileRegisters      Index = C.OMX_IndexParamBrcmConfigFileRegisters
	IndexParamBrcmConfigFileChunkRegisters Index = C.OMX_IndexParamBrcmConfigFileChunkRegisters
	IndexParamBrcmAttachLog                Index = C.OMX_IndexParamBrcmAttachLog
	IndexParamCameraZeroShutterLag         Index = C.OMX_IndexParamCameraZeroShutterLag

	// 0x7f0000d0
	IndexParamBrcmFpsRange                   Index = C.OMX_IndexParamBrcmFpsRange
	IndexParamCaptureExposureCompensation    Index = C.OMX_IndexParamCaptureExposureCompensation
	IndexParamBrcmVideoPrecodeForQP          Index = C.OMX_IndexParamBrcmVideoPrecodeForQP
	IndexParamBrcmVideoTimestampFifo         Index = C.OMX_IndexParamBrcmVideoTimestampFifo
	IndexParamSWSharpenDisable               Index = C.OMX_IndexParamSWSharpenDisable
	IndexConfigBrcmFlashRequired             Index = C.OMX_IndexConfigBrcmFlashRequired
	IndexParamBrcmVideoDrmProtectBuffer      Index = C.OMX_IndexParamBrcmVideoDrmProtectBuffer
	IndexParamSWSaturationDisable            Index = C.OMX_IndexParamSWSaturationDisable
	IndexParamBrcmVideoDecodeConfigVD3       Index = C.OMX_IndexParamBrcmVideoDecodeConfigVD3
	IndexConfigBrcmPowerMonitor              Index = C.OMX_IndexConfigBrcmPowerMonitor
	IndexParamBrcmZeroCopy                   Index = C.OMX_IndexParamBrcmZeroCopy
	IndexParamBrcmVideoEGLRenderDiscardMode  Index = C.OMX_IndexParamBrcmVideoEGLRenderDiscardMode
	IndexParamBrcmVideoAVC_VCLHRDEnable      Index = C.OMX_IndexParamBrcmVideoAVC_VCLHRDEnable
	IndexParamBrcmVideoAVC_LowDelayHRDEnable Index = C.OMX_IndexParamBrcmVideoAVC_LowDelayHRDEnable
	IndexParamBrcmVideoCroppingDisable       Index = C.OMX_IndexParamBrcmVideoCroppingDisable
	IndexParamBrcmVideoAVCInlineHeaderEnable Index = C.OMX_IndexParamBrcmVideoAVCInlineHeaderEnable

	// 0x7f0000f0
	IndexConfigBrcmAudioDownmixCoefficients    Index = C.OMX_IndexConfigBrcmAudioDownmixCoefficients
	IndexConfigBrcmAudioDownmixCoefficients8x8 Index = C.OMX_IndexConfigBrcmAudioDownmixCoefficients8x8
	IndexConfigBrcmAudioMaxSample              Index = C.OMX_IndexConfigBrcmAudioMaxSample
	IndexConfigCustomAwbGains                  Index = C.OMX_IndexConfigCustomAwbGains
	IndexParamRemoveImagePadding               Index = C.OMX_IndexParamRemoveImagePadding
	IndexParamBrcmVideoAVCInlineVectorsEnable  Index = C.OMX_IndexParamBrcmVideoAVCInlineVectorsEnable
	IndexConfigBrcmRenderStats                 Index = C.OMX_IndexConfigBrcmRenderStats
	IndexConfigBrcmCameraAnnotate              Index = C.OMX_IndexConfigBrcmCameraAnnotate
	IndexParamBrcmStereoscopicMode             Index = C.OMX_IndexParamBrcmStereoscopicMode
	IndexParamBrcmLockStepEnable               Index = C.OMX_IndexParamBrcmLockStepEnable
	IndexParamBrcmTimeScale                    Index = C.OMX_IndexParamBrcmTimeScale
	IndexParamCameraInterface                  Index = C.OMX_IndexParamCameraInterface
	IndexParamCameraClockingMode               Index = C.OMX_IndexParamCameraClockingMode
	IndexParamCameraRxConfig                   Index = C.OMX_IndexParamCameraRxConfig
	IndexParamCameraRxTiming                   Index = C.OMX_IndexParamCameraRxTiming
	IndexParamDynamicParameterConfig           Index = C.OMX_IndexParamDynamicParameterConfig

	// 0x7f000100
	IndexParamBrcmVideoAVCSPSTimingEnable Index = C.OMX_IndexParamBrcmVideoAVCSPSTimingEnable
	IndexParamBrcmBayerOrder              Index = C.OMX_IndexParamBrcmBayerOrder
	IndexParamBrcmMaxNumCallbacks         Index = C.OMX_IndexParamBrcmMaxNumCallbacks
	IndexParamBrcmJpegRestartInterval     Index = C.OMX_IndexParamBrcmJpegRestartInterval
	IndexParamBrcmSupportsSlices          Index = C.OMX_IndexParamBrcmSupportsSlices
	// IndexParamBrcmIspBlockOverride             Index = C.OMX_IndexParamBrcmIspBlockOverride
	// IndexParamBrcmSupportsUnalignedSliceheight Index = C.OMX_IndexParamBrcmSupportsUnalignedSliceheight
	// IndexParamBrcmLensShadingOverride          Index = C.OMX_IndexParamBrcmLensShadingOverride
	// IndexParamBrcmBlackLevel                   Index = C.OMX_IndexParamBrcmBlackLevel
	// IndexParamOutputShift                      Index = C.OMX_IndexParamOutputShift
	// IndexParamCcmShift                         Index = C.OMX_IndexParamCcmShift
	// IndexParamCustomCcm                        Index = C.OMX_IndexParamCustomCcm
	// IndexConfigCameraAnalogGain                Index = C.OMX_IndexConfigCameraAnalogGain
	// IndexConfigCameraDigitalGain               Index = C.OMX_IndexConfigCameraDigitalGain
	// IndexConfigBrcmDroppableRunLength          Index = C.OMX_IndexConfigBrcmDroppableRunLength
	// IndexParamMinimumAlignment                 Index = C.OMX_IndexParamMinimumAlignment
	IndexMax Index = C.OMX_IndexMax
)

func (v Index) String() string {
	switch v {
	case IndexComponentStartUnused:
		return "IndexComponentStartUnused"
	case IndexParamPriorityMgmt:
		return "IndexParamPriorityMgmt"
	case IndexParamAudioInit:
		return "IndexParamAudioInit"
	case IndexParamImageInit:
		return "IndexParamImageInit"
	case IndexParamVideoInit:
		return "IndexParamVideoInit"
	case IndexParamOtherInit:
		return "IndexParamOtherInit"
	case IndexParamNumAvailableStreams:
		return "IndexParamNumAvailableStreams"
	case IndexParamActiveStream:
		return "IndexParamActiveStream"
	case IndexParamSuspensionPolicy:
		return "IndexParamSuspensionPolicy"
	case IndexParamComponentSuspended:
		return "IndexParamComponentSuspended"
	case IndexConfigCapturing:
		return "IndexConfigCapturing"
	case IndexConfigCaptureMode:
		return "IndexConfigCaptureMode"
	case IndexAutoPauseAfterCapture:
		return "IndexAutoPauseAfterCapture"
	case IndexParamContentURI:
		return "IndexParamContentURI"
	case IndexParamCustomContentPipe:
		return "IndexParamCustomContentPipe"
	case IndexParamDisableResourceConcealment:
		return "IndexParamDisableResourceConcealment"
	case IndexConfigMetadataItemCount:
		return "IndexConfigMetadataItemCount"
	case IndexConfigContainerNodeCount:
		return "IndexConfigContainerNodeCount"
	case IndexConfigMetadataItem:
		return "IndexConfigMetadataItem"
	case IndexConfigCounterNodeID:
		return "IndexConfigCounterNodeID"
	case IndexParamMetadataFilterType:
		return "IndexParamMetadataFilterType"
	case IndexParamMetadataKeyFilter:
		return "IndexParamMetadataKeyFilter"
	case IndexConfigPriorityMgmt:
		return "IndexConfigPriorityMgmt"
	case IndexParamStandardComponentRole:
		return "IndexParamStandardComponentRole"

	case IndexPortStartUnused:
		return "IndexPortStartUnused"
	case IndexParamPortDefinition:
		return "IndexParamPortDefinition"
	case IndexParamCompBufferSupplier:
		return "IndexParamCompBufferSupplier"
	case IndexReservedStartUnused:
		return "IndexReservedStartUnused"

		/* Audio parameters and configurations */
	case IndexAudioStartUnused:
		return "IndexAudioStartUnused"
	case IndexParamAudioPortFormat:
		return "IndexParamAudioPortFormat"
	case IndexParamAudioPcm:
		return "IndexParamAudioPcm"
	case IndexParamAudioAac:
		return "IndexParamAudioAac"
	case IndexParamAudioRa:
		return "IndexParamAudioRa"
	case IndexParamAudioMp3:
		return "IndexParamAudioMp3"
	case IndexParamAudioAdpcm:
		return "IndexParamAudioAdpcm"
	case IndexParamAudioG723:
		return "IndexParamAudioG723"
	case IndexParamAudioG729:
		return "IndexParamAudioG729"
	case IndexParamAudioAmr:
		return "IndexParamAudioAmr"
	case IndexParamAudioWma:
		return "IndexParamAudioWma"
	case IndexParamAudioSbc:
		return "IndexParamAudioSbc"
	case IndexParamAudioMidi:
		return "IndexParamAudioMidi"
	case IndexParamAudioGsm_FR:
		return "IndexParamAudioGsm_FR"
	case IndexParamAudioMidiLoadUserSound:
		return "IndexParamAudioMidiLoadUserSound"
	case IndexParamAudioG726:
		return "IndexParamAudioG726"
	case IndexParamAudioGsm_EFR:
		return "IndexParamAudioGsm_EFR"
	case IndexParamAudioGsm_HR:
		return "IndexParamAudioGsm_HR"
	case IndexParamAudioPdc_FR:
		return "IndexParamAudioPdc_FR"
	case IndexParamAudioPdc_EFR:
		return "IndexParamAudioPdc_EFR"
	case IndexParamAudioPdc_HR:
		return "IndexParamAudioPdc_HR"
	case IndexParamAudioTdma_FR:
		return "IndexParamAudioTdma_FR"
	case IndexParamAudioTdma_EFR:
		return "IndexParamAudioTdma_EFR"
	case IndexParamAudioQcelp8:
		return "IndexParamAudioQcelp8"
	case IndexParamAudioQcelp13:
		return "IndexParamAudioQcelp13"
	case IndexParamAudioEvrc:
		return "IndexParamAudioEvrc"
	case IndexParamAudioSmv:
		return "IndexParamAudioSmv"
	case IndexParamAudioVorbis:
		return "IndexParamAudioVorbis"

	case IndexConfigAudioMidiImmediateEvent:
		return "IndexConfigAudioMidiImmediateEvent"
	case IndexConfigAudioMidiControl:
		return "IndexConfigAudioMidiControl"
	case IndexConfigAudioMidiSoundBankProgram:
		return "IndexConfigAudioMidiSoundBankProgram"
	case IndexConfigAudioMidiStatus:
		return "IndexConfigAudioMidiStatus"
	case IndexConfigAudioMidiMetaEvent:
		return "IndexConfigAudioMidiMetaEvent"
	case IndexConfigAudioMidiMetaEventData:
		return "IndexConfigAudioMidiMetaEventData"
	case IndexConfigAudioVolume:
		return "IndexConfigAudioVolume"
	case IndexConfigAudioBalance:
		return "IndexConfigAudioBalance"
	case IndexConfigAudioChannelMute:
		return "IndexConfigAudioChannelMute"
	case IndexConfigAudioMute:
		return "IndexConfigAudioMute"
	case IndexConfigAudioLoudness:
		return "IndexConfigAudioLoudness"
	case IndexConfigAudioEchoCancelation:
		return "IndexConfigAudioEchoCancelation"
	case IndexConfigAudioNoiseReduction:
		return "IndexConfigAudioNoiseReduction"
	case IndexConfigAudioBass:
		return "IndexConfigAudioBass"
	case IndexConfigAudioTreble:
		return "IndexConfigAudioTreble"
	case IndexConfigAudioStereoWidening:
		return "IndexConfigAudioStereoWidening"
	case IndexConfigAudioChorus:
		return "IndexConfigAudioChorus"
	case IndexConfigAudioEqualizer:
		return "IndexConfigAudioEqualizer"
	case IndexConfigAudioReverberation:
		return "IndexConfigAudioReverberation"
	case IndexConfigAudioChannelVolume:
		return "IndexConfigAudioChannelVolume"

		/* Image specific parameters and configurations */
	case IndexImageStartUnused:
		return "IndexImageStartUnused"
	case IndexParamImagePortFormat:
		return "IndexParamImagePortFormat"
	case IndexParamFlashControl:
		return "IndexParamFlashControl"
	case IndexConfigFocusControl:
		return "IndexConfigFocusControl"
	case IndexParamQFactor:
		return "IndexParamQFactor"
	case IndexParamQuantizationTable:
		return "IndexParamQuantizationTable"
	case IndexParamHuffmanTable:
		return "IndexParamHuffmanTable"
	case IndexConfigFlashControl:
		return "IndexConfigFlashControl"

		/* Video specific parameters and configurations */
	case IndexVideoStartUnused:
		return "IndexVideoStartUnused"
	case IndexParamVideoPortFormat:
		return "IndexParamVideoPortFormat"
	case IndexParamVideoQuantization:
		return "IndexParamVideoQuantization"
	case IndexParamVideoFastUpdate:
		return "IndexParamVideoFastUpdate"
	case IndexParamVideoBitrate:
		return "IndexParamVideoBitrate"
	case IndexParamVideoMotionVector:
		return "IndexParamVideoMotionVector"
	case IndexParamVideoIntraRefresh:
		return "IndexParamVideoIntraRefresh"
	case IndexParamVideoErrorCorrection:
		return "IndexParamVideoErrorCorrection"
	case IndexParamVideoVBSMC:
		return "IndexParamVideoVBSMC"
	case IndexParamVideoMpeg2:
		return "IndexParamVideoMpeg2"
	case IndexParamVideoMpeg4:
		return "IndexParamVideoMpeg4"
	case IndexParamVideoWmv:
		return "IndexParamVideoWmv"
	case IndexParamVideoRv:
		return "IndexParamVideoRv"
	case IndexParamVideoAvc:
		return "IndexParamVideoAvc"
	case IndexParamVideoH263:
		return "IndexParamVideoH263"
	case IndexParamVideoProfileLevelQuerySupported:
		return "IndexParamVideoProfileLevelQuerySupported"
	case IndexParamVideoProfileLevelCurrent:
		return "IndexParamVideoProfileLevelCurrent"
	case IndexConfigVideoBitrate:
		return "IndexConfigVideoBitrate"
	case IndexConfigVideoFramerate:
		return "IndexConfigVideoFramerate"
	case IndexConfigVideoIntraVOPRefresh:
		return "IndexConfigVideoIntraVOPRefresh"
	case IndexConfigVideoIntraMBRefresh:
		return "IndexConfigVideoIntraMBRefresh"
	case IndexConfigVideoMBErrorReporting:
		return "IndexConfigVideoMBErrorReporting"
	case IndexParamVideoMacroblocksPerFrame:
		return "IndexParamVideoMacroblocksPerFrame"
	case IndexConfigVideoMacroBlockErrorMap:
		return "IndexConfigVideoMacroBlockErrorMap"
	case IndexParamVideoSliceFMO:
		return "IndexParamVideoSliceFMO"
	case IndexConfigVideoAVCIntraPeriod:
		return "IndexConfigVideoAVCIntraPeriod"
	case IndexConfigVideoNalSize:
		return "IndexConfigVideoNalSize"

		/* Image & Video common Configurations */
	case IndexCommonStartUnused:
		return "IndexCommonStartUnused"
	case IndexParamCommonDeblocking:
		return "IndexParamCommonDeblocking"
	case IndexParamCommonSensorMode:
		return "IndexParamCommonSensorMode"
	case IndexParamCommonInterleave:
		return "IndexParamCommonInterleave"
	case IndexConfigCommonColorFormatConversion:
		return "IndexConfigCommonColorFormatConversion"
	case IndexConfigCommonScale:
		return "IndexConfigCommonScale"
	case IndexConfigCommonImageFilter:
		return "IndexConfigCommonImageFilter"
	case IndexConfigCommonColorEnhancement:
		return "IndexConfigCommonColorEnhancement"
	case IndexConfigCommonColorKey:
		return "IndexConfigCommonColorKey"
	case IndexConfigCommonColorBlend:
		return "IndexConfigCommonColorBlend"
	case IndexConfigCommonFrameStabilisation:
		return "IndexConfigCommonFrameStabilisation"
	case IndexConfigCommonRotate:
		return "IndexConfigCommonRotate"
	case IndexConfigCommonMirror:
		return "IndexConfigCommonMirror"
	case IndexConfigCommonOutputPosition:
		return "IndexConfigCommonOutputPosition"
	case IndexConfigCommonInputCrop:
		return "IndexConfigCommonInputCrop"
	case IndexConfigCommonOutputCrop:
		return "IndexConfigCommonOutputCrop"
	case IndexConfigCommonDigitalZoom:
		return "IndexConfigCommonDigitalZoom"
	case IndexConfigCommonOpticalZoom:
		return "IndexConfigCommonOpticalZoom"
	case IndexConfigCommonWhiteBalance:
		return "IndexConfigCommonWhiteBalance"
	case IndexConfigCommonExposure:
		return "IndexConfigCommonExposure"
	case IndexConfigCommonContrast:
		return "IndexConfigCommonContrast"
	case IndexConfigCommonBrightness:
		return "IndexConfigCommonBrightness"
	case IndexConfigCommonBacklight:
		return "IndexConfigCommonBacklight"
	case IndexConfigCommonGamma:
		return "IndexConfigCommonGamma"
	case IndexConfigCommonSaturation:
		return "IndexConfigCommonSaturation"
	case IndexConfigCommonLightness:
		return "IndexConfigCommonLightness"
	case IndexConfigCommonExclusionRect:
		return "IndexConfigCommonExclusionRect"
	case IndexConfigCommonDithering:
		return "IndexConfigCommonDithering"
	case IndexConfigCommonPlaneBlend:
		return "IndexConfigCommonPlaneBlend"
	case IndexConfigCommonExposureValue:
		return "IndexConfigCommonExposureValue"
	case IndexConfigCommonOutputSize:
		return "IndexConfigCommonOutputSize"
	case IndexParamCommonExtraQuantData:
		return "IndexParamCommonExtraQuantData"
	case IndexConfigCommonFocusRegion:
		return "IndexConfigCommonFocusRegion"
	case IndexConfigCommonFocusStatus:
		return "IndexConfigCommonFocusStatus"
	case IndexConfigCommonTransitionEffect:
		return "IndexConfigCommonTransitionEffect"

		/* Reserved Configuration range */
	case IndexOtherStartUnused:
		return "IndexOtherStartUnused"
	case IndexParamOtherPortFormat:
		return "IndexParamOtherPortFormat"
	case IndexConfigOtherPower:
		return "IndexConfigOtherPower"
	case IndexConfigOtherStats:
		return "IndexConfigOtherStats"

		/* Reserved Time range */
	case IndexTimeStartUnused:
		return "IndexTimeStartUnused"
	case IndexConfigTimeScale:
		return "IndexConfigTimeScale"
	case IndexConfigTimeClockState:
		return "IndexConfigTimeClockState"
	case IndexConfigTimeActiveRefClock:
		return "IndexConfigTimeActiveRefClock"
	case IndexConfigTimeCurrentMediaTime:
		return "IndexConfigTimeCurrentMediaTime"
	case IndexConfigTimeCurrentWallTime:
		return "IndexConfigTimeCurrentWallTime"
	case IndexConfigTimeCurrentAudioReference:
		return "IndexConfigTimeCurrentAudioReference"
	case IndexConfigTimeCurrentVideoReference:
		return "IndexConfigTimeCurrentVideoReference"
	case IndexConfigTimeMediaTimeRequest:
		return "IndexConfigTimeMediaTimeRequest"
	case IndexConfigTimeClientStartTime:
		return "IndexConfigTimeClientStartTime"
	case IndexConfigTimePosition:
		return "IndexConfigTimePosition"
	case IndexConfigTimeSeekMode:
		return "IndexConfigTimeSeekMode"

	case IndexKhronosExtensions:
		return "IndexKhronosExtensions"
		/* Vendor specific area */
	case IndexVendorStartUnused:
		return "IndexVendorStartUnused"
		/* Vendor specific structures should be in the range of 0x7F000000
		   to 0x7FFFFFFE.  This range is not broken out by vendor, so
		   private indexes are not guaranteed unique and therefore should
		   only be sent to the appropriate component. */

		/* used for ilcs-top communication */
	case IndexParamMarkComparison:
		return "IndexParamMarkComparison"
	case IndexParamPortSummary:
		return "IndexParamPortSummary"
	case IndexParamTunnelStatus:
		return "IndexParamTunnelStatus"
	case IndexParamBrcmRecursionUnsafe:
		return "IndexParamBrcmRecursionUnsafe"

		/* used for top-ril communication */
	case IndexParamBufferAddress:
		return "IndexParamBufferAddress"
	case IndexParamTunnelSetup:
		return "IndexParamTunnelSetup"
	case IndexParamBrcmPortEGL:
		return "IndexParamBrcmPortEGL"
	case IndexParamIdleResourceCount:
		return "IndexParamIdleResourceCount"

		/* used for ril-ril communication */
	case IndexParamImagePoolDisplayFunction:
		return "IndexParamImagePoolDisplayFunction"
	case IndexParamBrcmDataUnit:
		return "IndexParamBrcmDataUnit"
	case IndexParamCodecConfig:
		return "IndexParamCodecConfig"
	case IndexParamCameraPoolToEncoderFunction:
		return "IndexParamCameraPoolToEncoderFunction"
	case IndexParamCameraStripeFunction:
		return "IndexParamCameraStripeFunction"
	case IndexParamCameraCaptureEventFunction:
		return "IndexParamCameraCaptureEventFunction"

		/* used for client-ril communication */
	case IndexParamTestInterface:
		return "IndexParamTestInterface"

		// 0x7f000010
	case IndexConfigDisplayRegion:
		return "IndexConfigDisplayRegion"
	case IndexParamSource:
		return "IndexParamSource"
	case IndexParamSourceSeed:
		return "IndexParamSourceSeed"
	case IndexParamResize:
		return "IndexParamResize"
	case IndexConfigVisualisation:
		return "IndexConfigVisualisation"
	case IndexConfigSingleStep:
		return "IndexConfigSingleStep"
	case IndexConfigPlayMode:
		return "IndexConfigPlayMode"
	case IndexParamCameraCamplusId:
		return "IndexParamCameraCamplusId"
	case IndexConfigCommonImageFilterParameters:
		return "IndexConfigCommonImageFilterParameters"
	case IndexConfigTransitionControl:
		return "IndexConfigTransitionControl"
	case IndexConfigPresentationOffset:
		return "IndexConfigPresentationOffset"
	case IndexParamSourceFunctions:
		return "IndexParamSourceFunctions"
	case IndexConfigAudioMonoTrackControl:
		return "IndexConfigAudioMonoTrackControl"
	case IndexParamCameraImagePool:
		return "IndexParamCameraImagePool"
	case IndexConfigCameraISPOutputPoolHeight:
		return "IndexConfigCameraISPOutputPoolHeight"
	case IndexParamImagePoolSize:
		return "IndexParamImagePoolSize"

		// 0x7f000020
	case IndexParamImagePoolExternal:
		return "IndexParamImagePoolExternal"
	case IndexParamRUTILFifoInfo:
		return "IndexParamRUTILFifoInfo"
	case IndexParamILFifoConfig:
		return "IndexParamILFifoConfig"
	case IndexConfigCameraSensorModes:
		return "IndexConfigCameraSensorModes"
	case IndexConfigBrcmPortStats:
		return "IndexConfigBrcmPortStats"
	case IndexConfigBrcmPortBufferStats:
		return "IndexConfigBrcmPortBufferStats"
	case IndexConfigBrcmCameraStats:
		return "IndexConfigBrcmCameraStats"
	case IndexConfigBrcmIOPerfStats:
		return "IndexConfigBrcmIOPerfStats"
	case IndexConfigCommonSharpness:
		return "IndexConfigCommonSharpness"
	case IndexConfigCommonFlickerCancellation:
		return "IndexConfigCommonFlickerCancellation"
	case IndexParamCameraSwapImagePools:
		return "IndexParamCameraSwapImagePools"
	case IndexParamCameraSingleBufferCaptureInput:
		return "IndexParamCameraSingleBufferCaptureInput"
	case IndexConfigCommonRedEyeRemoval:
		return "IndexConfigCommonRedEyeRemoval"
	case IndexConfigCommonFaceDetectionControl:
		return "IndexConfigCommonFaceDetectionControl"
	case IndexConfigCommonFaceDetectionRegion:
		return "IndexConfigCommonFaceDetectionRegion"
	case IndexConfigCommonInterlace:
		return "IndexConfigCommonInterlace"

		// 0x7f000030
	case IndexParamISPTunerName:
		return "IndexParamISPTunerName"
	case IndexParamCameraDeviceNumber:
		return "IndexParamCameraDeviceNumber"
	case IndexParamCameraDevicesPresent:
		return "IndexParamCameraDevicesPresent"
	case IndexConfigCameraInputFrame:
		return "IndexConfigCameraInputFrame"
	case IndexConfigStillColourDenoiseEnable:
		return "IndexConfigStillColourDenoiseEnable"
	case IndexConfigVideoColourDenoiseEnable:
		return "IndexConfigVideoColourDenoiseEnable"
	case IndexConfigAFAssistLight:
		return "IndexConfigAFAssistLight"
	case IndexConfigSmartShakeReductionEnable:
		return "IndexConfigSmartShakeReductionEnable"
	case IndexConfigInputCropPercentages:
		return "IndexConfigInputCropPercentages"
	case IndexConfigStillsAntiShakeEnable:
		return "IndexConfigStillsAntiShakeEnable"
	case IndexConfigWaitForFocusBeforeCapture:
		return "IndexConfigWaitForFocusBeforeCapture"
	case IndexConfigAudioRenderingLatency:
		return "IndexConfigAudioRenderingLatency"
	case IndexConfigDrawBoxAroundFaces:
		return "IndexConfigDrawBoxAroundFaces"
	case IndexParamCodecRequirements:
		return "IndexParamCodecRequirements"
	case IndexConfigBrcmEGLImageMemHandle:
		return "IndexConfigBrcmEGLImageMemHandle"
	case IndexConfigPrivacyIndicator:
		return "IndexConfigPrivacyIndicator"

		// 0x7f000040
	case IndexParamCameraFlashType:
		return "IndexParamCameraFlashType"
	case IndexConfigCameraEnableStatsPass:
		return "IndexConfigCameraEnableStatsPass"
	case IndexConfigCameraFlashConfig:
		return "IndexConfigCameraFlashConfig"
	case IndexConfigCaptureRawImageURI:
		return "IndexConfigCaptureRawImageURI"
	case IndexConfigCameraStripeFuncMinLines:
		return "IndexConfigCameraStripeFuncMinLines"
	case IndexConfigCameraAlgorithmVersionDeprecated:
		return "IndexConfigCameraAlgorithmVersionDeprecated"
	case IndexConfigCameraIsoReferenceValue:
		return "IndexConfigCameraIsoReferenceValue"
	case IndexConfigCameraCaptureAbortsAutoFocus:
		return "IndexConfigCameraCaptureAbortsAutoFocus"
	case IndexConfigBrcmClockMissCount:
		return "IndexConfigBrcmClockMissCount"
	case IndexConfigFlashChargeLevel:
		return "IndexConfigFlashChargeLevel"
	case IndexConfigBrcmVideoEncodedSliceSize:
		return "IndexConfigBrcmVideoEncodedSliceSize"
	case IndexConfigBrcmAudioTrackGaplessPlayback:
		return "IndexConfigBrcmAudioTrackGaplessPlayback"
	case IndexConfigBrcmAudioTrackChangeControl:
		return "IndexConfigBrcmAudioTrackChangeControl"
	case IndexParamBrcmPixelAspectRatio:
		return "IndexParamBrcmPixelAspectRatio"
	case IndexParamBrcmPixelValueRange:
		return "IndexParamBrcmPixelValueRange"
	case IndexParamCameraDisableAlgorithm:
		return "IndexParamCameraDisableAlgorithm"

		// 0x7f000050
	case IndexConfigBrcmVideoIntraPeriodTime:
		return "IndexConfigBrcmVideoIntraPeriodTime"
	case IndexConfigBrcmVideoIntraPeriod:
		return "IndexConfigBrcmVideoIntraPeriod"
	case IndexConfigBrcmAudioEffectControl:
		return "IndexConfigBrcmAudioEffectControl"
	case IndexConfigBrcmMinimumProcessingLatency:
		return "IndexConfigBrcmMinimumProcessingLatency"
	case IndexParamBrcmVideoAVCSEIEnable:
		return "IndexParamBrcmVideoAVCSEIEnable"
	case IndexParamBrcmAllowMemChange:
		return "IndexParamBrcmAllowMemChange"
	case IndexConfigBrcmVideoEncoderMBRowsPerSlice:
		return "IndexConfigBrcmVideoEncoderMBRowsPerSlice"
	case IndexParamCameraAFAssistDeviceNumber_Deprecated:
		return "IndexParamCameraAFAssistDeviceNumber_Deprecated"
	case IndexParamCameraPrivacyIndicatorDeviceNumber_Deprecated:
		return "IndexParamCameraPrivacyIndicatorDeviceNumber_Deprecated"
	case IndexConfigCameraUseCase:
		return "IndexConfigCameraUseCase"
	case IndexParamBrcmDisableProprietaryTunnels:
		return "IndexParamBrcmDisableProprietaryTunnels"
	case IndexParamBrcmOutputBufferSize:
		return "IndexParamBrcmOutputBufferSize"
	case IndexParamBrcmRetainMemory:
		return "IndexParamBrcmRetainMemory"
	case IndexConfigCanFocus_Deprecated:
		return "IndexConfigCanFocus_Deprecated"
	case IndexParamBrcmImmutableInput:
		return "IndexParamBrcmImmutableInput"
	case IndexParamDynamicParameterFile:
		return "IndexParamDynamicParameterFile"

		// 0x7f000060
	case IndexParamUseDynamicParameterFile:
		return "IndexParamUseDynamicParameterFile"
	case IndexConfigCameraInfo:
		return "IndexConfigCameraInfo"
	case IndexConfigCameraFeatures:
		return "IndexConfigCameraFeatures"
	case IndexConfigRequestCallback:
		return "IndexConfigRequestCallback"
	case IndexConfigBrcmOutputBufferFullCount:
		return "IndexConfigBrcmOutputBufferFullCount"
	case IndexConfigCommonFocusRegionXY:
		return "IndexConfigCommonFocusRegionXY"
	case IndexParamBrcmDisableEXIF:
		return "IndexParamBrcmDisableEXIF"
	case IndexConfigUserSettingsId:
		return "IndexConfigUserSettingsId"
	case IndexConfigCameraSettings:
		return "IndexConfigCameraSettings"
	case IndexConfigDrawBoxLineParams:
		return "IndexConfigDrawBoxLineParams"
	case IndexParamCameraRmiControl_Deprecated:
		return "IndexParamCameraRmiControl_Deprecated"
	case IndexConfigBurstCapture:
		return "IndexConfigBurstCapture"
	case IndexParamBrcmEnableIJGTableScaling:
		return "IndexParamBrcmEnableIJGTableScaling"
	case IndexConfigPowerDown:
		return "IndexConfigPowerDown"
	case IndexConfigBrcmSyncOutput:
		return "IndexConfigBrcmSyncOutput"
	case IndexParamBrcmFlushCallback:
		return "IndexParamBrcmFlushCallback"

		// 0x7f000070
	case IndexConfigBrcmVideoRequestIFrame:
		return "IndexConfigBrcmVideoRequestIFrame"
	case IndexParamBrcmNALSSeparate:
		return "IndexParamBrcmNALSSeparate"
	case IndexConfigConfirmView:
		return "IndexConfigConfirmView"
	case IndexConfigDrmView:
		return "IndexConfigDrmView"
	case IndexConfigBrcmVideoIntraRefresh:
		return "IndexConfigBrcmVideoIntraRefresh"
	case IndexParamBrcmMaxFileSize:
		return "IndexParamBrcmMaxFileSize"
	case IndexParamBrcmCRCEnable:
		return "IndexParamBrcmCRCEnable"
	case IndexParamBrcmCRC:
		return "IndexParamBrcmCRC"
	case IndexConfigCameraRmiInUse_Deprecated:
		return "IndexConfigCameraRmiInUse_Deprecated"
	case IndexConfigBrcmAudioSource:
		return "IndexConfigBrcmAudioSource"
	case IndexConfigBrcmAudioDestination:
		return "IndexConfigBrcmAudioDestination"
	case IndexParamAudioDdp:
		return "IndexParamAudioDdp"
	case IndexParamBrcmThumbnail:
		return "IndexParamBrcmThumbnail"
	case IndexParamBrcmDisableLegacyBlocks_Deprecated:
		return "IndexParamBrcmDisableLegacyBlocks_Deprecated"
	case IndexParamBrcmCameraInputAspectRatio:
		return "IndexParamBrcmCameraInputAspectRatio"
	case IndexParamDynamicParameterFileFailFatal:
		return "IndexParamDynamicParameterFileFailFatal"

		// 0x7f000080
	case IndexParamBrcmVideoDecodeErrorConcealment:
		return "IndexParamBrcmVideoDecodeErrorConcealment"
	case IndexParamBrcmInterpolateMissingTimestamps:
		return "IndexParamBrcmInterpolateMissingTimestamps"
	case IndexParamBrcmSetCodecPerformanceMonitoring:
		return "IndexParamBrcmSetCodecPerformanceMonitoring"
	case IndexConfigFlashInfo:
		return "IndexConfigFlashInfo"
	case IndexParamBrcmMaxFrameSkips:
		return "IndexParamBrcmMaxFrameSkips"
	case IndexConfigDynamicRangeExpansion:
		return "IndexConfigDynamicRangeExpansion"
	case IndexParamBrcmFlushCallbackId:
		return "IndexParamBrcmFlushCallbackId"
	case IndexParamBrcmTransposeBufferCount:
		return "IndexParamBrcmTransposeBufferCount"
	case IndexConfigFaceRecognitionControl:
		return "IndexConfigFaceRecognitionControl"
	case IndexConfigFaceRecognitionSaveFace:
		return "IndexConfigFaceRecognitionSaveFace"
	case IndexConfigFaceRecognitionDatabaseUri:
		return "IndexConfigFaceRecognitionDatabaseUri"
	case IndexConfigClockAdjustment:
		return "IndexConfigClockAdjustment"
	case IndexParamBrcmThreadAffinity:
		return "IndexParamBrcmThreadAffinity"
	case IndexParamAsynchronousOutput:
		return "IndexParamAsynchronousOutput"
	case IndexConfigAsynchronousFailureURI:
		return "IndexConfigAsynchronousFailureURI"
	case IndexConfigCommonFaceBeautification:
		return "IndexConfigCommonFaceBeautification"

		// 0x7f000090
	case IndexConfigCommonSceneDetectionControl:
		return "IndexConfigCommonSceneDetectionControl"
	case IndexConfigCommonSceneDetected:
		return "IndexConfigCommonSceneDetected"
	case IndexParamDisableVllPool:
		return "IndexParamDisableVllPool"
	case IndexParamVideoMvc:
		return "IndexParamVideoMvc"
	case IndexConfigBrcmDrawStaticBox:
		return "IndexConfigBrcmDrawStaticBox"
	case IndexConfigBrcmClockReferenceSource:
		return "IndexConfigBrcmClockReferenceSource"
	case IndexParamPassBufferMarks:
		return "IndexParamPassBufferMarks"
	case IndexConfigPortCapturing:
		return "IndexConfigPortCapturing"
	case IndexConfigBrcmDecoderPassThrough:
		return "IndexConfigBrcmDecoderPassThrough"
	// case IndexParamBrcmDecoderPassThrough:
	// 	return "IndexParamBrcmDecoderPassThrough"
	case IndexParamBrcmMaxCorruptMBs:
		return "IndexParamBrcmMaxCorruptMBs"
	case IndexConfigBrcmGlobalAudioMute:
		return "IndexConfigBrcmGlobalAudioMute"
	case IndexParamCameraCaptureMode:
		return "IndexParamCameraCaptureMode"
	case IndexParamBrcmDrmEncryption:
		return "IndexParamBrcmDrmEncryption"
	case IndexConfigBrcmCameraRnDPreprocess:
		return "IndexConfigBrcmCameraRnDPreprocess"
	case IndexConfigBrcmCameraRnDPostprocess:
		return "IndexConfigBrcmCameraRnDPostprocess"
	case IndexConfigBrcmAudioTrackChangeCount:
		return "IndexConfigBrcmAudioTrackChangeCount"

		// 0x7f0000a0
	case IndexParamCommonUseStcTimestamps:
		return "IndexParamCommonUseStcTimestamps"
	case IndexConfigBufferStall:
		return "IndexConfigBufferStall"
	case IndexConfigRefreshCodec:
		return "IndexConfigRefreshCodec"
	case IndexParamCaptureStatus:
		return "IndexParamCaptureStatus"
	case IndexConfigTimeInvalidStartTime:
		return "IndexConfigTimeInvalidStartTime"
	case IndexConfigLatencyTarget:
		return "IndexConfigLatencyTarget"
	case IndexConfigMinimiseFragmentation:
		return "IndexConfigMinimiseFragmentation"
	case IndexConfigBrcmUseProprietaryCallback:
		return "IndexConfigBrcmUseProprietaryCallback"
	case IndexParamPortMaxFrameSize:
		return "IndexParamPortMaxFrameSize"
	case IndexParamComponentName:
		return "IndexParamComponentName"
	case IndexConfigEncLevelExtension:
		return "IndexConfigEncLevelExtension"
	case IndexConfigTemporalDenoiseEnable:
		return "IndexConfigTemporalDenoiseEnable"
	case IndexParamBrcmLazyImagePoolDestroy:
		return "IndexParamBrcmLazyImagePoolDestroy"
	case IndexParamBrcmEEDEEnable:
		return "IndexParamBrcmEEDEEnable"
	case IndexParamBrcmEEDELossRate:
		return "IndexParamBrcmEEDELossRate"
	case IndexParamAudioDts:
		return "IndexParamAudioDts"

		// 0x7f0000b0
	case IndexParamNumOutputChannels:
		return "IndexParamNumOutputChannels"
	case IndexConfigBrcmHighDynamicRange:
		return "IndexConfigBrcmHighDynamicRange"
	case IndexConfigBrcmPoolMemAllocSize:
		return "IndexConfigBrcmPoolMemAllocSize"
	case IndexConfigBrcmBufferFlagFilter:
		return "IndexConfigBrcmBufferFlagFilter"
	case IndexParamBrcmVideoEncodeMinQuant:
		return "IndexParamBrcmVideoEncodeMinQuant"
	case IndexParamBrcmVideoEncodeMaxQuant:
		return "IndexParamBrcmVideoEncodeMaxQuant"
	case IndexParamRateControlModel:
		return "IndexParamRateControlModel"
	case IndexParamBrcmExtraBuffers:
		return "IndexParamBrcmExtraBuffers"
	case IndexConfigFieldOfView:
		return "IndexConfigFieldOfView"
	case IndexParamBrcmAlignHoriz:
		return "IndexParamBrcmAlignHoriz"
	case IndexParamBrcmAlignVert:
		return "IndexParamBrcmAlignVert"
	case IndexParamColorSpace:
		return "IndexParamColorSpace"
	case IndexParamBrcmDroppablePFrames:
		return "IndexParamBrcmDroppablePFrames"
	case IndexParamBrcmVideoInitialQuant:
		return "IndexParamBrcmVideoInitialQuant"
	case IndexParamBrcmVideoEncodeQpP:
		return "IndexParamBrcmVideoEncodeQpP"
	case IndexParamBrcmVideoRCSliceDQuant:
		return "IndexParamBrcmVideoRCSliceDQuant"

		// 0x7f0000c0
	case IndexParamBrcmVideoFrameLimitBits:
		return "IndexParamBrcmVideoFrameLimitBits"
	case IndexParamBrcmVideoPeakRate:
		return "IndexParamBrcmVideoPeakRate"
	case IndexConfigBrcmVideoH264DisableCABAC:
		return "IndexConfigBrcmVideoH264DisableCABAC"
	case IndexConfigBrcmVideoH264LowLatency:
		return "IndexConfigBrcmVideoH264LowLatency"
	case IndexConfigBrcmVideoH264AUDelimiters:
		return "IndexConfigBrcmVideoH264AUDelimiters"
	case IndexConfigBrcmVideoH264DeblockIDC:
		return "IndexConfigBrcmVideoH264DeblockIDC"
	case IndexConfigBrcmVideoH264IntraMBMode:
		return "IndexConfigBrcmVideoH264IntraMBMode"
	case IndexConfigContrastEnhance:
		return "IndexConfigContrastEnhance"
	case IndexParamCameraCustomSensorConfig:
		return "IndexParamCameraCustomSensorConfig"
	case IndexParamBrcmHeaderOnOpen:
		return "IndexParamBrcmHeaderOnOpen"
	case IndexConfigBrcmUseRegisterFile:
		return "IndexConfigBrcmUseRegisterFile"
	case IndexConfigBrcmRegisterFileFailFatal:
		return "IndexConfigBrcmRegisterFileFailFatal"
	case IndexParamBrcmConfigFileRegisters:
		return "IndexParamBrcmConfigFileRegisters"
	case IndexParamBrcmConfigFileChunkRegisters:
		return "IndexParamBrcmConfigFileChunkRegisters"
	case IndexParamBrcmAttachLog:
		return "IndexParamBrcmAttachLog"
	case IndexParamCameraZeroShutterLag:
		return "IndexParamCameraZeroShutterLag"

		// 0x7f0000d0
	case IndexParamBrcmFpsRange:
		return "IndexParamBrcmFpsRange"
	case IndexParamCaptureExposureCompensation:
		return "IndexParamCaptureExposureCompensation"
	case IndexParamBrcmVideoPrecodeForQP:
		return "IndexParamBrcmVideoPrecodeForQP"
	case IndexParamBrcmVideoTimestampFifo:
		return "IndexParamBrcmVideoTimestampFifo"
	case IndexParamSWSharpenDisable:
		return "IndexParamSWSharpenDisable"
	case IndexConfigBrcmFlashRequired:
		return "IndexConfigBrcmFlashRequired"
	case IndexParamBrcmVideoDrmProtectBuffer:
		return "IndexParamBrcmVideoDrmProtectBuffer"
	case IndexParamSWSaturationDisable:
		return "IndexParamSWSaturationDisable"
	case IndexParamBrcmVideoDecodeConfigVD3:
		return "IndexParamBrcmVideoDecodeConfigVD3"
	case IndexConfigBrcmPowerMonitor:
		return "IndexConfigBrcmPowerMonitor"
	case IndexParamBrcmZeroCopy:
		return "IndexParamBrcmZeroCopy"
	case IndexParamBrcmVideoEGLRenderDiscardMode:
		return "IndexParamBrcmVideoEGLRenderDiscardMode"
	case IndexParamBrcmVideoAVC_VCLHRDEnable:
		return "IndexParamBrcmVideoAVC_VCLHRDEnable"
	case IndexParamBrcmVideoAVC_LowDelayHRDEnable:
		return "IndexParamBrcmVideoAVC_LowDelayHRDEnable"
	case IndexParamBrcmVideoCroppingDisable:
		return "IndexParamBrcmVideoCroppingDisable"
	case IndexParamBrcmVideoAVCInlineHeaderEnable:
		return "IndexParamBrcmVideoAVCInlineHeaderEnable"

		// 0x7f0000f0
	case IndexConfigBrcmAudioDownmixCoefficients:
		return "IndexConfigBrcmAudioDownmixCoefficients"
	case IndexConfigBrcmAudioDownmixCoefficients8x8:
		return "IndexConfigBrcmAudioDownmixCoefficients8x8"
	case IndexConfigBrcmAudioMaxSample:
		return "IndexConfigBrcmAudioMaxSample"
	case IndexConfigCustomAwbGains:
		return "IndexConfigCustomAwbGains"
	case IndexParamRemoveImagePadding:
		return "IndexParamRemoveImagePadding"
	case IndexParamBrcmVideoAVCInlineVectorsEnable:
		return "IndexParamBrcmVideoAVCInlineVectorsEnable"
	case IndexConfigBrcmRenderStats:
		return "IndexConfigBrcmRenderStats"
	case IndexConfigBrcmCameraAnnotate:
		return "IndexConfigBrcmCameraAnnotate"
	case IndexParamBrcmStereoscopicMode:
		return "IndexParamBrcmStereoscopicMode"
	case IndexParamBrcmLockStepEnable:
		return "IndexParamBrcmLockStepEnable"
	case IndexParamBrcmTimeScale:
		return "IndexParamBrcmTimeScale"
	case IndexParamCameraInterface:
		return "IndexParamCameraInterface"
	case IndexParamCameraClockingMode:
		return "IndexParamCameraClockingMode"
	case IndexParamCameraRxConfig:
		return "IndexParamCameraRxConfig"
	case IndexParamCameraRxTiming:
		return "IndexParamCameraRxTiming"
	case IndexParamDynamicParameterConfig:
		return "IndexParamDynamicParameterConfig"

		// 0x7f000100
	case IndexParamBrcmVideoAVCSPSTimingEnable:
		return "IndexParamBrcmVideoAVCSPSTimingEnable"
	case IndexParamBrcmBayerOrder:
		return "IndexParamBrcmBayerOrder"
	case IndexParamBrcmMaxNumCallbacks:
		return "IndexParamBrcmMaxNumCallbacks"
	case IndexParamBrcmJpegRestartInterval:
		return "IndexParamBrcmJpegRestartInterval"
	case IndexParamBrcmSupportsSlices:
		return "IndexParamBrcmSupportsSlices"
	// case IndexParamBrcmIspBlockOverride:
	// 	return "IndexParamBrcmIspBlockOverride"
	// case IndexParamBrcmSupportsUnalignedSliceheight:
	// 	return "IndexParamBrcmSupportsUnalignedSliceheight"
	// case IndexParamBrcmLensShadingOverride:
	// 	return "IndexParamBrcmLensShadingOverride"
	// case IndexParamBrcmBlackLevel:
	// 	return "IndexParamBrcmBlackLevel"
	// case IndexParamOutputShift:
	// 	return "IndexParamOutputShift"
	// case IndexParamCcmShift:
	// 	return "IndexParamCcmShift"
	// case IndexParamCustomCcm:
	// 	return "IndexParamCustomCcm"
	// case IndexConfigCameraAnalogGain:
	// 	return "IndexConfigCameraAnalogGain"
	// case IndexConfigCameraDigitalGain:
	// 	return "IndexConfigCameraDigitalGain"
	// case IndexConfigBrcmDroppableRunLength:
	// 	return "IndexConfigBrcmDroppableRunLength"
	// case IndexParamMinimumAlignment:
	// 	return "IndexParamMinimumAlignment"
	case IndexMax:
		return "IndexMax"
	}
	return fmt.Sprintf("UNKNOWN[%d]", int(v))
}
