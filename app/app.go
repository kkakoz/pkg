package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"time"
)

type App struct {
	servers     []Server
	stopTimeout time.Duration
	name        string
	ctx         context.Context
	cancel      func()
}

func NewApp(name string, servers ...Server) *App {
	return &App{servers: servers, name: name}
}

func (a *App) Start(ctx context.Context) error {

	//ctx, a.cancel = context.WithCancel(ctx)
	eg, ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}
	for _, server := range a.servers {
		cur := server
		eg.Go(func() error {
			<-ctx.Done()
			sctx, cancel := context.WithTimeout(ctx, a.stopTimeout)
			defer cancel()
			return cur.Stop(sctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return cur.Start(ctx)
		})
	}

	wg.Wait()

	c := make(chan os.Signal)
	signal.Notify(c, signals...)

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case <-c:
			fmt.Println("收到退出信号")
			return a.Stop()
		}
	})
	return eg.Wait()
}

func (a *App) Stop() error {
	a.cancel()
	return nil
}

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
