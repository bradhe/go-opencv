package  highgui

/*
#include <opencv/highgui.h>
*/
//import "C"
import "os"
import "image"
import "io"

//=============================================================================

// An RGBA is an in-memory image backed by a 2-D slice of RGBAColor values.
type IplImage struct {
	// The Pixel field's indices are y first, then x, so that At(x, y) == Pixel[y][x].
	Pixel [][]image.RGBAColor
}

func (p *IplImage) ColorModel() image.ColorModel { return image.NRGBAColorModel }

func (p *IplImage) Width() int {
	if len(p.Pixel) == 0 {
		return 0
	}
	return len(p.Pixel[0])
}

func (p *IplImage) Height() int { return len(p.Pixel) }

func (p *IplImage) At(x, y int) image.Color { return p.Pixel[y][x] }

//func (p *IplImage) Set(x, y int, c image.Color) { p.Pixel[y][x] = toRGBAColor(c).(RGBAColor) }

// NewRGBA returns a new RGBA with the given width and height.
func NewIplImage(w, h int) *IplImage {

	return &IplImage{}
}

/****************************************************************************************\
*                                  Basic GUI functions                                   *
\****************************************************************************************/

// 初始化系统
func InitSystem(args []string) {
	cvInitSystem(args)
}
// 启动窗口线程
func StartWindowThread()(err os.Error) {
	return cvStartWindowThread()
}

// 命名窗口
func NamedWindow(name string, autoSize bool)(err os.Error) {
 	return cvNamedWindow(name, autoSize)
}

// 设置窗口属性
func SetWindowProperty(name string, prop_id int, prop_value float)(err os.Error) {
	return nil
}
// 查询窗口属性
func GetWindowProperty(name string, prop_id int)(prop_value float, err os.Error) {
	return
}

// 显式图像
func ShowImage(name string, image image.Image)(err os.Error) {
	return
}

// 调整窗口大小
func ResizeWindow(name string, width, height int)(err os.Error) {
	return
}
// 移动窗口
func MoveWindow(name string, x, y int)(err os.Error) {
	return
}

// 销毁窗口
func DestroyWindow(name string)(err os.Error) {
	return
}
// 销毁全部窗口
func DestroyAllWindows() {
}

//=========================================================================
// 拖动条
type TrackbarCallback interface {
	OnChange(trackbar_name, window_name string, value, maxValue int)
}

// 创建拖动条
func CreateTrackbar(barName, winName string, value, maxValue int, callback TrackbarCallback)(err os.Error) {
	return
}

// 查询拖动条位置
func GetTrackbarPos(barName, winName string)(value int, err os.Error) {
	return
}
// 设置脱动条位置
func SetTrackbarPos(barName, winName string, pos int)(err os.Error) {
	return
}

// 窗口回调
type MouseCallback interface {
	OnMove(event, x, y, flags int)
}

// 消息循环
func WaitKey(ms int) {
	cvWaitKey(ms)
}

//=========================================================================

// 读图像
func LoadImage(name string)(img image.Image, err os.Error) {
	return
}

// 保存图像
func SaveImage(name string, img image.Image)(err os.Error) {
	return
}

// 解码图像
func DecodeImage(r io.Reader)(img image.Image, err os.Error) {
	return
}
// 编码图像
func EncodeImage(w io.Writer, img image.Image)(err os.Error) {
	return
}

// 转换图像
func ConvertImage(src, dst image.Image)(err os.Error) {
	return
}
                      
/****************************************************************************************\
*                         Working with Video Files and Cameras                           *
\****************************************************************************************/

// 模仿图片，构造一个视频接口


type CvCapture struct {
	a int
}

// 打开是否文件
func CreateFileCapture(name string)(cap CvCapture, err os.Error) {
	return
}

// 打开相机文件
func CreateCameraCapture(index int)(cap CvCapture, err os.Error) {
	return
}

// 获取一帧图像
func (cap *CvCapture)GrabFrame()(err os.Error) {
	return
}

// 获取之前抓取的帧
func RetrieveFrame()(img image.Image, err os.Error) {
	return
}
// 获取一帧图像
func (cap *CvCapture)QueryFrame()(img image.Image, err os.Error) {
	return
}

// 释放图像
func ReleaseCapture(){
}

// 属性
func GetCaptureProperty() {
}
func SetCaptureProperty() {
}

func GetCaptureDomain() {
}
//=========================================================================

// 创建写视频
func CreateVideoWriter() {
}

// 写一帧
func WriteFrame() {
}

// 关闭
func ReleaseVideoWriter() {
} 

//=========================================================================



