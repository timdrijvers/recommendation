package matrix

import (
	"gonum.org/v1/gonum/mat"
)

type vectorItem struct {
	index int
	value float64
}

// LabeledVector is a vector that keeps a reference
// from each value to it's original index. Common usecase is
// to sort the vector and get get all top indexes
type LabeledVector []*vectorItem

func (v LabeledVector) Len() int           { return len(v) }
func (v LabeledVector) Less(i, j int) bool { return v[i].value < v[j].value }
func (v LabeledVector) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

// NewLabelVector creates an empty indexed vector.
func NewLabelVector() LabeledVector {
	return LabeledVector{}
}

// AppendVector appends all non zero items to this vector
func (v *LabeledVector) AppendVector(vector mat.Vector) {
	for i := 0; i < vector.Len(); i++ {
		value := vector.AtVec(i)
		if value > 0 {
			v.Append(i, value)
		}
	}
}

// Append adds a single entry in the vector with index, value
func (v *LabeledVector) Append(index int, value float64) {
	*v = append(*v, &vectorItem{index, value})
}

// HeadLabeles return length indexes, 0 returns all
func (v LabeledVector) HeadLabeles(length int) []int {
	var slice LabeledVector
	if len(v) <= length || length == -1 {
		slice = v
	} else {
		slice = v[:length]
	}
	result := make([]int, len(slice))
	for index, entry := range slice {
		result[index] = entry.index
	}
	return result
}

// HeadValues return length values, 0 returns all
func (v LabeledVector) HeadValues(length int) []float64 {
	var slice LabeledVector
	if len(v) <= length || length == -1 {
		slice = v
	} else {
		slice = v[:length]
	}
	result := make([]float64, len(slice))
	for index, entry := range slice {
		result[index] = entry.value
	}
	return result
}

// Visit each entry in the vector with a callback
func (v LabeledVector) Visit(do func(index int, value float64)) {
	for _, vector := range v {
		do(vector.index, vector.value)
	}
}
