package main

import "fmt"

type result struct {
	id      int
	sqrtNum int
}

func worker(nums []int, jobs <-chan int, res chan<- result) {
	for job := range jobs {
		sqrtNum := nums[job] * nums[job]
		res <- result{job, sqrtNum}
	}
}

func main() {
	nums := []int{1, 2, 10, 5}

	const limit = 3
	jobs := make(chan int, len(nums))
	res := make(chan result, len(nums))

	for i := 0; i < limit; i++ {
		go worker(nums, jobs, res)
	}

	for job := range len(nums) {
		jobs <- job
	}
	close(jobs)

	list := make([]int, len(nums))
	for _ = range len(nums) {
		sNum := <-res
		list[sNum.id] = sNum.sqrtNum
	}

	for i := range len(nums) {
		fmt.Printf("%d squared is %d\n", nums[i], list[i])
	}
}
