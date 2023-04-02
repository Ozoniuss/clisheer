package color

import "fmt"

func Println(color Color, v ...any) (n int, err error) {
	fmt.Print(color)
	defer fmt.Print(Reset)
	return fmt.Println(v...)
}

func Printf(color Color, format string, v ...any) (n int, err error) {
	fmt.Print(color)
	defer fmt.Print(Reset)
	return fmt.Printf(format, v...)
}
