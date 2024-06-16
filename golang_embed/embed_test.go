package golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed a.jpg
var image []byte

//go:embed files/a.txt
//go"embed files/b.txt
var files embed.FS

func TestString(t *testing.T) {
	fmt.Println(version)
}

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("new.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestMultipleFIles(t *testing.T) {
	fmt.Println(image)
}