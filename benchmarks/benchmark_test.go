package main

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nfnt/resize"
)

func BenchmarkExample(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x++
	}
}

func BenchmarkImageResizing(b *testing.B) {
	f, err := os.Open("leipzig-gopher.png")
	if err != nil {
		b.Fatal(err.Error())
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		b.Fatal(err.Error())
	}

	// We cannot use io.Discard here since this was moved with Go 1.16 from ioutil into io.
	w := ioutil.Discard
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := resize.Resize(1000, 1000, img, resize.Lanczos3)
		err := png.Encode(w, m)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
