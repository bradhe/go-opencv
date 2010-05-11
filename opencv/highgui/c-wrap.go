
package highgui

/*
#include <stdio.h>
#include <opencv/highgui.h>

//=============================================================================
// opencv version

enum {
	CV_MAJOR_VERSION_		= CV_MAJOR_VERSION,
	CV_MINOR_VERSION_		= CV_MINOR_VERSION,
	CV_SUBMINOR_VERSION_	= CV_SUBMINOR_VERSION
};

// cvNamedWindow flags

enum {
	CV_WINDOW_AUTOSIZE_	= CV_WINDOW_AUTOSIZE
};

// CvMouseCallback flags

enum {
	CV_EVENT_MOUSEMOVE_		= CV_EVENT_MOUSEMOVE,
	CV_EVENT_LBUTTONDOWN_	= CV_EVENT_LBUTTONDOWN,
	CV_EVENT_RBUTTONDOWN_	= CV_EVENT_RBUTTONDOWN,
	CV_EVENT_MBUTTONDOWN_	= CV_EVENT_MBUTTONDOWN,
	CV_EVENT_LBUTTONUP_		= CV_EVENT_LBUTTONUP,
	CV_EVENT_RBUTTONUP_		= CV_EVENT_RBUTTONUP,
	CV_EVENT_MBUTTONUP_		= CV_EVENT_MBUTTONUP,
	CV_EVENT_LBUTTONDBLCLK_	= CV_EVENT_LBUTTONDBLCLK,
	CV_EVENT_RBUTTONDBLCLK_	= CV_EVENT_RBUTTONDBLCLK,
	CV_EVENT_MBUTTONDBLCLK_	= CV_EVENT_MBUTTONDBLCLK
};

enum {
	CV_EVENT_FLAG_LBUTTON_	= CV_EVENT_FLAG_LBUTTON,
	CV_EVENT_FLAG_RBUTTON_	= CV_EVENT_FLAG_RBUTTON,
	CV_EVENT_FLAG_MBUTTON_	= CV_EVENT_FLAG_MBUTTON,
	CV_EVENT_FLAG_CTRLKEY_	= CV_EVENT_FLAG_CTRLKEY,
	CV_EVENT_FLAG_SHIFTKEY_	= CV_EVENT_FLAG_SHIFTKEY,
	CV_EVENT_FLAG_ALTKEY_	= CV_EVENT_FLAG_ALTKEY
};

// cvLoadImage iscolor flags

enum {
	CV_LOAD_IMAGE_UNCHANGED_X	= (0-CV_LOAD_IMAGE_UNCHANGED)
};

enum {
	CV_LOAD_IMAGE_GRAYSCALE_	= CV_LOAD_IMAGE_GRAYSCALE,
	CV_LOAD_IMAGE_COLOR_		= CV_LOAD_IMAGE_COLOR,
	CV_LOAD_IMAGE_ANYDEPTH_		= CV_LOAD_IMAGE_ANYDEPTH,
	CV_LOAD_IMAGE_ANYCOLOR_		= CV_LOAD_IMAGE_ANYCOLOR
};

// cvConvertImage flags

enum {
	CV_CVTIMG_FLIP_		= CV_CVTIMG_FLIP,
	CV_CVTIMG_SWAP_RB_	= CV_CVTIMG_SWAP_RB
};

// cvCreateCameraCapture: index = camera_index + domain_offset (CV_CAP_*)

enum {
	CV_CAP_ANY_			= CV_CAP_ANY,

	CV_CAP_MIL_			= CV_CAP_MIL,

	CV_CAP_VFW_			= CV_CAP_VFW,
	CV_CAP_V4L_			= CV_CAP_V4L,
	CV_CAP_V4L2_		= CV_CAP_V4L2,

	CV_CAP_FIREWARE_	= CV_CAP_FIREWARE,
	CV_CAP_IEEE1394_	= CV_CAP_IEEE1394,
	CV_CAP_DC1394_		= CV_CAP_DC1394,
	CV_CAP_CMU1394_		= CV_CAP_CMU1394,

	CV_CAP_STEREO_		= CV_CAP_STEREO,
	CV_CAP_TYZX_		= CV_CAP_TYZX,
	CV_TYZX_LEFT_		= CV_TYZX_LEFT,
	CV_TYZX_RIGHT_		= CV_TYZX_RIGHT,
	CV_TYZX_COLOR_		= CV_TYZX_COLOR,
	CV_TYZX_Z_			= CV_TYZX_Z,

	CV_CAP_QT_			= CV_CAP_QT
};

// cvGetCaptureProperty: property_id

enum {
	CV_CAP_PROP_POS_MSEC_		= CV_CAP_PROP_POS_MSEC,
	CV_CAP_PROP_POS_FRAMES_		= CV_CAP_PROP_POS_FRAMES,
	CV_CAP_PROP_POS_AVI_RATIO_	= CV_CAP_PROP_POS_AVI_RATIO,
	CV_CAP_PROP_FRAME_WIDTH_	= CV_CAP_PROP_FRAME_WIDTH,
	CV_CAP_PROP_FRAME_HEIGHT_	= CV_CAP_PROP_FRAME_HEIGHT,
	CV_CAP_PROP_FPS_			= CV_CAP_PROP_FPS,
	CV_CAP_PROP_FOURCC_			= CV_CAP_PROP_FOURCC,
	CV_CAP_PROP_FRAME_COUNT_	= CV_CAP_PROP_FRAME_COUNT,
	CV_CAP_PROP_FORMAT_			= CV_CAP_PROP_FORMAT,
	CV_CAP_PROP_MODE_			= CV_CAP_PROP_MODE,
	CV_CAP_PROP_BRIGHTNESS_		= CV_CAP_PROP_BRIGHTNESS,
	CV_CAP_PROP_CONTRAST_		= CV_CAP_PROP_CONTRAST,
	CV_CAP_PROP_SATURATION_		= CV_CAP_PROP_SATURATION,
	CV_CAP_PROP_HUE_			= CV_CAP_PROP_HUE,
	CV_CAP_PROP_GAIN_			= CV_CAP_PROP_GAIN,
	CV_CAP_PROP_CONVERT_RGB_	= CV_CAP_PROP_CONVERT_RGB
};

// cvCreateVideoWriter: fourcc flags

unsigned int CV_FOURCC_(int c1, int c2, int c3, int c4) {
	return CV_FOURCC(c1,c2,c3,c4);
}

//=============================================================================
// 
static char*  argv_buf[128];
static char** argv = argv_buf;

static void argv_set(int i, char *v) {
	if(i < sizeof(argv_buf)/sizeof(argv_buf[0])) {
		argv[i] = v;
	}
}
//static char** toCharPP(void *p) {
//	return (char**)p;
//}

// define
static const int WINDOW_AUTOSIZE = CV_WINDOW_AUTOSIZE;

enum {
 GO_IPL_DEPTH_8U = IPL_DEPTH_8U
};
*/
import "C"
import "unsafe"
import "os"
import "image"

