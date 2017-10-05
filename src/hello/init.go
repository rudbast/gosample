package hello

import (
	"expvar"
	"net/http"

	"github.com/enicho/gosample/src/config"

	logging "gopkg.in/tokopedia/logging.v1"
)

type HelloWorldModule struct {
	cfg       *config.Config
	something string
	stats     *expvar.Int
}

func NewHelloWorldModule(cfgs *config.Config) *HelloWorldModule {
	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfgs.Server.Name)

	return &HelloWorldModule{
		cfg:       cfgs,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
	}

}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	hlm.stats.Add(1)
	w.Write([]byte("Hello " + hlm.something))
}
