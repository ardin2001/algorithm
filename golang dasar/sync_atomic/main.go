package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var point int64
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 1000; j++ {
				atomic.AddInt64(&point, 1)
			}
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(point)
}

// package main

// import (
//     "fmt"
//     "sync"
//     "sync/atomic"
// )

// func main() {

//     var ops uint64

//     var wg sync.WaitGroup

//     for i := 0; i < 50; i++ {
//         wg.Add(1)

//         go func() {
//             for c := 0; c < 1000; c++ {

//                 atomic.AddUint64(&ops, 1)
//             }
//             wg.Done()
//         }()
//     }

//     wg.Wait()

//     fmt.Println("ops:", ops)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func anjay(i int, group *sync.WaitGroup) {
// 	defer group.Done()
// 	fmt.Println("losss", i)
// }
// func main() {
// 	group := sync.WaitGroup{}
// 	for i := 1; i <= 100; i++ {
// 		group.Add(1)
// 		go anjay(i, &group)
// 	}
// 	group.Wait()
// 	fmt.Println("===========done============")
// }
