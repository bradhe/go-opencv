
package highgui

/*
#include <stdio.h>
#include <opencv/highgui.h>

// 
static char*  argv_buf[128];
static char** argv = argv_buf;

void argv_set(int i, char *v) {
	if(i < sizeof(argv_buf)/sizeof(argv_buf[0])) {
		argv[i] = v;
	}
}
char** toCharPP(void *p) {
	return (char**)p;
}

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

const IPL_DEPTH_8U = C.GO_IPL_DEPTH_8U

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

	C.cvReleaseImage(&pIplImage)
	return
}
func cvWaitKey(key int) {
	C.cvWaitKey(C.int(key))
}

//ddd
func init_wrap() {
	C.puts(C.CString("sss"))
}
