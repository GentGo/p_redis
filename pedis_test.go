package pedis

import (
	"fmt"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	//it :=pedis.NewStringOperation().MGet("name","age").Iter()
	//
	//for it.HasNext(){
	//	fmt.Println(it.Next())
	//}

	//fmt.Println(pedis.NewStringOperation().Set("name","today",pedis.WithExpire(time.Second*200)))

	fmt.Println(NewStringOperation().Set("name","today",
		WithExpire(time.Second*200),
		//WithXx(),
	))
}