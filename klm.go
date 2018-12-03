package klm

import (
	"github.com/cmodk/go-simplehttp"
	"github.com/sirupsen/logrus"
)

type KLM struct {
	lg    *logrus.Logger
	sh    simplehttp.SimpleHttp
	key   string
	debug bool
}

func New(k string, logger *logrus.Logger) *KLM {
	klm := KLM{
		lg:  logger,
		sh:  simplehttp.New("https://api.klm.com/opendata", logger),
		key: k,
	}

	klm.sh.AddHeader("Accept", "application/hal+json")
	klm.sh.AddHeader("Accept-Language", "en-US")
	klm.sh.AddHeader("api-key", k)

	return &klm

}

func (klm *KLM) SetDebug(d bool) {
	klm.debug = d
	klm.sh.SetDebug(d)
}
