// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo LDFLAGS: -lhighgui

#include <opencv/highgui.h>
#include <stdlib.h>
#include <string.h>

//-----------------------------------------------------------------------------
// cvCreateTrackbar2 wrap
struct TrackbarUserdata {
	schar* win_name;
	schar* bar_name;
};
static void trackbarCallback(int pos, void* userdata)
{
	extern void goTrackbarCallback(schar* pbarName0, schar* winName1, int pos);

	struct TrackbarUserdata *arg = (struct TrackbarUserdata*)userdata;
	goTrackbarCallback(arg->bar_name, arg->win_name, pos);
}
static int CreateTrackbar(
	const char* trackbar_name, const char* window_name,
	int value, int count
)
{
	struct TrackbarUserdata *userdata = malloc(sizeof(*userdata));
	userdata->win_name = (schar*)window_name;
	userdata->bar_name = (schar*)trackbar_name;

	return cvCreateTrackbar2(trackbar_name, window_name,
		&value, count,
		trackbarCallback, userdata
	);
}

//-----------------------------------------------------------------------------
// cvSetMouseCallback wrap
static void mouseCallback(int event, int x, int y, int flags, void* param)
{
	extern void goMouseCallback(schar* name, int x, int y, int flags);

	schar* name = (schar*)param;
	goMouseCallback(name, x, y, flags);
}
static void SetMouseCallback(const char* window_name)
{
	cvSetMouseCallback(window_name, mouseCallback, (void*)window_name);
}

//-----------------------------------------------------------------------------

// video writer args
unsigned int CV_FOURCC_(int c1, int c2, int c3, int c4) {
	return CV_FOURCC(c1,c2,c3,c4);
}

//-----------------------------------------------------------------------------
*/
import "C"
import (
//	"errors"
	"unsafe"
)

func init() {
	//
}

/*****************************************************************************\
*                     window/track wrap for go                                *
\*****************************************************************************/

// mouse callback
type MouseCallbackFunc func(event, x, y, flags int, param interface{})
// trackbar callback
type TrackbarCallbackFunc func(pos int)

type Window struct {
	name           string
	mouseHandle    MouseCallbackFunc
	param          interface{}
	trackbarHandle map[string]TrackbarCallbackFunc
	refCount       int
}
var allWindows = make(map[string]*Window, 1000)

/*****************************************************************************\
*                           Basic GUI functions                               *
/*****************************************************************************/

/* this function is used to set some external parameters in case of X Window */
func InitSystem(args []string) int {
	argc := C.int(len(args))
	argv := make([]*C.char, len(args))
	for i := 0; i < len(args); i++ {
		argv[i] = C.CString(args[i])
	}
	rv := C.cvInitSystem(argc, (**C.char)(unsafe.Pointer(&argv)))
	return int(rv)
}

func StartWindowThread() int {
	return int(C.cvStartWindowThread())
}

const WINDOW_AUTOSIZE = int(C.CV_WINDOW_AUTOSIZE)

/* create window */
func NamedWindow(name string, flags int) int {
	return int(C.cvNamedWindow(C.CString(name), C.int(flags)))
}

// ---------  YV ---------
const (
	WND_PROP_FULLSCREEN = C.CV_WND_PROP_FULLSCREEN
	WND_PROP_AUTOSIZE   = C.CV_WND_PROP_AUTOSIZE
	WINDOW_NORMAL       = C.CV_WINDOW_NORMAL
	WINDOW_FULLSCREEN   = C.CV_WINDOW_FULLSCREEN
)
/* Set and Get Property of the window */
func SetWindowProperty(name string, prop_id int, value float64) {
	C.cvSetWindowProperty(C.CString(name), C.int(prop_id), C.double(value))
}
func cvGetWindowProperty(name string, prop_id int) float64 {
	return float64(C.cvGetWindowProperty(C.CString(name), C.int(prop_id)))
}

/* display image within window (highgui windows remember their content) */
func ShowImage(name string, image unsafe.Pointer) {
	C.cvShowImage(C.CString(name), image)
}

/* resize/move window */
func ResizeWindow(name string, width, height int) {
	C.cvResizeWindow(C.CString(name), C.int(width), C.int(height))
}
func MoveWindow(name string, x, y int) {
	C.cvMoveWindow(C.CString(name), C.int(x), C.int(y))
}

/* destroy window and all the trackers associated with it */
func DestroyWindow(name string) {
	C.cvDestroyWindow(C.CString(name))
}
func DestroyAllWindows(name string) {
	C.cvDestroyAllWindows()
}

/* get native window handle (HWND in case of Win32 and Widget in case of X Window) */
func GetWindowHandle(name string) unsafe.Pointer {
	return C.cvGetWindowHandle(C.CString(name))
}

