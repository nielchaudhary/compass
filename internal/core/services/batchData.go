package core

// will try to maximise its utility across the codebase, hence using any as data type
func BatchData[T any](data []T) [20]T {
	var batch [20]T
	for i := 0; i < len(data) && i < 20; i++ {
		batch[i] = data[i]
	}

	return batch
}
