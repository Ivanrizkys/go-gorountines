package waitgroup

import (
	"log"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	// * menambahkan proses
	group.Add(1)

	time.Sleep(5 * time.Second)
	log.Println("Hai sayang")
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	go RunAsynchronus(group)

	log.Println("Jalan ke sini dulu")
	group.Wait()
}

// * output
// 2022/04/18 21:17:00 Jalan ke sini dulu
// 2022/04/18 21:17:05 Hai sayang