/* get name of highgui window given its native handle */
func GetWindowName(window_handle unsafe.Pointer) string {
	name := C.cvGetWindowName(window_handle)
	return C.GoString(name)
}

/* create trackbar and display it on top of given window, set callback */
func CreateTrackbar(bar_name, win_name string, value, count int) int {
	rv := C.CreateTrackbar(C.CString(bar_name), C.CString(win_name),
		C.int(value), C.int(count))
	return int(rv)
}
//export goTrackbarCallback
func goTrackbarCallback(barName_, winName_ *C.char, pos int) {
	winName := C.GoString(winName_)
	barName := C.GoString(barName_)
	if win, ok := allWindows[winName]; ok {
		if call, ok := win.trackbarHandle[barName]; ok {
			call(int(pos))
		}
	}
}

/* retrieve or set trackbar position */
func GetTrackbarPos(trackbar_name, window_name string) int {
	rv := C.cvGetTrackbarPos(C.CString(trackbar_name), C.CString(window_name))
	return int(rv)
}
func SetTrackbarPos(trackbar_name, window_name string, pos int) {
	C.cvSetTrackbarPos(C.CString(trackbar_name), C.CString(window_name), C.int(pos))
}

const (
	EVENT_MOUSEMOVE     = C.CV_EVENT_MOUSEMOVE
	EVENT_LBUTTONDOWN   = C.CV_EVENT_LBUTTONDOWN
	EVENT_RBUTTONDOWN   = C.CV_EVENT_RBUTTONDOWN
	EVENT_MBUTTONDOWN   = C.CV_EVENT_MBUTTONDOWN
	EVENT_LBUTTONUP     = C.CV_EVENT_LBUTTONUP
	EVENT_RBUTTONUP     = C.CV_EVENT_RBUTTONUP
	EVENT_MBUTTONUP     = C.CV_EVENT_MBUTTONUP
	EVENT_LBUTTONDBLCLK = C.CV_EVENT_LBUTTONDBLCLK
	EVENT_RBUTTONDBLCLK = C.CV_EVENT_RBUTTONDBLCLK
	EVENT_MBUTTONDBLCLK = C.CV_EVENT_MBUTTONDBLCLK

	EVENT_FLAG_LBUTTON  = C.CV_EVENT_FLAG_LBUTTON
	EVENT_FLAG_RBUTTON  = C.CV_EVENT_FLAG_RBUTTON
	EVENT_FLAG_MBUTTON  = C.CV_EVENT_FLAG_MBUTTON
	EVENT_FLAG_CTRLKEY  = C.CV_EVENT_FLAG_CTRLKEY
	EVENT_FLAG_SHIFTKEY = C.CV_EVENT_FLAG_SHIFTKEY
	EVENT_FLAG_ALTKEY   = C.CV_EVENT_FLAG_ALTKEY
)

/* assign callback for mouse events */
func SetMouseCallback(winName string) {
	C.SetMouseCallback(C.CString(winName))
}
//export goMouseCallback
func goMouseCallback(name *C.char, event, x, y, flags C.int) {
	winName := C.GoString(name)
	if win, ok := allWindows[winName]; ok {
		if win.mouseHandle != nil {
			win.mouseHandle(int(event), int(x), int(y), int(flags), win.param)
		}
	}
}

const (
	/* 8bit, color or not */
	LOAD_IMAGE_UNCHANGED = C.CV_LOAD_IMAGE_UNCHANGED
	/* 8bit, gray */
	LOAD_IMAGE_GRAYSCALE = C.CV_LOAD_IMAGE_GRAYSCALE
	/* ?, color */
	LOAD_IMAGE_COLOR     = C.CV_LOAD_IMAGE_COLOR
	/* any depth, ? */
	LOAD_IMAGE_ANYDEPTH  = C.CV_LOAD_IMAGE_ANYDEPTH
	/* ?, any color */
	LOAD_IMAGE_ANYCOLOR  = C.CV_LOAD_IMAGE_ANYCOLOR
)
/* load image from file
  iscolor can be a combination of above flags where CV_LOAD_IMAGE_UNCHANGED
  overrides the other flags
  using CV_LOAD_IMAGE_ANYCOLOR alone is equivalent to CV_LOAD_IMAGE_UNCHANGED
  unless CV_LOAD_IMAGE_ANYDEPTH is specified images are converted to 8bit
*/
func LoadImage(filename string, iscolor int) unsafe.Pointer {
	rv := C.cvLoadImage(C.CString(filename), C.int(iscolor))
	return unsafe.Pointer(rv)
}
func LoadImageM(filename string, iscolor int) unsafe.Pointer {
	rv := C.cvLoadImageM(C.CString(filename), C.int(iscolor))
	return unsafe.Pointer(rv)
}

