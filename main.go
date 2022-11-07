package main

import (
	"fmt"
	"strconv"
)

type option int8

const (
	opcionIngresarProducto option = iota + 1
	opcionImprimirAlmacen
	opcionBuscarProducto
	opcionActualizarProducto
	opcionEliminarProducto
	opcionCrearInventario
	opcionBuscarProductoPuntero
)

func main() {
	var (
		listadoAlmacenes []Almacen
		eleccion         option = 1
	)

	for eleccion >= 1 {
		fmt.Println("Bienvenido a Almacen")
		fmt.Println(listadoAlmacenes)

		fmt.Println(opcionIngresarProducto, ".Ingresar producto")
		fmt.Println(opcionImprimirAlmacen, ".Imprimir Almacen")
		fmt.Println(opcionBuscarProducto, ".Buscar producto")
		fmt.Println(opcionActualizarProducto, ".Actualizar producto")
		fmt.Println(opcionEliminarProducto, ".Eliminar producto")
		fmt.Println(opcionCrearInventario, ".Crear inventario")
		fmt.Println(opcionBuscarProductoPuntero, ".BuscarProducto(Puntero)")

		fmt.Print(">")
		fmt.Scan(&eleccion)

		switch eleccion {
		case opcionIngresarProducto:
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

			if len(listadoAlmacenes) == 0 {
				listadoAlmacenes = append(listadoAlmacenes, Almacen{})

				listadoAlmacenes[0].agregarProductos(ingresarProducto)

				break
			}

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

		case opcionImprimirAlmacen:
			var numAlmacen int

			fmt.Println("Ingrese el numero del almacen a buscar")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			fmt.Println(listadoAlmacenes[numAlmacen].presentacion())

		case opcionBuscarProducto:
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

		case opcionActualizarProducto:
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

		case opcionEliminarProducto:
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

		case opcionCrearInventario:
			listadoAlmacenes = append(listadoAlmacenes, Almacen{})

		case opcionBuscarProductoPuntero:
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
