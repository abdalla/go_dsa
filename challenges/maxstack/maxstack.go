package maxstack

type MaxStack struct {
	Stack []int
	Max   []int
}

func (ms *MaxStack) Top() int {
	return ms.Stack[len(ms.Stack)-1]
}

func (ms *MaxStack) GetMax() int {
	return ms.Max[len(ms.Max)-1]
}

func (ms *MaxStack) Push(value int) {
	ms.Stack = append(ms.Stack, value)

	maxSize := len(ms.Max)
	if maxSize > 0 {
		if ms.GetMax() < value {
			ms.Max = append(ms.Max, value)
		}
	} else {
		ms.Max = append(ms.Max, value)
	}
}

func (ms *MaxStack) Pop() {
	stackSize := len(ms.Stack)
	maxSize := len(ms.Max)

	if ms.Top() == ms.GetMax() {
		ms.Max = ms.Max[:maxSize-1]
	}

	ms.Stack = ms.Stack[:stackSize-1]
}
