package stack_array

const (
	MAXARRLEN   = 2
	MAXSTACKLEN = 10
)

type innerArray struct {
	next     *innerArray
	len      int
	innerArr []int
}

func (ia *innerArray) Add(item int) {
	ia.innerArr[ia.len] = item
	ia.len++
}

func (ia *innerArray) Get() int {
	ret := ia.innerArr[ia.len-1]
	ia.len--
	return ret
}

type StackArrayUnlim struct {
	len int
	arr *innerArray
}

func NewUnlim() *StackArrayUnlim {
	arr := innerArray{nil, 0, make([]int, MAXARRLEN)}
	ret := StackArrayUnlim{0, &arr}

	return &ret
}

func (sau *StackArrayUnlim) Pop() int {
	if sau.IsEmpty() {
		return -1
	}
	if sau.arr.len == 0 {
		sau.arr = sau.arr.next
	}
	ret := sau.arr.Get()
	sau.len--
	return ret
}

func (sau *StackArrayUnlim) Push(item int) {
	if sau.arr.len == MAXARRLEN {
		newarr := innerArray{sau.arr, 0, make([]int, MAXARRLEN)}
		sau.arr = &newarr
	}
	sau.arr.Add(item)
	sau.len++
}

func (sau *StackArrayUnlim) IsEmpty() bool {
	return sau.len == 0
}

func (sau *StackArrayUnlim) Len() int {
	return sau.len
}
