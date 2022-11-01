package main

import (
	"fmt"
	"strings"
)

func (a Almacen) buscarProducto(buscadorProducto string) {
	buscadorProducto = strings.ToLower(buscadorProducto)

	var found bool

	for i := 0; i < len(a); i++ {
		if a[i].nombre == buscadorProducto {
			found = true

			fmt.Println("Producto encontrado en el almacen, posicion:", i+1)
			a[i].presentacion()
			fmt.Println()

		}
	}

	if !found {
		fmt.Println("El producto solicitado, no esta registrado")
	}
}
