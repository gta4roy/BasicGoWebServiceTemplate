package webserver

import (
	"context"
	"gta4roy/app/log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/handlers"
)

func init() {

	//log.Trace.Println("Server is Starting UP ")
	/*Add any database connection start here*/
}

func CloseConnections() {
	log.Trace.Println("Server is shutting down ")
	/*Add any database connection close here*/
}

func processRequestURL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		log.Trace.Println("URL Path is ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func GraceFullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	sig := <-quit
	log.Trace.Println("Server is shutting down ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	CloseConnections()
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Trace.Println("Could not gracefull shutdown server 5v\n", err)
	}
	close(done)
}

func NewWebServer(listenAddr string) *http.Server {
	router := NewRouter()
	headers := handlers.AllowedHeaders([]string{"Host", "Origin", "Connection", "Upgrade", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "X-Requested-With", "Content-type", "Authorisation", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "access-control-allow-origin", "access-control-allow-headers"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	_, err := os.Getwd()
	if err != nil {
		log.Error.Println(err)
	}
	return &http.Server{
		Addr:         listenAddr,
		Handler:      handlers.CORS(headers, methods, origins, credentials)(processRequestURL(router)),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
}
