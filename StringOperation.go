package pedis

import (
	"context"
	"time"
)

type StringOperation struct{
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}



func (this *StringOperation)Set(key string,value interface{},attrs ...*OperationAttr) *InterfaceResult{
	exp :=OperationAttrs(attrs).Find(ATTR_EXPR)

	nx :=OperationAttrs(attrs).Find(ATTR_NX).Unwarp_Or(nil)

	if nx!=nil{
		return NewInterfaceResult(Redis().SetNX(this.ctx,key,value,exp.Unwarp_Or(time.Second*0).(time.Duration)).Result())
	}

	xx :=OperationAttrs(attrs).Find(ATTR_XX).Unwarp_Or(nil)

	if xx!=nil{
		return NewInterfaceResult(Redis().SetXX(this.ctx,key,value,exp.Unwarp_Or(time.Second*0).(time.Duration)).Result())
	}

	return NewInterfaceResult(Redis().Set(this.ctx,key,value,
		exp.Unwarp_Or(time.Second*0).(time.Duration)).Result())//exp.(time.Duration)断言操作
}

func(this *StringOperation)Get(key string) *StringResult{
	return NewStringResult(Redis().Get(this.ctx,key).Result())
}

func(this *StringOperation)MGet(key... string) *SliceResult{
	return NewSliceResult(Redis().MGet(this.ctx,key...).Result())
}