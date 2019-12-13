package basic

import (
	"github.com/tian-yuan/iot-common/basic/config"
)

var (
	pluginFuncs []func()
)

type Options struct {
	cfgOps []config.Option
}

type Option func(o *Options)

func Init(opts ...config.Option) {
	config.Init(opts...)

	for _, f := range pluginFuncs {
		f()
	}
}

func Register(f func()) {
	pluginFuncs = append(pluginFuncs, f)
}
