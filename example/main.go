package main

import (
	"fmt"
	"time"

	"github.com/arussellsaw/unicorn-go"
	"github.com/arussellsaw/unicorn-go/util"
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
	time.Sleep(1 * time.Second)
	circ := util.Circle(43, [2]int{0, 0}, util.Cyan)
	m := util.Matrix{}
	m.AddSupersample(circ)
	ps := unicorn.DeMatrix(m)
	err = c.SetAllPixels(ps)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Show()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 40; i > 0; i-- {
		err = c.SetBrightness(uint(i))
		if err != nil {
			fmt.Println(err)
			return
		}
		err = c.Show()
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = c.SetBrightness(40)
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
