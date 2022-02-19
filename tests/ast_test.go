package tests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

type V struct {

}

var emptySliceMap map[string]bool

func (visitor *V) Visit(node ast.Node) ast.Visitor {
	if node != nil {
		switch n := node.(type) {
		case *ast.AssignStmt:
			switch asRh := n.Rhs[0].(type) {
			case *ast.CallExpr:
				callF := asRh.Fun.(*ast.Ident)
				if callF.Name == "make" {
					switch makeT := asRh.Args[0].(type) {
					case *ast.ArrayType:
						if makeT.Len == nil {
							switch initLen := asRh.Args[1].(type) {
							case *ast.BasicLit:
								if initLen.Value == "0" {
									switch asVar := n.Lhs[0].(type) {
									case *ast.Ident:
										emptySliceMap[asVar.Name] = true
									}
								}
							}
						}
					}
				}
			}
		case *ast.RangeStmt:
			switch rangeVar := n.X.(type) {
			case *ast.Ident:
				for _, rangeBodyStmt := range n.Body.List {
					switch rbs := rangeBodyStmt.(type) {
					case *ast.AssignStmt:
						switch rbsRh := rbs.Rhs[0].(type) {
						case *ast.CallExpr:
							callF := rbsRh.Fun.(*ast.Ident)
							if callF.Name == "append" {
								switch apVar := rbsRh.Args[0].(type) {
								case *ast.Ident:
									if _, ok := emptySliceMap[apVar.Name]; ok {
										println(fmt.Sprintf("Missing init len of slice[%s], may use len(%s)", apVar.Name, rangeVar.Name))
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return visitor
}

func TestParseSlice(t *testing.T) {
	emptySliceMap = make(map[string]bool)

	fileSet := token.NewFileSet()
	expr, _ := parser.ParseFile(fileSet, "./../stubs/decl_slice.go", nil, parser.AllErrors)
	ast.Walk(&V{}, expr)

	for varName := range emptySliceMap {
		println(varName)
	}
}
