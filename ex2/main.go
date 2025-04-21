package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Green("hyola")

	var test Dispositivo
	fmt.Println(test.EstadoActual())
}

// ## ESTRUCTURA
type Dispositivo struct {
	Nombre string
	Estado bool
}

// ## INTERFAZ
type Controlable interface {
	Encender() error
	Apagar() error
	EstadoActual() string
}

// ## FUNCIONES
func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya está encendido")
	}
	d.Estado = false
	return errors.New("")
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo ya está apagado")
	}
	d.Estado = false
	return errors.New("")
}

func (d Dispositivo) EstadoActual() string {
	if !d.Estado {
		return "Apagado"
	} else {
		return "Encendido"
	}
}
