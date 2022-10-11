package gocc

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

/*
循環複雑度 | 状態
--------------------------------------------
1 - 10     | 安全なコードでテストしやすい
--------------------------------------------
11 - 20    | 少し複雑なコード
--------------------------------------------
21 - 40    | 複雑なコードでテストが難しくなる
--------------------------------------------
40 -       | テスト不可能
*/

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
