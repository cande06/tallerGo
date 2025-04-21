package main

import (
	"errors"
	"fmt"

	"math/rand"

	"github.com/fatih/color"
)

func main() {
	// color.Green("hyola")

	var devices = []Dispositivo{
		{"Lampara", true},
		{"TV", false},
		{"Estereo", false},
		{"Router", true},
		{"AC", true},
		{"Cerradura", true},
		{"Garage", false},
	}

	for _, device := range devices {
		fmt.Println("\n", device.Nombre)

		v := rand.Intn(10)
		if (v % 2) == 0 {
			fmt.Println("Encendiendo...")

			on := device.Encender()
			if on != nil {
				fmt.Println("error->", on)
			}

			fmt.Print("Estado: ")
			color.Green(device.EstadoActual())

		} else {
			fmt.Println("Apagando...")

			off := device.Apagar()
			if off != nil {
				fmt.Println("error->", off)
			}

			fmt.Print("Estado: ")
			color.Red(device.EstadoActual())

		}
	}
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

// Devolver estado del dispositivo
func (d Dispositivo) EstadoActual() string {
	if !d.Estado {
		return "Apagado"
	} else {
		return "Encendido"
	}
}

// Definir errores
func (d *Dispositivo) Encender() error {
	if d.Estado {
		return errors.New("el dispositivo ya se encuentra encendido")
	}
	d.Estado = true
	return nil
}

func (d *Dispositivo) Apagar() error {
	if !d.Estado {
		return errors.New("el dispositivo ya se encuentra apagado")
	}
	d.Estado = false
	return nil
}
