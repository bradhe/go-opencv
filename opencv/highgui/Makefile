# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=opencv/highgui

# Can have plain GOFILES too, but this example doesn't.

CGOFILES=\
	cv.go 
	#highgui.go 

GOFILES=doc.go highgui.go 

CGO_LDFLAGS=-lcv -lcxcore -lhighgui

# To add flags necessary for locating the library or its include files,
# set CGO_CFLAGS or CGO_LDFLAGS.  For example, to use an
# alternate installation of the library:
#	CGO_CFLAGS=-I/home/rsc/gmp32/include
#	CGO_LDFLAGS+=-L/home/rsc/gmp32/lib
# Note the += on the second line.

CGO_CFLAGS+=-I/usr/local/include
CGO_LDFLAGS+=-L/usr/local/lib
CLEANFILES+=hellocv

include $(GOROOT)/src/Make.pkg

# Simple test programs
%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O


