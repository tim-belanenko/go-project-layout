package usecase

import (
	"context"
	"layout/internal/pkg/ctxlogger"
	"sync"
)

type ModuleStatus string

const (
	StatusActive   ModuleStatus = "active"
	StatusInactive ModuleStatus = "inactive"
)

type Module struct {
	Module string       `json:"module"`
	Status ModuleStatus `json:"status"`
}

type HealthStatus struct {
	Modules []*Module `json:"modules"`
}

type IHealthcheck interface {
	Healthcheck(ctx context.Context) (*Module, error)
}

type healthcheck struct {
	modules []IHealthcheck
}

func New(modules ...IHealthcheck) *healthcheck {
	return &healthcheck{
		modules: modules,
	}
}

func (h *healthcheck) Status(ctx context.Context) (*HealthStatus, error) {
	log, err := ctxlogger.Logger(ctx)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	statusesC := make(chan *Module, len(h.modules))
	for _, module := range h.modules {
		wg.Add(1)
		go func(module IHealthcheck) {
			defer wg.Done()
			status, err := module.Healthcheck(ctx)
			if err != nil {
				log.Error("failed to check module status", err)
			}
			statusesC <- status
		}(module)
	}
	go func() {
		wg.Wait()
		close(statusesC)
	}()
	modules := make([]*Module, 0, len(h.modules))
	for module := range statusesC {
		modules = append(modules, module)
	}

	return &HealthStatus{
		Modules: modules,
	}, nil

}
