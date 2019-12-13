package config

import "github.com/micro/go-micro/config/source"

type Options struct {
	Sources []source.Source
}

type Option func(o *Options)

func WithSource(src source.Source) Option {
	return func(o *Options) {
		o.Sources = append(o.Sources, src)
	}
}
