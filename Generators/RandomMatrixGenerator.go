package generators

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(size int, maxValueForGeneration int) []int {
	rand.Seed(time.Now().UnixNano())

	matrix := make([]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = rand.Intn(maxValueForGeneration)
	}

	return matrix
}
