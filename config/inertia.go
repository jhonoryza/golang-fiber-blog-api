package config

import "github.com/theartefak/inertia-fiber"

var i *inertia.Engine

func InitInertia() {
	i = inertia.New()
}

func GetInertia() *inertia.Engine {
	if i == nil {
		InitInertia()
	}
	return i
}
