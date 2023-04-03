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
	counter := 0
	for {

		img, err := print.Take_Screenshot()
		if err != nil {
			log.Fatal(err)
		}

		var filename string = "out/image" + strconv.Itoa(counter) + ".png"
		fmt.Println(filename)
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		if err = png.Encode(file, img); err != nil {
			log.Fatal("SAVING IMAGE ERROR: ", err)
		}

		fmt.Println("IMAGE SAVED.")
		file.Close()
		counter++
		time.Sleep(time.Second * 5)
	}


}
