package main

import (
	"github.com/kidpeterpan/index-generator/generator"
)

func main() {
	g := generator.Generator{ InputFilePath: "./input/input.txt"}
	lines := g.ReadInputFile()
	generator.GenerateMDIndex(lines)
	generator.GenerateHTMLHeader(lines)
	generator.WriteOutputFile("hello")
}

