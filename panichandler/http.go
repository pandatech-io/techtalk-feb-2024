package panichandler

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func RunHttpServer() {
	port := 8000
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Create endpoint to test panic process and call HTTP wrapper function to wrap our process
	http.HandleFunc("/panic-without-recover", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n")) // Write message to the client
		panic("Panic happens")
	})
	http.HandleFunc("/panic-with-recover", panicHandlerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n")) // Write message to the client
		panic("Panic happens")
	}))
	httpServer.Handler = http.DefaultServeMux

	fmt.Println("HTTP server listening on port", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println("error when ListenAndServe")
		return
	}
}

func panicHandlerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Defer the process of recovery
		defer func() {
			// Recover from panic to stop termination of the application
			if err := recover(); err != nil {
				fmt.Printf("Panic message: %+v\n", err)
				fmt.Println("Function recovered from the panic")
				// use debug.PrintStack() if you want to trace the panic and print it
				debug.PrintStack()
			}
		}()
		// Execute HTTP function that has been wrapped
		next(w, r)
	}
}
