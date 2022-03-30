package player

import (
  "fmt"
  "math"
  "github.com/hajimehoshi/ebiten"
)

type Player struct {
	Image				*ebiten.Image
	Xpos, Ypos 	float64
	Speed 			float64
  Rot         float64
  RotSpeed    float64
  Acc         float64
  MaxSpeed    float64
}

// Keyboard events and movements
func (p *Player) MovePlayer(){
  if p.Acc <= p.MaxSpeed{
    if ebiten.IsKeyPressed(ebiten.KeyUp){
      p.Acc += p.Speed
    }
  }
  if p.Acc >= p.MaxSpeed{
  	if ebiten.IsKeyPressed(ebiten.KeyDown){
      p.Acc -= p.Speed
  	}
  }
	if ebiten.IsKeyPressed(ebiten.KeyLeft){
		p.Rot -= p.RotSpeed * p.Acc
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight){
    p.Rot += p.RotSpeed * p.Acc
	}

  if p.Acc != 0{
    // if p.Acc <= p.MaxSpeed || p.Acc >= -p.MaxSpeed{
    p.Ypos -= math.Cos(p.Rot*math.Pi/180) * p.Acc
    p.Xpos += math.Sin(p.Rot*math.Pi/180) * p.Acc

    if p.Acc > 0 {
      p.Acc -= 1
    }
    if p.Acc < 0 {
      p.Acc += 1
    }
  }
  fmt.Printf("%d\n", p.Acc)
}
