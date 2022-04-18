package deadlock

import (
	"log"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amout int) {
	user.Balance += amout
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amout int) {
	user1.Lock()
	log.Println("Lock user 1", user1.Name)
	user1.Change(-amout)

	time.Sleep(1 * time.Second)

	user2.Lock()
	log.Println("Lock user 2", user2.Name)
	user2.Change(amout)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Ivan Rizky Saputra",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Aisyah Nisrina Habibah",
		Balance: 2000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	log.Println("User ", user1.Name, ", Balance ", user1.Balance)
	log.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
