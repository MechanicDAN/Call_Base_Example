package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"test/parsers"
	"time"
)

func serve(ctx context.Context) (err error){


	r := GetRouter()

	srv := &http.Server{
		Addr: parsers.GetHost() + parsers.GetPort(),
		Handler: r,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server.serve: listen:%s\n", err)
		}

	}()
	log.Printf("server started at " + parsers.GetHost() + parsers.GetPort())
	<-ctx.Done()

	log.Printf("server.httpServer.serve: server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server.serve: server Shutdown Failed:%s", err)
	}

	log.Printf("server.httpServer.serve: server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

func StartServer() (err error){

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-c
		log.Printf("server.httpServer.StartServer: system call:%+v", osCall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("server.httpServer.StartServer: ailed to serve:+%v\n", err)
		return err
	}
	return nil
}



