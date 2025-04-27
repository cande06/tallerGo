package main

import (
	"fmt"
	"os"
	"proyecto_go/ex2/bases"

	"math/rand"

	"github.com/fatih/color"
)

func main() {
	// color.Green("hyola")
	const maxOpt = 4
	var opt int
	var devices []bases.Dispositivo

	for {
		fmt.Println("\tCONSOLA\n\n Ingrese una opción:")
		fmt.Println("[1] Crear un nuevo dispositivo")
		fmt.Println("[2] Encender o apagar un dispositivo")
		fmt.Println("[3] Ver todos los dispositivos")
		fmt.Println("[4] Test")
		fmt.Println("\n[0] Salir")
		fmt.Scanln(&opt)

		if opt < 0 || opt > maxOpt {
			fmt.Println("...valor inválido. Ingrese una opción...")
			continue
		}

		switch opt {
		case 1:
			newDevice(&devices)
			// break
		case 2:
			test()
		case 3:
		case 4:
			test()
		case 0:
			os.Exit(0)
		}
	}
}

func newDevice(list *[]bases.Dispositivo) {
	var d bases.Dispositivo
	var name string
	fmt.Println("--Crear nuevo dispositivo--")
	fmt.Print("Ingrese un nombre para el dispositivo: ")
	_, _ = fmt.Scanf("%s", &name)

	d = bases.Dispositivo{
		Nombre: name, Estado: false,
	}
	fmt.Print(d.Nombre)
	_ = append(*list, d)
}

func test() {
	var devices = []bases.Dispositivo{
		{Nombre: "Lampara", Estado: true},
		{Nombre: "TV", Estado: false},
		{Nombre: "Estereo", Estado: false},
		{Nombre: "Router", Estado: true},
		{Nombre: "AC", Estado: true},
		{Nombre: "Cerradura", Estado: true},
		{Nombre: "Garage", Estado: false},
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
