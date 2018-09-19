package main

import (
	"plugin"
	"fmt"
	"github.com/juhwany/go_server/apis"
)

type BarImpl struct {
}

func (b *BarImpl) DoSomething {
	fmt.Printf("DoSomething called!\n")
}

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	foo, err := p.Lookup("Foo")
	if err != nil {
		panic(err)
	}

	barImpl := &BarImpl{}
	foo.(func(apis.Bar))(barImpl)
}
