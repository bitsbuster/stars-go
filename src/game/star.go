package game

import (
	"github.com/bitsbuster/stars-go/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const ()

type Star struct {
	score    int16
	sprite   *ebiten.Image
	position Point
	movement Point
}

func NewStar(baseVelocity float64) *Star {

	pos := Point{
		X: float64((ScreenWidth-40-assets.Soil.Bounds().Dx()*2))*rnd.Float64() + 20 + float64(assets.Soil.Bounds().Dx()),
		Y: 10.0,
	}

	velocity := baseVelocity + rnd.Float64()*0.5

	movement := Point{
		X: 0,
		Y: velocity,
	}
	star := &Star{
		position: pos,
		movement: movement,
	}
	value := rnd.Intn(10)
	if value < 9 {
		star.sprite = assets.StarGold
		star.score = 1
	} else {
		star.sprite = assets.StarBronze
		star.score = -2
	}

	return star
}

func (s *Star) Update() {
	s.position.Y += s.movement.Y
}

func (s *Star) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(s.position.X, s.position.Y)

	screen.DrawImage(s.sprite, op)
}

func (s *Star) Collider() Rect {
	bounds := s.sprite.Bounds()

	return NewRect(
		s.position.X,
		s.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
