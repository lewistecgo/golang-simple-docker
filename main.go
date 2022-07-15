/**
go mod init golang
docker build -t my-golang-app .
docker run -it -v "$(pwd):/app" my-golang-app
**/

package main

import "fmt"

func main() {
	fmt.Println("Ejemplo de GO!")
	fmt.Println("Ejemplo de GO en Docker!")
	fmt.Println("Ejemplo de GO en Docker en Windows 11!")
	fmt.Println("Ejemplo de GO en Docker en Windows 11 & WSL!")

}
