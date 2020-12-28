package main

import (
	"fmt"
)

type figure struct {
	height, width int
}

func side(fn func(figure) figure) figure {
	return fn(figure{height: 20, width: 30})
}

func main() {
	var inc = func(f figure) figure {
		return (figure{height: f.height + 15, width: f.width + 15})
	}
	fmt.Println(inc(figure{height: 15, width: 50}))
	fmt.Println(side(inc))
}
