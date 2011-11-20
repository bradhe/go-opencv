// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo LDFLAGS: -lcxcore

#include <opencv/cxcore.h>
#include <stdlib.h>
#include <string.h>

//-----------------------------------------------------------------------------

// version
const char* CV_VERSION_ = CV_VERSION;

//-----------------------------------------------------------------------------

// IplImage
static int IPL_IMAGE_MAGIC_VAL_() {
	return IPL_IMAGE_MAGIC_VAL;
}
static const char* CV_TYPE_NAME_IMAGE_() {
	return CV_TYPE_NAME_IMAGE;
}
static int CV_IS_IMAGE_HDR_(void* img) {
	return CV_IS_IMAGE_HDR(img);
}
static int CV_IS_IMAGE_(void* img) {
	return CV_IS_IMAGE(img);
}

//-----------------------------------------------------------------------------

// type filed is go keyword
static int myGetMatType(const CvMat* mat) {
	return mat->type;
}

//-----------------------------------------------------------------------------
*/
import "C"
import (
	"unsafe"
)

//-----------------------------------------------------------------------------


// cvver.h

//-----------------------------------------------------------------------------

const (
	CV_MAJOR_VERSION    = int(C.CV_MAJOR_VERSION)
	CV_MINOR_VERSION    = int(C.CV_MINOR_VERSION)
	CV_SUBMINOR_VERSION = int(C.CV_SUBMINOR_VERSION)
)

var (
	CV_VERSION = C.GoString(C.CV_VERSION_)
)

//-----------------------------------------------------------------------------
// cxerror.h
//-----------------------------------------------------------------------------

const (
	CV_StsOk                     = C.CV_StsOk
	CV_StsBackTrace              = C.CV_StsBackTrace
	CV_StsError                  = C.CV_StsError
	CV_StsInternal               = C.CV_StsInternal
	CV_StsNoMem                  = C.CV_StsNoMem
	CV_StsBadArg                 = C.CV_StsBadArg
	CV_StsBadFunc                = C.CV_StsBadFunc
	CV_StsNoConv                 = C.CV_StsNoConv
	CV_StsAutoTrace              = C.CV_StsAutoTrace
	CV_HeaderIsNull              = C.CV_HeaderIsNull
	CV_BadImageSize              = C.CV_BadImageSize
	CV_BadOffset                 = C.CV_BadOffset
	CV_BadDataPtr                = C.CV_BadDataPtr
	CV_BadStep                   = C.CV_BadStep
	CV_BadModelOrChSeq           = C.CV_BadModelOrChSeq
	CV_BadNumChannels            = C.CV_BadNumChannels
	CV_BadNumChannel1U           = C.CV_BadNumChannel1U
	CV_BadDepth                  = C.CV_BadDepth
	CV_BadAlphaChannel           = C.CV_BadAlphaChannel
	CV_BadOrder                  = C.CV_BadOrder
	CV_BadOrigin                 = C.CV_BadOrigin
	CV_BadAlign                  = C.CV_BadAlign
	CV_BadCallBack               = C.CV_BadCallBack
	CV_BadTileSize               = C.CV_BadTileSize
	CV_BadCOI                    = C.CV_BadCOI
	CV_BadROISize                = C.CV_BadROISize
	CV_MaskIsTiled               = C.CV_MaskIsTiled
	CV_StsNullPtr                = C.CV_StsNullPtr
	CV_StsVecLengthErr           = C.CV_StsVecLengthErr
	CV_StsFilterStructContentErr = C.CV_StsFilterStructContentErr
	CV_StsKernelStructContentErr = C.CV_StsKernelStructContentErr
	CV_StsFilterOffsetErr        = C.CV_StsFilterOffsetErr
	CV_StsBadSize                = C.CV_StsBadSize
	CV_StsDivByZero              = C.CV_StsDivByZero
	CV_StsInplaceNotSupported    = C.CV_StsInplaceNotSupported
	CV_StsObjectNotFound         = C.CV_StsObjectNotFound
	CV_StsUnmatchedFormats       = C.CV_StsUnmatchedFormats
	CV_StsBadFlag                = C.CV_StsBadFlag
	CV_StsBadPoint               = C.CV_StsBadPoint
	CV_StsBadMask                = C.CV_StsBadMask
	CV_StsUnmatchedSizes         = C.CV_StsUnmatchedSizes
	CV_StsUnsupportedFormat      = C.CV_StsUnsupportedFormat
	CV_StsOutOfRange             = C.CV_StsOutOfRange
	CV_StsParseError             = C.CV_StsParseError
	CV_StsNotImplemented         = C.CV_StsNotImplemented
	CV_StsBadMemBlock            = C.CV_StsBadMemBlock
	CV_StsAssert                 = C.CV_StsAssert
	//CV_GpuNotSupported=          C.CV_GpuNotSupported
	//CV_GpuApiCallError=          C.CV_GpuApiCallError
	//CV_GpuNppCallError=          C.CV_GpuNppCallError
)

