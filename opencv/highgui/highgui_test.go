package highgui

import "testing"

func TestLoadAndSave(t *testing.T) {
   // ok
   var img Image
   img.Load("../../example/images/lena.jpg")
   img.Save("./1.bmp")
}

func TestError(t *testing.T) {
   t.Errorf("my error")
}
