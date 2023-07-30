package src

import (
	GameUtils "arrow-war/src/utils"
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	keys       []ebiten.Key
	x          float32
	y          float32
	speed      int
	screen     *ebiten.Image
	bullets    []*Bullet
	gun_angelX int
	gun_angelY int
}

func (player *Player) draw() {
	player.keys = inpututil.AppendPressedKeys(player.keys[:0])
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrintAt(player.screen, fmt.Sprintf("Bullets: %v", len(player.bullets)), 5, 0)
	ebitenutil.DebugPrintAt(player.screen, fmt.Sprintf("MouseX: %v MouseY: %v", x, y), 5, 20)
	ebitenutil.DebugPrintAt(player.screen, fmt.Sprintf("width: %v height: %v", GameUtils.SCREEN_WIDTH, GameUtils.SCREEN_HEIGHT), 5, 40)
	ebitenutil.DebugPrintAt(player.screen, fmt.Sprintf("canMove: %v ", player.canMove()), 5, 60)
	player.gun_angelX = x
	player.gun_angelY = y
	vector.StrokeCircle(player.screen, float32(player.x), float32(player.y), float32(25), float32(3), color.RGBA{255, 255, 255, 255}, false)

	// Get the angle between the player's position and the mouse position
	angle := math.Atan2(float64(player.y)-float64(y), float64(player.x)-float64(x))

	player.gun_angelX = x
	player.gun_angelY = y

	vector.StrokeCircle(player.screen, float32(player.x), float32(player.y), float32(25), float32(3), color.RGBA{255, 255, 255, 255}, false)

	// Calculate the endpoint of the line based on the angle
	gunLength := -30
	gunX := int(float64(player.x) + float64(gunLength)*math.Cos(angle))
	gunY := int(float64(player.y) + float64(gunLength)*math.Sin(angle))

	// Draw the line to represent the gun
	vector.StrokeLine(player.screen, float32(player.x), float32(player.y), float32(gunX), float32(gunY), float32(3), color.RGBA{255, 255, 255, 255}, false)
	for i, bullet := range player.bullets {
		if bullet.IsDone {
			if i < len(player.bullets) {
				player.bullets = append(player.bullets[:i], player.bullets[i+1:]...)
				i--
			}
		}
		bullet.draw()

	}
}

func (player *Player) Create(name string, screen *ebiten.Image) {
	player.screen = screen
	player.speed = 5
	// player.x = 200
	// player.y = 200

	for _, key := range player.keys {
		if key == ebiten.Key(GameUtils.KEY_W) {
			player.goUp()
		}
		if key == ebiten.Key(GameUtils.KEY_S) {
			player.goDown()
		}
		if key == ebiten.Key(GameUtils.KEY_D) {
			player.goRight()
		}
		if key == ebiten.Key(GameUtils.KEY_A) {
			player.goLeft()
		}

	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		player.fire()
	}
	player.draw()

}

func (player *Player) goLeft() {

	player.x += float32(player.speed)
}
func (player *Player) goRight() {

	player.x -= float32(player.speed)
}
func (player *Player) goUp() {

	player.y -= float32(player.speed)
}
func (player *Player) goDown() {

	player.y += float32(player.speed)
}
func (player *Player) canMove() bool {
	return (player.x < float32(GameUtils.SCREEN_WIDTH)) && (player.x != 0) || (player.y < float32(GameUtils.SCREEN_HEIGHT)) && (player.y != 0)
}
func (player *Player) fire() {
	bullet := Bullet{Screen: player.screen, owner: player}
	bullet.Create(player.x, player.y, float32(player.gun_angelX), float32(player.gun_angelY))
	player.bullets = append(player.bullets, &bullet)
}
