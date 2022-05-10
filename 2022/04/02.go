package main

import "github.com/Knetic/govaluate"

func main() {
	//demo1()
	//demo2()
	//demo3()
	demo4()
}

func demo1() {
	expression, _ := govaluate.NewEvaluableExpression("10 > 0")
	result, _ := expression.Eval(nil)
	println(result.(bool))
}

func demo2() {
	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = 1
	res, _ := expr.Evaluate(parameters)
	println(res.(bool))
}

func demo3() {
	expr, _ := govaluate.NewEvaluableExpression("a * b / 10")
	parameters := make(map[string]interface{}, 8)
	parameters["a"] = 1
	parameters["b"] = 100
	res, _ := expr.Evaluate(parameters)
	println(res.(float64))
}

func demo4() {
	functions := make(map[string]govaluate.ExpressionFunction, 4)
	functions["strlen"] = func(args ...interface{}) (interface{}, error) {
		length := len(args[0].(string))
		return (int)(length), nil
	}
	functions["sqrt"] = func(args ...interface{}) (interface{}, error) {
		x := args[0].(float64)
		return x * x, nil
	}

	exprStr1 := "3.14 * sqrt(2) < sqrt(4)"
	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(exprStr1, functions)
	res, _ := expr.Eval(nil)
	println(res.(bool))
}
