include $(GOROOT)/src/Make.inc

TARG=opencv

CGOFILES=\
	cxcore.go\
	highgui.go

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

