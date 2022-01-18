package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

func Success(msg string){
	colorOut(msg,"green")
}

func Error(msg string){
	colorOut(msg,"red")
}

func Warning(msg string){
	colorOut(msg,"yellow")
}

func Exit(msg string){
	Error(msg)
	os.Exit(1)
}

func ExitIf(err error){
	if err != nil {
		Exit(err.Error())
	}
}

//colorOut 设置高亮颜色
func colorOut(message,color string){
	fmt.Fprintln(os.Stdout,ansi.Color(message,color))
}

