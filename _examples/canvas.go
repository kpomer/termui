//go:build ignore
// +build ignore

package main

import (
	"image"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.SetLine(image.Pt(0, 0), image.Pt(10, 20), ui.ColorWhite)

	k := widgets.NewParagraph()
	k.Title = "Key Input"
	k.Text = "Press [any key](fg:red) to QUIT THE DEMO"
	k.SetRect(60, 0, 80, 5)
	k.BorderStyle.Bg = ui.ColorRed

	ui.Render(c, k)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
