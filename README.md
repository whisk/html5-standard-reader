# HTML5 Standard Reader

A Go command-line utility that parses the HTML5 Standard document and extracts all HTML element names defined in the specification.

## Overview

This tool reads the official HTML5 Standard document (in HTML format) and extracts element names. The document can be either the [one-page version](https://html.spec.whatwg.org/) or [the multipage version](https://html.spec.whatwg.org/multipage/).

## Installation

```bash
go build -o html5-standard-reader
```

## Usage

```bash
./html5-standard-reader [-debug] <path/to/HTML Standard.html>
```

### Example

```bash
# Download the HTML5 specification
curl -o html5-spec.html https://html.spec.whatwg.org/

# Extract element names
./html5-standard-reader html5-spec.html
```

## Output

The tool prints each discovered HTML element name on a separate line to standard output.

## License

MIT
