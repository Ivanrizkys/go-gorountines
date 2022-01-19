# Gorountine

Istilahnya gorountine adalah sebuah thread yang ringan yang berjalan di dalam thread sebuah komputer. Gorountine sendiri berjalan secara concurency dan bersifat non-blocking. Di golang sendiri, jumlah go rountine secara default adalah mengikuti jumlah core di sebuah sistem operasi. Selain itu goroutine juga sangat ringan, kita bisa membuat puluhan hinnga ribuan gorountine tanpa membuat boros memory.

## Membuat Gorountine

Untuk membuat gorountine, cukup meggunakan perintah **go** sebelum pemanggilan nama function. Saat function beralan secara gorountine, otomatis function itu berjalan secara asyncronus. Namun akan kurang cocok jika gorountine digunakan untuk function yang memiliki return value.

# Chanel

Chanel adalah tempat komunikasi secara syncronus di dalam gorountine. Saat kita membuat function yang memiliki return value, kita bisa menggunakan chanel untuk menangkap return value dari function tersebut. Untuk itu kita membutuhkan dua gorountine, satu untuk mengirim data dan satunya lagi untuk menangkap data. Chanel ini mirip seperti mekanisme **async** **await** di JavaScript.

Chanel hanya bisa menampung satu data, oleh karena itu saat kita ingin mengirimkan data lagi di chanel, data sebelumnya harus di ambil terlebih dahulu. Selain itu, chanel harus di close jika sudah tidak digunakan karena kalau tidak bisa menyebabkan memory leak.

## Membuat Chanel

Saat membuat chanel kita di haruskan untuk mengirim dan menerima data ke chanel. Jika hanya satu proses saja, maka program akan di block dan akan menyebabkan error.

```go
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
```

## Chanel as parameter

Kita juga bisa membuat chanel menjadi parameter sebuah function. Berbeda dengan variabel yang akan pass by value, jika kita memberikan parameter function dengan sebuah chanel, maka akan secara default pass by reference. Jadi kita tidak perlu lagi membuat pointer seperti halnya jika mengirimkan parameter dengan sebuah variabel.
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