//=============================================================================
// opencv version

const (
	CV_MAJOR_VERSION	= C.CV_MAJOR_VERSION_
	CV_MINOR_VERSION	= C.CV_MINOR_VERSION_
	CV_SUBMINOR_VERSION	= C.CV_SUBMINOR_VERSION_
)

// cvNamedWindow flags

const (
	CV_WINDOW_AUTOSIZE	= C.CV_WINDOW_AUTOSIZE_
)

// CvMouseCallback flags

const (
	CV_EVENT_MOUSEMOVE		= C.CV_EVENT_MOUSEMOVE_
	CV_EVENT_LBUTTONDOWN	= C.CV_EVENT_LBUTTONDOWN_
	CV_EVENT_RBUTTONDOWN	= C.CV_EVENT_RBUTTONDOWN_
	CV_EVENT_MBUTTONDOWN	= C.CV_EVENT_MBUTTONDOWN_
	CV_EVENT_LBUTTONUP		= C.CV_EVENT_LBUTTONUP_
	CV_EVENT_RBUTTONUP		= C.CV_EVENT_RBUTTONUP_
	CV_EVENT_MBUTTONUP		= C.CV_EVENT_MBUTTONUP_
	CV_EVENT_LBUTTONDBLCLK	= C.CV_EVENT_LBUTTONDBLCLK_
	CV_EVENT_RBUTTONDBLCLK	= C.CV_EVENT_RBUTTONDBLCLK_
	CV_EVENT_MBUTTONDBLCLK	= C.CV_EVENT_MBUTTONDBLCLK_
)

