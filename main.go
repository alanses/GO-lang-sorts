package main

import (
	Configs "Sortings/Configs"
	generators "Sortings/Generators"
	SortAlgorithms "Sortings/SortAlgoritms"
	SortAlgorithmsInterfaces "Sortings/SortAlgoritms/Interface"
	"fmt"
)

var (
	ConfigSortingGenerator *Configs.ConfigSortingGenerator
)

func init() {
	ConfigSortingGenerator = new(Configs.ConfigSortingGenerator)

	ConfigSortingGenerator.Size = 10
	ConfigSortingGenerator.MaxRandomValue = 1000
}

func main() {

	randomArray := generators.GenerateRandomArray(
		ConfigSortingGenerator.Size,
		ConfigSortingGenerator.MaxRandomValue,
	)

	sortAlgorithm := new(SortAlgorithms.BubbleSortAlgorithm)

	sortedArray := sort(randomArray, sortAlgorithm)

	fmt.Println(sortedArray)
}

func sort(randomArray []int, implementationObject SortAlgorithmsInterfaces.Sortable) []int {
	implementationObject.SetData(randomArray)
	return implementationObject.Sort()
}
