package pedis

type InterfaceResult struct {
	Result interface{}
	Err error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func(this *InterfaceResult) Unwarp() interface{}{
	if this.Err !=nil{
		panic(this.Err)
	}
	return this.Result
}

func(this *InterfaceResult) Unwarp_Or(str interface{}) interface{}{
	if this.Err !=nil{
		return str
	}
	return this.Result
}