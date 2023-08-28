package global

import "time"

//type TaskType int8
//const (
//	TaskTypeDelay TaskType = iota
//	TaskTypeDefer
//	TaskTypePeriod
//)
// TaskTypeDelay => 多长时间后执行; TaskTypeDefer => 无效; TaskTypePeriod => 执行的周期

// DeferTaskQueue 延迟任务队列
var DeferTaskQueue []*DeferTask

type Task interface {
	Execute()
}

// DelayTask 延迟任务
type DelayTask struct {
	Function func(...any)
	Params   []any
	Time     time.Duration
}

// NewDelayTask 创建延迟任务
func NewDelayTask(f func(...any), t time.Duration, params []any) *DelayTask {
	return &DelayTask{
		Function: f,
		Params:   params,
		Time:     t,
	}
}

func (t *DelayTask) Execute() {
	timer := time.NewTimer(t.Time)
	<-timer.C
	t.Function(t.Params...)
}

// PeriodTask 周期任务
type PeriodTask struct {
	Function func(...any)
	Params   []any
	Time     time.Duration // TaskTypeDelay => 多长时间后执行; TaskTypeDefer => 无效; TaskTypePeriod => 执行的周期
}

// NewPeriodTask 创建周期任务
func NewPeriodTask(f func(...any), t time.Duration, params []any) *PeriodTask {
	return &PeriodTask{
		Function: f,
		Params:   params,
		Time:     t,
	}
}

func (t *PeriodTask) Execute() {
	ticker := time.NewTicker(t.Time)
	for range ticker.C {
		t.Function(t.Params...)
	}
}

// DeferTask 释放资源
type DeferTask struct {
	Function func(...any)
	Params   []any
}

// NewDeferTask 创建资源释放任务
func NewDeferTask(f func(...any), params ...any) *DeferTask {
	return &DeferTask{
		Function: f,
		Params:   params,
	}
}

func (t *DeferTask) Execute() {
	t.Function(t.Params...)
}