//-----------------------------------------------------------------------------
// cxtypes.h
//-----------------------------------------------------------------------------

type Arr unsafe.Pointer

/*****************************************************************************\
*                      Common macros and inline functions                     *
\*****************************************************************************/

const (
	CV_PI   = 3.1415926535897932384626433832795
	CV_LOG2 = 0.69314718055994530941723212145818
)

func Round(value float64) int {
	rv := C.cvRound(C.double(value))
	return int(rv)
}
func Floor(value float64) int {
	rv := C.cvFloor(C.double(value))
	return int(rv)
}
func Ceil(value float64) int {
	rv := C.cvCeil(C.double(value))
	return int(rv)
}

func IsNaN(value float64) int {
	rv := C.cvIsNaN(C.double(value))
	return int(rv)
}
func IsInf(value float64) int {
	rv := C.cvIsInf(C.double(value))
	return int(rv)
}

/*************** Random number generation *******************/

type RNG C.CvRNG

func NewRNG(seed int64) RNG {
	rv := C.cvRNG(C.int64(seed))
	return RNG(rv)
}
func (rng *RNG) RandInt() uint32 {
	rv := C.cvRandInt((*C.CvRNG)(rng))
	return uint32(rv)
}
func (rng *RNG) RandReal() float64 {
	rv := C.cvRandReal((*C.CvRNG)(rng))
	return float64(rv)
}

/*****************************************************************************\
*                            Image type (IplImage)                            *
\*****************************************************************************/

/*
 * The following definitions (until #endif)
 * is an extract from IPL headers.
 * Copyright (c) 1995 Intel Corporation.
 */
const (
	IPL_DEPTH_SIGN = C.IPL_DEPTH_SIGN

	IPL_DEPTH_1U  = C.IPL_DEPTH_1U
	IPL_DEPTH_8U  = C.IPL_DEPTH_8U
	IPL_DEPTH_16U = C.IPL_DEPTH_16U
	IPL_DEPTH_32F = C.IPL_DEPTH_32F

	IPL_DEPTH_8S  = C.IPL_DEPTH_8S
	IPL_DEPTH_16S = C.IPL_DEPTH_16S
	IPL_DEPTH_32S = C.IPL_DEPTH_32S

	IPL_DATA_ORDER_PIXEL = C.IPL_DATA_ORDER_PIXEL
	IPL_DATA_ORDER_PLANE = C.IPL_DATA_ORDER_PLANE

	IPL_ORIGIN_TL = C.IPL_ORIGIN_TL
	IPL_ORIGIN_BL = C.IPL_ORIGIN_BL

	IPL_ALIGN_4BYTES  = C.IPL_ALIGN_4BYTES
	IPL_ALIGN_8BYTES  = C.IPL_ALIGN_8BYTES
	IPL_ALIGN_16BYTES = C.IPL_ALIGN_16BYTES
	IPL_ALIGN_32BYTES = C.IPL_ALIGN_32BYTES

	IPL_ALIGN_DWORD = C.IPL_ALIGN_DWORD
	IPL_ALIGN_QWORD = C.IPL_ALIGN_QWORD

	IPL_BORDER_CONSTANT  = C.IPL_BORDER_CONSTANT
	IPL_BORDER_REPLICATE = C.IPL_BORDER_REPLICATE
	IPL_BORDER_REFLECT   = C.IPL_BORDER_REFLECT
	IPL_BORDER_WRAP      = C.IPL_BORDER_WRAP
)

type IplImage C.IplImage

// normal fields
func (img *IplImage)Channels() int {
	return int(img.nChannels)
}
func (img *IplImage)Depth() int {
	return int(img.depth)
}
func (img *IplImage)Origin() int {
	return int(img.origin)
}
func (img *IplImage)Width() int {
	return int(img.width)
}
func (img *IplImage)Height() int {
	return int(img.height)
}
func (img *IplImage)WidthStep() int {
	return int(img.widthStep)
}
func (img *IplImage)ImageSize() int {
	return int(img.imageSize)
}
func (img *IplImage)ImageData() unsafe.Pointer {
	return unsafe.Pointer(img.imageData)
}

type IplROI C.IplROI

func (roi *IplROI)Init(coi, xOffset, yOffset, width, height int) {
	roi_c := (*C.IplROI)(roi)
	roi_c.coi = C.int(coi)
	roi_c.xOffset = C.int(xOffset)
	roi_c.yOffset = C.int(yOffset)
	roi_c.width  = C.int(width)
	roi_c.height = C.int(height)
}
func (roi *IplROI)Coi() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.coi)
}
func (roi *IplROI)XOffset() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.xOffset)
}
func (roi *IplROI)YOffset() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.yOffset)
}
func (roi *IplROI)Width() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.width)
}
func (roi *IplROI)Height() int {
	roi_c := (*C.IplROI)(roi)
	return int(roi_c.height)
}

type IplConvKernel C.IplConvKernel
type IplConvKernelFP C.IplConvKernelFP

