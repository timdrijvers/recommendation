Recommendation
[![Go Report Card](https://goreportcard.com/badge/github.com/timdrijvers/recommendation)](https://goreportcard.com/report/github.com/timdrijvers/recommendation)
=========================
A simple item-item and user-item recommendation engine using cosine similarity. 

## General overview
As input a user x item binary matrix is used. `gonum` Matrixes can be used or a 

- First a item x item similarity matrix is generated, using `NewCosineLabeledMatrix`. Each cell contains the cosine similarity between the corresponding row and column item.
- Using this matrix a second matrix can be generated containing the top-N most similar items per item. Taking each row (or column) sorting these to get the most similar items. `NewTopSimilaritiesFromMatrix` will help you to do so.
- Now using this matrix, it's possible to implement item x item, items x items and user x items recommendation. item x item, is simply the first element from the top-N matrix, the latter can be achieved by using the `ScoredSimilar` function.

## Example
The `example` folder implements a basic user -> items recommendation system. The input is a labeled binary CSV, where labels are present on the first row and column. Each row contains one user, each column is an item, the values are binary (either 0 or 1).

## Future ideas
- Calculate a cut off value, don't simply take the top-N similar items, but the top most similar based on the cut off value.