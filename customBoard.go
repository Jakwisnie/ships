package main

import (
	"fmt"
	"github.com/lxn/walk"
)

func NewCustomWidget() *CustomWidget {
	return &CustomWidget{
		colors: make(map[string]walk.Color),
	}
}

func (cw *CustomWidget) DrawFull(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	cellSize := 25
	boardSize := 10

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			rect := walk.Rectangle{
				X:      i * cellSize,
				Y:      j * cellSize,
				Width:  cellSize,
				Height: cellSize,
			}

			cellName := fmt.Sprintf("%c%d", 'A'+i, 10-j)

			var brush *walk.SolidColorBrush
			var err error
			if color, ok := cw.colors[cellName]; ok {
				brush, err = walk.NewSolidColorBrush(color)
			} else {
				brush, err = walk.NewSolidColorBrush(walk.RGB(0, 0, 0))
			}
			if err != nil {
				return err
			}
			defer brush.Dispose()

			if err := canvas.FillRectangle(brush, rect); err != nil {
				return err
			}

			pen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(128, 128, 128))
			if err != nil {
				return err
			}
			defer pen.Dispose()
			if err := canvas.DrawRectangle(pen, rect); err != nil {
				return err
			}
		}
	}
	return nil
}
