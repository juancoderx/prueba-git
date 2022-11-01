package main

import "fmt"

func (a Almacen) actualizarProducto(p Producto) {
	var found bool

	for i := 0; i < len(a); i++ {
		if a[i].nombre == p.nombre {
			found = true

			a[i] = p

			break
		}
	}

	if !found {
		fmt.Println("No se encontro el producto a actualizar")
	}
}
