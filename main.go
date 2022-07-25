package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/qeesung/asciiplayer/pkg/player"
	"github.com/qeesung/image2ascii/convert"
)

func main() {
	// get current working dir
	dir, err := os.Getwd()
	// if not able to do that panic
	if err != nil {
		panic(err)
	}
	// create a terminal player and pass path of the file in our working directory
	// in my case it's nyan.gif
	// if you create a cli filepath.Join() part should not be necessary
	// look up to docs thoıgh
	terminalPlayer, supported := player.NewTerminalPlayer(filepath.Join(dir, "nyan.gif"))
	// if file type is not supported then panic
	if !supported {
		panic("Terminal is not supported")
	}
	// necessary for image2ascii lib
	// look up https://github.com/qeesung/image2ascii
	// conver options part for best configuration
	convertOptions := convert.Options{
		Ratio:           1.0,
		FixedWidth:      -1,
		FixedHeight:     -1,
		FitScreen:       true,  // only work on terminal
		StretchedScreen: false, // only work on terminal
		Colored:         true,  // only work on terminal
		Reversed:        false,
	}
	// you pass Delay between frames
	// you can manipulate fps
	playOptions := player.PlayOptions{
		Delay: time.Duration(0.01*1000) * time.Millisecond,
	}
	// set playerOptions.Options to converOptions created before
	playOptions.Options = convertOptions
	// run this thing up baby 🔥
	terminalPlayer.Play(filepath.Join(dir, "nyan.gif"), &playOptions)
}
