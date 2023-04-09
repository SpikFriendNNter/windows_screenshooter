package main

import (
	"windows_screenshooter/print"

	"log"
	"fmt"
	"image/png"
	"os"
	"time"
	"strconv"
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
