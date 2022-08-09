package main

import (
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func Test_reflect(t *testing.T) {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())

		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		t.Logf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}

func Test_timeAfter(t *testing.T) {
	t.Log(time.Now())
	tchan := time.After(time.Second * 3)
	t.Log(time.Now())
	t.Log("get", <-tchan)
	t.Log(time.Now())
}
