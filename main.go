package main

import (
	"Estaci/scenes"
	"fmt"
	"fyne.io/fyne/v2"
)

func main() {
	main := scenes.NewMainScene("Estacionamiento del centro", fyne.NewSize(1240, 720))
	main.ShowScene()
	fmt.Println("Jalo")

}
