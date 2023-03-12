package main

import (
	"math/big"

	"golang.org/x/exp/constraints"
)

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
