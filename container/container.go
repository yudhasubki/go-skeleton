package container

import (
	"github.com/facebookgo/inject"
)

type ServiceInvoke interface {
	OnStart()
	OnShutdown()
}

type ServiceRegistry struct {
	injector inject.Graph
	services []*inject.Object
}

func (s *ServiceRegistry) Register(app string, svc interface{}) {
	switch obj := svc.(type) {
	case ServiceInvoke:
		obj.OnStart()
	}
	s.services = append(s.services, &inject.Object{Value: svc, Name: app})
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
	return &ServiceRegistry{}
}
