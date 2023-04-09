package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	myWindow.Resize(fyne.NewSize(500, 500))

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text := canvas.NewText("Hello World", green)
	textContainer := container.New(layout.NewCenterLayout(), text)

	button := widget.NewButton("Click me", func() {
		fmt.Println("Yes I am clicked!")
	})

	content := container.New(layout.NewVBoxLayout(), textContainer, button)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
