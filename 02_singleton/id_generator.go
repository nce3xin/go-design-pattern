package idgenerator

import "sync"

type IdGenerator struct {
}

var (
	instance *IdGenerator
	once     sync.Once
)

func GetInstance() *IdGenerator {
	once.Do(func() {
		instance = &IdGenerator{}
	})
	return instance
}
