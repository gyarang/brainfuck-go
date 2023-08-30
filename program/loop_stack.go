package program

import "fmt"

type LoopStack struct {
	data []int
}

func (l *LoopStack) Push(i int) {
	l.data = append(l.data, i)
}

func (l *LoopStack) Pop() (int, error) {
	if l.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	top := len(l.data) - 1
	data := l.data[top]
	l.data = l.data[:top]
	return data, nil
}

func (l *LoopStack) IsEmpty() bool {
	return len(l.data) == 0
}
