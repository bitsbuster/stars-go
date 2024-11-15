package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/bitsbuster/stars-go/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	starCreateTime    = 3 * time.Second
	starUpdateTime    = 10 * time.Millisecond
	baseStarVelocity  = 0.0005
	starSpeedUpAmount = 0.5
	starSpeedUpTime   = 5 * time.Second
)

var (
	rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMilli()))
)

type Game struct {
	player         *Player
	starCreateTime *Timer
	starUpdateTime *Timer
	stars          []*Star
	velocities     []float64
	// bullets          []*Bullet

	score int

	baseVelocity  float64
	velocityTimer *Timer
}

func NewGame() *Game {
	g := &Game{
		player:         NewPlayer(),
		starCreateTime: NewTimer(starCreateTime),
		starUpdateTime: NewTimer(starUpdateTime),
		velocityTimer:  NewTimer(starSpeedUpTime),
		baseVelocity:   baseStarVelocity,
		velocities:     make([]float64, 0),
	}

	for i := 0.0005; i < 0.0050; i += 0.0005 {
		g.velocities = append(g.velocities, i)
	}

	return g
}
func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		value := rnd.Intn(len(g.velocities))

		g.baseVelocity = baseStarVelocity + g.velocities[value]
	}

	g.player.Update()

	g.starCreateTime.Update()

	if g.starCreateTime.IsReady() {
		g.starCreateTime.Reset()

		g.stars = append(g.stars, NewStar(g.baseVelocity))

	}

	g.starUpdateTime.Update()
	if g.starUpdateTime.IsReady() {
		g.starUpdateTime.Reset()
		for _, m := range g.stars {
			m.Update()
		}
	}

	for i, s := range g.stars {
		if s.Collider().Intersects(g.player.Collider()) {
			fmt.Printf("Colisión")
			g.score += int(s.score)
			g.stars = g.removeStarFromArray(i)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x63, 0x85, 0x5b, 0xff})
	drawSoil(screen)

	for _, m := range g.stars {
		m.Draw(screen)
	}

	g.player.Draw(screen)

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func drawSoil(screen *ebiten.Image) {
	sprite := assets.Soil
	bounds := sprite.Bounds()
	i := 0
	for {
		y := float64(bounds.Dy() * i)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(20, y)
		screen.DrawImage(assets.Soil, op)
		op2 := &ebiten.DrawImageOptions{}
		op2.GeoM.Translate(ScreenWidth-20.0-float64(sprite.Bounds().Dx()), y)
		screen.DrawImage(sprite, op2)

		if y > ScreenHeight {
			break
		}
		i++
	}
}

func (g *Game) removeStarFromArray(index int) []*Star {
	if index < 0 || index >= len(g.stars) {
		return g.stars
	}

	return append(g.stars[:index], g.stars[index+1:]...)
}
