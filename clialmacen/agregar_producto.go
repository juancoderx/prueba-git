func (a *Almacen) agregarProductos(p Producto) {
	p.nombre = strings.ToLower(p.nombre)
	p.marca = strings.ToLower(p.marca)

	var found bool

	for i := 0; i < len(*a); i++ {
		if (*a)[i].nombre == p.nombre && (*a)[i].marca == p.marca {
			found = true

			(*a)[i].unidades += p.unidades

			if (*a)[i].precio != p.precio {
				(*a)[i].precio = p.precio
			}

			break
		}
	}

	if !found {
		*a = append(*a, p)
	}
}