package main

import (
	"testing"
)

func TestExtractHTMLElements(t *testing.T) {
	testContent := `
		<li><a href=#the-div-element><span class=secno>4.4.1</span> The <code>div</code> element</a>
		li><a href=#the-sub-and-sup-elements><span class=secno>4.5.19</span> The <code>sub</code> and <code>sup</code> elements</a>
		<li><a href=#the-h1-h2-h3-h4-h5-h6-elements><span class=secno>4.3.6</span> The <code>h1, h2, h3, h4, h5, and h6</code> elements</a>
		<li><a href=#the-iframe-element><span class=secno>4.8.5</span> The
  		<code>iframe</code> element</a>
	`

	elements := extractHTMLElements(testContent)

	expectedElements := []string{"div", "sub", "sup", "h1", "h2", "h3", "h4", "h5", "h6", "iframe"}

	if len(elements) < len(expectedElements) {
		t.Errorf("Expected at least %d elements, but got %d", len(expectedElements), len(elements))
	}

	// Convert to map for easier lookup
	foundElements := make(map[string]bool)
	for _, elem := range elements {
		foundElements[elem.Name] = true
	}

	// Check that all expected elements were found
	for _, expected := range expectedElements {
		if !foundElements[expected] {
			t.Errorf("Expected element '%s' was not found", expected)
		}
	}
}