const (
	CV_EVENT_FLAG_LBUTTON	= C.CV_EVENT_FLAG_LBUTTON_
	CV_EVENT_FLAG_RBUTTON	= C.CV_EVENT_FLAG_RBUTTON_
	CV_EVENT_FLAG_MBUTTON	= C.CV_EVENT_FLAG_MBUTTON_
	CV_EVENT_FLAG_CTRLKEY	= C.CV_EVENT_FLAG_CTRLKEY_
	CV_EVENT_FLAG_SHIFTKEY	= C.CV_EVENT_FLAG_SHIFTKEY_
	CV_EVENT_FLAG_ALTKEY	= C.CV_EVENT_FLAG_ALTKEY_
)

// cvLoadImage iscolor flags

const (
	CV_LOAD_IMAGE_UNCHANGED	= (0-C.CV_LOAD_IMAGE_UNCHANGED_X)

	CV_LOAD_IMAGE_GRAYSCALE	= C.CV_LOAD_IMAGE_GRAYSCALE_
	CV_LOAD_IMAGE_COLOR		= C.CV_LOAD_IMAGE_COLOR_
	CV_LOAD_IMAGE_ANYDEPTH	= C.CV_LOAD_IMAGE_ANYDEPTH_
	CV_LOAD_IMAGE_ANYCOLOR	= C.CV_LOAD_IMAGE_ANYCOLOR_
)

// cvConvertImage flags

const (
	CV_CVTIMG_FLIP		= C.CV_CVTIMG_FLIP_
	CV_CVTIMG_SWAP_RB	= C.CV_CVTIMG_SWAP_RB_
)

// cvCreateCameraCapture: index = camera_index + domain_offset (CV_CAP_*)

const (
	CV_CAP_ANY	= C.CV_CAP_ANY_

	CV_CAP_MIL	= C.CV_CAP_MIL_

	CV_CAP_VFW	= C.CV_CAP_VFW_
	CV_CAP_V4L	= C.CV_CAP_V4L_
	CV_CAP_V4L2	= C.CV_CAP_V4L2_

	CV_CAP_FIREWARE	= C.CV_CAP_FIREWARE_
	CV_CAP_IEEE1394	= C.CV_CAP_IEEE1394_
	CV_CAP_DC1394	= C.CV_CAP_DC1394_
	CV_CAP_CMU1394	= C.CV_CAP_CMU1394_

	CV_CAP_STEREO	= C.CV_CAP_STEREO_
	CV_CAP_TYZX		= C.CV_CAP_TYZX_
	CV_TYZX_LEFT	= C.CV_TYZX_LEFT_
	CV_TYZX_RIGHT	= C.CV_TYZX_RIGHT_
	CV_TYZX_COLOR	= C.CV_TYZX_COLOR_
	CV_TYZX_Z		= C.CV_TYZX_Z_

	CV_CAP_QT		= C.CV_CAP_QT_
)

// cvGetCaptureProperty: property_id

