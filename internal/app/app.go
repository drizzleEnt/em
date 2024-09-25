package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/drizzleent/em/internal/config"
	"github.com/drizzleent/em/internal/logger"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevel = flag.String("l", "info", "log level")

type App struct {
	serviceProvider *serviceProvider

	httpServer *http.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDebps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) Run() error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()
	wg.Wait()
	return nil
}

func (a *App) initDebps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServcer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServcer(ctx context.Context) error {
	logger.Init(getCore(getAtomicLevel()))

	srv := http.Server{
		Addr:           a.serviceProvider.HTTPConfig().Address(),
		Handler:        a.serviceProvider.Engine(ctx),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	a.httpServer = &srv
	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("server run on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()

	if err != nil {
		return err
	}
	return nil
}

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 3,
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	develomentCfg := zap.NewDevelopmentEncoderConfig()
	develomentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(develomentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

func getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level

	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
