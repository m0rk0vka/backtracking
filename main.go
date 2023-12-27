package main

import (
	"fmt"
	"log"
	"strconv"
)

const nums string = "9876543210"

func FindSolution(target int) string {
	var result []string

	var backtrack func(string, string)

	backtrack = func(nums string, curr string) {
		if len(nums) == 0 {
			answer, err := Solve(curr)
			if err != nil {
				log.Fatalf("Error while solving curr example: %v", err.Error())
			}
			if answer == target {
				result = append(result, curr)
			}
			return
		}
		newNums := nums[1:]
		operand := string(nums[0])
		if len(curr) == 0 {
			backtrack(newNums, operand)
		} else {
			backtrack(newNums, curr+operand)
			if len(nums) > 1 {
				backtrack(newNums, curr+"+"+operand)
				backtrack(newNums, curr+"-"+operand)
			}
		}
	}

	backtrack(nums, "")

	//for i := 0; i < len(result); i++ {
	//	fmt.Println(result[i])
	//}

	if len(result) == 0 {
		return ""
	}

	return result[0]
}

func Solve(example string) (int, error) {
	n := len(example)
	answer := 0
	right := n - 1
	left := n - 1
	for left >= 0 {
		switch string(example[left]) {
		case "+":
			operand, err := strconv.Atoi(example[left+1 : right+1])
			if err != nil {
				return 0, fmt.Errorf(`Error while strconv in case "+": %v`, err.Error())
			}
			answer += operand
			left -= 1
			right = left
		case "-":
			operand, err := strconv.Atoi(example[left+1 : right+1])
			if err != nil {
				return 0, fmt.Errorf(`Error while strconv in case "-": %v`, err.Error())
			}
			answer -= operand
			left -= 1
			right = left
		}
		left -= 1
	}
	firstOperand, err := strconv.Atoi(example[:right+1])
	if err != nil {
		return 0, fmt.Errorf(`Error while strconv first operand: %v`, err.Error())
	}

	answer += firstOperand

	//fmt.Printf("Example: %v, answer: %v\n", example, answer)

	return answer, nil
}

func main() {
	solution := FindSolution(200)
	fmt.Println(solution)
}
