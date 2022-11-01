package main

import (
	"fmt"
	"strings"
)

func (a *Almacen) eliminarProducto(eliminarNombre, eliminarMarca string) {
	eliminarNombre = strings.ToLower(eliminarNombre)
	eliminarMarca = strings.ToLower(eliminarMarca)

	var found bool

	for i := 0; i < len(*a); i++ {
		found = true

		if (*a)[i].nombre == eliminarNombre && (*a)[i].marca == eliminarMarca {
			*a = append((*a)[:i], (*a)[i+1:]...)
		}
	}

	if !found {
		fmt.Println("Producto a eliminar, no encontrado")
	}
}
