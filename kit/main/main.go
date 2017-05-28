package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	slog "log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/go-kit/kit/log"
	"github.com/myesui/uuid/kit/one"
	"github.com/myesui/uuid/kit/two"
	"github.com/myesui/uuid/kit/four"
	"github.com/myesui/uuid/kit/three"
	"github.com/myesui/uuid"
	"github.com/myesui/uuid/kit/five"
)

const (
	defaultPort              = "8080"
)

func main() {
	var (
		address  = envString("PORT", defaultPort)
		httpAddress          = flag.String("http.addr", ":"+address, "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	writer := log.NewSyncWriter(os.Stderr)

	logger = log.NewLogfmtLogger(writer)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	config := &uuid.GeneratorConfig{
		Logger: slog.New(writer, "", slog.LstdFlags),
	}

	v1Service := one.NewService(config).
		Add(one.NewLoggingMiddleware(log.With(logger, "component", "one"))).
		Add(one.NewInstrumentingMiddleware())

	v2Service := two.NewService(config).
		Add(two.NewLoggingMiddleware(log.With(logger, "component", "two"))).
		Add(two.NewInstrumentingMiddleware())

	v3Service := three.NewService(config).
		Add(three.NewLoggingMiddleware(log.With(logger, "component", "three"))).
		Add(three.NewInstrumentingMiddleware())

	v4Service := four.NewService(config).
		Add(four.NewLoggingMiddleware(log.With(logger, "component", "four"))).
		Add(four.NewInstrumentingMiddleware())

	v5Service := five.NewService(config).
		Add(five.NewLoggingMiddleware(log.With(logger, "component", "five"))).
		Add(five.NewInstrumentingMiddleware())

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/one/v1/", one.MakeHandler(v1Service, httpLogger))
	mux.Handle("/two/v1/", two.MakeHandler(v2Service, httpLogger))
	mux.Handle("/three/v1/", three.MakeHandler(v3Service, httpLogger))
	mux.Handle("/four/v1/", four.MakeHandler(v4Service, httpLogger))
	mux.Handle("/five/v1/", five.MakeHandler(v5Service, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)

	go func() {
		logger.Log("transport", "http", "address", *httpAddress, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddress, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

