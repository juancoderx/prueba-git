package main

import (
	"fmt"
	"strconv"
)

type option uint8

const (
	menuOption1 option = iota + 1
	menuOption2
	menuOption3
	menuOption4
	menuOption5
	menuOption6
	menuOption7
)

const (
	ingresarProducto int = iota + 1
	imprimirAlmacen
	buscarProducto
	actualizarProducto
	eliminarProducto
	crearInventario
	buscarProducto_Puntero
)

func main() {
	var listadoAlmacenes []Almacen

	eleccion := 1

	for eleccion >= 1 {
		fmt.Println("Bienvenido a Almacen")
		fmt.Println(listadoAlmacenes)

		fmt.Println(menuOption1, "Ingresar producto")
		fmt.Println(menuOption2, "Imprimir Almacen")
		fmt.Println(menuOption3, "Buscar producto")
		fmt.Println(menuOption4, "Actualizar producto")
		fmt.Println(menuOption5, "Eliminar producto")
		fmt.Println(menuOption6, "Crear inventario")
		fmt.Println(menuOption7, "BuscarProducto(Puntero)")

		fmt.Print(">")
		fmt.Scan(&eleccion)

		switch eleccion {
		case ingresarProducto:
			var ingresarProducto Producto

			fmt.Println("Ingrese el Nombre del Producto")
			fmt.Print(">")
			fmt.Scan(&ingresarProducto.nombre)

			fmt.Println("Ingresa la Marca del Producto")
			fmt.Print(">")
			fmt.Scan(&ingresarProducto.marca)

			fmt.Println("Ingrese el Precio del Producto")
			fmt.Print(">")
			fmt.Scan(&ingresarProducto.precio)

			fmt.Println("¿Cuantas unidades va a ingresar del Producto?")
			fmt.Print(">")
			fmt.Scan(&ingresarProducto.unidades)

			dispAlmacen := "Almacenes: "
			for i := 0; i < len(listadoAlmacenes); i++ {
				dispAlmacen += strconv.Itoa(i)

				if i < len(listadoAlmacenes)-1 {
					dispAlmacen += ", "
				}
			}

			var numAlmacen int

			fmt.Println("¿En que almacen guardara el producto?", dispAlmacen)
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Posicion de almacen no declarado")

				break
			}

			listadoAlmacenes[numAlmacen].agregarProductos(ingresarProducto)

		case imprimirAlmacen:
			var numAlmacen int

			fmt.Println("Ingrese el numero del almacen a buscar")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			fmt.Println(listadoAlmacenes[numAlmacen].presentacion())

		case buscarProducto:
			var (
				searchProduct string
				numAlmacen    int
			)

			fmt.Println("Ingrese el nombre del producto")
			fmt.Print(">")
			fmt.Scan(&searchProduct)

			fmt.Println("¿En cual almancen esta el producto?")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			infoProducto := listadoAlmacenes[numAlmacen].buscarProductos(searchProduct)

			if len(infoProducto) == 0 {
				fmt.Println("No se encontraron productos")

				break
			}

			for i := 0; i < len(infoProducto); i++ {
				fmt.Println(infoProducto[i].presentacion())
			}

		case actualizarProducto:
			var (
				updateProduct Producto
				numAlmacen    int
			)

			fmt.Println("Producto a actualizar")
			fmt.Print(">")
			fmt.Scan(&updateProduct.nombre)

			fmt.Println("Marca del producto")
			fmt.Print(">")
			fmt.Scan(&updateProduct.marca)

			fmt.Println("¿En cual almancen esta el producto?")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			fmt.Println("Ingrese nuevo precio")
			fmt.Print(">")
			fmt.Scan(&updateProduct.precio)

			fmt.Println("Ingrese las unidades existentes")
			fmt.Print(">")
			fmt.Scan(&updateProduct.unidades)

			if exist := listadoAlmacenes[numAlmacen].actualizarProducto(updateProduct); !exist {
				fmt.Println("No se encontro el producto a actualizar")
			}

		case eliminarProducto:
			var (
				eliminarNombre, eliminarMarca string
				numAlmacen                    int
			)

			fmt.Println("Indique el nombre del producto a eliminar")
			fmt.Print(">")
			fmt.Scan(&eliminarNombre)

			fmt.Println("Indique la marca del producto a eliminar")
			fmt.Print(">")
			fmt.Scan(&eliminarMarca)

			fmt.Println("¿En cual almancen esta el producto?")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			if exist := listadoAlmacenes[numAlmacen].eliminarProducto(eliminarNombre, eliminarMarca); !exist {
				fmt.Println("Producto a eliminar, no encontrado")
			}

		case crearInventario:
			listadoAlmacenes = append(listadoAlmacenes, Almacen{})

		case buscarProducto_Puntero:
			var (
				buscadorProducto string
				numAlmacen       int
			)
			fmt.Println("Ingrese producto")
			fmt.Print(">")
			fmt.Scan(&buscadorProducto)

			fmt.Println("Ingrese el almancen donde esta el producto")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			infoProducto := listadoAlmacenes[numAlmacen].buscadorProducto(buscadorProducto)

			if infoProducto == nil {
				fmt.Println("No se encontro el producto")

				break
			}

			fmt.Println(infoProducto.presentacion())

		}
	}
}
