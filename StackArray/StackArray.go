package stack_array

type StackArray struct {
	len    int
	maxlen int
	arr    []int
}

func New(arrLen int) *StackArray {

	ret := StackArray{0, arrLen, make([]int, arrLen)}
	return &ret
}

func (sa *StackArray) Pop() int {
	if sa.IsEmpty() {
		return -1
	}
	ret := sa.arr[sa.len-1]
	sa.len--
	return ret
}

func (sa *StackArray) Push(item int) {
	if sa.maxlen > sa.len {
		sa.arr[sa.len] = item
		sa.len++
	}
}

func (sa *StackArray) IsEmpty() bool {
	return sa.len == 0
}

func (sa *StackArray) Len() int {
	return sa.len
}
