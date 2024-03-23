package selfcare

import (
	"github.com/jffp113/selfcare/business/config"
	"github.com/jffp113/selfcare/business/controller/selfcare/operation"
)

type Selfcare struct {
	path    string
	config  config.Config
	baseUrl string
}

func New(opts ...Option) (*Selfcare, error) {
	var s Selfcare
	for _, o := range opts {
		o(&s)
	}

	cfg, err := config.ConfigFromFile(s.path)
	if err != nil {
		return nil, err
	}
	s.config = cfg

	return &s, nil
}

func (s *Selfcare) RegisterTimesheet(t config.Timesheet) error {
	op := operation.NewRegisterOp(s.config, t, s.baseUrl)
	return op.Do()
}
