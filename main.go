package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
// 1. Función para generar una matriz de tamaño aleatorio nxm

func generarArray(filas, columnas int) [][]int {
	matriz := make([][]int, filas)
	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}

// Función de imprimir la matriz

func imprimirArray(matriz [][]int) {
	for _, fila := range matriz {
		fmt.Println(fila)
	}
}

//Funcion de multiplicación concurrente

func multiplicarMatricesConcurrente(matrizA, matrizB [][]int, resultado chan<- [][]int) {
	filasA := len(matrizA)
	columnasA := len(matrizA[0])
	columnasB := len(matrizB[0])

	resultadoMatriz := make([][]int, filasA)
	for i := 0; i < filasA; i++ {
		resultadoMatriz[i] = make([]int, columnasB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < columnasB; j++ {
			go func(i, j int) {
				suma := 0
				for k := 0; k < columnasA; k++ {
					suma += matrizA[i][k] * matrizB[k][j]
				}
				resultadoMatriz[i][j] = suma
			}(i, j)
		}
	}
	//funcion de espera de Goroutines
	for x := 0; x < filasA; x++ {
		for z := 0; z < columnasB; z++ {
			<-time.After(10 * time.Millisecond)
		}
	}

	resultado <- resultadoMatriz
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//definir tamaño de matriz

	filasA := 3
	columnasA := 2
	filasB := 2
	columnasB := 3

	//generar matrices

	matrizA := generarArray(filasA, columnasA)
	matrizB := generarArray(filasB, columnasB)

	//imprimir las matrices iniciales

	fmt.Println("Matriz A: ")
	imprimirArray(matrizA)
	fmt.Println("Matriz B: ")
	imprimirArray(matrizB)

	// Crear canal resultado

	resultado := make(chan [][]int)

	//multiplicar matrices concurrentes

	go multiplicarMatricesConcurrente(matrizA, matrizB, resultado)

	//Espera e impresion del resultado

	fmt.Println("Resultado del producto cruz: ")
	fmt.Println(<-resultado)
}

*/

/*
//2. Productor-Consumidor con goroutines

// Funcion para generar aleatoriamente 10 notas de 0 al 10 y enviarlos por un canal
func enviarNotas(c chan int) {
	for i := 0; i < 10; i++ {
		notas := rand.Intn(10)
		c <- notas
	}
	close(c)
}

// Funcion para recibir las notas por canal y calcular su promedio
func calcularPromedio(c chan int) {
	acumulador := 0
	contador := 0
	promedio := 0.0
	for numero := range c {
		acumulador += numero
		contador += 1
	}
	fmt.Printf("La suma de los notas recibidas es: %d y su promedio es: %f.", acumulador, promedio)
}

func main() {
	// Creo un canal
	c := make(chan int)

	go enviarNotas(c)
	calcularPromedio(c)
}
*/

/*
//3. Sincronización de tareas paralelas

// Genera aletariamente 3 notas del 0 al 10 de un alumno y las envia por canal
func enviarNotas(c chan int) {
	for i := 0; i < 3; i++ {
		c <- rand.Intn(10)
	}
	close(c)
}

// Funcion para recibir las 3 notas por canal y calcular su promedio
func calcularPromedio(c chan int) {
	contador := 0
	acumulador := 0
	promedio := 0.0
	for nota := range c {
		contador += 1
		acumulador += nota
		fmt.Printf("Orden de la nota: %d, nota: %d.\n", contador, nota)
	}
	promedio = float64(acumulador) / float64(contador)
	// Llamo a la funcion para determinar si esta aprobado o no
	esAprobado(promedio)
}

// Funcion para determinar si esta aprobado o no
func esAprobado(nota float64) {
	if nota >= 7 {
		fmt.Printf("El promedio es: %f, el alumno ha aprobado", nota)
	} else {
		fmt.Printf("El promedio es: %f, el alumno ha reprobado", nota)
	}
}

func main() {
	// Creo un canal
	c := make(chan int)

	// Llamada a las funciones gorutine
	go enviarNotas(c)
	calcularPromedio(c)
}
*/

//4. Simulación de carrera con Goroutines

// Funciones corredor, el prmero en terminar cierra el canal
func corredorUno(c chan int) {
	c <- 1
	rand.Seed(time.Now().UnixNano()) // Semilla de tiempo individual para cada corredor
	espera := rand.Intn(500)
	time.Sleep(time.Duration(espera) * time.Millisecond)
	close(c)
}

func corredorDos(c chan int) {
	c <- 2
	rand.Seed(time.Now().UnixNano())
	espera := rand.Intn(500)
	time.Sleep(time.Duration(espera) * time.Millisecond)
	close(c)
}

func corredorTres(c chan int) {
	c <- 3
	rand.Seed(time.Now().UnixNano())
	espera := rand.Intn(500)
	time.Sleep(time.Duration(espera) * time.Millisecond)
	close(c)
}

func corredorCuatro(c chan int) {
	c <- 4
	rand.Seed(time.Now().UnixNano())
	espera := rand.Intn(500)
	time.Sleep(time.Duration(espera) * time.Millisecond)
}

// Funcion arbitro, determina el ganador que es el primer en cerrar el canal
func arbitro(c chan int) {
	ganador := <-c
	fmt.Printf("El ganador de la carrera es el corredor número %d\n", ganador)
}

func main() {

	// Canal que simula la pista para todos los corredores
	pista := make(chan int)

	// Simulo 10 carreras
	for i := 0; i < 10; i++ {
		go corredorUno(pista)
		go corredorDos(pista)
		go corredorTres(pista)
		go corredorCuatro(pista)
		arbitro(pista)
	}
}
