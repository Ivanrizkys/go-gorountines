package pool

import (
	"log"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	// * untuk memberikan return value saat pool kosong
	var pool sync.Pool = sync.Pool{
		New: func() any {
			return "new"
		},
	}
	var group sync.WaitGroup

	// * digunakan untuk menambah data ke pool
	pool.Put("Ivan")
	pool.Put("Rizky")
	pool.Put("Saputra")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			// * digunakan untuk mengambil data dari pool
			data := pool.Get()
			log.Println(data)
			// * jika kita sudah menggunakan data dari pool maka kita
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
}
