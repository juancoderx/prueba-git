package main

import (
	"strings"
)

func (a Almacen) buscarProductos(buscadorProductos string) (productos []Producto) {
	buscadorProductos = strings.ToLower(buscadorProductos)

	for i := 0; i < len(a); i++ {
		if a[i].nombre == buscadorProductos {
			productos = append(productos, a[i])
		}
	}

	return productos
}

func (a Almacen) buscadorProducto(buscadorProducto string) (producto *Producto) {
	buscadorProducto = strings.ToLower(buscadorProducto)

	for i := 0; i < len(a); i++ {
		if (a)[i].nombre == buscadorProducto {
			return &a[i]
		}
	}

	return nil
}
