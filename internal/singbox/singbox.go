package singbox

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	box "github.com/sagernet/sing-box"

	"github.com/sagernet/sing-box/experimental/deprecated"
	"github.com/sagernet/sing-box/include"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
	"github.com/sagernet/sing/common/json"
	"github.com/sagernet/sing/service"
)

func Run(ctx context.Context) error {
	ctx = include.Context(service.ContextWith(ctx, deprecated.NewStderrManager(log.StdLogger())))
	instance, err := create(ctx)
	if err != nil {
		return fmt.Errorf("can't create box: %w", err)
	}

	<-ctx.Done()
	err = instance.Close()
	if err != nil {
		return fmt.Errorf("can't close box: %w", err)
	}
	return nil
}

func create(ctx context.Context) (*box.Box, error) {
	options, err := readConfig(ctx)
	if err != nil {
		return nil, err
	}
	instance, err := box.New(box.Options{
		Context: ctx,
		Options: options,
	})
	if err != nil {
		return nil, E.Cause(err, "create service")
	}

	err = instance.Start()
	if err != nil {
		return nil, E.Cause(err, "start service")
	}
	return instance, nil
}

// readConfig reads config.json in dir with executable
func readConfig(ctx context.Context) (option.Options, error) {
	ex, err := os.Executable()
	if err != nil {
		return option.Options{}, err
	}
	dir := filepath.Dir(ex)
	path := filepath.Join(dir, "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return option.Options{}, E.Cause(err, "read config at ", path)
	}
	options, err := json.UnmarshalExtendedContext[option.Options](ctx, data)
	if err != nil {
		return option.Options{}, E.Cause(err, "decode config at ", path)
	}

	return options, nil
}
