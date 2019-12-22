package base

import (
	"infra/v1"

	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

type PropsStarter struct {
	infra.BaseStarter
}

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	//file := kvs.GetCurrentFilePath("config.ini", 1)
	props = ini.NewIniFileCompositeConfigSource("config.ini")
}

func (p *PropsStarter) StartBlocking() bool {
	return false
}
