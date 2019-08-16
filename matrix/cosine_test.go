package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCosineMatrixBasics(t *testing.T) {
	base := NewLabeledMatrix([]float64{1, 0, 0, 1}, LabelMap{"col1", "col2"}, LabelMap{"row1", "row2"})
	cosine := NewCosineLabeledMatrix(base)

	assert.Equal(t, 2, cosine.Width())
	assert.Equal(t, []float64{1, 0, 0, 1}, cosine.Dense.RawMatrix().Data)
	assert.Equal(t, LabelMap{"col1", "col2"}, cosine.Label)
}

func TestCosineMatrix(t *testing.T) {
	base := NewLabeledMatrix([]float64{4, 2, 2, 3}, LabelMap{"col1", "col2"}, LabelMap{"row1", "row2"})
	cosine := NewCosineLabeledMatrix(base)

	assert.InDeltaSlice(t, []float64{1, 0.866, 0.866, 1}, cosine.Dense.RawMatrix().Data, 0.001)
}
