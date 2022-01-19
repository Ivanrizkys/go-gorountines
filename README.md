# Gorountine

Istilahnya gorountine adalah sebuah thread yang ringan yang berjalan di dalam thread sebuah komputer. Gorountine sendiri berjalan secara concurency dan bersifat non-blocking. Di golang sendiri, jumlah go rountine secara default adalah mengikuti jumlah core di sebuah sistem operasi. Selain itu goroutine juga sangat ringan, kita bisa membuat puluhan hinnga ribuan gorountine tanpa membuat boros memory.

## Membuat Gorountine

Untuk membuat gorountine, cukup meggunakan perintah **go** sebelum pemanggilan nama function. Saat function beralan secara gorountine, otomatis function itu berjalan secara asyncronus. Namun akan kurang cocok jika gorountine digunakan untuk function yang memiliki return value.

# Chanel

Channel adalah tempat komunikasi secara syncronus di dalam gorountine. Saat kita membuat function yang memiliki return value, kita bisa menggunakan channel untuk menanngkap return value dari function tersebut. Untuk itu kita membutuhkan dua gorountine, satu untuk mengirim data dan satunya lagi untuk menangkap data. Channel ini mirip seperti mekanisme **async** **await** di JavaScript.

Channel hanya bisa menampung satu data, oleh karena itu saat kita ingin mengirimkan data lagi di channel, data sebelumnya harus di ambil terlebih dahulu. Selain itu, channel harus di close jika sudah tidak digunakan karena kalau tidak bisa menyebabkan memory leak.

## Membuat Channel

Saat membuat chanel kita di haruskan untuk mengirim dan menerima data ke channel. Jika hanya satu proses saja, maka program akan di block dan akan menyebabkan error.

```go
import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	// * membuat channel
	// * channel memiliki tipe data chan
	// * di dalam channel hanya boleh menampung tipe data string
	chanel := make(chan string)
	// * untuk close chanel
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)
		// * mengirim data ke sebuah channel
		chanel <- "Hai doddy"
		fmt.Println("Berhasil mengirimkan data ke channel")
	}()

	// * mengambil data dari channel
	// * di ambil dan disimpan di variabel
	data := <-chanel
	fmt.Println(data)
	time.Sleep(4 * time.Second)
}
```

## Channel as parameter

Kita juga bisa membuat channel menjadi parameter sebuah function. Berbeda dengan variabel yang akan pass by value, jika kita memberikan parameter function dengan sebuah channel, maka akan secara default pass by reference. Jadi kita tidak perlu lagi membuat pointer seperti halnya jika mengirimkan parameter dengan sebuah variabel.
```go
import (
	"fmt"
	"testing"
	"time"
)

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
```

## Channel In dan Channel Out

Saat kita mengirimkan parameter function dengan sebuah channel, maka secara default function tersebut bisa digunakan untuk mengirim atau menerima channel. Jika kita hanya ingin menjadikan function tersebut hanya bisa untuk mengirimkan data atau menerima data dari channel, Maka kita bisa melakukan nya dengan seperti berikut.

```go
// * function ini hanya bisa digunakan untuk mengirimkan value ke sebuah channel
// * chan<- membuat function hanya bisa untuk mengirimkan value ke calam channel
func OnlyIn(channel chan<- string) {
	channel <- "Hallo semua apa kabar"
}

// * function ini hanya bisa digunakan untuk memperoleh data yang dikirimkan dari sebuah channel
// * <-chan membuat function hanya bisa memperoleh data dari sebuah channel
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
```

## Buffered Channel

Secara default saat kita membuat channel, kita hanya bisa mengisi channel tersebut dengan satu data. Untuk merubahya kita bisa menyimpan data tersebut di dalam buffer yang berada di dalam channel.

```go
func TestBufferedChannel(t *testing.T) {
	// * channel bisa menerima 3 data dalam bentuk string
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Ivan"
	fmt.Println("Succes")
}
```

Jika kita membuat kode seperti di atas dan channel kita tidak memiliki buffer/default. Maka kode tersebut akan error/terkena blocking karena isi dari channel tersebut tidak diambil. Namun, jika menggunakan buffer hal tersebut tidak terjadi.

Untuk mekanisme sederhananya seperti di bawah ini:

```go
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
```

## Range Channel

Jika kita tidak bisa memperkirakan channel tersebut bisa memnerima berapa data / pengirim mengirimkan data channel secara terus menerus. Maka kita bisa melakukan dengan menggunakan range channel.
```go
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
```

## Select Channel

Jika kita memiliki lebih dari satu channel dan ingin mengambil data dari sebuah channel, Kita tidak bisa mengambilnya dengan menggunakan range channel dengan menggunakan for range biasa. Sebagai gantinya, kita bisa menggunakan select.

```go
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeChanelValue(channel1)
	go GiveMeChanelValue(channel2)

	// * mengambil data dari channel 1 atau 2
	// * tergantung mana yang lebih cepat
	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1 " + data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2 " + data)
	}

	// * mengambil data dari channel 1 atau 2
	// * tergantung mana yang lebih cepat
	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1 " + data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2 " + data)
	}
}
```

Dalam kode diatas, Jika kita hanya melakukan satu select. Kode tersebut akan error karena salah satu channel tidak diambil value/data nya. Namun, ada cara lain agar kita bisa membuat kode jadi lebih ringkas, yaitu dengan menggunakan perulangan for.

```go
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
		}
		if counter == 2 {
			break
		}
	}
}
```

Sebelum channel yang akan di select mempunyai data maka secara deafult select akan menunggu data terlebih dahulu selama proses yang ada di function **GiveMeChannelResult()**. Saat proses menunggu tersebut kita bisa melakukan sesuatu di dalam **default**. Contohnya seperti di bawah ini.

```go
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
```