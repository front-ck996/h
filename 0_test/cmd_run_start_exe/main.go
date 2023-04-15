package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/front-ck996/csy"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Getwd())
	fmt.Println(mahonia.NewDecoder("utf8").ConvertString("你好"))
	// 创建一个命令
	//cmd := exec.Command("cmd.exe", "/C", "C:\\Users\\Administrator\\Desktop\\xUltimate-d9pc-x86\\xUltimate-d9pc.exe")
	cmd := exec.Command("cmd.exe", "/C", "aifgabfife ")

	// 获取标准输入流
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("获取标准输入流时发生错误：", err)
		return
	}

	// 获取标准输出流
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("获取标准输出流时发生错误：", err)
		return
	}

	// 获取标准错误流
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("获取标准错误流时发生错误：", err)
		return
	}

	// 将标准输出流连接到当前程序的流
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			//fmt.Println("12345")
			//fmt.Println(scanner.Text())
			utf8, _ := csy.ConvCharsetToUtf8(scanner.Bytes())
			//csy.ConvCharsetToUtf8(scanner.Text(), "")
			fmt.Println(utf8)
		}
	}()

	// 将标准错误流连接到当前程序的流
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			toUtf8, _ := csy.GbkToUtf8(scanner.Bytes())
			fmt.Println(string(toUtf8))
		}
	}()

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("启动命令时发生错误：", err)
		return
	}

	// 读取用户输入并发送到标准输入流中
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintln(stdin, text)
		if err != nil {
			fmt.Println("向标准输入流发送数据时发生错误：", err)
			return
		}
	}

	// 关闭标准输入流
	err = stdin.Close()
	if err != nil {
		fmt.Println("关闭标准输入流时发生错误：", err)
		return
	}

	// 等待命令完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("等待命令完成时发生错误：", err)
		return
	}
}
