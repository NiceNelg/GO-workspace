package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window
	buttonMin   *gtk.Button
	buttonClose *gtk.Button
	lableBlack  *gtk.Label
	lableWhite  *gtk.Label
	lableTime   *gtk.Label
	imgBlack    *gtk.Image
	imgWhite    *gtk.Image
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

func ButtonSetImgFromFile(butto *gtk.Button, filename string) {
	//获取按钮大小
	w, h := 0, 0
	butto.GetSizeRequest(&w, &h)

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)
	//创建image
	image := gtk.NewImageFromPixbuf(pixbuf)
	//释放pixbuf
	pixbuf.Unref()
	//给按钮设置图片
	butto.SetImage(image)
	//去掉按钮的焦距
	butto.SetCanFocus(false)
}

func (obj *Chessboard) CreateWindow() {
	//读取glade
	builder := gtk.NewBuilder()
	builder.AddFromFile("./ui.glade")

	/*窗口相关*/
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

	/*按钮相关*/
	obj.buttonMin = gtk.ButtonFromObject(builder.GetObject("buttonMin"))
	ButtonSetImgFromFile(obj.buttonMin, "./images/min.png")
	obj.buttonClose = gtk.ButtonFromObject(builder.GetObject("buttonClose"))
	ButtonSetImgFromFile(obj.buttonClose, "./images/close.png")

	/*标签相关*/
	lableBlack := gtk.LabelFromObject(builder.GetObject("lableBlack"))
	lableWhite := gtk.LabelFromObject(builder.GetObject("lableWhite"))
	lableTime := gtk.LabelFromObject(builder.GetObject("lableTime"))
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

func PaintEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard)
	if ok == false {
		fmt.Println("MousePressEvent Chessboard error")
		return
	}

	//获取画板
	painter := obj.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(painter)
	//新建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./images/bg.jpg", obj.w, obj.h, false)
	//画图
	painter.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
	//释放资源
	pixbuf.Unref()
}

func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标移动事件
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)

	//关闭按钮
	obj.buttonClose.Clicked(func() {
		gtk.MainQuit()
	})

	//最小化按钮
	obj.buttonMin.Clicked(func() {
		obj.window.Iconify()
	})

	/*绘图相关*/
	//大小改变事件
	obj.window.Connect("configure_event", func() {
		//重新刷图
		obj.window.QueueDraw()
	})

	//绘图事件
	obj.window.Connect("expose-event", PaintEvent, obj)

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
