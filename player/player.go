package player

import (
  "math"
  "github.com/hajimehoshi/ebiten"
)

type Player struct {
	Image				*ebiten.Image
	Xpos, Ypos 	float64
	Speed 			float64
  Rot         float64
  RotSpeed    float64
}

// Keyboard events and movements
func (player *Player) MovePlayer(){
	if ebiten.IsKeyPressed(ebiten.KeyUp){
		player.Ypos -= math.Cos(player.Rot*math.Pi/180) * player.Speed
    player.Xpos += math.Sin(player.Rot*math.Pi/180) * player.Speed
  }
	if ebiten.IsKeyPressed(ebiten.KeyDown){
    player.Ypos += math.Cos(player.Rot*math.Pi/180) * player.Speed
    player.Xpos -= math.Sin(player.Rot*math.Pi/180) * player.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		player.Rot -= player.RotSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
    player.Rot += player.RotSpeed
	}
}
