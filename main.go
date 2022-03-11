package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth, screenHeight = 640, 480
)

// Create the player class
type player struct {
	image				*ebiten.Image
	xpos, ypos 	float64
	speed 			float64
}

// vars
var (
	err					error
	background 	*ebiten.Image
	spaceShip		*ebiten.Image
	playerOne		player
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
	playerOne = player{spaceShip, screenWidth/2.0, screenHeight/2.0, 4}
}
// Keyboard events and movements
func movePlayer(x float64, y float64){
	if ebiten.IsKeyPressed(ebiten.KeyUp){
		playerOne.ypos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown){
		playerOne.ypos += playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		playerOne.xpos -= playerOne.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
		playerOne.xpos += playerOne.speed
	}
}

// Update
func update (screen *ebiten.Image) error {
	movePlayer(playerOne.xpos, playerOne.ypos)

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	// Draw our background
	bgOp := &ebiten.DrawImageOptions{}
	bgOp.GeoM.Translate(0, 0)
	screen.DrawImage(background, bgOp)
	// Draw our player
	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Scale(0.1, 0.1)
	playerOp.GeoM.Translate(playerOne.xpos, playerOne.ypos)
	screen.DrawImage(playerOne.image, playerOp)

	return nil
}

// Main Loop
func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "My Game!"); err != nil {
		log.Fatal(err)
	}
}