const (
	IMWRITE_JPEG_QUALITY    = C.CV_IMWRITE_JPEG_QUALITY
	IMWRITE_PNG_COMPRESSION = C.CV_IMWRITE_PNG_COMPRESSION
	IMWRITE_PXM_BINARY      = C.CV_IMWRITE_PXM_BINARY
)

/* save image to file */
func SaveImage(filename string, image unsafe.Pointer, params int) int {
	params_c := C.int(params)
	rv := C.cvSaveImage(C.CString(filename), (image), &params_c)
	return int(rv)
}

/* decode image stored in the buffer */
func DecodeImage(buf unsafe.Pointer, iscolor int) unsafe.Pointer {
	rv := C.cvDecodeImage((*C.CvMat)(buf), C.int(iscolor))
	return unsafe.Pointer(rv)
}
func DecodeImageM(buf unsafe.Pointer, iscolor int) unsafe.Pointer {
	rv := C.cvDecodeImageM((*C.CvMat)(buf), C.int(iscolor))
	return unsafe.Pointer(rv)
}

/* encode image and store the result as a byte vector (single-row 8uC1 matrix) */
func EncodeImage(ext string, image unsafe.Pointer, params int) unsafe.Pointer {
	params_c := C.int(params)
	rv := C.cvEncodeImage(C.CString(ext), (image), &params_c)
	return unsafe.Pointer(rv)
}

const (
	CVTIMG_FLIP    = C.CV_CVTIMG_FLIP
	CVTIMG_SWAP_RB = C.CV_CVTIMG_SWAP_RB
)

/* utility function: convert one image to another with optional vertical flip */
func ConvertImage(src, dst unsafe.Pointer, flags int) {
	C.cvConvertImage(src, dst, C.int(flags))
}
/* wait for key event infinitely (delay<=0) or for "delay" milliseconds */
func WaitKey(delay int) {
	C.cvWaitKey(C.int(delay))
}

/*****************************************************************************\
*                  Working with Video Files and Cameras                       *
/*****************************************************************************/

/* start capturing frames from video file */
func CreateFileCapture(filename string) unsafe.Pointer {
	rv := C.cvCreateFileCapture(C.CString(filename))
	return unsafe.Pointer(rv)
}

const (
	CAP_ANY      = C.CV_CAP_ANY		// autodetect

	CAP_MIL      = C.CV_CAP_MIL		// MIL proprietary drivers

	CAP_VFW      = C.CV_CAP_VFW		// platform native
	CAP_V4L      = C.CV_CAP_V4L
	CAP_V4L2     = C.CV_CAP_V4L2

	CAP_FIREWARE = C.CV_CAP_FIREWARE	// IEEE 1394 drivers
	CAP_FIREWIRE = C.CV_CAP_FIREWIRE
	CAP_IEEE1394 = C.CV_CAP_IEEE1394
	CAP_DC1394   = C.CV_CAP_DC1394
	CAP_CMU1394  = C.CV_CAP_CMU1394

	CAP_STEREO  = C.CV_CAP_STEREO	// TYZX proprietary drivers
	CAP_TYZX    = C.CV_CAP_TYZX
	TYZX_LEFT   = C.CV_TYZX_LEFT
	TYZX_RIGHT  = C.CV_TYZX_RIGHT
	TYZX_COLOR  = C.CV_TYZX_COLOR
	TYZX_Z      = C.CV_TYZX_Z

	CAP_QT      = C.CV_CAP_QT		// QuickTime

	CAP_UNICAP  = C.CV_CAP_UNICAP	// Unicap drivers

	CAP_DSHOW   = C.CV_CAP_DSHOW	// DirectShow (via videoInput)

	CAP_PVAPI   = C.CV_CAP_PVAPI	// PvAPI, Prosilica GigE SDK
)

/* start capturing frames from camera: index = camera_index + domain_offset (CV_CAP_*) */
func CreateCameraCapture(index int) unsafe.Pointer {
	rv := C.cvCreateCameraCapture(C.int(index))
	return unsafe.Pointer(rv)
}

/* grab a frame, return 1 on success, 0 on fail.
  this function is thought to be fast               */
func GrabFrame(capture unsafe.Pointer) int {
	rv := C.cvGrabFrame((*C.CvCapture)(capture))
	return int(rv)
}

/* get the frame grabbed with cvGrabFrame(..)
  This function may apply some frame processing like
  frame decompression, flipping etc.
  !!!DO NOT RELEASE or MODIFY the retrieved frame!!! */
func RetrieveFrame(capture unsafe.Pointer, streamIdx int) unsafe.Pointer {
	rv := C.cvRetrieveFrame((*C.CvCapture)(capture), C.int(streamIdx))
	return unsafe.Pointer(rv)
}

