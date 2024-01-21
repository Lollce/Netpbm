package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           int
}

func main() {
	Pbmcall, err := ReadPGM("Imagepgm.pbm")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier PBM:", err)
		return
	}
	display(Pbmcall.data)
}

func ReadPGM(filename string) (*PGM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var width, height, max int
	var data [][]uint8

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	magicNumber := scanner.Text()
	if magicNumber != "P2" && magicNumber != "P5" {
		return nil, errors.New("type de fichier non pris en charge")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			_, err := fmt.Sscanf(line, "%d %d", &width, &height)
			if err == nil {
				break
			} else {
				fmt.Println("Largeur ou hauteur invalide :", err)
			}
		}
	}

	scanner.Scan()
	max, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, errors.New("valeur maximale de pixel invalide")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P2" {
			row := make([]uint8, 0)
			for _, char := range strings.Fields(line) {
				pixel, err := strconv.Atoi(char)
				if err != nil {
					fmt.Println("Erreur de conversion en entier :", err)
				}
				if pixel >= 0 && pixel <= max {
					row = append(row, uint8(pixel))
				} else {
					fmt.Println("Valeur de pixel invalide :", pixel)
				}
			}
			data = append(data, row)
		}
	}

	return &PGM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
		max:         max,
	}, nil
}

func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

// At returns the value of the pixel at (x, y).
func (pgm *PGM) At(x, y int) uint8 {
	return pgm.data[x][y]
}

// Set sets the value of the pixel at (x, y).
func (pgm *PGM) Set(x, y int, value uint8) {
	pgm.data[x][y] = value
}

// Save saves the PGM image to a file and returns an error if there was a problem.
func (pgm *PGM) Save(filename string) error {
	return nil
}

// Invert inverts the colors of the PGM image.
func (pgm *PGM) Invert() {
	for i := 0; i < len(pgm.data); i++ {
		for j := 0; j < len(pgm.data[i]); j++ {
			pgm.data[i][j] = 255 - pgm.data[i][j]
		}
	}
}

// Flip flips the PGM image horizontally.
func (pgm *PGM) Flip() {
	NumRows := pgm.width
	Numcolums := pgm.height
	for i := 0; i < NumRows; i++ {
		for j := 0; j < Numcolums/2; j++ {
			pgm.data[i][j], pgm.data[i][Numcolums-j-1] = pgm.data[i][Numcolums-j-1], pgm.data[i][j]
		}
	}
}

// Flop flops the PGM image vertically.
func (pgm *PGM) Flop() {
	NumRows := pgm.width
	Numcolums := pgm.height

}

func display(data [][]uint8) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			fmt.Print(data[i][j], " ")
		}
		fmt.Println()
	}
}
