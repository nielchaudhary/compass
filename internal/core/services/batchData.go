package core

// will try to maximise its utility across the codebase, hence using any as data type
func BatchData[T any](data []T, size int) []T {
	batch := make([]T, size)
	for i := 0; i < len(data) && i < size; i++ {
		batch[i] = data[i]
	}

	return batch
}