const (
	IPL_IMAGE_HEADER = C.IPL_IMAGE_HEADER
	IPL_IMAGE_DATA   = C.IPL_IMAGE_DATA
	IPL_IMAGE_ROI    = C.IPL_IMAGE_ROI
)

/* extra border mode */
var (
	IPL_IMAGE_MAGIC_VAL = C.IPL_IMAGE_MAGIC_VAL_()
	CV_TYPE_NAME_IMAGE  = C.CV_TYPE_NAME_IMAGE_()
)

func CV_IS_IMAGE_HDR(img unsafe.Pointer) bool {
	rv := C.CV_IS_IMAGE_HDR_(img)
	return (int(rv) != 0)
}
func CV_IS_IMAGE(img unsafe.Pointer) bool {
	rv := C.CV_IS_IMAGE_(img)
	return (int(rv) != 0)
}

const (
	IPL_DEPTH_64F = int(C.IPL_DEPTH_64F)
)

/****************************************************************************************\
*                                  Matrix type (CvMat)                                   *
\****************************************************************************************/

type Mat C.CvMat

func (mat *Mat)Type() int {
	return int(C.myGetMatType((*C.CvMat)(mat)))
}
func (mat *Mat)Step() int {
	return int(mat.step)
}

func (mat *Mat)Rows() int {
	return int(mat.rows)
}
func (mat *Mat)Cols() int {
	return int(mat.cols)
}

func CV_IS_MAT_HDR(mat interface{}) bool {
	return false
}
func CV_IS_MAT(mat interface{}) bool {
	return false
}
func CV_IS_MASK_ARR() bool {
	return false
}
func CV_ARE_TYPE_EQ() bool {
	return false
}

/****************************************************************************************\
*                       Multi-dimensional dense array (CvMatND)                          *
\****************************************************************************************/

type MatND C.CvMatND

/****************************************************************************************\
*                      Multi-dimensional sparse array (CvSparseMat)                      *
\****************************************************************************************/

type SparseMat C.CvSparseMat

/**************** iteration through a sparse array *****************/

type SparseNode C.CvSparseNode
type SparseMatIterator C.CvSparseMatIterator

/****************************************************************************************\
*                                         Histogram                                      *
\****************************************************************************************/

type HistType C.CvHistType
type Histogram C.CvHistogram

/****************************************************************************************\
*                      Other supplementary data type definitions                         *
\****************************************************************************************/

/*************************************** CvRect *****************************************/

type Rect C.CvRect

func (r *Rect)X() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.x)
}
func (r *Rect)Y() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.y)
}
func (r *Rect)Width() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.width)
}
func (r *Rect)Height() int {
	r_c := (*C.CvRect)(r)
	return int(r_c.height)
}

func (r *Rect)ToROI(coi int) IplROI {
	r_c := (*C.CvRect)(r)
	return (IplROI)(C.cvRectToROI(*r_c, C.int(coi)))
}

func (roi *IplROI)ToRect() Rect {
	r := C.cvRect(
		C.int(roi.XOffset()),
		C.int(roi.YOffset()),
		C.int(roi.Width()),
		C.int(roi.Height()),
	)
	return Rect(r)
}

/*********************************** CvTermCriteria *************************************/

type TermCriteria C.CvTermCriteria

/******************************* CvPoint and variants ***********************************/

type Point struct {
	X int
	Y int
}

type Point2D32f struct {
	X float32
	Y float32
}
type Point3D32f struct {
	X float32
	Y float32
	Z float32
}

//type Point2D32f C.CvPoint2D32f


//type Point3D32f C.CvPoint3D32f



type Point2D64f C.CvPoint2D64f
type Point3D64f C.CvPoint3D64f

/******************************** CvSize's & CvBox **************************************/

type Size struct {
	Width  int
	Height int
}

type Size2D32f C.CvSize2D32f

type Box2D C.CvBox2D

type LineIterator C.CvLineIterator

/************************************* CvSlice ******************************************/

type Slice C.CvSlice

/************************************* CvScalar *****************************************/

type Scalar C.CvScalar

func ScalarAll(val0 float64) Scalar {
	rv := C.cvScalarAll(C.double(val0))
	return (Scalar)(rv)
}

/****************************************************************************************\


*                                   Dynamic Data structures                              *


\****************************************************************************************/

/******************************** Memory storage ****************************************/

type MemBlock C.CvMemBlock
type MemStoragePos C.CvMemStoragePos

/*********************************** Sequence *******************************************/

type SeqBlock C.CvSeqBlock

/*************************************** Set ********************************************/

/************************************* Graph ********************************************/

/*********************************** Chain/Countour *************************************/

/****************************************************************************************\


*                                    Sequence types                                      *


\****************************************************************************************/

/****************************************************************************************/
/*                            Sequence writer & reader                                  */
/****************************************************************************************/

/****************************************************************************************/
/*                                Operations on sequences                               */
/****************************************************************************************/

/****************************************************************************************\


*             Data structures for persistence (a.k.a serialization) functionality        *


\****************************************************************************************/

/*****************************************************************************\


*                                 --- END ---                                 *


\*****************************************************************************/

