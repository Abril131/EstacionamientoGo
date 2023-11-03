package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Car struct {
	ini        int
	status     bool
	imgC       *canvas.Image
	imgC2      *canvas.Image
	posy, posx int
	//height float32
}

func (c *Car) IniC() int {
	return c.ini
}

func (c *Car) SetIniC(ini int) {
	c.ini = ini
}

func (c *Car) Posx() int {
	return c.posx
}

func (c *Car) SetPosx(posx int) {
	c.posx = posx
}

func (c *Car) Posy() int {
	return c.posy
}

func (c *Car) SetPosy(posy int) {
	c.posy = posy
}

func (c *Car) Status() bool {
	return c.status
}

func (c *Car) SetStatus(status bool) {
	c.status = status
}

func (c *Car) Img() *canvas.Image {
	return c.imgC
}

func (c *Car) SetImg(imgC *canvas.Image) {
	c.imgC = imgC
}

func NewCar() *Car {
	return &Car{
		status: false,
		imgC:   canvas.NewImageFromFile("assets/c1.png"),
		imgC2:  canvas.NewImageFromFile("assets/c1.png"),
		posy:   0,
		posx:   0,
	}
}
