package recommendation

import (
	"sort"

	"gonum.org/v1/gonum/mat"

	"github.com/timdrijvers/recommendation/matrix"
	"github.com/yourbasic/bit"
)

// Similarities maps a item to a list of similar items
type Similarities []*bit.Set

func intSliceBuilder(set *bit.Set) []int {
	ret := []int{}
	if set == nil {
		return ret
	}
	set.Visit(func(n int) bool {
		ret = append(ret, n)
		return false
	})
	return ret
}

// NewSimilarities creates a new empty similaritites strucuture
func NewSimilarities(entries int) Similarities {
	return make(Similarities, entries)
}

// NewTopSimilaritiesFromMatrix creates a filled similaritites structure with max top items per item
func NewTopSimilaritiesFromMatrix(m *matrix.CosineLabeledMatrix, top int) Similarities {
	width := m.Width()
	itemItemSimilarities := NewSimilarities(width)
	for row := 0; row < width; row++ {
		indexRowVector := matrix.NewLabelVector()
		indexRowVector.AppendVector(m.RowView(row))
		sort.Sort(sort.Reverse(indexRowVector))
		itemItemSimilarities.Set(row, indexRowVector.HeadLabeles(top))
	}
	return itemItemSimilarities
}

// Set a list of entries at index
func (m Similarities) Set(index int, entries []int) {
	m[index] = bit.New(entries...)
}

// Get a list of entries at index
func (m Similarities) Get(index int) []int {
	return intSliceBuilder(m[index])
}

// GetAll returns a merged list of entries at indexes
func (m Similarities) GetAll(indexes []int) []int {
	var mergedSet *bit.Set
	for _, index := range indexes {
		if m[index] == nil {
			continue
		}
		if mergedSet == nil {
			mergedSet = m[index]
		} else {
			mergedSet.SetOr(mergedSet, m[index])
		}
	}

	return intSliceBuilder(mergedSet)
}

// GetAllVector returns a merged list of entries, for items in where vector is set
func (m Similarities) GetAllVector(vector mat.Vector) []int {
	var mergedSet *bit.Set
	for index := 0; index < vector.Len(); index++ {
		if vector.AtVec(index) == 0 {
			continue
		}
		if m[index] == nil {
			continue
		}
		if mergedSet == nil {
			mergedSet = m[index]
		} else {
			mergedSet.SetOr(mergedSet, m[index])
		}
	}

	return intSliceBuilder(mergedSet)
}

func contains(slice []int, value int) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}

// ScoredSimilar creates a vector of recommended items based on a binary vector of already liked items
func (m Similarities) ScoredSimilar(cosineMatrix *mat.Dense, vector mat.Vector) matrix.LabeledVector {
	similarLikes := m.GetAllVector(vector)
	scoreVector := matrix.NewLabelVector()

	for _, item := range similarLikes {
		// Skip already liked
		if vector.AtVec(item) > 0 {
			continue
		}

		itemScore := float64(0)
		itemSum := float64(0)
		row := cosineMatrix.RawRowView(item)
		for j, cosine := range row {
			if contains(similarLikes, j) {
				itemScore += cosine * vector.AtVec(j)
				itemSum += cosine
			}
		}
		scoreVector.Append(item, itemScore/itemSum)
	}

	sort.Sort(sort.Reverse(scoreVector))
	return scoreVector
}
