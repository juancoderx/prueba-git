package main

import (
	"strconv"
)

func (p Producto) presentacion() (presenProducto string) {
	// fmt.Println("Nombre del Producto:", p.nombre)
	// fmt.Println("Marca:", p.marca)
	// fmt.Println("Precio del Producto:", p.precio)
	// fmt.Println("Unidades disponible:", p.unidades)
	presenProducto = "Nombre: " + p.nombre + "\n"
	presenProducto += "Marca: " + p.marca + "\n"
	presenProducto += "Precio: " + strconv.Itoa(p.precio) + "\n"
	presenProducto += "Unidades: " + strconv.Itoa(p.unidades) + "\n"

	return presenProducto
}

func (a Almacen) presentacion() (presenAlmacen string) {
	for i := 0; i < len(a); i++ {
		presenAlmacen += a[i].presentacion()
		presenAlmacen += "-------------------" + "\n"

	}
	return presenAlmacen
}
