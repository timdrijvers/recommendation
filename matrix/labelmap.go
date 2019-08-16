package matrix

import "gonum.org/v1/gonum/mat"

// LabelMap is a slice of strings
type LabelMap []string

// IndexOf searches for label in the LabelMap
func (im *LabelMap) IndexOf(search string) int {
	for i, v := range *im {
		if v == search {
			return i
		}
	}
	return -1
}

// Filter creates a new LabelMap only returning items at indexes
func (im *LabelMap) Filter(indexes []int) LabelMap {
	ret := make(LabelMap, len(indexes))
	for i, value := range indexes {
		ret[i] = (*im)[value]
	}
	return ret
}

// FilterVector creates a new LabelMap only returning items at vector
func (im *LabelMap) FilterVector(vector mat.Vector) LabelMap {
	ret := make(LabelMap, 0)
	for i := 0; i < vector.Len(); i++ {
		if vector.AtVec(i) > 0 {
			ret = append(ret, (*im)[i])
		}
	}
	return ret
}
