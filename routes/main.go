package routes

import "net/http"

func ListenAndServe() {
	var server = http.Server{
		Addr:              ":9999",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       10,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	server.ListenAndServe()
}
