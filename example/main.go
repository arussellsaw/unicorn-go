package main

import (
	"fmt"

	"github.com/arussellsaw/unicorn-go"
)

func main() {
	c := unicorn.Client{Path: unicorn.SocketPath}
	err := c.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	pixels := [64]unicorn.Pixel{}
	for i := range pixels {
		pixels[i] = unicorn.Pixel{200, 0, 0}
	}
	err = c.SetBrightness(40)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.SetAllPixels(pixels)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.SetPixel(3, 3, 0, 255, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Show()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Clear()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Show()
	if err != nil {
		fmt.Println(err)
		return
	}
}
