package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Car2 struct {
	ini        int
	status     bool
	imgC       *canvas.Image
	posy, posx int
	//height float32
}

func (c2 *Car2) IniC2() int {
	return c2.ini
}

func (c2 *Car2) SetIniC2(ini int) {
	c2.ini = ini
}

func (c2 *Car2) Posx2() int {
	return c2.posx
}

func (c2 *Car2) SetPosx2(posx int) {
	c2.posx = posx
}

func (c2 *Car2) Posy2() int {
	return c2.posy
}

func (c2 *Car2) SetPosy2(posy int) {
	c2.posy = posy
}

func (c2 *Car2) Status2() bool {
	return c2.status
}

func (c2 *Car2) SetStatus2(status bool) {
	c2.status = status
}

func (c2 *Car2) Img2() *canvas.Image {
	return c2.imgC
}

func (c2 *Car2) SetImg2(imgC *canvas.Image) {
	c2.imgC = imgC
}

func NewCar2() *Car2 {
	return &Car2{
		status: false,
		imgC:   canvas.NewImageFromFile("assets/c2.png"),
		posy:   0,
		posx:   0,
	}
}
