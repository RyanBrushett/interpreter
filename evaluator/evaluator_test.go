package evaluator

import (
	"testing"
	"github.com/ryanbrushett/interpreter/object"
	"github.com/ryanbrushett/interpreter/lexer"
	"github.com/ryanbrushett/interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct{
		input string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("Object is not an integer, it's a %T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("Object has wrong value. Expected %d, got %d", result.Value, expected)
		return false
	}
	return true
}
