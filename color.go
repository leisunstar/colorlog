package colorlog

import (
    "fmt"
    "runtime"
)

const (
    TextBlack = iota + 30
    TextRed
    TextGreen
    TextYellow
    TextBlue
    TextMagenta
    TextCyan
    TextWhite
)

func Black(str string) string {
    return textColor(TextBlack, str)
}

func Red(str string) string {
    return textColor(TextRed, str)
}


func Green(str string) string {
    return textColor(TextGreen, str)
}


func Yellow(str string) string {
    return textColor(TextYellow, str)
}

func Blue(str string) string {
    return textColor(TextBlue, str)
}

func Magenta(str string) string {
    return textColor(TextMagenta, str)
}

func Cyan(str string) string {
    return textColor(TextCyan, str)
}

func White(str string) string {
    return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
    if IsWindows() {
        return str
    }

    switch color {
        case TextBlack:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
        case TextRed:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
        case TextGreen:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
        case TextYellow:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
        case TextBlue:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
        case TextMagenta:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
        case TextCyan:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
        case TextWhite:
        return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
        default:
        return str
    }
}

type ColorLog struct {

}

func NewColorLog()(*ColorLog){
    return &ColorLog{}
}

func (this *ColorLog)Info(format string, a ...interface{}){
    fmt.Println(textColor(TextWhite, fmt.Sprintf("info:%s", fmt.Sprintf(format, a))))
}

func (this *ColorLog)Debug(format string, a ...interface{}){
    fmt.Println(textColor(TextGreen, fmt.Sprintf("debug:%s", fmt.Sprintf(format, a))))
}

func (this *ColorLog)Warning(format string, a ...interface{}){
    fmt.Println(textColor(TextYellow, fmt.Sprintf("warning:%s", fmt.Sprintf(format, a))))
}

func (this *ColorLog) Error(format string, a ...interface{}){
    fmt.Println(textColor(TextRed, fmt.Sprintf("error:%s", fmt.Sprintf(format, a))))
}

func IsWindows() bool {
    if runtime.GOOS == "windows" {
        return true
    } else {
        return false
    }
}

