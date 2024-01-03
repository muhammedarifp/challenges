// This is a problem one solution

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	fileNames = []string{
		"small_testcase1.txt",
		"small_testcase2.txt",
		"small_testcase3.txt",
		"small_testcase4.txt",
		"small_testcase5.txt",
		"large_testcase1.txt",
		"large_testcase2.txt",
		"large_testcase3.txt",
		"large_testcase4.txt",
		"large_testcase5.txt",
	}
)

func main() {
	for _, fileName := range fileNames {
		data, err := os.ReadFile("./" + fileName)
		if err != nil {
			fmt.Println("Error reading file : ", err)
			return
		}

		fileDatas := strings.Split(string(data), "\n")
		//var exp int
		var arr = []int{}
		var tab map[int][]int

		if err := json.Unmarshal([]byte(fileDatas[1]), &arr); err != nil {
			fmt.Println("Error unmarshalling array : ", err)
			return
		}
		if err := json.Unmarshal([]byte(fileDatas[2]), &tab); err != nil {
			fmt.Println("Error unmarshalling array : ", err)
			return
		}

		res := MaximumOccurences(arr, tab)
		fmt.Println("Result : ", res)
	}
}

func MaximumOccurences(arr []int, tab map[int][]int) int {
	maxCount := 0
	maxKey := -1

	// Create a set to store the elements of the input array
	arrSet := make(map[int]bool)
	for _, num := range arr {
		arrSet[num] = true
	}

	// Iterate over the keys of the map
	for key, val := range tab {

		count := 0
		for _, val := range val {
			// Check if the value is present in the set
			if arrSet[val] {
				count++
			}
		}

		// Update the maximum count and key
		if count > maxCount {
			maxCount = count
			maxKey = key
		}
	}

	return maxKey
}
