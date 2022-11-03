package main

import (
	"strings"
)

func (a *Almacen) eliminarProducto(eliminarNombre, eliminarMarca string) (exist bool) {
	eliminarNombre = strings.ToLower(eliminarNombre)
	eliminarMarca = strings.ToLower(eliminarMarca)

	for i := 0; i < len(*a); i++ {
		if (*a)[i].nombre == eliminarNombre && (*a)[i].marca == eliminarMarca {
			*a = append((*a)[:i], (*a)[i+1:]...)

			return true
		}
	}

	return false
}
