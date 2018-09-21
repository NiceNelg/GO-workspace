package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window
	buttonMin   *gtk.Button
	buttonClose *gtk.Button
}

//控件属性结构体
type ChessInfo struct {
	w, h int
	x, y int
}

//黑白棋结构体
type Chessboard struct {
	ChessWidget
	ChessInfo
}

func (obj *Chessboard) CreateWindow() {
	//读取glade
	builder := gtk.NewBuilder()
	builder.AddFromFile("./ui.glade")
	//窗口相关
	obj.window = gtk.WindowFromObject(builder.GetObject("window"))
	//允许绘图
	obj.window.SetAppPaintable(true)
	//居中显示
	obj.window.SetPosition(gtk.WIN_POS_CENTER)
	obj.w, obj.h = 800, 480
	obj.window.SetSizeRequest(obj.w, obj.h)
	obj.window.SetDecorated(false)
	//设置事件
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
}

func MousePressEvent(ctx *glib.CallbackContext) {
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard)
	if ok == false {
		fmt.Println("MousePressEvent Chessboard error")
		return
	}
	//保存点击的x,y坐标
	obj.x, obj.y = int(event.X), int(event.Y)
}

func MouseMoveEvent(ctx *glib.CallbackContext) {
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard)
	if ok == false {
		fmt.Println("MousePressEvent Chessboard error")
		return
	}

	x, y := int(event.XRoot)-obj.x, int(event.YRoot)-obj.y
	obj.window.Move(x, y)
}

func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标移动事件
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)
}

func main() {
	//初始化GTK库
	gtk.Init(&os.Args)

	var obj Chessboard
	obj.CreateWindow()
	obj.HandleSignal()
	obj.window.Show()

	//运行主事件
	gtk.Main()
}
