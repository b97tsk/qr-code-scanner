package main

import (
	"log"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/vova616/screenshot"
)

const (
	title   = "QR Code Scanner"
	tooltip = title
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconData)
	systray.SetTitle(title)
	systray.SetTooltip(tooltip)

	scan := systray.AddMenuItem("Scan", "Capture and scan screen")
	scanDelay := systray.AddMenuItem("Scan after 4s", "Scan after 4 seconds")
	quit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-scan.ClickedCh:
				onScan()
			case <-scanDelay.ClickedCh:
				time.Sleep(4 * time.Second)
				onScan()
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {}

func onScan() {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		alert(err.Error())
		systray.Quit()

		return
	}

	qrReader := qrcode.NewQRCodeReader()

	src := gozxing.NewLuminanceSourceFromImage(img)
	bmp, _ := gozxing.NewBinaryBitmap(gozxing.NewHybridBinarizer(src))

	result, _ := qrReader.Decode(bmp, nil)
	if result == nil {
		alert("No QR code found on screen.")
		return
	}

	text := result.GetText()
	log.Println(text)
	_ = clipboard.WriteAll(text)
	_ = beeep.Notify(title, text, "")
}

func alert(message string) {
	log.Println(message)
	_ = beeep.Alert(title, message, "")
}
