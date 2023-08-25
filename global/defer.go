package global

type DeferFuncListT []func()

var DeferFuncList DeferFuncListT

// Push 向延迟执行函数列表中添加函数
func (d *DeferFuncListT) Push(f ...func()) {
	DeferFuncList = append(DeferFuncList, f...)
}

// Run 运行延迟执行函数列表中的所有函数
func (d *DeferFuncListT) Run() {
	for _, v := range DeferFuncList {
		go func(v func()) {
			v()
		}(v)
	}
}
