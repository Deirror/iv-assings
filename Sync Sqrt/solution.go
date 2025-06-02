package main

import (
	"fmt"
	"sync"
)

type sqrtRes struct {
	id  int
	num int
}

func main() {
	nums := []int{1, 2, 10, 5}

	wg := &sync.WaitGroup{}
	ch := make(chan sqrtRes, len(nums))

	for id, num := range nums {
		wg.Add(1)
		go func(id, num int) {
			defer wg.Done()
			sqrt := num * num
			ch <- sqrtRes{id, sqrt}
		}(id, num)
	}

	wg.Wait()
	close(ch)

	list := make([]int, len(nums))
	for res := range ch {
		list[res.id] = res.num
		if len(ch) == 0 {
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d squared is %d\n", nums[i], list[i])
	}
}
