package model

type MyContext struct {
	Task   Task
	Result chan any
	Error  error
}

func NewMyCtx(task Task) *MyContext {
	return &MyContext{
		Task:   task,
		Result: make(chan any, 1),
	}
}
