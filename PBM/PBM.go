package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

/* main() {
	pbmCall, err := ReadPBM("Image.pbm")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier PBM:", err)
		return
	}
	display(pbmCall.data)
}

*/

func ReadPBM(filename string) (*PBM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	magicNumber := scanner.Text()
	if magicNumber != "P1" && magicNumber != "P4" {
		return nil, errors.New("unsupported file type")
	}

	scanner.Scan()
	dimensions := strings.Fields(scanner.Text())
	if len(dimensions) != 2 {
		return nil, errors.New("invalid image dimensions")
	}

	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	var data [][]bool
	for scanner.Scan() {
		line := scanner.Text()
		if magicNumber == "P1" {
			row := make([]bool, width)
			for i, char := range strings.Fields(line) {
				pixel, _ := strconv.Atoi(char)
				row[i] = pixel == 1
			}
			data = append(data, row)
		} else if magicNumber == "P4" {
			reader := bufio.NewReader(file)
			reader.Discard(width % 8)
			for y := 0; y < height; y++ {
				row := make([]bool, width)
				for x := 0; x < width; x += 8 {
					b, err := reader.ReadByte()
					if err != nil {
						return nil, err
					}
					for i := 0; i < 8; i++ {
						row[x+i] = b&(1<<(7-i)) != 0
					}
				}
				data = append(data, row)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &PBM{
		data:        data,
		width:       width,
		height:      height,
		magicNumber: magicNumber,
	}, nil
}

func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

func (pbm *PBM) At(x, y int) bool {
	if len(pbm.data) == 0 || x < 0 || y < 0 || x >= pbm.width || y >= pbm.height {
		return pbm.data[y][x]
	}
	return false

}

func (pbm *PBM) Set(x, y int, value bool) {
	pbm.data[y][x] = value
}

func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n%d %d\n", pbm.magicNumber, pbm.width, pbm.height)

	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Fprint(file, "1 ")
			} else {
				fmt.Fprint(file, "0 ")
			}
		}
		fmt.Fprintln(file)
	}
	fmt.Println("Données écrites avec succès dans le fichier PBM.")
	return nil
}

func (pbm *PBM) Invert() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width; j++ {
			pbm.data[i][j] = !pbm.data[i][j]
		}
	}
}

func (pbm *PBM) Flip() {
	for i := 0; i < pbm.height; i++ {
		for j := 0; j < pbm.width/2; j++ {
			pbm.data[i][j], pbm.data[i][pbm.width-j-1] = pbm.data[i][pbm.width-j-1], pbm.data[i][j]
		}
	}
}

func (pbm *PBM) Flop() {
	for i := 0; i < pbm.height/2; i++ {
		pbm.data[i], pbm.data[pbm.height-i-1] = pbm.data[pbm.height-i-1], pbm.data[i]
		if numRows == 0 {
			return
		}
		for i := 0; i < numRows/2; i++ {
			pbm.data[i], pbm.data[numRows-i-1] = pbm.data[numRows-i-1], pbm.data[i]
		}
	}
}

func (pbm *PBM) SetMagicNumber(magicNumber string) {
	pbm.magicNumber = magicNumber
	if magicNumber != "P1" && magicNumber != "P4" {
		fmt.Println("unsupported format")
	} else {
		fmt.Println(magicNumber)
	}
}

func ReadimagePBM() {
	imagePBM, err := pbm.ReadPBM("test.pbm")
	if err != nil {
		fmt.Println("Erreur lors de la lecture de l'image :", err)
		return imagePBM
	}
}
