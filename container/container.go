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
	objects  []*inject.Object
	services []interface{}
	status   chan bool
}

func (s *ServiceRegistry) Register(app string, svc interface{}) {
	go s.Ready(svc, s.status)
	s.objects = append(s.objects, &inject.Object{Value: svc, Name: app})
	if <-s.status {
		log.Printf("Invoke %v finished...", reflect.TypeOf(svc).Elem())
	}
}

func (s *ServiceRegistry) Bind() error {
	for _, svc := range s.objects {
		err := s.injector.Provide(svc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceRegistry) Ready(svc interface{}, status chan bool) {
	switch obj := svc.(type) {
	case ServiceInvoke:
		obj.OnStart()
		s.services = append(s.services, svc)
		status <- true
	}
	status <- false
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

func (s *ServiceRegistry) Shutdown() {
	if len(s.services) == 0 {
		return
	}

	log.Println("Starting shutdown...")
	for _, svc := range s.services {
		switch obj := svc.(type) {
		case ServiceInvoke:
			obj.OnShutdown()
		}
	}
}

func NewContainer() *ServiceRegistry {
	return &ServiceRegistry{status: make(chan bool)}
}
