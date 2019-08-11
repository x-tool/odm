package core

import "runtime"

type functionCall struct {
	filePath     string
	line         int
	inputValues  []interface{}
	functionName string
}

func runtimeFunctionCall(functionName string, values ...interface{}) (call *functionCall) {

	stackLayer := 2 // user use database method stack layer
	_, filePath, line, ok := runtime.Caller(stackLayer)
	return &functionCall{
		filePath:     filePath,
		line:         line,
		inputValues:  values,
		functionName: functionName,
	}
}
