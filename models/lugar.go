package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
	"time"
)

type Lugar struct {
	Cap    int
	occ    int
	ent    sync.Mutex
	Space  []bool
	spacex []int
	spacey []int
	count  int
}

func NewLugar(cap int) *Lugar {
	return &Lugar{
		Cap:   cap,
		Space: make([]bool, cap),
		spacex: []int{
			272,
			370,
			560,
			755,
			945,
			150,
			345,
			440,
			633,
			730,
			930,
			1025,
		},
		spacey: []int{
			100,
			380,
			560,
		},
	}
}

/*
func (l *Lugar) Entrada(car1 *Car, car2 *Car2, imgC1 *canvas.Image, imgC2 *canvas.Image, sceneC *fyne.Container) {
	l.ent.Lock()
	defer l.ent.Unlock()

	time.Sleep(time.Second)
	for {
		if l.occ < l.Cap {
			break
		}
		sceneC.Remove(imgC1)
		sceneC.Remove(imgC2)
		sceneC.Refresh()
		return
	}

	for i := 0; i < l.Cap; i++ {
		if !l.Space[i] {
			l.Space[i] = true
			l.occ++
			if i < 5 {
				for j := car2.posy + 100; j < car1.posy; j += 2 {
					imgC1.Move(fyne.NewPos(float32(0), float32(j)))
					imgC2.Move(fyne.NewPos(float32(0), float32(j)))
					time.Sleep(time.Millisecond)
				}
				imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(l.spacey[0])))
				imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(l.spacey[1])))
				car1.SetIniC(i)
				car2.SetIniC2(i)
				break
			}
			for j := car1.posx; j < car1.posx+100; j += 2 {
				imgC1.Move(fyne.NewPos(float32(0), float32(j)))
				imgC2.Move(fyne.NewPos(float32(0), float32(j)))
				time.Sleep(time.Millisecond)
			}
			imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(l.spacey[0])))
			imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(l.spacey[1])))
			car1.SetIniC(i)
			car2.SetIniC2(i)
			break
		}
	}
}

func (l *Lugar) Salida(car1 *Car, car2 *Car2, imgC1 *canvas.Image, imgC2 *canvas.Image, sceneC *fyne.Container) {
	l.ent.Lock()
	defer l.ent.Unlock()
	for i := l.Cap - 1; i >= 0; i-- {
		if i == car1.ini {
			l.Space[i] = false
			l.occ--
			if i < 5 {
				for j := l.spacey[0]; j < l.spacey[0]+100; j += 2 {
					imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
					imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
					time.Sleep(time.Millisecond)
				}
				sceneC.Remove(imgC1)
				sceneC.Remove(imgC2)
				sceneC.Refresh()
				break
			}
			for j := l.spacey[1]; j > l.spacey[1]-100; j -= 2 {
				imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
				imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
				time.Sleep(time.Millisecond)
			}
			sceneC.Remove(imgC1)
			sceneC.Remove(imgC2)
			sceneC.Refresh()
			break
		}
	}
}*/

func (l *Lugar) Entrada(car1 *Car, car2 *Car2, imgC1 *canvas.Image, imgC2 *canvas.Image, sceneC *fyne.Container) {
	l.ent.Lock()
	defer l.ent.Unlock()

	time.Sleep(time.Second)

	// Calcular la fila actual
	filaActual := l.count / 5

	if filaActual >= 3 {
		sceneC.Remove(imgC1)
		sceneC.Remove(imgC2)
		sceneC.Refresh()
		return
	}

	for i := 0; i < l.Cap; i++ {
		if !l.Space[i] {
			l.Space[i] = true
			l.occ++
			l.count++

			// Calcular la posición y la fila para mover el carro
			posY := l.spacey[filaActual]
			for j := car2.posy + 100; j < car1.posy; j += 2 {
				imgC1.Move(fyne.NewPos(float32(0), float32(j)))
				imgC2.Move(fyne.NewPos(float32(0), float32(j)))
				time.Sleep(time.Millisecond)
			}
			imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(posY)))
			imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(posY)))
			car1.SetIniC(i)
			car2.SetIniC2(i)
			break
		}
	}
}

func (l *Lugar) Salida(car1 *Car, car2 *Car2, imgC1 *canvas.Image, imgC2 *canvas.Image, sceneC *fyne.Container) {
	l.ent.Lock()
	defer l.ent.Unlock()
	for i := l.Cap - 1; i >= 0; i-- {
		if i == car1.ini {
			l.Space[i] = false
			l.occ--

			// Calcular la fila actual
			filaActual := l.count / 5

			// Calcular la posición en la fila
			posY := l.spacey[filaActual]

			for j := posY + 100; j < posY+200; j += 2 {
				imgC1.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
				imgC2.Move(fyne.NewPos(float32(l.spacex[i]), float32(j)))
				time.Sleep(time.Millisecond)
			}
			sceneC.Remove(imgC1)
			sceneC.Remove(imgC2)
			sceneC.Refresh()
			break
		}
	}
}
