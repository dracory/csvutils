# CSV Utils

[![Tests Status](https://github.com/dracory/csvutils/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/dracory/csvutils/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/csvutils)](https://goreportcard.com/report/github.com/dracory/csvutils)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dracory/csvutils)](https://pkg.go.dev/github.com/dracory/csvutils)

Simple Go utilities for reading CSV files into arrays or maps.

## Installation

```bash
go get github.com/dracory/csvutils
```

## Usage

### ToArrays

Reads a CSV file and returns an array with each line as an array of strings.

```go
package main

import (
    "fmt"
    "log"
    "github.com/dracory/csvutils"
)

func main() {
    records, err := csvutils.ToArrays("data.csv")
    if err != nil {
        log.Fatal(err)
    }

    // records[0] contains headers
    // records[1:] contains data rows
    for _, row := range records {
        fmt.Println(row)
    }
}
```

### ToMaps

Reads a CSV file and returns an array of maps using the header row as keys.

```go
package main

import (
    "fmt"
    "log"
    "github.com/dracory/csvutils"
)

func main() {
    // Optional: replace header names
    replacements := map[string]string{
        "First Name": "first_name",
        "Last Name":  "last_name",
    }

    records, err := csvutils.ToMaps("data.csv", replacements)
    if err != nil {
        log.Fatal(err)
    }

    for _, row := range records {
        fmt.Printf("Name: %s %s\n", row["first_name"], row["last_name"])
    }
}
```

## Features

- `ToArrays` - Returns CSV as `[][]string` for raw access
- `ToMaps` - Returns CSV as `[]map[string]string` with header-based keys
- Header trimming and name replacement support
- Standard Go `encoding/csv` compliance (handles quoted fields)

## License

MIT
