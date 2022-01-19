package gorountine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWord() {
	fmt.Println("Hello Word")
}

func TestWithNotGorountine(t *testing.T) {
	// * ini akan di jalankan terlebih dahulu
	HelloWord()
	fmt.Println("Hai")
}

func TestWithGorountine(t *testing.T) {
	go HelloWord()
	// * ini yang akan di jalankan terlebih dahulu
	fmt.Println("Hai")

	// * menunggu selama 1 detik untuk memastikan function Hellowrod dieksekusi
	time.Sleep(1 * time.Second)
}
