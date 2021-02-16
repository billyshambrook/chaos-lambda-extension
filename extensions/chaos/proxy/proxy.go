package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Start(logger *log.Logger, port string) {
	go startServer(logger, port)
}

func startServer(logger *log.Logger, port string) {
	server := &http.Server{
		Addr: ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handleTunneling(w, r)
			}
		}),
	}
	logger.Fatal(server.ListenAndServe())
}

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	chaos()

	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func chaos() {
	latency := getenvInt("CHAOS_LATENCY_MS")
	time.Sleep(time.Duration(latency) * time.Millisecond)
}

func getenvInt(key string) int {
	s := os.Getenv(key)
	if len(s) == 0 {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}
