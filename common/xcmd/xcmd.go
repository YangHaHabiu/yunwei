/*
@Time : 2021-11-10 10:27
@Author : acool
@File : xcmd
*/
package xcmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

type JobResult struct {
	OutMsg    string
	ErrMsg    string
	IsOk      bool
	IsTimeout bool
}

//本地执行命令
func NewCommandJob(timeout time.Duration, command string) *JobResult {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("CMD", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	//将pid和pgid设置相同的进程
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Start()
	err, isTimeout := runCmdWithTimeout(cmd, timeout)
	jobresult := new(JobResult)
	jobresult.OutMsg = bufOut.String()
	jobresult.ErrMsg = bufErr.String()
	jobresult.IsOk = true
	if err != nil {
		jobresult.ErrMsg = err.Error()
		jobresult.IsOk = false
	}
	jobresult.IsTimeout = isTimeout

	return jobresult
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

		if err = syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
			fmt.Printf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
		}
		// if err = cmd.Process.Kill(); err != nil {
		// 	fmt.Printf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
		// }
		return err, true
	case err = <-done:
		return err, false
	}
}
