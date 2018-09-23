package main

import (
	"fmt"
	"os"
	"strconv"
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
	labelBlack  *gtk.Label
	labelWhite  *gtk.Label
	labelTime   *gtk.Label
	imgBlack    *gtk.Image
	imgWhite    *gtk.Image
}

//控件属性结构体
type ChessInfo struct {
	w, h           int
	x, y           int
	startX, startY int
	gridW, gridH   int
}

//黑白棋结构体
type Chessboard struct {
	ChessWidget
	ChessInfo

	CurrentRole int
	tipTimeId   int
	endTimeId   int
	timeNum     int

	chess [8][8]int
}

//枚举,标识黑白子状态
const (
	Empty = iota
	Black
	White
)

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

func ImgSetPicFromFile(image *gtk.Image, filename string) {
	//获取image大小
	w, h := 0, 0
	image.GetSizeRequest(&w, &h)
	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)
	//给image设置图片
	image.SetFromPixbuf(pixbuf)
	//释放图片
	pixbuf.Unref()
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
	obj.labelBlack = gtk.LabelFromObject(builder.GetObject("labelBlack"))
	obj.labelWhite = gtk.LabelFromObject(builder.GetObject("labelWhite"))
	obj.labelTime = gtk.LabelFromObject(builder.GetObject("labelTime"))

	//设置字体大小
	obj.labelBlack.ModifyFontSize(50)
	obj.labelWhite.ModifyFontSize(50)
	obj.labelTime.ModifyFontSize(30)

	obj.labelBlack.SetText("0")
	obj.labelWhite.SetText("0")
	obj.labelTime.SetText("20")

	obj.labelBlack.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelWhite.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelTime.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))

	obj.imgBlack = gtk.ImageFromObject(builder.GetObject("imageBlack"))
	obj.imgWhite = gtk.ImageFromObject(builder.GetObject("imageWhite"))

	ImgSetPicFromFile(obj.imgBlack, "./images/black.png")
	ImgSetPicFromFile(obj.imgWhite, "./images/white.png")

	obj.startX, obj.startY = 200, 60
	obj.gridW, obj.gridH = 50, 40
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
	x := (obj.x - obj.startX) / obj.gridW
	y := (obj.y - obj.startY) / obj.gridH

	if x >= 0 && x <= 7 && y >= 0 && y <= 7 {
		obj.chess[x][y] = obj.CurrentRole
		obj.window.QueueDraw()
		//改变角色
		obj.ChangeRole()
	}
}

func (obj *Chessboard) ChangeRole() {
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))
	obj.imgWhite.Hide()
	obj.imgBlack.Hide()
	if obj.CurrentRole == Black {
		obj.CurrentRole = White
	} else {
		obj.CurrentRole = Black
	}
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

	//画棋子
	blackPixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./images/black.png", obj.gridW, obj.gridH, false)
	whitePixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("./images/white.png", obj.gridW, obj.gridH, false)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				painter.DrawPixbuf(gc, blackPixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				painter.DrawPixbuf(gc, whitePixbuf, 0, 0, (obj.startX + i*obj.gridW), (obj.startY + j*obj.gridH), -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}

	//释放资源
	pixbuf.Unref()
	blackPixbuf.Unref()
	whitePixbuf.Unref()
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

func showTip(obj *Chessboard) {
	if obj.CurrentRole == Black {
		//隐藏白子
		obj.imgWhite.Hide()
		if obj.imgBlack.GetVisible() == true {
			obj.imgBlack.Hide()
		} else {
			obj.imgBlack.Show()
		}
	} else {
		obj.imgBlack.Hide()
		if obj.imgWhite.GetVisible() == true {
			obj.imgWhite.Hide()
		} else {
			obj.imgWhite.Show()
		}
	}
}

func (obj *Chessboard) InitChess() {
	//初始化棋盘
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			obj.chess[i][j] = Empty
		}
	}
	//默认各方有两个棋子
	obj.chess[3][3] = Black
	obj.chess[4][4] = Black
	obj.chess[4][3] = White
	obj.chess[3][4] = White
	obj.imgBlack.Hide()
	obj.imgWhite.Hide()
	//黑子先下
	obj.CurrentRole = Black
	//启动定时器
	obj.tipTimeId = glib.TimeoutAdd(500, func() bool {
		showTip(obj)
		return true
	})
	//倒计时定时器
	obj.timeNum = 20
	obj.labelTime.SetText(strconv.Itoa(obj.timeNum))

	//启动下子定时器
	obj.endTimeId = glib.TimeoutAdd(1000, func() bool {
		obj.timeNum--
		obj.labelTime.SetText(strconv.Itoa(obj.timeNum))
		if obj.timeNum == 0 {
			obj.ChangeRole()
		}
		return true
	})
}

func main() {
	//初始化GTK库
	gtk.Init(&os.Args)

	var obj Chessboard
	obj.CreateWindow()
	obj.HandleSignal()
	obj.window.Show()
	obj.InitChess()
	//运行主事件
	gtk.Main()
}
