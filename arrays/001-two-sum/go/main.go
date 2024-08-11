package main

func twoSum(nums []int, target int) []int {
	// Initialize map to store the complement number as the key and its index as the value.
	hash := make(map[int]int)

	for i, num := range nums {

		// Check if the complement number exists in the map and return its index if it does.
		value, exists := hash[num]
		if exists {
			return []int{value, i}
		} else {
			// If the complement number doesn't exist, add the current index number to the map.
			complement := target - num
			hash[complement] = i
		}
	}

	return nil
}
