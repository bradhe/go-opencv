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

/****************************************************************************************\
* Array allocation, deallocation, initialization and access to elements       *
\****************************************************************************************/

func Alloc(size int) unsafe.Pointer {
	return unsafe.Pointer(C.cvAlloc(C.size_t(size)))
}
func Free(p unsafe.Pointer) {
	C.cvFree_(p)
}

/******** matrix iterator: used for n-ary operations on dense arrays *********/




/****************************************************************************************\
*                   Arithmetic, logic and comparison operations               *
\****************************************************************************************/



/****************************************************************************************\
*                                Math operations                              *
\****************************************************************************************/


/****************************************************************************************\
*                                Matrix operations                            *
\****************************************************************************************/



/****************************************************************************************\
*                                    Array Statistics                         *
\****************************************************************************************/



/****************************************************************************************\
*                      Discrete Linear Transforms and Related Functions       *
\****************************************************************************************/




/****************************************************************************************\
*                              Dynamic data structures                        *
\****************************************************************************************/



/****************************************************************************************\
*                                     Drawing                                 *
\****************************************************************************************/



/****************************************************************************************\
*                                    System functions                         *
\****************************************************************************************/


/****************************************************************************************\
*                                    Data Persistence                         *
\****************************************************************************************/







