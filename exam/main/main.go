package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cppdebug/windev"
)

func main() {

	for i := 0; i < 10; i++ {

		// 按Q键退出
		if windev.KeyDownUp(windev.VK_CHARQ) == 1 {
			fmt.Println("exit")
			os.Exit(-1)
		}

		// 设置鼠标的位置
		if windev.SetCursorPos(910, 300) {
			fmt.Println("set pos")
		}

		// 鼠标的滑轮事件, 页面向下滑动
		windev.MouseEvent(windev.MOUSEEVENTF_WHEEL, 0, 0, -100, 0)

		// 鼠标左键按下和松开事件
		windev.MouseEvent(windev.MOUSEEVENTF_LEFTDOWN|windev.MOUSEEVENTF_LEFTUP, 0, 0, 0, 0)

		time.Sleep(time.Second)
	}
}
