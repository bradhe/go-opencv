// copyright 2011 <chaishushan@gmail.com>. all rights reserved.
// use of this source code is governed by a bsd-style
// license that can be found in the LICENSE file.

package opencv

/*
#include "opencv.h"
#cgo linux  pkg-config: opencv
#cgo darwin pkg-config: opencv
#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
*/
import "C"

import (
	"reflect"
	"unsafe"
)

//

/****************************************************************************************\
*                                  Eigen objects                                         *
\****************************************************************************************/



/****************************************************************************************\
*                                       1D/2D HMM                                        *
\****************************************************************************************/


/*********************************** Embedded HMMs *************************************/


/****************************************************************************************\
*               A few functions from old stereo gesture recognition demosions            *
\****************************************************************************************/


/****************************************************************************************\
*                           Additional operations on Subdivisions                        *
\****************************************************************************************/



/****************************************************************************************\
*                           More operations on sequences                                 *
\****************************************************************************************/
func (self Seq) ForEach(t reflect.Type, callback func(interface{})) {
	c_seq := (*C.CvSeq)(&self)
	seqLen := int(c_seq.total)

	for i := 0; i < seqLen; i++ {
		c_obj := C.cvGetSeqElem(c_seq, C.int(i))
		obj := reflect.NewAt(t, unsafe.Pointer(c_obj))
		callback(obj.Interface())
	}
}

/*******************************Stereo correspondence*************************************/


/*****************************************************************************************/
/************ Epiline functions *******************/


/****************************************************************************************\
*                                   Contour Morphing                                     *
\****************************************************************************************/


/****************************************************************************************\
*                                    Texture Descriptors                                 *
\****************************************************************************************/



/****************************************************************************************\
*                                  Face eyes&mouth tracking                              *
\****************************************************************************************/
func releaseHaarClassifier(classifier *HaarClassifierCascade) {
	// TODO: Figure out how to release the classifier. Oh god.
}

func NewHaarClassifier(filename string) *HaarClassifierCascade {
	// TODO: Figure out an interface for managing the three parameters
	// at the end of this function.
	classifier := C.cvLoad(C.CString(filename), nil, nil, nil)
	return (*HaarClassifierCascade)(classifier)
}

func (self *HaarClassifierCascade) DetectObjects(image *IplImage, storage *MemStorage, scale float64, minNeighbors, flags int, minSize, maxSize Size) []*Rect {
	c_storage := (*C.CvMemStorage)(storage)
	c_self := (*C.CvHaarClassifierCascade)(self)

	c_seq := C.cvHaarDetectObjects(
		unsafe.Pointer(image),
		c_self,
		c_storage,
		C.double(scale),
		C.int(minNeighbors),
		C.int(flags),
		C.cvSize(C.int(minSize.Width), C.int(minSize.Height)),
		C.cvSize(C.int(maxSize.Width), C.int(maxSize.Height)),
	)

	seq := (*Seq)(c_seq)

	var arr []*Rect
	var r Rect

	seq.ForEach(reflect.TypeOf(r), func(obj interface{}) {
		arr = append(arr, obj.(*Rect))
	})

	return arr
}

func (self *HaarClassifierCascade) Release() {
	c_self := unsafe.Pointer(self)
	C.cvRelease(&c_self)
}



/****************************************************************************************\
*                                         3D Tracker                                     *
\****************************************************************************************/




/****************************************************************************************\
*                           Skeletons and Linear-Contour Models                          *
\****************************************************************************************/


/****************************************************************************************\
*                           Background/foreground segmentation                           *
\****************************************************************************************/


/****************************************************************************************\
*                                   Calibration engine                                   *
\****************************************************************************************/



/*****************************************************************************\
*                                 --- END ---                                 *
\*****************************************************************************/


