package gocc

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: [file path]")
	}

	filePath := os.Args[1]

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		log.Fatal("Failed to parse file")
	}

}
