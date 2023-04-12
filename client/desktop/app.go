package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("SysTray")

	//if desk, ok := a.(desktop.App); ok {
	//	m := fyne.NewMenu("MyApp",
	//		fyne.NewMenuItem("Show", func() {
	//			fmt.Println("Showing...")
	//		}))
	//
	//	desk.SetSystemTrayMenu(m)
	//}

	w.Resize(fyne.NewSize(1920, 1080))

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.ShowAndRun()
}
