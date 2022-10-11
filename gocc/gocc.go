package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
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
--------------------------------------------
*/

type Report struct {
	Pos struct {
		Line int
		Col  int
	}
	FuncName string
	CC       int
}

var fset *token.FileSet
var reports []Report

func main() {

	var filePath = flag.String("path", "", "File path")
	flag.Parse()
	fmt.Println(*filePath)

	fset = token.NewFileSet()

	f, err := parser.ParseFile(fset, *filePath, nil, 0)
	if err != nil {
		log.Fatal("Failed to parse file")
	}

	reports = make([]Report, 0, fset.Position(f.Pos()).Line)

	walkAll(f)

	for _, v := range reports {
		if v.CC > 10 {
			fmt.Println(v)
		}
	}
}

func walkAll(root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			walkFunc(n.(*ast.FuncDecl))
		}
		return true
	})
}

func walkFunc(root *ast.FuncDecl) {
	position := fset.Position(root.Pos())
	report := Report{
		Pos: struct {
			Line int
			Col  int
		}{
			Line: position.Line,
			Col:  position.Column,
		},
		FuncName: root.Name.Name,
		CC:       1,
	}
	fmt.Println(report)

	ast.Inspect(root, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			return false
		case *ast.IfStmt:
			report.CC++
			if n.(*ast.IfStmt).Else != nil {
				report.CC++
			}
		case *ast.SwitchStmt:
			report.CC++
		case *ast.TypeSwitchStmt:
			report.CC++
		case *ast.CaseClause:
			report.CC++
		case *ast.SelectStmt:
			report.CC++
		case *ast.CommClause:
			report.CC++
		case *ast.ForStmt:
			report.CC++
		case *ast.RangeStmt:
			report.CC++
		}
		return true
	})
	reports = append(reports, report)
}
