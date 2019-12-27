package util

type Options struct {
	RegistryUrls string
	TracerUrl    string
}

type Option func(o *Options)

func WithRegistryUrls(registryUrls string) Option {
	return func(o *Options) {
		o.RegistryUrls = registryUrls
	}
}

func WithTracerUrl(tracerUrl string) Option {
	return func(o *Options) {
		o.TracerUrl = tracerUrl
	}
}
