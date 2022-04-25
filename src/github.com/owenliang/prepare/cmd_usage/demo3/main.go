package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	//执行一个command,让他在一个协程里去执行，让他执行两秒
	//一秒的时候，将它提前杀死
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// 创建一个结果队列
	resultChan = make(chan *result, 1000)

	// context:chan byte
	// cancelFunc:close(chan byte)

	// ctx用来感知 cancelFunc是取消功能
	ctx, cancelFunc = context.WithCancel(context.TODO())
	//将函数丢进协程
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2; echo hello;")
		//通道关闭 context里的select{case <- ctx.Done()}监听到通道被关了 就会kill bash程序

		//执行任务，捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果，传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	// 继续往下走
	time.Sleep(1 * time.Second)

	// 取消上下文,一旦调用cancelFunc,chan会关闭
	cancelFunc()

	// 在main协程里，等待子协程的退出，并打印任务执行结果，子协程通过chan与主协(main)程通信
	res = <-resultChan
	fmt.Println(res.err, string(res.output))

}
