package kqstatd_test

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rickyninja/kqstat/mock/kqstatd"
)

func ExampleReplay() {
	reader := strings.NewReader(`![k[alive],v[12:39:14 PM]]!
![k[playerKill],v[750,861,1,2,Queen]]!
![k[playerKill],v[801,860,1,10,Worker]]!
![k[playerKill],v[830,860,1,4,Worker]]!
![k[playerKill],v[830,860,1,6,Worker]]!
`)
	logger := new(simpleLogger)
	replay, err := kqstatd.NewReplay(reader, logger)
	if err != nil {
		logger.Logf("Failed NewReplay: %s", err)
	}
	http.ListenAndServe(":8081", replay)
}

type simpleLogger struct{}

func (s *simpleLogger) Logf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
