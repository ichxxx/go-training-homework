package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"week03/internal/server"
)

func main() {
	errGroup, errCtx := errgroup.WithContext(context.Background())

	// 启动 http server
	errGroup.Go(func() (err error) {
		httpServer := server.NewHttpServer()
		go func() {
			log.Println("server: start")
			_ = httpServer.ListenAndServe()
		}()

		// 监听 linux signal 信号或者 errgroup 中其它 goroutine 的错误
		select {
		case <- listenShutdown():
			err = errors.New("receive shutdown signal")
		case <- errCtx.Done():
		}
		shutdown(httpServer)
		return
	})

	// 启动其它任务
	errGroup.Go(someTask)

	if err := errGroup.Wait(); err != nil {
		log.Println("closed:", err.Error())
	}
}

// 一个不发生或随机在10秒内发生错误的任务
func someTask() error {
	randInt := rand.Intn(10)
	if randInt % 2 == 0 {
		log.Println("task: success")
		return nil
	}
	time.Sleep(time.Duration(randInt) * time.Second)
	log.Println("task: failed")
	return errors.New("task: some error")
}

func listenShutdown() <- chan os.Signal {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	// 随机在10秒发送退出信号
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		shutdownSignal <- syscall.SIGQUIT
	}()

	return shutdownSignal
}

func shutdown(server *http.Server) {
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Println("server: shutdown failed,", err.Error())
	}
	log.Println("server: shutdown")
}