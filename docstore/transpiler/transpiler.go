package transpiler

import (
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/ast"
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/compiler"
	"strconv"
	"fmt"
	"bytes"
	"github.com/robertkrimen/otto/file"
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
	funcLiteral, err := parser.ParseFunction("", input)
	if err != nil {
		return "", err
	}
	output := []byte(`
package main
import "github.com/v2pro/quokka/docstore/runtime"
import "fmt"

func Fn(doc interface{}, req interface{}) (ret interface{}) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			jsErr, _ := recovered.(*runtime.JavascriptError)
			if jsErr != nil {
				ret = recovered
				if _, isThrown := jsErr.Recovered.(runtime.Thrown); isThrown {
					ret = jsErr.Recovered
				}
			} else {
				ret = fmt.Errorf("unknown error: %v", recovered)
			}
		}
	}()
	return handle(doc, req)
}
`)
	translators := map[string]*translator{}
	for _, funcStmtObj := range funcLiteral.Body.(*ast.BlockStatement).List {
		funcStmt := funcStmtObj.(*ast.FunctionStatement)
		funcName := funcStmt.Function.Name.Name
		tl := &translator{
			input: input[funcStmt.Function.Body.Idx0()-14: funcStmt.Function.Body.Idx1()-16],
		}
		paramNames := []string{}
		for _, param := range funcStmt.Function.ParameterList.List {
			paramNames = append(paramNames, param.Name)
		}
		tl.translate(funcName, paramNames, bytes.Count(output, []byte{'\n'}))
		if tl.err != nil {
			return "", tl.err
		}
		output = append(output, tl.output...)
		translators[funcName] = tl
	}
	for funcName, tl := range translators {
		output = append(output, `var `...)
		output = append(output, funcName...)
		output = append(output, `_JS_SOURCE=`...)
		output = append(output, '`')
		output = append(output, tl.input...)
		output = append(output, '`', '\n')
		output = append(output, `var `...)
		output = append(output, funcName...)
		output = append(output, `_GO_SOURCE=`...)
		output = append(output, '`')
		output = append(output, tl.goSrc...)
		output = append(output, '`', '\n')
	}
	return string(output), nil
}

type translator struct {
	input          string
	output         []byte
	jsOffsets      []int
	goOffsets      []int
	goSrc          []byte
	err            error
	hasReturnValue bool
	bodyStartIdx   file.Idx
}

func (tl *translator) reportError(node ast.Node, errmsg string) {
	tl.err = errors.New(errmsg)
}

func (tl *translator) translate(funcName string, paramNames []string, absoluteLineNo int) {
	funcLiteral, err := parser.ParseFunction("", tl.input)
	if err != nil {
		tl.err = err
		return
	}
	tl.bodyStartIdx = funcLiteral.Body.Idx0() + 2
	tl.translateStatement(funcLiteral.Body)
	tl.goSrc = tl.output
	tl.output = []byte(`func `)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, '(')
	for i, paramName := range paramNames {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, paramName...)
		tl.output = append(tl.output, " interface{}"...)
	}
	tl.output = append(tl.output, ") "...)
	if tl.hasReturnValue {
		tl.output = append(tl.output, "interface{} "...)
	}
	tl.output = append(tl.output, `{
		defer func() {
			recovered := recover()
			if recovered != nil {
				runtime.ReportError("`...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, `", `...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, "_JS_SOURCE, []int{"...)
	for i, offset := range tl.jsOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, "}, "...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, "_GO_SOURCE, []int{"...)
	for i, offset := range tl.goOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, `}, `...)
	tl.output = append(tl.output, strconv.Itoa(absoluteLineNo)...)
	tl.output = append(tl.output, `, recovered)
		}
	}()
`...)
	tl.output = append(tl.output, tl.goSrc...)
	tl.output = append(tl.output, "}\n"...)
}

func (tl *translator) translateStatement(stmt ast.Statement) {
	switch typedStmt := stmt.(type) {
	case *ast.BlockStatement:
		tl.translateBlockStatement(typedStmt)
	case *ast.ExpressionStatement:
		tl.translateExpression(typedStmt.Expression)
	case *ast.ReturnStatement:
		tl.hasReturnValue = true
		tl.translateReturnStatement(typedStmt)
	case *ast.ThrowStatement:
		tl.hasReturnValue = true
		tl.translateThrowStatement(typedStmt)
	default:
		tl.reportError(stmt, "can not handle "+reflect.TypeOf(stmt).String())
	}
}
func (tl *translator) translateReturnStatement(stmt *ast.ReturnStatement) {
	tl.output = append(tl.output, "return "...)
	tl.translateExpression(stmt.Argument)
}

func (tl *translator) translateThrowStatement(stmt *ast.ThrowStatement) {
	tl.output = append(tl.output, "panic(runtime.Thrown{"...)
	tl.translateExpression(stmt.Argument)
	tl.output = append(tl.output, "})"...)
}

func (tl *translator) translateBlockStatement(stmt *ast.BlockStatement) {
	for _, child := range stmt.List {
		tl.jsOffsets = append(tl.jsOffsets, int(child.Idx0()-tl.bodyStartIdx))
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
	case *ast.DotExpression:
		tl.translateDotExpression(typedExpr)
	case *ast.ObjectLiteral:
		tl.translateObjectLiteral(typedExpr)
	case *ast.NumberLiteral:
		tl.output = append(tl.output, fmt.Sprintf("%v", typedExpr.Value)...)
	case *ast.NullLiteral:
		tl.output = append(tl.output, "nil"...)
	case *ast.CallExpression:
		tl.translateCallExpression(typedExpr)
	case *ast.BinaryExpression:
		tl.translateBinaryExpression(typedExpr)
	default:
		tl.reportError(typedExpr, "can not handle "+reflect.TypeOf(typedExpr).String())
	}
}

func (tl *translator) translateDotExpression(expr *ast.DotExpression) {
	tl.output = append(tl.output, `runtime.AsObj(`...)
	tl.translateExpression(expr.Left)
	tl.output = append(tl.output, `).Get("`...)
	tl.output = append(tl.output, expr.Identifier.Name...)
	tl.output = append(tl.output, `")`...)
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
	case *ast.DotExpression:
		tl.output = append(tl.output, `runtime.AsObj(`...)
		tl.translateExpression(leftExpr.Left)
		tl.output = append(tl.output, `).Set("`...)
		tl.output = append(tl.output, leftExpr.Identifier.Name...)
		tl.output = append(tl.output, `", `...)
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

func (tl *translator) translateCallExpression(expr *ast.CallExpression) {
	tl.translateExpression(expr.Callee)
	tl.output = append(tl.output, '(')
	for i, arg := range expr.ArgumentList {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.translateExpression(arg)
	}
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateBinaryExpression(expr *ast.BinaryExpression) {
	operator := expr.Operator.String()
	switch operator {
	case "+":
		tl.output = append(tl.output, "runtime.Add("...)
		tl.translateExpression(expr.Left)
		tl.output = append(tl.output, ", "...)
		tl.translateExpression(expr.Right)
		tl.output = append(tl.output, ')')
	case "-":
		tl.output = append(tl.output, "runtime.Subtract("...)
		tl.translateExpression(expr.Left)
		tl.output = append(tl.output, ", "...)
		tl.translateExpression(expr.Right)
		tl.output = append(tl.output, ')')
	default:
		tl.reportError(expr, "do not support operator "+operator)
	}
}
