package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	//me intriguei com a origem dos valores de entrada
	//até ver as 'const' dx e dy declaradas na lib
	//fmt.Println(dx, dy)

	img := make([][]uint8, dy)

	for x := range img {
		img[x] = make([]uint8, dx)
		for y := range img[x] {
			//funções sugeridas:
			//img[x][y] = uint8(x ^ y)
			//img[x][y] = uint8((x + y) / 2)
			img[x][y] = uint8(x * y)
		}
	}
	return img
}

// retorna uma imagem em base64
// remova o "IMAGE:" do retorno e veja a imagem num conversor base64 => imagem
func main() {
	pic.Show(Pic)
}
