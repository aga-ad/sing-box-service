package service

import (
	"context"

	"golang.org/x/sys/windows/svc"
)

// Run start service with specific name in Service Control mode
func Run(name string, run func(ctx context.Context)) error {
	executer := executer{run: run}
	err := svc.Run(name, &executer)
	if err != nil {
		return err
	}
	return nil
}

type executer struct {
	run func(ctx context.Context)
}

func (e *executer) Execute(args []string, r <-chan svc.ChangeRequest, status chan<- svc.Status) (bool, uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	status <- svc.Status{State: svc.StartPending}
	stopedC := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		e.run(ctx)
		close(stopedC)
	}()
	status <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	running := true
	for running {
		select {
		case <-stopedC:
			running = false
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				status <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				running = false
			}
		}
	}
	cancel()
	status <- svc.Status{State: svc.StopPending}
	<-stopedC
	status <- svc.Status{State: svc.Stopped}
	return false, 0
}

func IsWindowsService() (bool, error) {
	return svc.IsWindowsService()
}
