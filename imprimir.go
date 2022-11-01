package main

import "fmt"

func (p Producto) imprimir() {
	fmt.Println("Nombre del Producto:", p.nombre)
	fmt.Println("Marca:", p.marca)
	fmt.Println("Precio del Producto:", p.precio)
	fmt.Println("Unidades disponible:", p.unidades)
}

func (a Almacen) imprimir() {
	for i := 0; i < len(a); i++ {
		fmt.Println("Productos registrados en Almacen:", i)
		a[i].imprimir()
		fmt.Println("----------------------------")
	}
}
