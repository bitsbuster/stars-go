package game

import (
	"math"

	"github.com/bitsbuster/stars-go/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite   *ebiten.Image
	position Point
	rotation float64
}

func NewPlayer() *Player {

	sprite := assets.PlayerSprite
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	// halfH := float64(bounds.Dy()) / 2

	pos := Point{
		X: ScreenWidth/2 - halfW,
		Y: float64(ScreenHeight - 20 - sprite.Bounds().Dy()),
	}
	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) UpdateNotUsed() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}
func (p *Player) Update() {
	speed := 5.0

	var delta Point

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.position.X += delta.X
	p.position.Y += delta.Y

	// p.updateRotation()
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	// Tamaño deseado
	newWidth := 75.0  // ancho deseado
	newHeight := 75.0 // alto deseado

	// Escala calculada en base al tamaño original
	scaleX := newWidth / float64(p.sprite.Bounds().Dx())
	scaleY := newHeight / float64(p.sprite.Bounds().Dy())
	op.GeoM.Scale(scaleX, scaleY)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(assets.PlayerSprite, op)
}

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
