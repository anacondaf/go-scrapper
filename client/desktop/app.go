package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	a := app.New()
	window := a.NewWindow("Hello World")

	window.Resize(fyne.NewSize(400, 400))

	window.SetContent(
		container.NewVBox(
			redButton(),
			redButton(),
			redButton(),
			redButton(),
		))

	window.ShowAndRun()
}

// first colored button
func redButton() *fyne.Container {
	btn := widget.NewButton("Visit", nil)

	btnColor := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 0, A: 255})

	container1 := container.New(
		layout.NewMaxLayout(),
		btnColor,
		btn,
	)

	// our button is ready
	return container1
}
