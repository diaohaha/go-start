package channel_examples

import (
	"sync"
	"time"
)

// 使用channel实现生产者消费者任务队列

// 任务结构
type Task struct {
	TaskDataItem1 string
	TaskDataItem2 string
	Retry         int
}

type Worker struct {
	TaskQueue chan Task
	LimitChan chan int
	//mu sync.RWMutex
	Wg sync.WaitGroup
}

// consumer 函数
func (w *Worker) Consumer() {
	for true {
		task := <-w.TaskQueue
		w.LimitChan <- 1
		//w.Wg.Add(1)
		go func() {
			defer func() {
				<-w.LimitChan
			}()
			// 执行任务
			succ := false
			// 任务重试
			if task.Retry < 3 && !succ {
				task.Retry += 1
				w.TaskQueue <- task
			}
			//w.Wg.Done()
		}()
	}
}

// producer 函数
func (w *Worker) Producer() {
	// 可能是一个文件一次输出
	// 可能是定时任务定时产生
	ticker := time.NewTicker(time.Minute * 1)

	// 每天6-7点每分钟触发一次任务
	for range ticker.C {
		if time.Now().Hour() == 6 {
			task := Task{
				Retry:         0,
				TaskDataItem1: "params1",
				TaskDataItem2: "params2",
			}
			w.TaskQueue <- task
		}
	}
}

// 启动函数
func (w *Worker) Run() {
	// 任务队列大小
	w.TaskQueue = make(chan Task, 2000)
	// 并行协程大小
	w.LimitChan = make(chan int, 10)
	go w.Consumer()
	go w.Producer()
}

var GWorker Worker

func InitWorker() {
	GWorker.Run()
}
