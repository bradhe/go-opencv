
package highgui

/*
#include <stdio.h>
#include <opencv/highgui.h>

void* argv_create(int len) {
	void *p = malloc(sizeof(char*)*len);
	return p;
}
void grgv_release(void *p) {
	if(p) free(p);
}
void argv_set(void *p, int i, char *v, int len) {
	if(i < len) {
		char **argv = (char**)p;
		argv[i] = (char*)v;
	}
}
char** toCharPP(void *p) {
	return (char**)p;
}
*/
import "C"
//import "unsafe"
import "os"



// 初始化系统
func cvInitSystem(args []string) {
	argc := C.int(len(args))
	argv := C.argv_create(argc)
	
	for i := 0; i < len(args); i++ {
		C.argv_set(argv, C.int(i), C.CString(args[i]), argc)
	}
	
	C.cvInitSystem(argc, C.toCharPP(argv))
	C.grgv_release(argv)
}
// 启动窗口线程
func cvStartWindowThread()(err os.Error) {
	C.cvStartWindowThread()
	return
}

func cvWaitKey(key int) {
	C.cvWaitKey(C.int(key))
}

//ddd
func init_wrap() {
	C.puts(C.CString("sss"))
}
