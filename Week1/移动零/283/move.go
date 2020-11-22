package main

import (
	"fmt"
	"time"
)

func moveZeroes(nums []int)  {
	notZero:=0
	for i, num := range nums {
		if num!=0{
			if i!=notZero{
				nums[notZero]=num
			}
			notZero++
			continue
		}
	}
	for i := notZero; i < len(nums); i++ {
		nums[i]=0
	}
	fmt.Println(nums)
	time.Sleep(10*time.Microsecond)
}