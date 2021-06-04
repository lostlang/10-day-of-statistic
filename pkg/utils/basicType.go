package utils

type IntSlice64 struct {
	data []int64
}

func (s IntSlice64) Len() int {
	return len(s.data)
}

func (s IntSlice64) Less(i, j int) bool {
	return s.data[i] < s.data[j]
}

func (s IntSlice64) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s IntSlice64) SetData(arr []int64) {
	s.data = arr
}

func (s IntSlice64) GetData() []int64 {
	return s.data
}