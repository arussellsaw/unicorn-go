package unicorn

import (
	"net"

	"github.com/lunixbochs/struc"
)

var (
	SocketPath = "/var/run/unicornd.socket"
)

type Client struct {
	Path string
	sock net.Conn
}

func (c *Client) Connect() error {
	var err error
	c.sock, err = net.Dial("unix", SocketPath)
	return err
}

func (c *Client) SetBrightness(v uint) error {
	b := brightness{
		Code: CMDSetBrightness,
		Val:  v,
	}
	return struc.Pack(c.sock, &b)
}

func (c *Client) SetPixel(x, y, r, g, b uint) error {
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

func (c *Client) SetAllPixels(ps [64]Pixel) error {
	sp := setAll{
		Code:   CMDSetAllPixels,
		Pixels: ps,
	}
	return struc.Pack(c.sock, &sp)
}

func (c *Client) Show() error {
	s := show{
		Code: CMDShow,
	}
	return struc.Pack(c.sock, &s)
}

func (c *Client) Clear() error {
	sp := setAll{
		Code: CMDSetAllPixels,
	}
	return struc.Pack(c.sock, &sp)
}

// DeMatrix converts an 8x8 Pixel grid into a [64]Pixel for use by Client.SetAllPixels
func DeMatrix(m [8][8]Pixel) [64]Pixel {
	return [64]Pixel{}
}
