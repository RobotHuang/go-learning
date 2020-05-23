package taillog

// 收集日志模块

import (
	"context"
	"fmt"
	"go_learning/logagent/kafka"
	"os"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

// TailTask 一个日志收集任务
type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.Init()
	return
}

func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println(err)
	}
	go t.Run()
	return
}

/* func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
} */

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

func (t *TailTask) Run() {
	select {
	case <-t.ctx.Done():
		fmt.Printf("task: %s finish\n", t.path+t.topic)
		return
	case line := <-t.instance.Lines:
		kafka.SendToKafka(t.topic, line.Text)
		//kafka.SendToChan(t.topic, line.Text)
	}
}
