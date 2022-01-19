package chanel

import (
	"fmt"
	"strconv"
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

// * function ini hanya bisa digunakan untuk mengirimkan value ke sebuah channel
func OnlyIn(channel chan<- string) {
	channel <- "Hallo semua apa kabar"
}

// * function ini hanya bisa digunakan untuk memperoleh data yang dikirimkan dari sebuah channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

// * buffered channel
func TestBufferedChannel(t *testing.T) {
	// * channel bisa menerima 3 data dalam bentuk string
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Ivan"
	channel <- "Rizky"
	channel <- "Saputra"

	// * akan mencetak channel yang dikirimkan pertaman ("Ivan")
	fmt.Println(<-channel)
	// * akan mencetak channel yang dikirimkan kedua ("Rizky")
	fmt.Println(<-channel)
	// * akan mencetak channel yang dikirimkan ketiga ("Saputra")
	fmt.Println(<-channel)
	fmt.Println("Succes")
}

// * range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Channel ke: " + strconv.Itoa(i)
		}
		// * saat di close maka perulangan untuk menerima data dari channel juga akan berhenti
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data " + data)
	}
}

// * Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeChanelValue(channel1)
	go GiveMeChanelValue(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 " + data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 " + data)
			counter++
		default:
			// * ini akan dicetak terus menerus selama channel belum berisi data
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
