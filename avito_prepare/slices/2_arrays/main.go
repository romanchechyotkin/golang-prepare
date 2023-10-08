package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	arr := [2]int{30}
	arr2 := [3]int{40, 50}

	fmt.Println(arr, unsafe.Sizeof(arr))   // [30 0] 16
	fmt.Println(arr2, unsafe.Sizeof(arr2)) // [40 50 0] 24

	mergedArr := [len(arr) + len(arr2)]int{}
	counter := 0
	for i := 0; i < len(arr); i++ {
		mergedArr[i] = arr[i]
		counter++
	}
	for i := 0; i < len(arr2); i++ {
		mergedArr[counter+i] = arr2[i]
	}
	fmt.Println(mergedArr, unsafe.Sizeof(mergedArr))

	fmt.Println("---------------------------------")

	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats: %d\n", stats.HeapAlloc)
	// stats: 150640

	bigArray := [1 << 25]int{}
	_ = bigArray
	runtime.ReadMemStats(stats)
	fmt.Printf("stats after created bigArray: %d\n", stats.HeapAlloc)
	// stats after created bigArray: 268591872

	go funcWithArray(bigArray)
	<-time.After(time.Second)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats bigArray copied to goroutine stack: %d\n", stats.HeapAlloc)
	// stats bigArray copied to goroutine stack: 537030864

	go funcWithArrayPtr(&bigArray)
	<-time.After(time.Second)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats bigArray passed by ptr: %d\n", stats.HeapAlloc)
	// stats bigArray passed by ptr: 537033168

}

func funcWithArray(arr [1 << 25]int) {
	time.Sleep(time.Hour)
	_ = arr
}

func funcWithArrayPtr(arr *[1 << 25]int) {
	time.Sleep(time.Hour)
	_ = arr
}
