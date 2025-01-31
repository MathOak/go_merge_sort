package main

import (
	"fmt"
	"sync"
)

// Função principal do Merge Sort
func mergeSort(arr []int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	// Se o array tiver um ou nenhum elemento, já está ordenado
	if len(arr) <= 1 {
		return arr
	}

	// Divide o array ao meio
	mid := len(arr) / 2

	// Cria WaitGroups para as goroutines
	var leftWg, rightWg sync.WaitGroup
	leftWg.Add(1)
	rightWg.Add(1)

	// Ordena recursivamente as duas metades em goroutines
	var left, right []int
	go func() {
		left = mergeSort(arr[:mid], &leftWg)
	}()
	go func() {
		right = mergeSort(arr[mid:], &rightWg)
	}()

	// Espera as goroutines terminarem
	leftWg.Wait()
	rightWg.Wait()

	// Junta as duas metades ordenadas
	return merge(left, right)
}

// Função que junta dois arrays ordenados
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	/* Aplicar logica nesse bloco de codigo para tratar com Objetos diversos */
	// Compara os elementos dos dois arrays e adiciona o menor ao resultado
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Adiciona os elementos restantes de 'left' ao resultado
	result = append(result, left[i:]...)
	// Adiciona os elementos restantes de 'right' ao resultado
	result = append(result, right[j:]...)

	return result
}

func main() {
	// Array de exemplo a ser ordenado
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)

	// Cria um WaitGroup para a goroutine principal do mergeSort
	var wg sync.WaitGroup
	wg.Add(1)

	// Chama a função de ordenação em uma goroutine
	var sortedArr []int
	go func() {
		sortedArr = mergeSort(arr, &wg)
	}()

	// Espera a goroutine principal terminar
	wg.Wait()

	// Imprime o array ordenado
	fmt.Println("Sorted array:", sortedArr)
}
