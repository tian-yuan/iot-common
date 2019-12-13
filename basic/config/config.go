package config

import (
	"sync"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
)

var (
	m      sync.RWMutex
	inited bool

	c = &configurator{}
)

type Configurator interface {
	GetConfig() config.Config
}

type configurator struct {
	conf config.Config
}

func C() Configurator {
	return c
}

func (c *configurator) init(ops Options) (err error) {
	m.Lock()
	defer m.Unlock()

	if inited {
		return
	}

	c.conf = config.NewConfig()

	err = c.conf.Load(ops.Sources...)
	if err != nil {
		log.Fatal(err)
	}

	inited = true
	return
}

func (c *configurator) GetConfig() config.Config {
	return c.conf
}

func Init(opts ...Option) {
	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	c = &configurator{}
	c.init(ops)
}
