package once

import (
	"log"
	"sync"
	"testing"
)

var counter = 0

// * function ini hanya akan di eksekusi sekali
func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			// * hanya boleh memasukkan function yang tidak memiliki parameter
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	log.Println(counter)
}
