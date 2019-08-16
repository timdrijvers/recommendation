package main

import (
	"fmt"
	"os"

	"github.com/timdrijvers/recommendation/matrix"
	"github.com/timdrijvers/recommendation/recommendation"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintln(os.Stderr, "example FILE ROW")
		fmt.Fprintln(os.Stderr, "FILE a csv file with header and first column labels")
		fmt.Fprintln(os.Stderr, "ROW a label from the first column (row label)")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	baseMatrix := matrix.LabelMatrixFromCSV(file)
	rowLabel := baseMatrix.RowLabel.IndexOf(os.Args[2])
	if rowLabel == -1 {
		fmt.Fprintf(os.Stderr, "Can not find row %s\n", os.Args[2])
		os.Exit(2)
	}

	cosineMatrix := matrix.NewCosineLabeledMatrix(baseMatrix)

	// Create recommendation/similarity matrix, containing top 10 items recommendation for a item
	itemItemSimilarities := recommendation.NewTopSimilaritiesFromMatrix(cosineMatrix, 10)

	// user - item recommendation
	userLikes := baseMatrix.RowView(rowLabel)
	fmt.Println("user likes", baseMatrix.ColumnLabel.FilterVector(userLikes))

	fmt.Println("Ranked recommended ")
	recommend := itemItemSimilarities.ScoredSimilar(cosineMatrix.Dense, userLikes)
	recommend.Visit(func(index int, value float64) {
		fmt.Println(baseMatrix.ColumnLabel[index], value)
	})
}
