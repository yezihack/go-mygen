package gocolor

import (
	"fmt"
)

//定义不同颜色值
var (
	greenBg   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green     = string([]byte{27, 91, 51, 50, 109})
	white     = string([]byte{27, 91, 51, 55, 109})
	yellow    = string([]byte{27, 91, 51, 51, 109})
	red       = string([]byte{27, 91, 51, 49, 109})
	blue      = string([]byte{27, 91, 51, 52, 109})
	magenta   = string([]byte{27, 91, 51, 53, 109})
	cyan      = string([]byte{27, 91, 51, 54, 109})
	reset     = string([]byte{27, 91, 48, 109})
)

func colorPrint(format string, color string, s ...interface{}) string {
	if len(s) == 0 {
		return fmt.Sprint(color + format + reset)
	}
	data := fmt.Sprintf(format, s...)
	return fmt.Sprint(color + data + reset)
}

/************************直接打印不同颜色字符串**************************/

//绿色字体
func Green(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, green, s...))
}

//白色字体
func White(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, white, s...))
}

//黄色字体
func Yellow(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, yellow, s...))
}

//红色字体
func Red(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, red, s...))
}

//蓝色字体
func Blue(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, blue, s...))
}

//品红字体
func Magenta(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, magenta, s...))
}

//蓝绿色
func Cyan(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, cyan, s...))
}

/************************返回不同颜色字符串**************************/

//绿色字体
func SGreen(format string, s ...interface{}) string {
	return colorPrint(format, green, s...)
}

//白色字体
func SWhite(format string, s ...interface{}) string {
	return colorPrint(format, white, s...)
}

//黄色字体
func SYellow(format string, s ...interface{}) string {
	return colorPrint(format, yellow, s...)
}

//红色字体
func SRed(format string, s ...interface{}) string {
	return colorPrint(format, red, s...)
}

//蓝色字体
func SBlue(format string, s ...interface{}) string {
	return colorPrint(format, blue, s...)
}

//品红字体
func SMagenta(format string, s ...interface{}) string {
	return colorPrint(format, magenta, s...)
}

//蓝绿色
func SCyan(format string, s ...interface{}) string {
	return colorPrint(format, cyan, s...)
}

/************************直接打印不同背景颜色字符串**************************/

//绿色字体
func GreenBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, greenBg, s...))
}

//白色字体
func WhiteBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, whiteBg, s...))
}

//黄色字体
func YellowBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, yellowBg, s...))
}

//红色字体
func RedBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, redBg, s...))
}

//蓝色字体
func BlueBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, blueBg, s...))
}

//品红字体
func MagentaBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, magentaBg, s...))
}

//蓝绿色背景色
func CyanBG(format string, s ...interface{}) {
	fmt.Print(colorPrint(format, cyanBg, s...))
}

/************************返回不同背景颜色字符串**************************/

//绿色字体
func SGreenBG(format string, s ...interface{}) string {
	return colorPrint(format, greenBg, s...)
}

//白色字体
func SWhiteBG(format string, s ...interface{}) string {
	return colorPrint(format, whiteBg, s...)
}

//黄色字体
func SYellowBG(format string, s ...interface{}) string {
	return colorPrint(format, yellowBg, s...)
}

//红色字体
func SRedBG(format string, s ...interface{}) string {
	return colorPrint(format, redBg, s...)
}

//蓝色字体
func SBlueBG(format string, s ...interface{}) string {
	return colorPrint(format, blueBg, s...)
}

//品红字体
func SMagentaBG(format string, s ...interface{}) string {
	return colorPrint(format, magentaBg, s...)
}

//蓝绿色背景色
func SCyanBG(format string, s ...interface{}) string {
	return colorPrint(format, cyanBg, s...)
}
