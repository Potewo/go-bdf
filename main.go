package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Glyph struct {
	charCode string
	bitMap []string
}


func main() {
	filename := "font.bit"
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	glyphs := map[string]Glyph{}
	currentCharCode := ""
	isScanning := false

	counter := 0
	for scanner.Scan() {
		counter++
		if counter > 200 {
			break
		}
		row := scanner.Text()
		fmt.Println(row)
		codes := strings.Split(row, " ")
		if codes[0] == "ENDCHAR" {
			isScanning = false
		} else if isScanning {
			g := glyphs[currentCharCode]
			g.bitMap = append(g.bitMap, codes[0])
			glyphs[currentCharCode] = g
		} else if codes[0] == "STARTCHAR" {
			currentCharCode = codes[1]
			glyphs[currentCharCode] = Glyph{charCode: currentCharCode}
		} else if codes[0] == "BITMAP" {
			isScanning = true
		}
	}

	fmt.Printf("%#v", glyphs)

	if err = scanner.Err(); err != nil {
		return
	}
}
