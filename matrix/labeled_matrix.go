package matrix

import (
	"gonum.org/v1/gonum/mat"

	"encoding/csv"
	"io"
	"strconv"
)

// LabeledMatrix is a dense matrix where columns and rows are indexed using label -> index mapping
type LabeledMatrix struct {
	*mat.Dense
	ColumnLabel LabelMap
	RowLabel    LabelMap
}

// NewLabeledMatrix creates a new Dense labeled matrix
func NewLabeledMatrix(data []float64, columnLabel LabelMap, rowLabel LabelMap) *LabeledMatrix {
	if len(columnLabel)*len(rowLabel) != len(data) {
		panic("data should be length len(columnLabel) * len(rowLabel)")
	}
	return &LabeledMatrix{mat.NewDense(len(rowLabel), len(columnLabel), data), columnLabel, rowLabel}
}

// LabelMatrixFromCSV creates a new LabeledMatrix from a CSV file
// The file needs to contain both a header and a first column containing unique labels
func LabelMatrixFromCSV(r io.Reader) *LabeledMatrix {
	reader := csv.NewReader(r)

	rowsLabel := LabelMap{}
	columnsLabel := readHeader(reader)

	matrixData := []float64{}
	columnsLen := len(columnsLabel)

	for {
		rawRecord, readErr := reader.Read()
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			panic(readErr)
		}
		rowsLabel = append(rowsLabel, rawRecord[0])
		record := make([]float64, columnsLen)
		for i, v := range rawRecord[1:] {
			record[i], _ = strconv.ParseFloat(v, 64)
		}
		matrixData = append(matrixData, record...)
	}

	return NewLabeledMatrix(matrixData, columnsLabel, rowsLabel)
}

func readHeader(reader *csv.Reader) LabelMap {
	header, readErr := reader.Read()
	if readErr != nil {
		panic(readErr)
	}
	return header[1:]
}
