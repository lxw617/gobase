package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	// 打印在终端有颜色

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

	fmt.Println("// ------------------------------------")
	const escape = "\x1b" //ascii码表中对应escape的编码
	f := fmt.Sprintf("%s[%sm", escape, "34")
	fmt.Fprint(os.Stdout, f)
	fmt.Fprintln(os.Stdout, "Hello World in blue")
}
