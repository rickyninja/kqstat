package kqstat

import (
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/rickyninja/kqstat/event"
	"github.com/rickyninja/kqstat/mock/kqstatd"
)

func TestReplay_PastEOF(t *testing.T) {
	r := strings.NewReader(`![k[playernames],v[alpha,beta,gamma,delta,epsilon,zêta,êta,thêta,iota,kappa]]!
![k[spawn],v[10,False]]!
![k[playerKill],v[1190,860,8,5,Worker]]!
`)
	cl, _ := clientServer(t, r)
	for i := 0; i < 10; i++ {
		ev, err := cl.GetEvent()
		if err != nil {
			t.Fatal(err)
		}
		r := i % 3
		switch r {
		case 0:
			v, ok := ev.(event.PlayerNames)
			if !ok {
				t.Errorf("expected playernames got %#v", ev)
				t.Log(v)
			}
		case 1:
			v, ok := ev.(event.Spawn)
			if !ok {
				t.Errorf("expected spawn got %#v", ev)
				t.Log(v)
			}
		case 2:
			v, ok := ev.(event.PlayerKill)
			if !ok {
				t.Errorf("expected playerKill got %#v", ev)
				t.Log(v)
			}
		}
	}
}

func clientServer(t *testing.T, r io.Reader) (*Client, *kqstatd.Replay) {
	t.Helper()
	replay, err := kqstatd.NewReplay(r)
	if err != nil {
		t.Fatal(err)
	}
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	//t.Logf("listener address: %s", l.Addr())
	go http.Serve(l, replay)
	cl, err := NewClient(l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	return cl, replay
}
