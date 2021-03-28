package pedis

type StringResult struct {
	Result string
	Err error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func(this *StringResult) Unwarp() string{
	if this.Err !=nil{
		panic(this.Err)
	}
	return this.Result
}

func(this *StringResult) Unwarp_Or(str string) string{
	if this.Err !=nil{
		return str
	}
	return this.Result
}