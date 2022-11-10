package ratio

import (
	"fmt"
	"image/png"
	"net/http"
)

func Run() {
	url := "https://raw.githubusercontent.com/mouredev/mouredev/master/mouredev_github_profile.png"
	width, height := getImageResolution(url)
	xRatio, yRatio := CalcRatio(width, height)
	fmt.Printf("%d:%d", xRatio, yRatio)
}

func CalcRatio(width int, height int) (int, int) {
	var xRatio, yRatio int
	for true {
		if width%2 == 0 && height%2 == 0 {
			xRatio = width / 2
			yRatio = height / 2
			return CalcRatio(xRatio, yRatio)
		} else if width%3 == 0 && height%3 == 0 {
			xRatio = width / 3
			yRatio = height / 3
			return CalcRatio(xRatio, yRatio)
		} else if width%5 == 0 && height%5 == 0 {
			xRatio = width / 5
			yRatio = height / 5
			return CalcRatio(xRatio, yRatio)
		} else {
			return width, height
		}
	}
	return xRatio, yRatio
}

func getImageResolution(url string) (width int, height int) {
	resp, getErr := http.Get(url)
	if getErr != nil {
		fmt.Printf("Couldn't get -> %s", getErr)
	}
	defer resp.Body.Close()
	imageData, decodeErr := png.DecodeConfig(resp.Body)
	if decodeErr != nil {
		fmt.Printf("Couldn't decode -> %s", getErr)
	}
	return imageData.Width, imageData.Height
}
