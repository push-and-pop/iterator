package iterator

import (
	"fmt"
	"reflect"
)

type Iterator interface {
	HasNext() bool
	Next()
	Pre()
	CurrentItem() interface{}
}

func Toiter(container interface{}) Iterator {
	va := mustBeSlice(container)
	fmt.Println(va, va.Len())
	return &iterator{
		len:   va.Len(),
		index: 0,
		value: va,
	}
}

func mustBeSlice(i interface{}) reflect.Value {
	va := reflect.ValueOf(i)
	if va.Kind() != reflect.Slice {
		panic("must be slice!")
	}
	return va
}

type iterator struct {
	len   int
	index int
	value reflect.Value
}

func (i *iterator) HasNext() bool {
	return i.index < i.len
}

func (i *iterator) Next() {
	i.index++
}

func (i *iterator) Pre() {
	i.index--
}

func (i *iterator) CurrentItem() interface{} {
	return i.value.Index(i.index)
}

func (i *iterator) Filter() *iterator {
	return i
}

func (i *iterator) Collect() interface{} {
	return i
}
