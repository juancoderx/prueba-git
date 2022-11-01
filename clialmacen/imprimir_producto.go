package main

import "fmt"

func (a Almacen) imprimir() {
	for i := 0; i < len(a); i++ {
		fmt.Println("Productos registrados en Almacen:", i)
		a[i].imprimir()
		fmt.Println("----------------------------")
	}
}
