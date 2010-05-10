
package main

//#include <stdio.h>
//import "C"

import (
	"fmt";
	"image/png"
	"image"
	"os"
	//"opencv"
)

func readPng(filename string) (image.Image, os.Error) {
	f, err := os.Open(filename, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	print("read 2\n")
	defer f.Close()
	return png.Decode(f)
}
func writePng(filename string, img image.Image) (os.Error) {
	f, err := os.Open(filename, os.O_WRONLY , 0666)
	if err != nil {
		print("create err")
		return err
	}

	print("w 2\n")
	defer f.Close()
	return png.Encode(f, img)
}

func main() {
	//C.puts(C.CString("aa"));
	fmt.Printf("Hello, 世界!\n");
	
	img, err := readPng("./images/lena.png")
	if err != nil {
		print("save\n")
		if writePng("./111.png", img) != nil {
			print("ok")
		}
	}

	//img := new(opencv.Image);
	//img.Load("lena.jpg");
	//img.Save("11");
}

