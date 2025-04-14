package main

import "fmt"

func main() {
	// cómo ingresan los datos? -> como en C !!

	//## TEST ARRAYS
	// scoringList := []int{2, 2, 3, 5, 4, 1, 2, 3, 4, 1}
	// scoringList := []int{4, 3, 5, 5, 4, 3, 2, 5, 4, 1}

	//## INGRESO MANUAL
	var score int
	var scoringList []int //como definir tamaño? -> las listas son infinitas daa
	total := 10           //cantidad de respuestas a procesar

	for i := 0; i < total; i++ {
		fmt.Println("En una escala del 1 al 5 (siendo 5 el mejor)\n ¿Que tan satisfecho se encuentra con el servicio? :)")
		fmt.Scanln(&score)
		scoringList = append(scoringList, score)
		// fmt.Println(score)
	}

	//contar concurrencia de puntajes + contar cant de pos vs neg = 2 contadores?
	//como contar concurrencia? usar array??
	contPos := 0
	contNeg := 0

	var conc []int
	for i := 0; i < 5; i++ { //5 opciones de puntaje
		conc = append(conc, 0) // Agrega elementos al slice dinámicamente
	}

	for _, score := range scoringList {
		switch score {
		case 1: //puntaje 1
			contNeg++
			conc[0] += 1
		case 2: //puntaje 2
			contNeg++
			conc[1] += 1
		case 3: //puntaje 3
			contNeg++
			conc[2] += 1
		case 4: //puntaje 4
			contPos++
			conc[3] += 1
		case 5: //puntaje 5
			contPos++
			conc[4] += 1
		}
	}

	//## MOSTRAR CONCURRENCIA
	fmt.Println("-RESULTADOS-")
	for key, score := range conc {
		fmt.Printf("%d: Votado %d veces\n", key+1, score)
	}

	//## MENSAJE FINAL
	if contPos > contNeg {
		fmt.Println("\n== ¡Buen resultado!")
	} else {
		fmt.Println("\n==Resultado mejorable")
	}
}
