package tool

import (
	"encoding/json"
	"fmt"
	"github.com/TestsLing/aj-captcha-go/model/vo"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestGetAllFile(t *testing.T) {
	var s []string
	s, _ = GetAllFile("D:\\goCode\\modeltools\\result\\aa", s)
	fmt.Println(s)

}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello World!")
}

func Test1(t *testing.T) {

	p, _ := ants.NewPool(10)
	defer p.Release()
	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Submit(syncCalculateSum)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	//p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
	//	myFunc(i)
	//	wg.Done()
	//})
	//defer p.Release()
	//// Submit tasks one by one.
	//for i := 0; i < runTimes; i++ {
	//	wg.Add(1)
	//	_ = p.Invoke(int32(i))
	//}
	//wg.Wait()
	//fmt.Printf("running goroutines: %d\n", p.Running())
	//fmt.Printf("finish all tasks, result is %d\n", sum)
}

func TestWriteFileAppend(t *testing.T) {
	filepathNew := filepath.Join("./taskProcessList", time.Now().Format("20060102"))
	if !IsExist(filepathNew) {
		CreateMutiDir(filepathNew)
	}
	files := filepath.Join(filepathNew, time.Now().Format("15")+".txt")
	content := fmt.Sprintf("%s %s %d\n", time.Now().Format("2006-01-02 15:04:05"), "xxxxxxxx", 12)
	WriteFileAppend(files, content)

}

func TestAbpath(t *testing.T) {
	//fmt.Println(CurrentAbPath())
	aa := `{"x":139,"y":5}`
	cachePoint := &vo.PointVO{}
	//userPoint := &vo.PointVO{}
	err := json.Unmarshal([]byte(aa), cachePoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cachePoint)
}

func CurrentAbPath() (dir string) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir, _ = filepath.EvalSymlinks(filepath.Dir(exePath))
	tempDir := os.Getenv("TEMP")
	if tempDir == "" {
		tempDir = os.Getenv("TMP")
	}
	tDir, _ := filepath.EvalSymlinks(tempDir)
	if strings.Contains(dir, tDir) {
		//return getCurrentAbPathByCaller()
		var abPath string
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abPath = path.Dir(filename)
		}
		return abPath
	}
	return dir
}

func TestGetAllFile2(t *testing.T) {
	files := make([]string, 0)
	files, _ = GetAllFile("./util", files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
		// 获取文件原来的访问时间，修改时间

		finfo, _ := os.Stat("./util/" + files[i])
		fmt.Println(finfo.ModTime())

		//// windows下代码如下
		//winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		//fmt.Println("文件创建时间：", SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9))

		//linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))

	}
	//fmt.Println(file)
}
