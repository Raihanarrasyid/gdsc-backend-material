package helper

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	a := make(chan string)
	defer close(a)
	go func() {
		a <- "hello"
	}()
	time.Sleep(2*time.Second)
	fmt.Println()
}