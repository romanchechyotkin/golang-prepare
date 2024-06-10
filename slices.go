package main

import (
	"fmt"
	"log"
	"sync"
)

type account struct {
	value int
}

func main() {
	a := make([]int32, 0)
	log.Println(a, cap(a), len(a))
	a = append(a, []int32{1, 2, 3}...)
	log.Println(a, len(a), cap(a))
	a = append(a, 4)
	log.Println(a, len(a), cap(a))

	a2 := make([]int64, 0)
	log.Println(a2, cap(a2), len(a2))
	a2 = append(a2, []int64{1, 2, 3}...)
	log.Println(a2, len(a2), cap(a2))
	a2 = append(a2, 4)
	log.Println(a2, len(a2), cap(a2))

	sl1 := make([]account, 0, 2)
	sl1 = append(sl1, account{})  // sl1 len 1 cap 2
	sl2 := append(sl1, account{}) // sl2 len 2 cap 2 | sl1 len 1 cap 2

	log.Println(sl1, len(sl1), cap(sl1))
	log.Println(sl2, len(sl2), cap(sl2))

	acc := &sl2[0]
	acc.value = 100

	// оба имеют acc 100
	log.Println(sl1, len(sl1), cap(sl1))
	log.Println(sl2, len(sl2), cap(sl2))

	sl1 = append(sl2, account{}) // len 3 cap 4 sl1
	acc.value += 100
	log.Println(sl1, len(sl1), cap(sl1)) // acc.value 100
	log.Println(sl2, len(sl2), cap(sl2)) // acc.value = 200

	var foo []int
	var bar []int

	foo = append(foo, 1)
	log.Println(len(foo), cap(foo))
	foo = append(foo, 2)
	log.Println(len(foo), cap(foo))

	foo = append(foo, 3)
	log.Println(len(foo), cap(foo))

	bar = append(foo, 4)
	log.Println(len(foo), cap(foo))
	log.Println(len(bar), cap(bar))
	log.Println(bar)
	log.Println(foo)

	foo = append(foo, 5)
	log.Println(len(foo), cap(foo))

	log.Println(bar)
	log.Println(bar)

	log.Println("--------------------------")

	slice := []int{1, 2, 3, 4, 5}

	log.Println(slice)

	slice = append(slice, 6, 7)

	log.Println(slice)

	test(slice)

	log.Println(slice)

	x := newX(true)
	log.Println(x.get())
	x = newX(false)
	//log.Println(x.get())

	log.Println("---------------")

	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		log.Println("channel", <-ch)
	}()

	ch <- 2

	log.Println("---------------")

	ch2 := make(chan int)

	for i2 := 0; i2 < 10; i2++ {
		go func() {
			ch2 <- (i2 + 1) * 2
		}()
	}

	go func() {
		for {
			fmt.Println(<-ch2)
		}
	}()

	res := make(chan int)
	m := map[int]int{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			m[i] = i
			res <- i
			mu.Unlock()
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			log.Println(<-res)
		}
	}()

	wg.Wait()
	close(res)

}

func test(slice []int) {
	slice = append(slice, 1)
}

type x struct {
	value int
}

func (x x) get() int {
	return x.value
}

type i interface {
	get() int
}

func newX(t bool) i {
	var a *x
	log.Println(a == nil)

	if t {
		a = &x{value: 1}
	}

	return a
}
