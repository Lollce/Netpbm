package main

import (
	"fmt"
)

func main() {

	imagePBM, err := pbm.ReadPBM("PBM.go")
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'image :", err)
		return imagePBM
	}
}
