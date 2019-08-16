package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestLabelOf(t *testing.T) {
	data := LabelMap{"one", "two", "three"}
	assert.Equal(t, data.IndexOf("one"), 0)
	assert.Equal(t, data.IndexOf("four"), -1)
}

func TestFilter(t *testing.T) {
	data := LabelMap{"one", "two", "three"}
	assert.Equal(t, data.Filter([]int{0, 2}), LabelMap{"one", "three"})
}

func TestFilterVector(t *testing.T) {
	data := LabelMap{"one", "two", "three"}
	assert.Equal(t, data.FilterVector(mat.NewVecDense(3, []float64{1, 0, 1})), LabelMap{"one", "three"})
}
