package src

import (
	"arrow-war/src/utils"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	owner      *Player
	Screen     *ebiten.Image
	X, Y       float32
	dirX, dirY float32
	IsDone     bool
	color      color.Color
}

func (bullet *Bullet) Create(x float32, y float32, dirX float32, dirY float32) Bullet {
	bullet.X = x
	bullet.Y = y
	bullet.dirX = dirX
	bullet.dirY = dirY
	bullet.color = color.RGBA{255, 0, 0, 255}
	return *bullet
}
func (bullet *Bullet) update() {
	var speed float64 = 22
	deltaX := float64(bullet.dirX - bullet.X)
	deltaY := float64(bullet.dirY - bullet.Y)

	// Normalize the vector to get its unit direction
	length := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	unitDeltaX := float32(deltaX) / float32(length)
	unitDeltaY := float32(deltaY) / float32(length)

	// Add the velocity to the current velocity of the bullet
	bullet.X += unitDeltaX * float32(speed)
	bullet.Y += unitDeltaY * float32(speed)

	// Update the direction of the bullet towards the mouse cursor
	bullet.dirX += unitDeltaX * float32(speed)
	bullet.dirY += unitDeltaY * float32(speed)

	if bullet.X < 0 || bullet.X > float32(utils.SCREEN_WIDTH) || bullet.Y < 0 || bullet.Y > float32(utils.SCREEN_HEIGHT) {
		bullet.IsDone = true
	}
}
func (bullet *Bullet) draw() {
	bullet.update()
	vector.DrawFilledCircle(bullet.Screen, bullet.X, bullet.Y, float32(5), bullet.color, true)
}
