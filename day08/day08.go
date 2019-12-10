package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func buildLayers(imgData string, w int, h int) []string {
	layers := []string{}
	for i := 0; i < len(imgData); i += (w * h) {
		layers = append(layers, imgData[i:i+w*h])
	}
	return layers
}

func decode(layers []string) string {
	topLayer := make([]string, len(layers[0]))
	for i := range layers[0] {
		layer := 0
		var color string
		for {
			if layer < len(layers) {
				color = string(layers[layer][i])
			} else {
				color = string(layers[len(layers)-1][i])
			}
			if color == "0" || color == "1" {
				topLayer[i] = color
				break
			} else {
				layer++
			}
		}
	}
	return strings.Join(topLayer, "")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := 25
	h := 6

	// part 1
	var layers []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		imgData := scanner.Text()
		layers = buildLayers(imgData, w, h)
	}

	layer := 0
	zeroCount := w * h
	for i, l := range layers {
		layerCount := 0
		for _, pixel := range l {
			if string(pixel) == "0" {
				layerCount++
			}
		}
		if layerCount < zeroCount {
			zeroCount = layerCount
			layer = i
		}
	}

	oneCount := 0
	twoCount := 0
	for _, d := range layers[layer] {
		if string(d) == "1" {
			oneCount++
		} else if string(d) == "2" {
			twoCount++
		}
	}

	fmt.Printf("Day 08, Part 1: %v", oneCount*twoCount)
	fmt.Println()

	// part 2
	image := decode(layers)
	fmt.Println("Day 08, Part 2:")
	for i := 0; i < len(image); i += w {
		re := regexp.MustCompile(`0`)
		fmt.Println(re.ReplaceAllString(image[i:i+w], " "))
	}
	fmt.Println()
}
