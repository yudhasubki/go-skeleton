package container

import (
	"log"
	"reflect"

	"github.com/facebookgo/inject"
)

type ServiceInvoke interface {
	OnStart()
	OnShutdown()
}

type ServiceRegistry struct {
	injector inject.Graph
	services []*inject.Object
	status   chan bool
}

func (s *ServiceRegistry) Register(app string, svc interface{}) {
	go s.HasImplementServiceInvoke(svc, s.status)
	s.services = append(s.services, &inject.Object{Value: svc, Name: app})
	if <-s.status {
		log.Printf("Invoke %v finished...", reflect.TypeOf(svc).Elem())
	}
}

func (s *ServiceRegistry) HasImplementServiceInvoke(svc interface{}, status chan bool) {
	switch obj := svc.(type) {
	case ServiceInvoke:
		obj.OnStart()
		status <- true
	}
	status <- false
}

func (s *ServiceRegistry) Bind() error {
	for _, svc := range s.services {
		err := s.injector.Provide(svc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceRegistry) Start() error {
	err := s.Bind()
	if err != nil {
		return err
	}

	err = s.injector.Populate()
	if err != nil {
		return err
	}
	return nil
}

func NewContainer() *ServiceRegistry {
	return &ServiceRegistry{status: make(chan bool)}
}
