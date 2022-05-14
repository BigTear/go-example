package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// 不同系统下换行符不同。。。
	// 其实可以直接简单暴力trim \n后trim\r
	// 还是判断一下系统吧
	switch runtime.GOOS {
	case "linux":
		input = strings.TrimSuffix(input, "\n")
	case "windows":
		input = strings.TrimSuffix(input, "\r\n")
	}

	// 分析参数
	args := strings.Split(input, " ")

	// 执行内置命令
	switch args[0] {
	case "cd":
		if len(args) != 2 {
			return errors.New("请输入一个有效的路径。")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
