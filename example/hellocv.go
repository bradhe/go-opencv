
package main

//#include <stdio.h>
//import "C"

import (
	"fmt";
	"opencv"
)

func main() {
	//C.puts(C.CString("aa"));
	fmt.Printf("Hello, 世界!\n");

	img := new(opencv.Image);
	img.Load("lena.jpg");
	img.Save("11");
}

