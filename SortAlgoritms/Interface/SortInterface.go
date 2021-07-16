package Interface

type Sortable interface {
	Sort() []int
	SetData(unsortedArray []int)
	GetData() []int
}
