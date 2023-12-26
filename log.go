package main

import (
	"fmt"
	"time"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Reset   = "\033[0m"
)

func P(a ...any) (n int, err error) {
	fmt.Print(White)
	currentTime := time.Now().Format("2023-01-01")
	n, err = fmt.Println(currentTime, a)
	fmt.Print(Reset) // 重置颜色
	return
}

func E(a ...any) (n int, err error) {
	fmt.Print(Red)
	currentTime := time.Now().Format("2023-01-01")
	n, err = fmt.Println(currentTime, a)
	fmt.Print(Reset) // 重置颜色
	return
}

func W(a ...any) (n int, err error) {
	fmt.Print(Yellow)
	currentTime := time.Now().Format("2023-01-01")
	n, err = fmt.Println(currentTime, a)
	fmt.Print(Reset) // 重置颜色
	return
}
