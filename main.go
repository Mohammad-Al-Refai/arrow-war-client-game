package main

import (
	src "arrow-war/src/player"
	utils "arrow-war/src/utils"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

var player = src.Player{}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Set(500, 500, color.RGBA{255, 25, 255, 255})
	player.Create("Mohammad", screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	utils.SCREEN_WIDTH = outsideWidth
	utils.SCREEN_HEIGHT = outsideHeight
	return outsideWidth, outsideHeight

}

func main() {
	ebiten.SetWindowTitle("Arrow War")

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
