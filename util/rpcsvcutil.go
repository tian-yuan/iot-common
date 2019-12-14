package util

import (
	"io"
	"strings"
	"sync"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/zookeeper"
	"github.com/tian-yuan/iot-common/plugins/tracer"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
)

const (
	REGISTER_MANAGER_SVC   = "iot.register.manager"
	TOPIC_MANAGER_SVC      = "iot.topic.manager"
	MESSAGE_DISPATCHER_SVC = "iot.message.dispatcher"
	TOPIC_ACL_SVC          = "iot.topic.acl"
	PUBLISH_ENGINE_SVC     = "iot.publish.engine"
	CONTROLLER_SVC         = "iot.controller"
)

var once sync.Once

type RpcCtx struct {
	// The following fields is rpc service variable
	// {{
	RegisterSvc          micro.Service
	TopicManagerSvc      micro.Service
	MessageDispatcherSvc micro.Service
	PubEngineSvc         micro.Service
	TopicAclSvc          micro.Service
	ControllerSvc        micro.Service
	// }}
	Options
	Closer map[string]io.Closer
}

var Ctx RpcCtx

func Init(opts ...Option) {
	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}
	Ctx.Options = ops
	Ctx.Closer = make(map[string]io.Closer)
}

func (ctx *RpcCtx) InitRegisterSvc() {
	clientSvcName := REGISTER_MANAGER_SVC + ".client"
	Ctx.RegisterSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) CloseRegisterSvc() {
	clientSvcName := REGISTER_MANAGER_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) InitPubEngineSvc() {
	clientSvcName := PUBLISH_ENGINE_SVC + ".client"
	Ctx.PubEngineSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) ClosePubEngineSvc() {
	clientSvcName := PUBLISH_ENGINE_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) InitTopicManagerSvc() {
	clientSvcName := TOPIC_MANAGER_SVC + ".client"
	Ctx.TopicManagerSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) CloseTopicManagerSvc() {
	clientSvcName := TOPIC_MANAGER_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) InitTopicAclSvc() {
	clientSvcName := TOPIC_ACL_SVC + ".client"
	Ctx.TopicAclSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) CloseTopicAclSvc() {
	clientSvcName := TOPIC_ACL_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) InitMessageDispatcherSvc() {
	clientSvcName := MESSAGE_DISPATCHER_SVC + ".client"
	Ctx.MessageDispatcherSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) CloseMessageDispatcherSvc() {
	clientSvcName := MESSAGE_DISPATCHER_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) InitControllerSvc() {
	clientSvcName := CONTROLLER_SVC + ".client"
	Ctx.ControllerSvc = ctx.initClientSvc(clientSvcName)
}

func (ctx *RpcCtx) CloseControllerSvc() {
	clientSvcName := CONTROLLER_SVC + ".client"
	ctx.closeClientSvc(clientSvcName)
}

func (ctx *RpcCtx) initClientSvc(clientSvcName string) (svc micro.Service) {
	once.Do(func() {
		optFunc := func(opt *registry.Options) {
			opt = &registry.Options{
				Addrs: strings.Split(Ctx.Options.ZkUrls, ";"),
			}
		}
		registry := zookeeper.NewRegistry(optFunc)

		t, io, err := tracer.NewTracer(clientSvcName, Ctx.Options.TracerUrl)
		if err != nil {
			log.Errorf("init tracer failed, err : %s", err.Error())
			svc = micro.NewService(
				micro.Name(clientSvcName),
				micro.Registry(registry),
			)

		} else {
			Ctx.Closer[clientSvcName] = io

			svc = micro.NewService(
				micro.Name(clientSvcName),
				micro.Registry(registry),
				micro.WrapClient(ocplugin.NewClientWrapper(t)),
			)

		}
		svc.Init()
	})
	return
}

func (ctx *RpcCtx) closeClientSvc(clientSvcName string) {
	once.Do(func() {
		if Ctx.Closer[clientSvcName] != nil {
			Ctx.Closer[clientSvcName].Close()
		}
	})
}