const (
	CV_CAP_PROP_POS_MSEC	= C.CV_CAP_PROP_POS_MSEC_
	CV_CAP_PROP_POS_FRAMES	= C.CV_CAP_PROP_POS_FRAMES_
	CV_CAP_PROP_POS_AVI_RATIO	= C.CV_CAP_PROP_POS_AVI_RATIO_
	CV_CAP_PROP_FRAME_WIDTH		= C.CV_CAP_PROP_FRAME_WIDTH_
	CV_CAP_PROP_FRAME_HEIGHT	= C.CV_CAP_PROP_FRAME_HEIGHT_
	CV_CAP_PROP_FPS				= C.CV_CAP_PROP_FPS_
	CV_CAP_PROP_FOURCC			= C.CV_CAP_PROP_FOURCC_
	CV_CAP_PROP_FRAME_COUNT		= C.CV_CAP_PROP_FRAME_COUNT_
	CV_CAP_PROP_FORMAT			= C.CV_CAP_PROP_FORMAT_
	CV_CAP_PROP_MODE			= C.CV_CAP_PROP_MODE_
	CV_CAP_PROP_BRIGHTNESS		= C.CV_CAP_PROP_BRIGHTNESS_
	CV_CAP_PROP_CONTRAST		= C.CV_CAP_PROP_CONTRAST_
	CV_CAP_PROP_SATURATION		= C.CV_CAP_PROP_SATURATION_
	CV_CAP_PROP_HUE				= C.CV_CAP_PROP_HUE_
	CV_CAP_PROP_GAIN			= C.CV_CAP_PROP_GAIN_
	CV_CAP_PROP_CONVERT_RGB		= C.CV_CAP_PROP_CONVERT_RGB_
)

//=============================================================================

// 初始化系统
func cvInitSystem(args []string) {
	for i := 0; i < len(args); i++ {
		C.argv_set(C.int(i), C.CString(args[i]))
	}
	C.cvInitSystem(C.int(len(args)), C.argv)
}
// 启动窗口线程
func cvStartWindowThread()(err os.Error) {
	C.cvStartWindowThread()
	return
}

// 命名窗口
func cvNamedWindow(name string, autoSize bool)(err os.Error) {
	if autoSize {
		C.cvNamedWindow(C.CString(name), C.WINDOW_AUTOSIZE)
	} else {
		C.cvNamedWindow(C.CString(name), C.int(0))
	}
 	return nil
}
// 设置窗口属性
func cvSetWindowProperty(name string, prop_id int, prop_value float)(err os.Error) {
	return nil
}
// 查询窗口属性
func cvGetWindowProperty(name string, prop_id int)(prop_value float, err os.Error) {
	return
}

// 显式图像
func cvShowImage(name string, image image.Image)(err os.Error) {
	// make a copy
	size := C.cvSize(C.int(image.Width()), C.int(image.Height()))
	pIplImage := C.cvCreateImage(size, C.GO_IPL_DEPTH_8U, 3)

	for i := 0; i < image.Height(); i++ {
		for j := 0; j < image.Width(); j++ {
			color := image.At(i, j)
			r, g, b, _ := color.RGBA()
			rgb := C.cvScalar(C.double(b), C.double(g), C.double(r), 0)
			C.cvSet2D(unsafe.Pointer(pIplImage), C.int(i), C.int(j), rgb);
		}
	}

	C.cvShowImage(C.CString(name), unsafe.Pointer(pIplImage))
	C.cvReleaseImage(&pIplImage)
	return
}

// 调整窗口大小
func cvResizeWindow(name string, width, height int)(err os.Error) {
	C.cvResizeWindow(C.CString(name), C.int(width), C.int(height))
	return
}
// 移动窗口
func cvMoveWindow(name string, x, y int)(err os.Error) {
	C.cvMoveWindow(C.CString(name), C.int(x), C.int(y))
	return
}

// 销毁窗口
func cvDestroyWindow(name string)(err os.Error) {
	C.cvDestroyWindow(C.CString(name))
	return
}
// 销毁全部窗口
func cvDestroyAllWindows() {
	C.cvDestroyAllWindows()
}

func cvWaitKey(key int) {
	C.cvWaitKey(C.int(key))
}

//=============================================================================
//ddd
func init_wrap() {
	C.puts(C.CString("sss"))
}
