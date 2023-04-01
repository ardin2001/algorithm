// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type userBalance struct {
// 	sync.Mutex
// 	Name    string
// 	Balance int
// }

// func (user *userBalance) Lock() {
// 	user.Mutex.Lock()
// }
// func (user *userBalance) Unlock() {
// 	user.Mutex.Unlock()
// }
// func (user *userBalance) Change(amount int) {
// 	user.Balance += amount
// }
// func Transfer(user1 *userBalance, user2 *userBalance, amount int, group *sync.WaitGroup) {
// 	defer group.Done()
// 	user1.Lock()
// 	fmt.Println(user1.Name, "loc user1 :", user1.Name)
// 	user1.Change(-amount)

// 	time.Sleep(2 * time.Second)

// 	user2.Lock()
// 	fmt.Println(user1.Name, "loc user2 :", user2.Name)
// 	user2.Change(amount)

// 	time.Sleep(2 * time.Second)

// 	user1.Unlock()
// 	user2.Unlock()

// }

// func main() {
// 	group := sync.WaitGroup{}
// 	user1 := userBalance{
// 		Name:    "ardin",
// 		Balance: 1000,
// 	}
// 	user2 := userBalance{
// 		Name:    "nugraha",
// 		Balance: 1000,
// 	}
// 	group.Add(1)
// 	go Transfer(&user1, &user2, 50, &group)
// 	group.Wait()
// 	group.Add(1)
// 	go Transfer(&user2, &user1, 150, &group)
// 	group.Wait()

// 	fmt.Println("user", user1.Name, "balance", user1.Balance)
// 	fmt.Println("user", user2.Name, "balance", user2.Balance)
// }
