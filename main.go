package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	lua "github.com/yuin/gopher-lua"
	"time"
)

var (
	minSleep = 0
)

func main() {
	doLua()
}

func doLua() {
	L := lua.NewState()
	defer L.Close()

	//方法注册
	L.SetGlobal("click", L.NewFunction(click))
	L.SetGlobal("input", L.NewFunction(input))
	L.SetGlobal("keyTap", L.NewFunction(keyTap))
	L.SetGlobal("keyTaps", L.NewFunction(keyTaps))
	L.SetGlobal("right", L.NewFunction(right))
	L.SetGlobal("sleep", L.NewFunction(sleep))
	L.SetGlobal("getRgb", L.NewFunction(getRGB))
	L.SetGlobal("hasRgb", L.NewFunction(hasRGB))
	L.SetGlobal("move", L.NewFunction(move))
	L.SetGlobal("addEvent", L.NewFunction(addEvent))
	L.SetGlobal("eventSleep", L.NewFunction(eventSleep))

	err := L.DoFile("script.lua")
	if err != nil {
		fmt.Println(err.Error())
	}
}

//每次操作睡眠时间
func eventSleep(L *lua.LState) int {
	minSleep = L.ToInt(1)
	return 0
}

//单击
func click(L *lua.LState) int {
	eventSleeps()
	x := L.ToInt(1)
	y := L.ToInt(2)
	double := L.ToBool(3)
	robotgo.MoveClick(x, y, "left", double)
	return 0
}

//输入
func input(L *lua.LState) int {
	eventSleeps()
	robotgo.TypeStr(L.ToString(1))
	return 0
}

//键盘
func keyTap(L *lua.LState) int {
	eventSleeps()
	if err := robotgo.KeyTap(L.ToString(1)); err != nil {
		fmt.Println("按键错误：", err.Error())
	}
	return 0
}

func right(L *lua.LState) int {
	eventSleeps()
	x := L.ToInt(1)
	y := L.ToInt(2)
	double := L.ToBool(3)
	robotgo.MoveClick(x, y, "right", double)
	return 0
}

//睡眠
func sleep(L *lua.LState) int {
	if L.ToBool(2) {
		robotgo.MilliSleep(L.ToInt(1))
	} else {
		robotgo.Sleep(L.ToInt(1))
	}
	return 0
}

//获取坐标点颜色
func getRGB(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	rgb := robotgo.GetPixelColor(x, y)
	L.Push(lua.LString(rgb))
	return 1
}

//判断坐标点是否为指定颜色
func hasRGB(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	color := L.ToString(3)
	timeout := L.ToInt(4)
	step := L.ToInt(5)
	if color == "" { //获取指定颜色
		L.Push(lua.LBool(false))
		return 1
	}
	if timeout == 0 {
		has := robotgo.GetPixelColor(x, y) == color
		L.Push(lua.LBool(has))
		return 1
	}
	if step == 0 {
		step = 50
	}

	tc := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		tc <- false
	}()

	go func() {
		for {
			rgb := robotgo.GetPixelColor(x, y)
			if rgb == color {
				tc <- true
				return
			}
			time.Sleep(time.Duration(step) * time.Millisecond)
		}
	}()

	L.Push(lua.LBool(<-tc))
	return 1

}

//鼠标移动
func move(L *lua.LState) int {
	eventSleeps()
	x := L.ToInt(1)
	y := L.ToInt(2)
	robotgo.Move(x, y)
	return 0
}

//添加监听事件
func addEvent(L *lua.LState) int {
	eventSleeps()
	event := L.ToString(1)
	hook.AddEvent(event)
	return 0
}

//键盘输入
func keyTaps(L *lua.LState) int {
	eventSleeps()
	taps := L.ToTable(1)
	key := ""
	var keys []interface{}
	taps.ForEach(func(value lua.LValue, value2 lua.LValue) {
		if value.String() == "1" {
			key = value2.String()
		} else {
			keys = append(keys, value2.String())
		}
	})
	robotgo.KeyTap(key, keys...)
	return 0
}

func eventSleeps() {
	if minSleep > 0 {
		robotgo.MilliSleep(minSleep)
	}
}
