package matrix

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestHead(t *testing.T) {
	vector := NewLabelVector()
	assert.Equal(t, vector.HeadLabeles(2), []int{})
	assert.Equal(t, vector.HeadValues(2), []float64{})
	vector.Append(0, 1)
	assert.Equal(t, vector.HeadLabeles(2), []int{0})
	assert.Equal(t, vector.HeadValues(2), []float64{1})
	vector.Append(1, 10)
	assert.Equal(t, vector.HeadLabeles(2), []int{0, 1})
	assert.Equal(t, vector.HeadValues(2), []float64{1, 10})
}

func TestAppendVector(t *testing.T) {
	vector := NewLabelVector()
	vector.AppendVector(mat.NewVecDense(4, []float64{10, 0, 20, 0}))
	assert.Equal(t, vector.HeadLabeles(-1), []int{0, 2})
	assert.Equal(t, vector.HeadValues(-1), []float64{10, 20})
}

func TestVisit(t *testing.T) {
	vector := NewLabelVector()
	vector.AppendVector(mat.NewVecDense(2, []float64{20, 10}))
	vector.Visit(func(index int, value float64) {
		correct := (index == 0 && value == 20) || (index == 1 && value == 10)
		assert.True(t, correct)
	})
}

func TestSort(t *testing.T) {
	vector := NewLabelVector()
	vector.AppendVector(mat.NewVecDense(2, []float64{20, 10}))

	assert.Equal(t, vector.HeadValues(-1), []float64{20, 10})
	sort.Sort(vector)
	assert.Equal(t, vector.HeadValues(-1), []float64{10, 20})
}
