# tightywhities [![GoDoc](https://godoc.org/github.com/chewxy/tightywhities?status.svg)](https://godoc.org/github.com/chewxy/tightywhities) [![Build Status](https://travis-ci.org/chewxy/tightywhities.svg?branch=master)](https://travis-ci.org/chewxy/tightywhities) [![Coverage Status](https://coveralls.io/repos/github/chewxy/tightywhities/badge.png)](https://coveralls.io/github/chewxy/tightywhities) [![Go Report Card](https://goreportcard.com/badge/github.com/chewxy/tightywhities)](https://goreportcard.com/report/github.com/chewxy/tightywhities)

package `tightywhities` plots charts in ASCII. It's mainly useful for simple examples for godoc.

# Usage #

This package is go-gettable: `go get -u github.com/chewxy/tightywhities`. There are NO external dependencies.

To use, import the package. Currently only line charts are supported

```
func main() {
	// generate
	var data []float64
	for i := 0.0; i <= 360; i += 5 {
		data = append(data, math.Sin(i))
	}

	l := NewLine(nil, data)
	l.Plot(30, 15, os.Stdout)
}
```

You should see an output like this:

```
│
│                        ╭╮   ╭╮   ╭╮   ╭╮   ╭╮   ╭╮                      ╭
│     ╭╮   ╭╮  ╭╮   ╭╮   ││   ││   ││   ││   ││   ││   ╭╮   ╭╮  ╭╮   ╭╮   │
│     ││  ╭╯│  │╰╮  │╰╮  ││   ││   ││   ││   ││   ││  ╭╯│  ╭╯│  │╰╮  ││   │
│    ╭╯│  │ │  │ │  │ │  │╰╮  ││   ││   ││   ││  ╭╯│  │ │  │ │  │ │  │╰╮  │
│    │ │  │ │  │ │  │ │  │ │  │╰╮  ││   ││  ╭╯│  │ │  │ │  │ │  │ │  │ │  │
│    │ │  │ │  │ │  │ │  │ │  │ │  │╰╮ ╭╯│  │ │  │ │  │ │  │ │  │ │  │ │  │
│    │ │  │ │  │ │  │ │  │ │  │ │ ╭╯ │ │ ╰╮ │ │  │ │  │ │  │ │  │ │  │ │  │
│+ ╮ │ │  │ │  │ │  │ │  │ │ ╭╯ │ │  │ │  │ │ ╰╮ │ │  │ │  │ │  │ │  │ │  │
│  │ │ ╰╮ │ │  │ │  │ │ ╭╯ │ │  │ │  │ │  │ │  │ │ ╰╮ │ │  │ │  │ │ ╭╯ │ ╭╯
│  │ │  │ │ ╰╮ │ │ ╭╯ │ │  │ │  │ │  │ │  │ │  │ │  │ │ ╰╮ │ │ ╭╯ │ │  │ │
│  │ │  │ │  │╭╯ ╰╮│  │ │  │ │  │ │  │ │  │ │  │ │  │ │  │╭╯ ╰╮│  │ │  │ │
│  │ │  │╭╯  ││   ││  ╰╮│  │ │  │ │  │ │  │ │  │ │  │╭╯  ││   ││  ╰╮│  │ │
│  │╭╯  ││   ││   ││   ││  ╰╮│  │ │  │ │  │ │  │╭╯  ││   ││   ││   ││  ╰╮│
│  ││   ││   ││   ││   ││   ││  ╰╮│  │ │  │╭╯  ││   ││   ││   ││   ││   ││
│  ││   ││   ││   ││   ││   ╰╯   ╰╯  ╰─╯  ╰╯   ╰╯   ││   ││   ││   ││   ││
│  ╰╯   ╰╯   ╰╯   ╰╯   ╰╯                           ╰╯   ╰╯   ╰╯   ╰╯   ╰
```

# Future #

Right now only line charts are supported (I wrote this to scratch an itch about a package that had graphical examples). In the future, more may be supported. Or you could just send a pull request.


# Contributing #

1. File an issue (bug or feature request)
2. Send a pull request

Be nice. Don't be a dick.

# Licence # 

This package is MIT licenced

# Trivia #

The name of the package is due to @bassjacob and @johndagostino - I promised them that I'll name my next package "tightywhiteys" after I complained about stupid package names. :). 

The point drawing was inspired by [TermUI](https://github.com/gizak/termui) (originally only points were drawn, but it was difficult to see the points for the sine curve)
