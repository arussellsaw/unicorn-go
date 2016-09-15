package unicorn

import (
	"fmt"
	"net"

	"github.com/lunixbochs/struc"
)

var (
	SocketPath = "/var/run/unicornd.socket"
)

// Client is the unix socket client for unicornd
type Client struct {
	Path string
	sock net.Conn
}

// Connect opens a connection to the Client.Path
func (c *Client) Connect() error {
	var err error
	c.sock, err = net.Dial("unix", SocketPath)
	return err
}

// SetBrightness of the display, 0..255
func (c *Client) SetBrightness(v uint) error {
	if v > 255 {
		return fmt.Errorf("brightness must be 0..255, passed: %v", v)
	}
	b := brightness{
		Code: CMDSetBrightness,
		Val:  v,
	}
	return struc.Pack(c.sock, &b)
}

// SetPixel sets the color of an individual pixel
func (c *Client) SetPixel(x, y, r, g, b uint) error {
	if x > 7 || y > 7 || r > 255 || g > 255 || b > 255 {
		return fmt.Errorf("x, y must be 0..7, r,g,b must be 0..255, passed: x:%v y:%v r:%v g:%v b:%v", x, y, r, g, b)
	}
	sp := setPixel{
		Code: CMDSetPixel,
		Pos: pos{
			X: x,
			Y: y,
		},
		Col: Pixel{
			R: r,
			G: g,
			B: b,
		},
	}
	return struc.Pack(c.sock, &sp)
}

// SetAllPixels allows you to set the color of all pixels in a single call
func (c *Client) SetAllPixels(ps [64]Pixel) error {
	for i := range ps {
		if ps[i].R > 255 || ps[i].G > 255 || ps[i].B > 255 {
			return fmt.Errorf("pixel %v out of range, r,g,b must be 0..255, got: r:%v g:%v b:%v", ps[i].R, ps[i].G, ps[i].B)
		}
	}
	sp := setAll{
		Code:   CMDSetAllPixels,
		Pixels: ps,
	}
	return struc.Pack(c.sock, &sp)
}

// Show the pixels written to the buffer
func (c *Client) Show() error {
	s := show{
		Code: CMDShow,
	}
	return struc.Pack(c.sock, &s)
}

// Sets all pixels to 0,0,0 wrapper for Client.SetAllPixels([64]Pixel{})
func (c *Client) Clear() error {
	return c.SetAllPixels([64]Pixel{})
}

// DeMatrix converts an 8x8 Pixel grid into a [64]Pixel for use by Client.SetAllPixels
func DeMatrix(m [8][8]Pixel) [64]Pixel {
	ps := [64]Pixel{}
	n := 0
	for i := range m {
		for j := range m[i] {
			ps[n] = m[i][j]
			n++
		}
	}
	return ps

}
