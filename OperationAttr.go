package pedis

import (
	"fmt"
	"time"
)
type empty struct{}
const(
	ATTR_EXPR = "expr"
	ATTR_NX = "nx" //setnx
	ATTR_XX = "xx" //setnx
)

//属性
type OperationAttr struct {
	Name string
	Vlaue interface{}
}

type OperationAttrs []*OperationAttr

func (this OperationAttrs) Find(name string) *InterfaceResult {
	for _,attr :=range this{
		if attr.Name==name{
			return NewInterfaceResult(attr.Vlaue,nil)
		}
	}
	return NewInterfaceResult(nil,fmt.Errorf("OperationAttrs found error:%s",name))
}

func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{Name:ATTR_EXPR,Vlaue:t}
}


func WithNx() *OperationAttr {
	return &OperationAttr{Name:ATTR_NX,Vlaue:empty{}}
}

func WithXx() *OperationAttr {
	return &OperationAttr{Name:ATTR_XX,Vlaue:empty{}}
}