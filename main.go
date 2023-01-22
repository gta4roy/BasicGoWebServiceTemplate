package main

import (
	"flag"
	"gta4roy/app/log"
	"gta4roy/app/util"
	"gta4roy/app/webserver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	listenAddr string
)

func init() {
	// Parse log level from command line
	logLevel := util.GetProperty(util.LogLevel)
	// Calling the SetLogLevel with the command-line argument
	//log.SetLogLevel(logLevel, "logs.txt")
	log.SetLogLevelForStdOUT(logLevel)

	log.Trace.Println("Loging initialised")
	flag.StringVar(&listenAddr, "listen-addr", util.GetProperty(util.Host)+":"+util.GetProperty(util.Port), "server listen address")
	flag.Parse()

}

func main() {
	log.Trace.Println("Application started ")

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	server := webserver.NewWebServer(listenAddr)

	go webserver.GraceFullShutdown(server, quit, done)
	log.Trace.Println("Server is ready to handle request %s", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error.Println("Could not listen on the %s : %v\n", listenAddr, err)
	}

	<-done

	log.Trace.Println("Server stopped....")
}
