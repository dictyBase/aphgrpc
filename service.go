package aphgrpc

import (
	"context"

	"github.com/fatih/structs"
)

// AssignFieldsToStructs copy fields value
// between structure
func AssignFieldsToStructs(from interface{}, to interface{}) {
	toR := structs.New(to)
	for _, f := range structs.New(from).Fields() {
		if !f.IsZero() {
			if nf, ok := toR.FieldOk(f.Name()); ok {
				nf.Set(f.Value())
			}
		}
	}
}

type ServiceOptions struct {
	Topics   map[string]string
	Params   map[string]string
	Resource string
}

type Option func(*ServiceOptions)

func TopicsOption(t map[string]string) Option {
	return func(so *ServiceOptions) {
		so.Topics = t
	}
}

type Service struct {
	Resource string
	Context  context.Context
	Topics   map[string]string
	Params   map[string]string
}

func (s *Service) GetResourceName() string {
	return s.Resource
}
