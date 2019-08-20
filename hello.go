package main

import (
	"gocv.io/x/gocv"
)

func helloWorld() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello cv")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
