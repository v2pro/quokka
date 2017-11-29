package runtime

import (
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/thrift"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
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

func ThriftSchemas(thriftIDL string) map[string]*Schema {
	charStream := antlr.NewInputStream(thriftIDL)
	lexer := thrift.NewAntlrThriftLexer(charStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := thrift.NewAntlrThriftParser(tokenStream)
	schemas := map[string]*Schema{}
	for charStream.Index() < len(thriftIDL) {
		structDef := parser.StructDef()
		structName := structDef.GetChild(1).GetPayload().(*antlr.CommonToken).GetText()
		schema := &Schema{
			Fields: map[string]*Schema{},
		}
		for i := 3; i < structDef.GetChildCount(); i++ {
			field, _ := structDef.GetChild(i).(*thrift.FieldContext)
			if field == nil {
				break
			}
			fillField(schema, field)
		}
		schemas[structName] = schema
	}
	return schemas
}

var thriftTypes = map[string]*Schema{
	"string": mustBeString,
}

func fillField(schema *Schema, field *thrift.FieldContext) {
	_, isFieldType := field.GetChild(3).(*thrift.FieldTypeContext)
	idx := 3
	if !isFieldType {
		idx = 2
	}
	fieldType := field.GetChild(idx).GetChild(0).GetChild(0).GetPayload().(*antlr.CommonToken).GetText()
	fieldName := field.GetChild(idx + 1).GetPayload().(*antlr.CommonToken).GetText()
	schema.Fields[fieldName] = thriftTypes[fieldType]
}

func Validate(obj interface{}, schema *Schema) (err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err, _ = recovered.(error)
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", recovered))
			}
		}
	}()
	if schema == nil {
		return nil
	}
	dobj, _ := obj.(*DObject)
	if dobj == nil {
		return fmt.Errorf("%v is not DObject", obj)
	}
	dobj.Schema = schema
	dobj.validate()
	return nil
}