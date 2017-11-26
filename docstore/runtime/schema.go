package runtime

import (
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/thrift"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Schema struct {
	Fields    map[string]*Schema
	Element   *Schema
	Validator func(value interface{}) error
}

var stringType = reflect.TypeOf((*string)(nil)).Elem()
var mustBeString = &Schema{
	Validator: func(value interface{}) error {
		if reflect.TypeOf(value) != stringType {
			return errors.New("must be string")
		}
		return nil
	},
}

func NewSchemaByThriftIDL(thriftIDL string) *Schema {
	charStream := antlr.NewInputStream(thriftIDL)
	lexer := thrift.NewAntlrThriftLexer(charStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := thrift.NewAntlrThriftParser(tokenStream)
	docStruct := parser.StructDef()
	schema := &Schema{
		Fields: map[string]*Schema{},
	}
	for i := 3; i < docStruct.GetChildCount(); i++ {
		field, _ := docStruct.GetChild(i).(*thrift.FieldContext)
		if field == nil {
			break
		}
		fillField(schema, field)
	}
	return schema
}

var thriftTypes = map[string]*Schema{
	"string": mustBeString,
}

func fillField(schema *Schema, field *thrift.FieldContext) {
	fieldType := field.GetChild(3).GetChild(0).GetChild(0).GetPayload().(*antlr.CommonToken).GetText()
	fieldName := field.GetChild(4).GetPayload().(*antlr.CommonToken).GetText()
	schema.Fields[fieldName] = thriftTypes[fieldType]
}
