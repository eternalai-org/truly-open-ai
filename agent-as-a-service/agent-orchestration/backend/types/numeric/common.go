package numeric

func copyData(data []byte) []byte {
	newData := make([]byte, len(data))
	copy(newData, data)
	return newData
}
