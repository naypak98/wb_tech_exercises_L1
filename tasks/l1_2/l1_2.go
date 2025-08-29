package l1_2

import "fmt"

func square(number int, ch chan int) {
	ch <- number * number
}

func Run() {
	array := []int{2, 4, 6, 8, 10}

	results := make(chan int, len(array))

	for _, n := range array {
		go square(n, results)
	}

	fmt.Println("Задание L1.2: ")

	for i := 0; i < len(array); i++ {
		result := <-results
		fmt.Println(result)

	}

}
