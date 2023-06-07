package main

import (
	"github.com/SpikFriendNNter/windows_screenshooter/print"
	"log"
	"image/png"
	"os"
)

func main() {
	img, err := print.Take_Screenshot()
	file, err := os.OpenFile("screenshot.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	if err = png.Encode(file, img); err != nil {
		log.Fatal("ERROR SAVING IMAGE: ", err)
	}
}
