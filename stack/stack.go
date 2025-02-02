package stack

type Stack struct {
	items []int
}

func (stk *Stack) Push(val int) {
	stk.items = append(stk.items, val)
}

func (stk *Stack) Pop() (int, bool) {
	if len(stk.items) == 0 {
		return -1, false
	}

	// Get the last index
	lastIndex := len(stk.items) - 1
	top := stk.items[lastIndex]

	// Remove the last element
	stk.items = stk.items[:lastIndex]

	return top, true
}

func (stk *Stack) Peek() (int, bool) {
	if len(stk.items) == 0 {
		return 0, false
	}
	return stk.items[len(stk.items)-1], true
}

func (stk *Stack) IsEmpty() bool {
	return len(stk.items) == 0
}

func (stk *Stack) Size() int {
	return len(stk.items)
}
