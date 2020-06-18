package example

import (
	conf "github.com/yudhasubki/go-skeleton/config"
)

type Service struct {
	Repo   *Repository  `inject:"repository"`
	Config *conf.Config `inject:"config"`
}

func (s *Service) OnStart() {}

func (s *Service) OnShutdown() {}
