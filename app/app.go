package app

import (
	"context"
	"fmt"
	"github.com/kkakoz/pkg/logger"
	"go.uber.org/zap"
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

	a.stopTimeout = time.Second * 5

	ctx, a.cancel = context.WithCancel(ctx)
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
			err := cur.Start(ctx)
			if err != nil {
				logger.Error("启动失败", zap.Error(err))
			}
			return err
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
