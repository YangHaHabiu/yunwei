/**
* @Author: Aku
* @Description:
* @Email: 271738303@qq.com
* @File: init.
* @Date: 2021-6-28 16:54
 */
package jobs

import (
	"fmt"
	"os/exec"
	"sync"
	"syscall"
	"time"
	"ywadmin-v3/service/yunwei/rpc/internal/svc"
)

var (
	chLimit chan bool
	lock    sync.Mutex
	wg      sync.WaitGroup
	stopCh  chan int64
	timeOut time.Duration
)

func init() {

	chLimit = make(chan bool, 10)
	stopCh = make(chan int64, 1)
	timeOut = 160
}

// 系统启动循环执行
func InitJobs(svx *svc.ServiceContext) {
	//异步开启
	for {
		rangeRand := RangeRand(3, 10)
		time.Sleep(time.Duration(rangeRand) * time.Second)
		go NewJobFromTask(svx)
	}

}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		fmt.Println(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		// if err = cmd.Process.Kill(); err != nil {
		// 	fmt.Printf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
		// }

		if err = syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
			fmt.Printf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
