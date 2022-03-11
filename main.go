package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err 				error
	background 	*ebiten.Image
	pa 					*ebiten.Image
)

func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	pa, _, err = ebitenutil.NewImageFromFile("assets/pa.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}
func update (screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
	op.GeoM.Translate(0, 50)
	background.DrawImage(pa, op)
	ebitenutil.DebugPrint(pa, "Heyya")
	return nil
}


func main() {
	if err := ebiten.Run(update, 640, 480, 1, "My Game!"); err != nil {
		log.Fatal(err)
	}
}
