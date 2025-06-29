package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Element represents an HTML element with its name and description
type Element struct {
	Name   string
	Source string
}

var showDebug bool
var showHelp bool

func init() {
}

func main() {
	flag.BoolVar(&showDebug, "debug", false, "Enable debug mode")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.Parse()

	if showHelp || len(flag.Args()) == 0 || len(flag.Args()) > 1 {
		fmt.Println("Usage:\n\thtml5-standard-reader [-debug] <path/to/HTML Standard.html>")
		fmt.Println("Reads the HTML 5 Standard file and prints found HTML elements.")
		os.Exit(0)
	}

	sourceFile := flag.Args()[0]

	// Read the HTML Standard file from the root directory
	file, err := os.Open(sourceFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Extract HTML elements from the content
	elements := extractHTMLElements(string(content))

	// Sort elements alphabetically by name
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Name < elements[j].Name
	})

	for _, element := range elements {
		if showDebug {
			fmt.Printf("%-20s %s\n", element.Name, element.Source)
		} else {
			fmt.Println(element.Name)
		}
	}

	if showDebug {
		fmt.Printf("\nTotal elements found: %d\n", len(elements))
	}
}

// extractHTMLElements parses the HTML content and extracts all HTML element definitions
func extractHTMLElements(content string) []Element {
	var elements []Element
	elementSet := make(map[string]bool) // To avoid duplicates

	// Table of contents entries like "The <code>element</code> element"
	pattern := regexp.MustCompile(`<a href=\S*?#the-([a-z0-9,-]+)-elements?[^>]*><span class=secno>[^<]*</span>\s*The\s*<code>`)
	matches := pattern.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) < 2 {
			panic("Invalid regexp match. This is a bug in the code.")
		}
		elementsSlug := strings.TrimSpace(match[1])

		// Split by continuous non-word symbols
		parts := regexp.MustCompile(`[\W]+`).Split(elementsSlug, -1)
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" && part != "and" && !elementSet[part] {
				// the word "and" can appear in the list of elements, but we don't want to include it
				elements = append(elements, Element{Name: part, Source: "a " + elementsSlug + " (" + match[0] + ")"})
				elementSet[part] = true
			}
		}
	}

	return elements
}
