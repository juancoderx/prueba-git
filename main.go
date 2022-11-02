package main

import (
	"fmt"
	"strconv"
)

func main() {
	var listadoAlmacenes []Almacen

	eleccion := 1

	for eleccion >= 1 {
		fmt.Println("Bienvenido a Almacen")
		fmt.Println(listadoAlmacenes)

		fmt.Println("1.Ingresar producto")
		fmt.Println("2.Imprimir Almacen")
		fmt.Println("3.Buscar producto")
		fmt.Println("4.Actualizar producto")
		fmt.Println("5.Eliminar producto")
		fmt.Println("6.Crear inventario")

		fmt.Print(">")
		fmt.Scan(&eleccion)

		switch eleccion {
		case 1:
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

		case 2:
			var numAlmacen int

			fmt.Println("Ingrese el numero del almacen a buscar")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			fmt.Println(listadoAlmacenes[numAlmacen].presentacion())

		case 3:
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

			if listadoAlmacenes[numAlmacen].buscarProducto(searchProduct) == nil {
				fmt.Println("No se encontraron productos")

				break
			}

			fmt.Println(listadoAlmacenes[numAlmacen].presentacion())

		case 4:
			var (
				editProduct Producto
				numAlmacen  int
			)

			fmt.Println("Producto a actualizar")
			fmt.Print(">")
			fmt.Scan(&editProduct.nombre)

			fmt.Println("¿En cual almancen esta el producto?")
			fmt.Print(">")
			fmt.Scan(&numAlmacen)

			if numAlmacen >= len(listadoAlmacenes) {
				fmt.Println("Numero de almacen inexistente")

				break
			}

			if !listadoAlmacenes[numAlmacen].actualizarProducto(editProduct) {
				fmt.Println("No se encontro el producto a actualizar")

				break
			}

			fmt.Println("Ingrese marca")
			fmt.Print(">")
			fmt.Scan(&editProduct.marca)

			fmt.Println("Ingrese nuevo precio")
			fmt.Print(">")
			fmt.Scan(&editProduct.precio)

			fmt.Println("Ingrese las unidades existentes")
			fmt.Print(">")
			fmt.Scan(&editProduct.unidades)

			fmt.Println(listadoAlmacenes[numAlmacen].actualizarProducto(editProduct))

		case 5:
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

			if !listadoAlmacenes[numAlmacen].eliminarProducto(eliminarNombre, eliminarMarca) {
				fmt.Println("Producto a eliminar, no encontrado")

				break
			}

			fmt.Println(listadoAlmacenes[numAlmacen].eliminarProducto(eliminarNombre, eliminarMarca))

		case 6:
			listadoAlmacenes = append(listadoAlmacenes, Almacen{})

		}
	}
}
