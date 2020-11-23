// Author: d1y<chenhonzhou@gmail.com>

package res

import (
	"io/ioutil"

	"fyne.io/fyne"
	"github.com/jsxxzy/inetx/conf"
)

// 该字体未正版授权, 本软件并不会用作商业用途, 侵权必删
//
// http://zitixiazai.taofont.com/方正有猫在简体.html
var ttfName = conf.FontTTF

// ByteData 字体文件
var ByteData []byte

// FontStatic 字体文件
var FontStatic fyne.StaticResource

func init() {
	data, err := ioutil.ReadFile(ttfName)
	if err != nil {
		panic(err)
	}
	ByteData = data

	// 中文乱码: https://github.com/andydotxyz/beebui
	// 参考: https://github.com/fyne-io/fyne/issues/598
	// github issues: https://github.com/fyne-io/fyne/issues/604
	FontStatic = fyne.StaticResource{
		StaticName:    "FZYouMaoZaiS-R-GB.ttf",
		StaticContent: ByteData,
	}

}
