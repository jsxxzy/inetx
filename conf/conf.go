// Author: d1y<chenhonzhou@gmail.com>

package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Appname 文件名称
const Appname = `校园网登录小工具`

// FromTIP 表单提示
const FromTIP = `填入账号密码即可登录`

// QuitText 退出文字
const QuitText = `退出软件`

// LoginText 登录
const LoginText = `登录`

// LogoutText 注销
const LogoutText = `注销`

// UseInfoText 使用信息
const UseInfoText = `使用信息`

// UsernameText u
var UsernameText = `用户名`

// PasswordText p
var PasswordText = `密码`

var filename = `方正有猫在简体.ttf`

var devPATH = fmt.Sprintf(`./static/%s`, filename)

var releasePATH = fmt.Sprintf(`./%s`, filename)

// FontTTF 字体文件
var FontTTF = devPATH

// GetLocalAuth 获取本地用户名密码
//
//
func GetLocalAuth() (a string, b string, r bool) {
	easyGetLocalConfig(&a, &b)
	return a, b, len(a) >= 1 && len(b) >= 1
}

// SetLocalAuth 设置本地用户名和密码
//
//
func SetLocalAuth(u, p string) error {
	return setConfigProfile(u, p)
}

func init() {
	if !checkFileNotExit(devPATH) {
		FontTTF = devPATH
	} else if !checkFileNotExit(releasePATH) {
		FontTTF = releasePATH
	}
}

// checkFileNotExit 判断文件是否存在
func checkFileNotExit(f string) bool {
	if _, err := os.Stat(f); err != nil {
		return true
	}
	return false
}

// =====

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func getHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}

// 设置本地缓存
func setConfigProfile(u, p string) error {
	var f = getConfigFile()
	if len(u) == 0 && len(p) == 0 {
		return ioutil.WriteFile(f.Name(), []byte(""), 0777)
	}
	var parseStr = fmt.Sprintf("%v,%v", u, p)
	var b = []byte(parseStr)
	return ioutil.WriteFile(f.Name(), b, 0777)
}

// 获取配置文件, 不安全的方法, 切勿使用!!!
func getConfigFile() *os.File {
	homeDir, err := getHomeDir()
	if err != nil {
		panic(err)
	}
	var configfile = ".inetconfig"
	var file = filepath.Join(homeDir, configfile)
	if !exists(file) {
		var f, _ = os.Create(file)
		return f
	}
	var f, _ = os.Open(file)
	return f
}

// 解析配置文件
func parseConfig(b *os.File) (string, string, error) {
	var p = b.Name()
	var a, e = ioutil.ReadFile(p)
	if e != nil {
		return "", "", errors.New("get config file is error")
	}
	var c = string(a)
	var arr = strings.Split(c, ",")
	if len(arr) <= 1 {
		return "", "", errors.New("解析失败")
	}
	return arr[0], arr[1], nil
}

func easyGetLocalConfig(u, p *string) {
	var username, password, err = parseConfig(getConfigFile())
	if err != nil || len(username) <= 1 || len(password) <= 1 {
		return
	}
	*u = username
	*p = password
}

// =====
