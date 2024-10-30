package utils

// PipeCloser类型定义，是一个无参数且返回值为布尔值的函数类型
type PipeCloser func() bool

// AssignPipeCloser函数用于给PipeCloser类型赋值，接收一个通道作为参数并返回一个判断管道是否关闭的函数
func AssignPipeCloser(ch chan string) PipeCloser {
	return func() bool {
		// 使用select语句实现非阻塞接收操作来判断管道状态
		select {
		case _, ok := <-ch:
			return !ok
		default:
			// 通道无数据且未关闭时，返回当前认为管道未关闭的状态（这里假设未收到明确关闭信号就认为未关闭）
			return false
		}
	}
}
