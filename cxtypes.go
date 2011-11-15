// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

/*
#cgo LDFLAGS: -lcxcore

#include <opencv/cxcore.h>
#include <stdlib.h>
#include <string.h>

*/
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

type Arr C.IplImage

func Alloc2(size int) unsafe.Pointer {
	return unsafe.Pointer(C.cvAlloc(C.size_t(size)))
}
func Free2(p unsafe.Pointer) {
	C.cvFree_(p)
}




