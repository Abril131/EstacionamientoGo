package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainScene struct {
	window fyne.Window
	app    fyne.App
	cont   *fyne.Container
}

func NewMainScene(title string, size fyne.Size) *MainScene {
	mainA := app.New()
	mainW := mainA.NewWindow(title)
	mainW.CenterOnScreen()
	mainW.SetFixedSize(true)
	mainW.Resize(size)
	return &MainScene{
		app:    mainA,
		window: mainW,
		cont:   container.NewWithoutLayout(),
	}
}

func (m *MainScene) ShowScene() {
	image := canvas.NewImageFromFile("./assets/fondo.jpeg")
	image.Resize(fyne.NewSize(1240, 720))
	image.Move(fyne.NewPos(0, 0))
	m.cont.Add(image)
	m.window.SetContent(m.cont)
	scene := NewEstScenes(m.window, m.app, m.cont)

	go scene.Status()
	m.window.ShowAndRun()

}
