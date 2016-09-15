package unicorn

// types pulled and made Go compatible from https://github.com/pimoroni/unicorn-hat/tree/master/library_c/unicornd

/*
UNICORND_CMD_SET_BRIGHTNESS = 0
UNICORND_CMD_SET_PIXEL      = 1
UNICORND_CMD_SET_ALL_PIXELS = 2
UNICORND_CMD_SHOW           = 3
*/
var (
	CMDSetBrightness uint = 0
	CMDSetPixel      uint = 1
	CMDSetAllPixels  uint = 2
	CMDShow          uint = 3
)

// brightness
type brightness struct {
	//uint8_t code; // set to 0
	Code uint `struc:"uint8"`
	//double  val;
	Val uint `struc:"uint8"`
}

//set pixel

type setPixel struct {
	//uint8_t code; // set to 1
	Code uint `struc:"uint8"`
	//pos_t pos;
	Pos pos
	//col_t col;
	Col Pixel
}

//where pos_t is a struct like this:

type pos struct {
	//uint8_t x;
	X uint `struc:"uint8"`
	//uint8_t y;
	Y uint `struc:"uint8"`
}

//and where col_t is a struct like this:

type Pixel struct {
	//uint8_t r;
	R uint `struc:"uint8"`
	//uint8_t g;
	G uint `struc:"uint8"`
	//uint8_t b;
	B uint `struc:"uint8"`
}

//set all pixels

type setAll struct {
	//uint8_t code; // set to 2
	Code uint `struc:"uint8"`
	//col_t pixels[64];
	Pixels [64]Pixel
}

//show

type show struct {
	//uint8_t code; // set to 3
	Code uint `struc:"uint8"`
}
