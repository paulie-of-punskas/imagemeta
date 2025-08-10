package png

import (
	// "encoding/binary"
	// "io"

	// "github.com/evanoberholster/imagemeta/meta"
	// "github.com/evanoberholster/imagemeta/meta/utils"

	"os"
	"testing"
)

var (
	dir           = "../testImages/"
	benchmarksPNG = []struct {
		fileName  string
		noExifErr bool
	}{
		{"PNG_boy-217417_1280.png", true},
		{"PNG_fruit-1234657_1280.png", true},
		{"PNG_pub-2992808_1280.png", true},
		{"basn3p08.png", true},
	}
)

func TestScanPNG(t *testing.T) {
	t.Run("Read PNG files, expect no Exif", func(t *testing.T) {
		want := [2]string{"ByteOrder: UnknownEndian, Ifd: UnknownIfd, Offset: 0x0000 TiffOffset: 0x0000 Length: 0 Imagetype: application/octet-stream", "error no Exif"}

		for j := 0; j < len(benchmarksPNG); j++ {
			file, err := os.Open(dir + benchmarksPNG[j].fileName)
			defer file.Close()
			if err != nil {
				t.Fatal("Error while reading:", err)
			}
			head, meta := ScanPngHeader(file)

			if want[0] != head.String() {
				t.Fatalf("Header is different for %s: wanted %s, got %s", benchmarksPNG[j].fileName, want, head)
			}

			if want[1] != meta.Error() {
				t.Fatalf("Error message differs for %s: wanted %s, got %s", benchmarksPNG[j].fileName, want, meta)
			}
		}
	})
}
