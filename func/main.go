package main

import (
	"fmt"
	"net/http"
)

type Pwd struct {
	name string
}

type Interceptor func(r *http.Request) Pwd

type Transporter struct {
	interceptor Interceptor
}

var DefaultInterceptor Interceptor = func(r *http.Request) Pwd {
	return Pwd{}
}

func NewTransporter(interceptor ...Interceptor) Transporter {
	var fn = DefaultInterceptor
	if len(interceptor) > 0 && interceptor[0] != nil {
		fn = interceptor[0]
	}

	return Transporter{
		interceptor: fn,
	}
}

func (t *Transporter) Print() {
	fmt.Printf("%+v", t.interceptor(nil))
}

func main() {
	a := NewTransporter(func(r *http.Request) Pwd {
		return Pwd{"def"}
	})
	b := NewTransporter()

	fmt.Printf("%+v \n", a)
	fmt.Printf("%+v \n", b)

	a.Print()
	b.Print()

}
