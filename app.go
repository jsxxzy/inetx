// Author: d1y<chenhonzhou@gmail.com>
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/jsxxzy/inet"
	"github.com/jsxxzy/inetx/conf"
	"github.com/jsxxzy/inetx/itheme"
)

type inetapp struct {
	app           fyne.App       // 窗口上层实例
	window        fyne.Window    // 单个窗口实例
	tipText       *widget.Label  // 提示文字
	form          *widget.Form   // 表单
	usernameInput *widget.Entry  // 表单用户名
	pwdInput      *widget.Entry  // 表单密码
	actionButton  *widget.Button // action-button
	reloadButton  *widget.Button // 刷新按钮
	quitButton    *widget.Button // 退出按钮
	infoBox       *widget.Card   // 信息弹窗
	isLogin       bool           // 是否登录
}

const (
	remindLogin = "请先登录"
	hasLogin    = "已登录"
	unknownX    = "未知"
)

// Init 初始化
func (i inetapp) Init() {

	i.app = app.New()

	i.LoadFontTTF()

	i.window = i.app.NewWindow(conf.Appname)

	i.SetWindowDefaultSize()

	i.tipText = widget.NewLabel(remindLogin)
	i.tipText.Alignment = fyne.TextAlignCenter

	i.usernameInput = widget.NewEntry()
	i.pwdInput = widget.NewPasswordEntry()

	i.form = widget.NewForm(&widget.FormItem{
		Text:   conf.UsernameText,
		Widget: i.usernameInput,
	}, &widget.FormItem{
		Text:   conf.PasswordText,
		Widget: i.pwdInput,
	})

	i.actionButton = widget.NewButton(conf.LoginText, func() {
		var tty = i.isLogin
		if tty {
			log.Println("已登录, 准备注销")
			var err = i.Logout()
			if err != nil {
				log.Println("注销错误")
			}
			i.SetInfo()
		} else {
			log.Println("未登录, 登录中")
			i.Login()
		}
	})

	i.reloadButton = widget.NewButton("刷新", func() {
		i.SetInfo()
	})

	i.quitButton = widget.NewButton(conf.QuitText, func() {
		i.app.Quit()
	})

	i.infoBox = widget.NewCard(conf.UseInfoText, "", widget.NewLabel(""))

	var ctx = widget.NewVBox(
		i.tipText,
		i.infoBox,
		i.form,
		i.actionButton,
		i.reloadButton,
		i.quitButton,
	)

	i.window.SetContent(ctx)

	i.SetInfo()
	i.SetFormField()

	i.Loop()

}

// SetLoginFlag 设置是否登录
func (i *inetapp) SetLoginFlag(f bool) {
	i.isLogin = f
	if f {
		i.form.Hide()
		i.infoBox.Show()
		i.tipText.SetText(hasLogin)
		i.actionButton.SetText(conf.LogoutText)
		return
	}
	i.actionButton.SetText(conf.LoginText)
	i.tipText.SetText(remindLogin)
	i.form.Show()
	i.infoBox.Hide()
}

// SetInfo 设置
func (i *inetapp) SetInfo() error {
	data, err := i.QueryInfo()
	if err != nil || data.Error() != nil {
		i.SetLoginFlag(false)
		return err
	}
	i.SetLoginFlag(true)
	xTime, _ := strconv.Atoi(data.Time)
	// 使用时长
	var time = getHumanTime(xTime)
	// 流量
	var flow = getHumanFlow(data.Flow)
	// 用户ID
	var uid = data.UID

	// v4ip
	var v4ip = data.V4ip
	i.infoBox.SetSubTitle(v4ip)
	var content = fmt.Sprintf("使用流量: %s\n使用时长: %s\nUID: %s", flow, time, uid)
	i.infoBox.SetContent(widget.NewLabel(content))
	return nil
}

var suffixes [5]string

// SetFormField 设置表单
func (i inetapp) SetFormField() {
	var a, b, c = i.GetAuth()
	if !c {
		return
	}
	i.usernameInput.SetText(a)
	i.pwdInput.SetText(b)
}

// SetAuth 设置本地用户
func (i *inetapp) SetAuth(u, p string) bool {
	if len(u) >= 1 && len(p) >= 1 {
		return conf.SetLocalAuth(u, p) == nil
	}
	return false
}

// GetAuth 获取本地用户
func (i *inetapp) GetAuth() (a string, b string, r bool) {
	return conf.GetLocalAuth()
}

// Round round offset
//
// https://gist.github.com/anikitenko/b41206a49727b83a530142c76b1cb82d
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// 获取格式化好的时间
func getHumanTime(h int) string {
	if h < 60 {
		return fmt.Sprintf("%v分钟", h)
	}
	if h == 60 {
		return "1小时"
	}
	m := h % 60
	var p float64 = 60
	b := float64(h) / p
	c := math.Floor(b)
	return fmt.Sprintf("%v小时%v分钟", c, m)
}

// getHumanFlow 转换流量格式转为阳间格式
//
// https://gist.github.com/anikitenko/b41206a49727b83a530142c76b1cb82d
func getHumanFlow(f float64) (r string) {
	defer func() {
		if err := recover(); err != nil {
			r = "0kb"
		}
	}()
	size := f * 1024 * 1024 // This is in bytes
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	var result = strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
	r = result
	return result
}

// LoadFontTTF 加载字体文件(主题相关)
func (i inetapp) LoadFontTTF() {
	i.app.Settings().SetTheme(&itheme.MyThem{})
}

// SetWindowSize 设置默认窗口大小
func (i inetapp) SetWindowDefaultSize() {
	var w, h = 420, 210
	i.SetWindowSize(w, h)
}

// SetWindowSize 设置窗口大小
func (i inetapp) SetWindowSize(w, h int) {
	i.window.Resize(fyne.NewSize(w, h))
}

// GetUsername 用户名
func (i inetapp) GetUsername() string {
	return i.usernameInput.Text
}

// GetPassword 密码
func (i inetapp) GetPassword() string {
	return i.pwdInput.Text
}

// Login 登录
func (i inetapp) Login() bool {
	var u = i.GetUsername()
	var p = i.GetPassword()
	if verifyForm(u, p) {
		i.actionButton.Disable()
		var loginInfo, err = inet.Login(u, p)
		i.actionButton.SetText("登录中")
		i.SetInfo()
		if err != nil {
			return false
		}
		i.SetAuth(u, p)
		i.actionButton.Enable()
		var msg = loginInfo.GetMsg()
		log.Println(msg)
	} else {
		var e = errors.New("账号密码填写错误")
		log.Println(e)
		dialog.ShowError(e, i.window)
	}
	return false
}

// 验证表单
func verifyForm(u, p string) bool {

	// 然而一个账号最起码长度为 5+
	// 学校的账号后缀为: @jszy
	// 而密码为6位数
	// 作者注解: 2020/11/23
	return len(u) >= 1 && len(p) >= 1

}

// Logout 注销
func (i inetapp) Logout() error {
	return inet.Logout()
}

// Check 判断是否登录
func (i inetapp) Check() bool {
	return inet.HasLogin()
}

// QueryInfo 查询信息
func (i inetapp) QueryInfo() (inet.QueryInfoData, error) {
	return inet.QueryInfo()
}

// Loop 死循环
func (i inetapp) Loop() {
	i.window.ShowAndRun()
}
