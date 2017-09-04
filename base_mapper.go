package mappers

import (
	"fmt"
)

type IRepository interface {
	Close() error
}

type BaseMapper struct {
	Repository IRepository
}

func (bm *BaseMapper) Log(data ...interface{}) {
	fmt.Println(data...)
}
func (bm *BaseMapper) Error(data ...interface{}) {
	fmt.Println(data...)
}

func (bm *BaseMapper) Close() error {
	return bm.Repository.Close()
}
