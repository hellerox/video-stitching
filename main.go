package main

import "fmt"

func main() {

	clips := [][]int{
		{0, 1},
		{0, 2},
		{0, 1},
		{1, 4},
		{4, 6},
		{8, 10},
		{1, 9},
		{1, 5},
		{5, 9},
	}

	videos := videoStitching(clips, 10)
	fmt.Print(videos)
}

func videoStitching(clips [][]int, T int) int {

	result := bubbleSort(clips)

	fmt.Println("resultado: ", result)

	bound, nextBound := 0, 0
	clipCount := 1

	for _, video := range result {
		fmt.Println(video, bound, nextBound, clipCount)
		if video[0] <= bound {
			if video[1] > nextBound {
				nextBound = video[1]
			}
		} else if video[0] <= nextBound {
			clipCount++
			bound = nextBound

			if video[1] > nextBound {
				nextBound = video[1]
			}
		} else {
			return -1
		}

		if nextBound >= T {
			return clipCount
		}
	}

	return -1

}

// function to sort number using Bubblesort Algorithm
func bubbleSort(numbers [][]int) [][]int {
	if numbers != nil && len(numbers) > 0 {
		for i := 0; i < len(numbers); i++ {
			for j := 0; j < len(numbers); j++ {
				if numbers[i][0] < numbers[j][0] {
					// swapping the numbers
					temp := numbers[i]
					numbers[i] = numbers[j]
					numbers[j] = temp
				}
			}
		}
	} else {
		fmt.Println("Empty or Null Array")
	}
	return numbers
}
