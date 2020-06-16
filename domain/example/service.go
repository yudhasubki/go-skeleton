package example

import (
	"fmt"

	conf "github.com/yudhasubki/go-skeleton/config"
)

type Service struct {
	Repo   *Repository  `inject:"repository"`
	Config *conf.Config `inject:"config"`
}

func (s *Service) OnStart() {
	fmt.Println("invoke...")
}

func (s *Service) OnShutdown() {}
