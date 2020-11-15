package main

import (
	"github.com/kidpeterpan/index-generator/generator"
)

func main() {
	g := generator.Generator{ InputFilePath: "./input/unit-testing.txt", OutputFilePath: "./output/output.txt"}
	lines := g.ReadInputFile()
	mdIndex := generator.GenerateMDIndex(lines)
	htmlHeader :=  generator.GenerateHTMLHeader(lines)
	g.WriteOutputFile(mdIndex.String() + htmlHeader.String())
}

