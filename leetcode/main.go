package main

import "fmt"

func main() {
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))	
    fmt.Println(reverseDigits(212))
}

func reverseDigits(num int) int {
    revNum := 0;
    for {
        if num == 0 {
            break
        }
        revNum = revNum * 10 + num % 10;
        num /= 10;
    }
    return revNum
}

func isPalindrome(x int) bool {
    var palindrome bool;
    
    return palindrome
}

func twoSum(nums []int, target int) []int {
    for i:= 0; i<len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j};
            }
			fmt.Println("total :", nums[i] + nums[j])
        }
    }
    return []int{};
}