package matrix

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	matrix := LabelMatrixFromCSV(strings.NewReader(
		`row_id,one,two,three
Eindhoven,1,2,3
Amsterdam,4,5,6`))
	assert.Equal(t, matrix.ColumnLabel, LabelMap{"one", "two", "three"})
	assert.Equal(t, matrix.RowLabel, LabelMap{"Eindhoven", "Amsterdam"})
	assert.Equal(t, matrix.RawMatrix().Data, []float64{1, 2, 3, 4, 5, 6})
}
