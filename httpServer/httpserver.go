package httpserver

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func StartServer() {
	// http.HandleFunc("/", index)
	// http.ListenAndServe(":10001", nil)
	start := time.Now()     //记录开始时间
	ch := make(chan string) //创建一个字符信道

	for _, arg := range os.Args[1:] {
		//根据参数开启协程
		if strings.HasPrefix(arg, "http://") {
			go new_routine(arg, ch)
		} else {
			arg = "http://" + arg
			go new_routine(arg, ch)
		}
	}
	//读取信道数据
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	end := time.Since(start).Seconds() //结束时间
	fmt.Printf("time used :%.2fs\n", end)
}

func get_txt(arg string) {
	res, err := http.Get(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out, err := os.Create("/tmp/buf_file2.txt")
	// 初始化一个 io.Writer
	wt := bufio.NewWriter(out)
	result, err := io.Copy(wt, res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
	wt.Flush()
}

func get_code(arg string) {
	res, err := http.Get(arg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	status := res.Status
	fmt.Println("Http Code :", status)
}
func new_routine(arg string, ch chan<- string) {
	start_time := time.Now()
	res, err := http.Get(arg)
	if err != nil {
		ch <- fmt.Sprintf("err:%v", err)
		return
	}
	size_bytes, err := io.Copy(io.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("reading usl:%v", err)
		return
	}
	end_time := time.Since(start_time).Seconds()
	ch <- fmt.Sprintf("%.2fs %10d %s", end_time, size_bytes, arg)
}
