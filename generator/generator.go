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
func GenerateMDIndex(lines []string) {
	fmt.Println("![xxx_logo](/media/cover/xxx_logo.png)")
	fmt.Print("\n")
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
			fmt.Println(chString)
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
			fmt.Println(subChString)
		}
	}
}

// GenerateHTMLHeader generate HTML content header for link with markdown index
func GenerateHTMLHeader(lines []string) {
	for _, line := range lines {
		match, err := regexp.Match("(^\\d+.*)", []byte(line))
		if err != nil {
			fmt.Println("err occurred when matching first line of chapter: ", err)
		}
		if match {
			fmt.Print("\n")
			fmt.Println("---")
			fmt.Print("\n")
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
			fmt.Println(chString)
			fmt.Print("\n")
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
			fmt.Println(subChString)
			fmt.Print("\n")
			fmt.Println("// TODO: add of",line)
			fmt.Print("\n")
		}
	}
}
