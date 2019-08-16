package matrix

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// CosineMatrix is an item x item matrix where the cells contain
// the cosine similarity between the two items
type CosineMatrix struct {
	*mat.Dense
}

// CosineLabeledMatrix is a CosineMatrix including a label -> row/col index mapping
type CosineLabeledMatrix struct {
	*mat.Dense
	Label LabelMap
}

// NewCosineMatrix creates a new cosine similarity matrix of size col x col
func NewCosineMatrix(matrix *mat.Dense) *CosineMatrix {
	return &CosineMatrix{cosineMatrix(matrix)}
}

// NewCosineLabeledMatrix creates a new indexed cosine similarity matrix of size col x col
func NewCosineLabeledMatrix(matrix *LabeledMatrix) *CosineLabeledMatrix {
	return &CosineLabeledMatrix{
		cosineMatrix(matrix.Dense),
		matrix.ColumnLabel}
}

// Width returns the width of the matrix (matrix is size width x width)
func (matrix *CosineLabeledMatrix) Width() int {
	return len(matrix.Label)
}

func cosineMatrix(matrix *mat.Dense) *mat.Dense {
	rows, columns := matrix.Dims()

	// L2 normalize, loop over the raw data
	for i := 0; i < rows; i++ {
		row := matrix.RawRowView(i)
		length := hypotSlice(row)
		for j := range row {
			row[j] = row[j] / length
		}
	}

	// L2 normalize
	matrixT := mat.DenseCopyOf(matrix.T())
	rows, columns = columns, rows
	for row := 0; row < rows; row++ {
		length := hypotVector(matrixT.RowView(row))
		for col := 0; col < columns; col++ {
			matrixT.Set(row, col, matrixT.At(row, col)/length)
		}
	}
	var product mat.Dense
	product.Mul(matrixT, matrixT.T())
	return &product
}

//Square root of sum of squares (hypotenuse)
func hypotVector(vector mat.Vector) (res float64) {
	for i := 0; i < vector.Len(); i++ {
		v := vector.AtVec(i)
		res += v * v
	}
	return math.Sqrt(res)
}

//Square root of sum of squares (hypotenuse)
func hypotSlice(vector []float64) (res float64) {
	for _, v := range vector {
		res += v * v
	}
	return math.Sqrt(res)
}
