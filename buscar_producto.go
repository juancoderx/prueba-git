package main

import (
	"strings"
)

func (a Almacen) buscarProductos(buscadorProducto string) (productos []Producto) {
	buscadorProducto = strings.ToLower(buscadorProducto)

	for i := 0; i < len(a); i++ {
		if a[i].nombre == buscadorProducto {
			productos = append(productos, a[i])
		}
	}

	return productos
}

func (a Almacen) buscadorProducto(buscadorProducto string) (productos *Producto) {
	buscadorProducto = strings.ToLower(buscadorProducto)

	for i := 0; i < len(a); i++ {
		if (a)[i].nombre == buscadorProducto {
			productos = &a[i]

			return productos
		}
	}

	return productos
}
