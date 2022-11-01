package minstack

type MinStack struct {
	Stack []int
	Min   []int
}

func (ms *MinStack) Top() int {
	return ms.Stack[len(ms.Stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.Min[len(ms.Min)-1]
}

func (ms *MinStack) Push(value int) {
	ms.Stack = append(ms.Stack, value)

	minSize := len(ms.Min)
	if minSize > 0 {
		if ms.GetMin() > value {
			ms.Min = append(ms.Min, value)
		}
	} else {
		ms.Min = append(ms.Min, value)
	}
}

func (ms *MinStack) Pop() {
	stackSize := len(ms.Stack)
	minSize := len(ms.Min)

	if ms.Top() == ms.GetMin() {
		ms.Min = ms.Min[:minSize-1]
	}

	ms.Stack = ms.Stack[:stackSize-1]
}
