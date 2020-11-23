// Author: d1y<chenhonzhou@gmail.com>

package itheme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"github.com/jsxxzy/inetx/res"
)

type MyThem struct{}

// return bundled font resource
func (MyThem) TextFont() fyne.Resource {
	return &res.FontStatic
}
func (MyThem) TextBoldFont() fyne.Resource {
	return &res.FontStatic
}

func (MyThem) BackgroundColor() color.Color      { return theme.DarkTheme().BackgroundColor() }
func (MyThem) ButtonColor() color.Color          { return theme.DarkTheme().ButtonColor() }
func (MyThem) DisabledButtonColor() color.Color  { return theme.DarkTheme().DisabledButtonColor() }
func (MyThem) IconColor() color.Color            { return theme.DarkTheme().IconColor() }
func (MyThem) DisabledIconColor() color.Color    { return theme.DarkTheme().DisabledIconColor() }
func (MyThem) HyperlinkColor() color.Color       { return theme.DarkTheme().HyperlinkColor() }
func (MyThem) TextColor() color.Color            { return theme.DarkTheme().TextColor() }
func (MyThem) DisabledTextColor() color.Color    { return theme.DarkTheme().DisabledTextColor() }
func (MyThem) HoverColor() color.Color           { return theme.DarkTheme().HoverColor() }
func (MyThem) PlaceHolderColor() color.Color     { return theme.DarkTheme().PlaceHolderColor() }
func (MyThem) PrimaryColor() color.Color         { return theme.DarkTheme().PrimaryColor() }
func (MyThem) FocusColor() color.Color           { return theme.DarkTheme().FocusColor() }
func (MyThem) ScrollBarColor() color.Color       { return theme.DarkTheme().ScrollBarColor() }
func (MyThem) ShadowColor() color.Color          { return theme.DarkTheme().ShadowColor() }
func (MyThem) TextSize() int                     { return theme.DarkTheme().TextSize() }
func (MyThem) TextItalicFont() fyne.Resource     { return theme.DarkTheme().TextItalicFont() }
func (MyThem) TextBoldItalicFont() fyne.Resource { return theme.DarkTheme().TextBoldItalicFont() }
func (MyThem) TextMonospaceFont() fyne.Resource  { return theme.DarkTheme().TextMonospaceFont() }
func (MyThem) Padding() int                      { return theme.DarkTheme().Padding() }
func (MyThem) IconInlineSize() int               { return theme.DarkTheme().IconInlineSize() }
func (MyThem) ScrollBarSize() int                { return theme.DarkTheme().ScrollBarSize() }
func (MyThem) ScrollBarSmallSize() int           { return theme.DarkTheme().ScrollBarSmallSize() }