/* Just a combination of cvGrabFrame and cvRetrieveFrame
   !!!DO NOT RELEASE or MODIFY the retrieved frame!!!      */
func QueryFrame(capture unsafe.Pointer) unsafe.Pointer {
	rv := C.cvQueryFrame((*C.CvCapture)(capture))
	return unsafe.Pointer(rv)
}

/* stop capturing/reading and free resources */
func ReleaseCapture(capture unsafe.Pointer) {
	cap_c := (*C.CvCapture)(capture)
	C.cvReleaseCapture(&cap_c)
}

const (
	CAP_PROP_POS_MSEC      = C.CV_CAP_PROP_POS_MSEC
	CAP_PROP_POS_FRAMES    = C.CV_CAP_PROP_POS_FRAMES
	CAP_PROP_POS_AVI_RATIO = C.CV_CAP_PROP_POS_AVI_RATIO
	CAP_PROP_FRAME_WIDTH   = C.CV_CAP_PROP_FRAME_WIDTH
	CAP_PROP_FRAME_HEIGHT  = C.CV_CAP_PROP_FRAME_HEIGHT
	CAP_PROP_FPS           = C.CV_CAP_PROP_FPS
	CAP_PROP_FOURCC        = C.CV_CAP_PROP_FOURCC
	CAP_PROP_FRAME_COUNT   = C.CV_CAP_PROP_FRAME_COUNT
	CAP_PROP_FORMAT        = C.CV_CAP_PROP_FORMAT
	CAP_PROP_MODE          = C.CV_CAP_PROP_MODE
	CAP_PROP_BRIGHTNESS    = C.CV_CAP_PROP_BRIGHTNESS
	CAP_PROP_CONTRAST      = C.CV_CAP_PROP_CONTRAST
	CAP_PROP_SATURATION    = C.CV_CAP_PROP_SATURATION
	CAP_PROP_HUE           = C.CV_CAP_PROP_HUE
	CAP_PROP_GAIN          = C.CV_CAP_PROP_GAIN
	CAP_PROP_EXPOSURE      = C.CV_CAP_PROP_EXPOSURE
	CAP_PROP_CONVERT_RGB   = C.CV_CAP_PROP_CONVERT_RGB
	CAP_PROP_WHITE_BALANCE = C.CV_CAP_PROP_WHITE_BALANCE
	CAP_PROP_RECTIFICATION = C.CV_CAP_PROP_RECTIFICATION
)

/* retrieve or set capture properties */
func GetCaptureProperty(capture unsafe.Pointer, property_id int) float64 {
	rv := C.cvGetCaptureProperty((*C.CvCapture)(capture), C.int(property_id))
	return float64(rv)
}
func SetCaptureProperty(capture unsafe.Pointer, property_id int, value float64) int {
	rv := C.cvSetCaptureProperty((*C.CvCapture)(capture), C.int(property_id), C.double(value))
	return int(rv)
}

// Return the type of the capturer (eg, CV_CAP_V4W, CV_CAP_UNICAP), which is unknown if created with CV_CAP_ANY
func GetCaptureDomain(capture unsafe.Pointer) int {
	rv := C.cvGetCaptureDomain((*C.CvCapture)(capture))
	return int(rv)
}

// Prototype for CV_FOURCC so that swig can generate wrapper without mixing up the define
func FOURCC(c1, c2, c3, c4 int8) int {
	rv := C.CV_FOURCC_(C.int(c1),C.int(c2),C.int(c3),C.int(c4))
	return int(rv)
}
const (
	FOURCC_PROMPT = C.CV_FOURCC_PROMPT	/* Open Codec Selection Dialog (Windows only) */
	//FOURCC_DEFAULT = FOURCC('I', 'Y', 'U', 'V')	/* Use default codec for specified filename (Linux only) */
)

/* initialize video file writer */
func CreateVideoWriter(filename string, fourcc int, fps float32,
	frame_width, frame_height, is_color int) unsafe.Pointer {
	size := C.cvSize(C.int(frame_width), C.int(frame_height))
	rv := C.cvCreateVideoWriter(C.CString(filename), C.int(fourcc), C.double(fps), size, C.int(is_color))
	return unsafe.Pointer(rv)
}

/* write frame to video file */
func WriteFrame(writer, image unsafe.Pointer) int {
	rv := C.cvWriteFrame((*C.CvVideoWriter)(writer), (*C.IplImage)(image))
	return int(rv)
}

/* close video file writer */
func ReleaseVideoWriter(writer unsafe.Pointer) {
	writer_c := (*C.CvVideoWriter)(writer)
	C.cvReleaseVideoWriter(&writer_c)
}

/*****************************************************************************\
*                                END                                          *
/*****************************************************************************/


