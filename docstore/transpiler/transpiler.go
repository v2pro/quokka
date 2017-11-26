package transpiler

import (
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/ast"
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/compiler"
	"strconv"
	"fmt"
)

func Compile(input string) (func(doc interface{}, req interface{}) interface{}, error) {
	output, err := translate(input)
	if err != nil {
		return nil, err
	}
	plugin, err := compiler.DynamicCompile(output)
	if err != nil {
		return nil, err
	}
	fn, err := plugin.Lookup("Fn")
	if err != nil {
		return nil, err
	}
	return fn.(func(doc interface{}, req interface{}) interface{}), nil
}

func translate(input string) (string, error) {
	tl := &translator{
		input: input,
	}
	tl.translate()
	if tl.err != nil {
		return "", tl.err
	}
	return string(tl.output), nil
}

type config struct {
}

type translator struct {
	input     string
	output    []byte
	jsOffsets []int
	goOffsets []int
	cfg       config
	err       error
}

func (tl *translator) reportError(node ast.Node, errmsg string) {
	tl.err = errors.New(errmsg)
}

func (tl *translator) translate() {
	funcLiteral, err := parser.ParseFunction("", tl.input)
	if err != nil {
		tl.err = err
		return
	}
	tl.translateStatement(funcLiteral.Body)
	goSrc := tl.output
	tl.output = []byte(`
package main
import "github.com/v2pro/quokka/docstore/runtime"
func Fn(doc interface{}, req interface{}) interface{} {
	defer func() {
		recovered := recover()
		if recovered != nil {
			runtime.ReportError(`)
	tl.output = append(tl.output, "JS_SOURCE, []int{"...)
	for i, offset := range tl.jsOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, "}, GO_SOURCE, []int{"...)
	for i, offset := range tl.goOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, `}, recovered)
		}
	}()
`...)

	tl.output = append(tl.output, goSrc...)
	tl.output = append(tl.output, "}\n"...)
	tl.output = append(tl.output, `var JS_SOURCE=`...)
	tl.output = append(tl.output, '`')
	tl.output = append(tl.output, tl.input...)
	tl.output = append(tl.output, '`', '\n')
	tl.output = append(tl.output, `var GO_SOURCE=`...)
	tl.output = append(tl.output, '`')
	tl.output = append(tl.output, goSrc...)
	tl.output = append(tl.output, '`', '\n')
}

func (tl *translator) translateStatement(stmt ast.Statement) {
	switch typedStmt := stmt.(type) {
	case *ast.BlockStatement:
		tl.translateBlockStatement(typedStmt)
	case *ast.ExpressionStatement:
		tl.translateExpression(typedStmt.Expression)
	case *ast.ReturnStatement:
		tl.translateReturnStatement(typedStmt)
	default:
		tl.reportError(stmt, "can not handle "+reflect.TypeOf(stmt).String())
	}
}
func (tl *translator) translateReturnStatement(stmt *ast.ReturnStatement) {
	tl.output = append(tl.output, "return "...)
	tl.translateExpression(stmt.Argument)
}

func (tl *translator) translateBlockStatement(stmt *ast.BlockStatement) {
	for _, child := range stmt.List {
		tl.jsOffsets = append(tl.jsOffsets, int(child.Idx0()))
		tl.goOffsets = append(tl.goOffsets, len(tl.output))
		tl.translateStatement(child)
		tl.output = append(tl.output, '\n')
	}
}

func (tl *translator) translateExpression(expr ast.Expression) {
	switch typedExpr := expr.(type) {
	case *ast.AssignExpression:
		tl.translateAssignExpression(typedExpr)
	case *ast.StringLiteral:
		tl.translateStringLiteral(typedExpr)
	case *ast.Identifier:
		tl.output = append(tl.output, typedExpr.Name...)
	case *ast.BracketExpression:
		tl.translateBracketExpression(typedExpr)
	case *ast.ObjectLiteral:
		tl.translateObjectLiteral(typedExpr)
	case *ast.NumberLiteral:
		tl.output = append(tl.output, fmt.Sprintf("%v", typedExpr.Value)...)
	case *ast.NullLiteral:
		tl.output = append(tl.output, "nil"...)
	default:
		tl.reportError(typedExpr, "can not handle "+reflect.TypeOf(typedExpr).String())
	}
}

func (tl *translator) translateBracketExpression(expr *ast.BracketExpression) {
	tl.output = append(tl.output, `runtime.AsObj(`...)
	tl.translateExpression(expr.Left)
	tl.output = append(tl.output, ").Get("...)
	tl.translateExpression(expr.Member)
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateAssignExpression(expr *ast.AssignExpression) {
	switch leftExpr := expr.Left.(type) {
	case *ast.BracketExpression:
		tl.output = append(tl.output, `runtime.AsObj(`...)
		tl.translateExpression(leftExpr.Left)
		tl.output = append(tl.output, ").Set("...)
		tl.translateExpression(leftExpr.Member)
		tl.output = append(tl.output, ", "...)
		tl.translateExpression(expr.Right)
		tl.output = append(tl.output, ')')
	default:
		tl.reportError(expr.Left, "can not handle "+reflect.TypeOf(expr.Left).String())
	}
}

func (tl *translator) translateStringLiteral(expr *ast.StringLiteral) {
	tl.output = append(tl.output, '"')
	tl.output = append(tl.output, expr.Value...)
	tl.output = append(tl.output, '"')
}

func (tl *translator) translateObjectLiteral(expr *ast.ObjectLiteral) {
	tl.output = append(tl.output, `runtime.NewObject(`...)
	for i, prop := range expr.Value {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, '"')
		tl.output = append(tl.output, prop.Key...)
		tl.output = append(tl.output, `", `...)
		tl.translateExpression(prop.Value)
	}
	tl.output = append(tl.output, ')')
}
