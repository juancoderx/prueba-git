package main

import (
	"strings"
)

func (a Almacen) buscarProducto(buscadorProducto string) (productos []Producto) {
	buscadorProducto = strings.ToLower(buscadorProducto)

	for i := 0; i < len(a); i++ {
		if a[i].nombre == buscadorProducto {
			productos = append(productos, a[i])
		}
	}

	return productos
}
