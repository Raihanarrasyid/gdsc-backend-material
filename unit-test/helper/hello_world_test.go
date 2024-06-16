package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Tests are about to run")
	m.Run()
	fmt.Println("Tests have finished running")
}

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Hello World", HelloWorld("World"))
}

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2))
}