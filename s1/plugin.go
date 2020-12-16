package plugin

type Task struct {
	TaskID   string // 任务ID
	TaskBody []byte // json 不同plugin 处理JSON结构不一
}

type Plugin interface {
	Register() string           // 返回当前插件名称
	Assignment(task Task) error // 下发任务
}
