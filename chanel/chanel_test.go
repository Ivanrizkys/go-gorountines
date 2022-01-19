package chanel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	// * membuat chanel
	// * chanel memiliki tipe data chan
	// * di dalam chanel hanya boleh menampung tipe data string
	chanel := make(chan string)
	// * untuk close chanel
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		// * mengirim data ke sebuah chanel
		chanel <- "Hai doddy"
		fmt.Println("Berhasil mengirimkan data ke chanel")
	}()

	// * mengambil data dari chanel
	// * di ambil dan disimpan di variabel
	data := <-chanel
	fmt.Println(data)
	time.Sleep(4 * time.Second)
}

func GiveMeChanelValue(chanel chan string) {
	chanel <- "Hai doddy apa kabar , kamu sehat kan"
}

func TestChanelAsParameter(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)

	go GiveMeChanelValue(chanel)
	data := <-chanel
	fmt.Println(data)
}
