package  highgui

//#include <opencv/highgui.h>
import "C"
import "unsafe"

//=========================================================================

// 读图象
func (img *Image)Load(name string) int {
	img.cptr = C.cvLoadImage(C.CString(name), C.int(1));
	return 1;
}
// 保存图像
func (img *Image)Save(name string) {
	C.cvSaveImage(C.CString(name), unsafe.Pointer(img.cptr));
}

//=========================================================================

type Window struct {
	name	string
	isInit	bool
}

// 新建窗口

func (win *Window)Create(winName string) {
	if !win.isInit {
		C.cvNamedWindow(C.CString(win.name), 1);
		win.isInit = true;
	}
}

// 销毁窗口

func (win *Window)Destroy() {
	if win.isInit {
		C.cvDestroyWindow(C.CString(win.name));
		win.isInit = false;
	}
}

// 销毁所有窗口

func DestroyAllWindows() {
	C.cvDestroyAllWindows();
}

// 调整窗口大小

func (win *Window)Resize(width, height int){
	if win.isInit {
		C.cvResizeWindow(C.CString(win.name), C.int(width), C.int(height));
	}
}

// 移动窗口

func (win *Window)Move(x, y int) {
	if win.isInit {
		C.cvMoveWindow(C.CString(win.name), C.int(x), C.int(y));
	}
}

// 获取窗口名字

func (win *Window)Name() string {
	return win.name;
}

// 创建窗口

func (win *Window)NamedWindow() {
	C.cvNamedWindow(C.CString(win.name), 1);
}

// 销毁窗口

func (win *Window)DestroyWindow() {
	C.cvDestroyWindow(C.CString(win.name));
}

// 显示图形

func (win *Window)ShowImage(img *Image) {
	C.cvShowImage(C.CString(win.name), unsafe.Pointer(img.cptr));
}


// 消息循环

func WaitKey(key int) {
	C.cvWaitKey(C.int(key));
}


//=========================================================================

type Capture struct {
	cptr	*C.CvCapture
}

// 打开视频文件
func (cap *Capture)OpenFile(name string) {
	if cap.cptr != nil { cap.Release(); }
	cap.cptr = C.cvCreateFileCapture(C.CString(name));
}

// 打开摄像头
func (cap *Capture)OpenCamera(index int) {
	if cap.cptr != nil { cap.Release(); }
	cap.cptr = C.cvCreateCameraCapture(C.int(index));
}

// 关闭视频
func (cap *Capture)Release() {
	if cap.cptr != nil {
		C.cvReleaseCapture(&cap.cptr);
	}
}

// 抓取图象
func (cap *Capture)QueryFrame() *Image {
	return &Image{ cptr: C.cvQueryFrame(cap.cptr) }
}
//=========================================================================



/*
CV_EXTERN_C_FUNCPTR( void (*CvTrackbarCallback)(int pos) );

int cvCreateTrackbar( const char* trackbar_name, const char* window_name,
                      int* value, int count, CvTrackbarCallback on_change );
*/

// 创建脱动条

func (win *Window)CreateTrackbar(trackbar_name string) {
	//
}

/*
cvGetTrackbarPos

Retrieves trackbar position

int cvGetTrackbarPos( const char* trackbar_name, const char* window_name );

trackbar_name
    Name of trackbar. 
window_name
    Name of the window which is the parent of trackbar. 

The function cvGetTrackbarPos returns the ciurrent position of the specified trackbar.
cvSetTrackbarPos

Sets trackbar position

void cvSetTrackbarPos( const char* trackbar_name, const char* window_name, int pos );

trackbar_name
    Name of trackbar. 
window_name
    Name of the window which is the parent of trackbar. 
pos
    New position. 

The function cvSetTrackbarPos sets the position of the specified trackbar.

*/


// 鼠标回调函数

/*
cvSetMouseCallback

Assigns callback for mouse events

#define CV_EVENT_MOUSEMOVE      0
#define CV_EVENT_LBUTTONDOWN    1
#define CV_EVENT_RBUTTONDOWN    2
#define CV_EVENT_MBUTTONDOWN    3
#define CV_EVENT_LBUTTONUP      4
#define CV_EVENT_RBUTTONUP      5
#define CV_EVENT_MBUTTONUP      6
#define CV_EVENT_LBUTTONDBLCLK  7
#define CV_EVENT_RBUTTONDBLCLK  8
#define CV_EVENT_MBUTTONDBLCLK  9

#define CV_EVENT_FLAG_LBUTTON   1
#define CV_EVENT_FLAG_RBUTTON   2
#define CV_EVENT_FLAG_MBUTTON   4
#define CV_EVENT_FLAG_CTRLKEY   8
#define CV_EVENT_FLAG_SHIFTKEY  16
#define CV_EVENT_FLAG_ALTKEY    32

CV_EXTERN_C_FUNCPTR( void (*CvMouseCallback )(int event, int x, int y, int flags, void* param) );

void cvSetMouseCallback( const char* window_name, CvMouseCallback on_mouse, void* param=NULL );

*/




func init() {
	// init 
}

//=========================================================================
