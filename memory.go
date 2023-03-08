package main

type memoryCell uint64

type memoryCellCollection []memoryCell

func (collection memoryCellCollection) Len() int {
	return len(collection)
}

func (collection memoryCellCollection) Swap(i, j int) {
	collection[i], collection[j] = collection[j], collection[i]
}

func (collection memoryCellCollection) Less(i, j int) bool {
	return collection[i] < collection[j]
}
