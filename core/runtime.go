package core

import "runtime"

type functionCall struct {
	filePath     string
	line         int
	inputValues  []interface{}
	functionName string
}

func runtimeFunctionCall(values ...interface{}) (call *functionCall) {

	stackLayer := 2 // user use database method stack layer
	pcptr, filePath, line, ok := runtime.Caller(stackLayer)
	_func := runtim.FuncForPC(pcptr)
	return &functionCall{
		filePath:     filePath,
		line:         line,
		inputValues:  values,
		functionName: _func.Name(),
	}
}
