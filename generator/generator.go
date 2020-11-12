package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Generator struct {
	InputFilePath string
	OutputFilePath string
}

// ReadInputFile read input file and return slice of content line
func (g Generator) ReadInputFile() []string {
	f, err := os.Open(g.InputFilePath)
	if err != nil {
		fmt.Println("err occurred when open an input file: ", err)
		os.Exit(1)
	}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var lines []string
	for fScanner.Scan() {
		lines = append(lines, fScanner.Text())
	}

	return lines
}

// GenerateMDIndex generate index for markdown via input file
func GenerateMDIndex(lines []string) strings.Builder {
	var mdIndex strings.Builder
	mdIndex.WriteString("![xxx_logo](/media/cover/xxx_logo.png)\n")
	mdIndex.WriteString("\n")
	for _, line := range lines {
		match, err := regexp.Match("(^\\d+.*)", []byte(line))
		if err != nil {
			fmt.Println("err occurred when matching first line of chapter: ", err)
		}
		if match {
			words := strings.Split(line, " ")
			id := "#"
			for i, word := range words {
				if i == 0 {
					continue
				}
				id += strings.ToLower(word)
				if (i + 1) != len(words) {
					id += "-"
				}
			}
			chString := "- [" + line + "](" + id + ")"
			mdIndex.WriteString(chString + "\n")
		} else {
			words := strings.Split(line, " ")
			id := "#"
			for i, word := range words {
				id +=  strings.ToLower(word)
				if (i + 1) != len(words) {
					id += "-"
				}
			}
			subChString := "\t+ [" + line + "](" + id + ")"
			mdIndex.WriteString(subChString + "\n")
		}
	}
	return mdIndex
}

// WriteOutputFile write an output result to an output file
func (g Generator) WriteOutputFile(content string) {
	f, err := os.OpenFile(g.OutputFilePath,os.O_RDWR,0755)
	if err != nil {
		fmt.Println("error occurred when opening the output file:",err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("error occurred when writing output file:",err)
	}
}

// GenerateHTMLHeader generate HTML content header for link with markdown index
func GenerateHTMLHeader(lines []string) strings.Builder {
	var htmlHeader strings.Builder
	for _, line := range lines {
		match, err := regexp.Match("(^\\d+.*)", []byte(line))
		if err != nil {
			fmt.Println("err occurred when matching first line of chapter: ", err)
		}
		if match {
			htmlHeader.WriteString("\n")
			htmlHeader.WriteString("---\n")
			htmlHeader.WriteString("\n")
			words := strings.Split(line, " ")
			id := ""
			for i, word := range words {
				if i == 0 {
					continue
				}
				id += strings.ToLower(word)
				if (i + 1) != len(words) {
					id += "-"
				}
			}
			chString := "<h2 id=" + "\"" +id + "\">" + line + "</h2>"
			htmlHeader.WriteString(chString + "\n")
			htmlHeader.WriteString("\n")
		} else {
			words := strings.Split(line, " ")
			id := ""
			for i, word := range words {
				id +=  strings.ToLower(word)
				if (i + 1) != len(words) {
					id += "-"
				}
			}
			subChString := "<h3 id=" + "\"" +id + "\">" + line + "</h3>"
			htmlHeader.WriteString(subChString +"\n")
			htmlHeader.WriteString("\n")
			htmlHeader.WriteString("// TODO: add of " + line + "\n")
			htmlHeader.WriteString("\n")
		}
	}
	return htmlHeader
}


