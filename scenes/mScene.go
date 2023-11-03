package scenes

import (
	"Estaci/models"
	"fyne.io/fyne/v2"
	"math/rand"
	"time"
)

const capM = 10

type EstScenes struct {
	window fyne.Window
	app    fyne.App
	lugar  *models.Lugar
	cont   *fyne.Container
	// Agregar un semáforo
	semaphore chan struct{}
}

func NewEstScenes(window fyne.Window, app fyne.App, cont *fyne.Container) *EstScenes {
	return &EstScenes{
		window:    window,
		app:       app,
		cont:      cont,
		lugar:     models.NewLugar(capM),
		semaphore: make(chan struct{}, 1), // Inicializar el semáforo con capacidad 1
	}
}

/*func (me *EstScenes) Wi() {
	botonIniciar := widget.NewButton("Start Game", func() {
		me.Status(me.window)
	})
	botonIniciar.Resize(fyne.NewSize(150, 30))
	botonIniciar.Move(fyne.NewPos(100, 10))

	botonDetener := widget.NewButton("Stop Game", me.Status)
	botonDetener.Resize(fyne.NewSize(150, 30))
	botonDetener.Move(fyne.NewPos(300, 10))

	me.window.SetContent(container.NewWithoutLayout(botonIniciar, botonDetener))

}*/

func (me *EstScenes) Status() {
	go func() {
		for i := 0; i < 100; i++ {
			go me.Show(nil)
			time.Sleep(time.Duration(rand.ExpFloat64()/0.5) * time.Second)
		}
	}()
}

func (me *EstScenes) Show(fyne.Window) {
	c := models.NewCar()
	c2 := models.NewCar2()
	c.SetPosx(0)
	c.SetPosy(280)
	intn := rand.Intn(2)
	if intn == 1 {
		c.Img().Resize(fyne.NewSize(96, 96))
		c.Img().Move(fyne.NewPos(0, 600))
		me.cont.Add(c.Img())
	} else {
		c2.Img2().Resize(fyne.NewSize(96, 96))
		c2.Img2().Move(fyne.NewPos(0, 600))
		me.cont.Add(c2.Img2())
	}
	me.cont.Refresh()

	me.lugar.Entrada(c, c2, c.Img(), c2.Img2(), me.cont)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	me.lugar.Salida(c, c2, c.Img(), c2.Img2(), me.cont)
}
