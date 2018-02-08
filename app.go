package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tokopedia/gosample/bigproject"
	"github.com/tokopedia/gosample/bigproject/src/handler"
	"github.com/tokopedia/gosample/hello"
	"github.com/tokopedia/logging/tracer"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(&agent.Options{}); err != nil {
		log.Fatal(err)
	}

	hwm := hello.NewHelloWorldModule()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	go logging.StatsLog()

	tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	// big project http handler
	bigproject.InitBP()
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/get-user", handler.GetUser)
	http.HandleFunc("/get-visitor-count", handler.GetVisitorCount)

	http.Handle("/public/",
		http.StripPrefix("/public/", http.FileServer(http.Dir("bigproject/files/templates/public/"))))

	log.Fatal(grace.Serve(":9000", nil))
}
