package main

import (
	"math/big"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
)

type memoryCell uint64

type memoryCellCollection []memoryCell

func memoryCellFromIntegerConstraint[T constraints.Integer](n T) memoryCell {
	return memoryCell(n)
}

func memoryCellFromBigInt(n *big.Int) (m memoryCell, err error) {
	if !n.IsUint64() {
		err = ErrMemoryCellConversionFailed
		return
	}

	m = memoryCellFromIntegerConstraint(n.Uint64())

	return
}

func (cell memoryCell) Uint64() uint64 {
	return uint64(cell)
}

func (cell memoryCell) String() string {
	return strconv.FormatUint(cell.Uint64(), 10)
}

func (collection memoryCellCollection) Len() int {
	return len(collection)
}

func (collection memoryCellCollection) Swap(i, j int) {
	collection[i], collection[j] = collection[j], collection[i]
}

func (collection memoryCellCollection) Less(i, j int) bool {
	return collection[i] < collection[j]
}

func (collection memoryCellCollection) SortAscending() {
	sort.Sort(collection)
}

func (collection memoryCellCollection) SortDescending() {
	collection.SortAscending()
	collection.Reverse()
}

func (collection memoryCellCollection) Reverse() {
	for i, j := 0, len(collection)-1; i < j; i, j = i+1, j-1 {
		collection[i], collection[j] = collection[j], collection[i]
	}
}
