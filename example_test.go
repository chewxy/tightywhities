package tightywhities

import (
	"math"
	"os"
)

func Example_Connective() {
	// generate
	var data []float64
	for i := 0.0; i <= 360; i += 5 {
		data = append(data, math.Sin(i))
	}

	l := NewLine(nil, data)
	l.Plot(30, 15, os.Stdout)

	// Output:
	// │
	// │                        ╭╮   ╭╮   ╭╮   ╭╮   ╭╮   ╭╮                      ╭
	// │     ╭╮   ╭╮  ╭╮   ╭╮   ││   ││   ││   ││   ││   ││   ╭╮   ╭╮  ╭╮   ╭╮   │
	// │     ││  ╭╯│  │╰╮  │╰╮  ││   ││   ││   ││   ││   ││  ╭╯│  ╭╯│  │╰╮  ││   │
	// │    ╭╯│  │ │  │ │  │ │  │╰╮  ││   ││   ││   ││  ╭╯│  │ │  │ │  │ │  │╰╮  │
	// │    │ │  │ │  │ │  │ │  │ │  │╰╮  ││   ││  ╭╯│  │ │  │ │  │ │  │ │  │ │  │
	// │    │ │  │ │  │ │  │ │  │ │  │ │  │╰╮ ╭╯│  │ │  │ │  │ │  │ │  │ │  │ │  │
	// │    │ │  │ │  │ │  │ │  │ │  │ │ ╭╯ │ │ ╰╮ │ │  │ │  │ │  │ │  │ │  │ │  │
	// │+ ╮ │ │  │ │  │ │  │ │  │ │ ╭╯ │ │  │ │  │ │ ╰╮ │ │  │ │  │ │  │ │  │ │  │
	// │  │ │ ╰╮ │ │  │ │  │ │ ╭╯ │ │  │ │  │ │  │ │  │ │ ╰╮ │ │  │ │  │ │ ╭╯ │ ╭╯
	// │  │ │  │ │ ╰╮ │ │ ╭╯ │ │  │ │  │ │  │ │  │ │  │ │  │ │ ╰╮ │ │ ╭╯ │ │  │ │
	// │  │ │  │ │  │╭╯ ╰╮│  │ │  │ │  │ │  │ │  │ │  │ │  │ │  │╭╯ ╰╮│  │ │  │ │
	// │  │ │  │╭╯  ││   ││  ╰╮│  │ │  │ │  │ │  │ │  │ │  │╭╯  ││   ││  ╰╮│  │ │
	// │  │╭╯  ││   ││   ││   ││  ╰╮│  │ │  │ │  │ │  │╭╯  ││   ││   ││   ││  ╰╮│
	// │  ││   ││   ││   ││   ││   ││  ╰╮│  │ │  │╭╯  ││   ││   ││   ││   ││   ││
	// │  ││   ││   ││   ││   ││   ╰╯   ╰╯  ╰─╯  ╰╯   ╰╯   ││   ││   ││   ││   ││
	// │  ╰╯   ╰╯   ╰╯   ╰╯   ╰╯                           ╰╯   ╰╯   ╰╯   ╰╯   ╰╯

}

func Example_Point() {
	// generate
	var data []float64
	for i := 0.0; i <= 360; i += 5 {
		data = append(data, math.Sin(i))
	}

	l := NewLine(nil, data)
	l.Mode = Point
	l.Plot(30, 15, os.Stdout)

	// Output:
	// │
	// │                        ⡠⠢   ⡠⠢   ⡠⠢   ⡠⠢   ⡠⠢   ⡠⠢                      ⡠
	// │     ⡠⠢   ⡠⠢  ⡠⠢   ⡠⠢   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡠⠢   ⡠⠢  ⡠⠢   ⡠⠢   ⡇
	// │     ⡇⡇  ⡠⠊⡇  ⡇⠑⠢  ⡇⠑⠢  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇  ⡠⠊⡇  ⡠⠊⡇  ⡇⠑⠢  ⡇⡇   ⡇
	// │    ⡠⠊⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⠑⠢  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇  ⡠⠊⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⠑⠢  ⡇
	// │    ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⠑⠢  ⡇⡇   ⡇⡇  ⡠⠊⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇
	// │    ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⠑⠢ ⡠⠊⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇
	// │    ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⡠⠊ ⡇ ⡇ ⠑⠢ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇
	// │+ ⠢ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⡠⠊ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⠑⠢ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇
	// │  ⡇ ⡇ ⠑⠢ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⡠⠊ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⠑⠢ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⡠⠊ ⡇ ⡠⠊
	// │  ⡇ ⡇  ⡇ ⡇ ⠑⠢ ⡇ ⡇ ⡠⠊ ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇ ⠑⠢ ⡇ ⡇ ⡠⠊ ⡇ ⡇  ⡇ ⡇
	// │  ⡇ ⡇  ⡇ ⡇  ⡇⡠⠊ ⠑⠢⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⡠⠊ ⠑⠢⡇  ⡇ ⡇  ⡇ ⡇
	// │  ⡇ ⡇  ⡇⡠⠊  ⡇⡇   ⡇⡇  ⠑⠢⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⡠⠊  ⡇⡇   ⡇⡇  ⠑⠢⡇  ⡇ ⡇
	// │  ⡇⡠⠊  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇  ⠑⠢⡇  ⡇ ⡇  ⡇ ⡇  ⡇ ⡇  ⡇⡠⠊  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇  ⠑⠢⡇
	// │  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇  ⠑⠢⡇  ⡇ ⡇  ⡇⡠⠊  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇
	// │  ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⠑⠊   ⠑⠊  ⠑⠤⠊  ⠑⠊   ⠑⠊   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇   ⡇⡇
	// │  ⠑⠊   ⠑⠊   ⠑⠊   ⠑⠊   ⠑⠊                           ⠑⠊   ⠑⠊   ⠑⠊   ⠑⠊   ⠑⠊

}