package main

func (a Almacen) actualizarProducto(p Producto) (exist bool) {
	for i := 0; i < len(a); i++ {
		if a[i].nombre == p.nombre {
			a[i] = p

			return true
		}
	}

	return false
}
