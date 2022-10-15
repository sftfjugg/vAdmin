package deploy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func AsyncLog(reader io.ReadCloser, ResultChan chan string, ExitChan chan int)error {
	cache := "" //缓存不足一行的日志信息
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err!=io.EOF{
			return err
		}
		if err == io.EOF {
			ExitChan <- 1  // 记录协程管斌
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
			fmt.Printf("%s%s\n", cache, line)
			ResultChan <- line
			cache = s[len(s)-1]
		}
	}
	return nil
}

func Execute(CMD string,ResultExecChan chan string, ExitExecChan chan int) error {
	var (
		ResultChan = make(chan string, 1000)
		ExitChan   = make(chan int)
	)
	fmt.Println("CMD:",CMD)
	//cmd := exec.Command("sh", "-c", CMD)
	cmd := exec.Command("cmd", "/C", "dir D: /s/b")
	// 命令的错误输出和标准输出都连接到同一个管道
	// stdout, err := cmd.StdoutPipe()
	// cmd.Stderr = cmd.Stdout

	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	//stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	go AsyncLog(stdout,ResultChan,ExitChan)
	//go AsyncLog(stderr,ResultChan,ExitChan)

	go func() {
		for i := 0; i < 1; i++ {  // 从exitchan管道中获取到1次goroute完成完毕记录才放行
			a := <-ExitChan // 没有记录的话会堵塞，会等下次的记录被插入才会放行
			if a == 1 {
				fmt.Println("AsyncLog goroute结束")
			}
		}
		close(ResultChan)  // 1次循环结束代表1个收集结果的goroute全部完成完毕了，此时才能关闭resultchan，为的是让下面for循环取完resultchan管道中值的时可以正常退出
	}()

	for v := range ResultChan {
		//fmt.Println(v)
		ResultExecChan <- v
	}

	ExitExecChan <- 1  // 记录协程管斌

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
		return err
	}
	return nil
}

func ExecuteToLog(CMD string, Filename string, ExitExecChan chan int) error {
	// open the out file for writing
	outfile, err := os.Create(Filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	fmt.Println("CMD:",CMD)
	sysType := runtime.GOOS
	switch sysType {
		case "linux":
			cmd := exec.Command("sh", "-c", CMD)
			stdout, _ := cmd.StdoutPipe()
			cmd.Stderr = cmd.Stdout
			//stderr, _ := cmd.StderrPipe()

			if err := cmd.Start(); err != nil {
				log.Printf("Error starting command: %s......", err.Error())
				return err
			}
			writer := bufio.NewWriter(outfile)
			go io.Copy(writer, stdout)

			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for command execution: %s......", err.Error())
				return err
			}
		case "windows":
			//cmd := exec.Command("cmd", "/C", "del", "D:\\a.txt")
			cmd := exec.Command("cmd", "/C", "dir D: /s/b")
			stdout, _ := cmd.StdoutPipe()
			cmd.Stderr = cmd.Stdout
			//stderr, _ := cmd.StderrPipe()

			if err := cmd.Start(); err != nil {
				log.Printf("Error starting command: %s......", err.Error())
				return err
			}
			writer := bufio.NewWriter(outfile)
			go io.Copy(writer, stdout)

			if err := cmd.Wait(); err != nil {
				log.Printf("Error waiting for command execution: %s......", err.Error())
				return err
			}
	}
	// 命令的错误输出和标准输出都连接到同一个管道
	// stdout, err := cmd.StdoutPipe()
	// cmd.Stderr = cmd.Stdout

	ExitExecChan <- 1  // 记录协程管斌
	return nil
}