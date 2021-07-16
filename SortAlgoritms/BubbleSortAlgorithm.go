package SortAlgoritms

type BubbleSortAlgorithm struct {
	Data []int
}

func (bubbleSortAlgorithm *BubbleSortAlgorithm) SetData(unsortedArray []int) {
	bubbleSortAlgorithm.Data = unsortedArray
}

func (bubbleSortAlgorithm *BubbleSortAlgorithm) GetData() []int {
	return bubbleSortAlgorithm.Data
}

func (bubbleSortAlgorithm *BubbleSortAlgorithm) Sort() []int {
	unsortedArray := bubbleSortAlgorithm.Data

	for i := 0; i < len(unsortedArray)-1; i++ {
		for j := 0; j < len(unsortedArray)-i-1; j++ {
			if unsortedArray[j] > unsortedArray[j+1] {
				unsortedArray[j], unsortedArray[j+1] = unsortedArray[j+1], unsortedArray[j]
			}
		}
	}

	return unsortedArray
}
