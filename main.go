package main

import (
	"math"
	"log"
	"ebiten/player"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)


const (
	screenWidth, screenHeight = 640, 480
)


// vars
var (
	err					error
	background 	*ebiten.Image
	spaceShip		*ebiten.Image
	playerOne		player.Player
)

// Error Check
func checkForErr() {
	if err != nil {
		log.Fatal(err)
	}
}

// This code will be ran once at startup
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	checkForErr()
	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceShip.png", ebiten.FilterDefault)
	checkForErr()
	playerOne = player.Player{spaceShip, screenWidth/2.0, screenHeight/2.0, 4, 0, 2.5}
}

// Update
func update (screen *ebiten.Image) error {
	playerOne.MovePlayer()
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw our background
	bgOp := &ebiten.DrawImageOptions{}
	bgOp.GeoM.Translate(0, 0)
	screen.DrawImage(background, bgOp)
	// Draw our player
	w, h := playerOne.Image.Size()
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(-float64(w)/2, -float64(h)-2)
	playerOp.GeoM.Scale(0.1, 0.1)
	playerOp.GeoM.Rotate(playerOne.Rot*math.Pi/180)
	playerOp.GeoM.Translate(playerOne.Xpos, playerOne.Ypos)
	screen.DrawImage(playerOne.Image, playerOp)

	return nil
}


// Main Loop
func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "My Game!"); err != nil {
		log.Fatal(err)
	}

	// ebiten.SetWindowSize(screenWidth, screenHeight)
	// ebiten.SetWindowTitle("just a game")
	// if err := ebiten.RunGame(&Game{}); err != nil {
	// 	log.Fatal(err)
	// }

}
