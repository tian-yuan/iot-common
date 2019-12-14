package util

type Options struct {
	ZkUrls    string
	TracerUrl string
}

type Option func(o *Options)

func WithZkUrls(zkUrls string) Option {
	return func(o *Options) {
		o.ZkUrls = zkUrls
	}
}

func WithTracerUrl(tracerUrl string) Option {
	return func(o *Options) {
		o.TracerUrl = tracerUrl
	}
}
