package main

import "ginFrame/command"

func main() {
	//ginFrame.New() // 直接启动
	command.Init() // 通过命令行参数启动
}
