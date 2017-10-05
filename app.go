package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"

	"github.com/enicho/gosample/src/config"
	"github.com/enicho/gosample/src/hello"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	var configs *config.Config
	configs = InitConfig()

	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal(err)
	}

	hwm := hello.NewHelloWorldModule(configs)

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	go logging.StatsLog()

	log.Fatal(grace.Serve(":9000", nil))
}

func InitConfig() *config.Config {
	var cfg config.Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	return &cfg
}
