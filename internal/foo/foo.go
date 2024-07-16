package foo

import "errors"

type Foo map[string]string

type FooCache map[Id]Foo

type Id int

func InitFooCache() FooCache {
	return map[Id]Foo{
		Id(10): map[string]string{
			"entry1": "this is entry1",
		},
	}
}

func (j *FooCache) GetFoo(id Id) (Foo, error) {
	if foo, ok := (*j)[id]; ok {
		return foo, nil
	}
	return nil, ErrIdNotFound
}

var ErrIdNotFound = errors.New("ID was not in the cache")
