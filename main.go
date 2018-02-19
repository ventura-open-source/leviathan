package main

import (
    "github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
    "fmt"
	"strings"
	"strconv"
    //"gopkg.in/h2non/bimg.v1"
)

var sizes = [17]string{
    "2048x1536",
    "640x480",
    "480x360",
    "226x170",
    "192x144",
    "181x136",
    "173x130",
    "160x120",
    "108x81",
    "96x72",
    "60x45",
    // new sizes for responsive web site
    "1920x1440",
    "1200x900",
    "992x744",
    "768x576",
    "544x408",
    "136x102",
}

func main() {
	// open "test.jpg"
	file, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

    for i, v := range sizes {
        //get the width and height 
        size := strings.Split(v, "x")
        w, _ := strconv.ParseUint(size[0], 10, 64)
        //h, _ := strconv.ParseUint(size[1], 10, 64)
        fmt.Printf("%d. resizing to %s \n", i, v)

        // resize to especific size using Nearest-neighbor interpolation
        // and preserve aspect ratio
        m := resize.Resize(uint(w), 0, img, resize.NearestNeighbor)

        out, err := os.Create("test_resized_" + v + ".jpg")
        if err != nil {
            log.Fatal(err)
        }
        defer out.Close()

        // write new image to file
        jpeg.Encode(out, m, nil)
    }
}
