package canvas

import "fyne.io/fyne/v2"

type CanvasMixin struct {
	canvas fyne.Canvas
}

func (c *CanvasMixin) SetCanvas(canvas fyne.Canvas, setup func()) {
	isNew := c.canvas != canvas
	c.canvas = canvas
	if isNew {
		setup()
	}
}

func (c *CanvasMixin) Canvas() fyne.Canvas {
	return c.canvas
}
