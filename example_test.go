package kqstat_test

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rickyninja/kqstat"
	"github.com/rickyninja/kqstat/event"
)

func ExampleClient() {
	logger := newMylog(log.New(os.Stderr, "", 0))
	cl, err := kqstat.NewClient(":12749", logger)
	if err != nil {
		log.Fatal(err)
	}
	for {
		ev, err := cl.GetEvent()
		if err != nil {
			logger.Logf("GetEvent: %s", err)
			continue
		}
		switch v := ev.(type) {
		case event.Alive:
			fmt.Printf("alive time: %s\n", v.Time)
		case event.BerryDeposit:
			fmt.Printf("berry deposit x: %d\n", v.X)
			fmt.Printf("berry deposit y: %d\n", v.Y)
			fmt.Printf("berry depositor : %s\n", v.Who)
		case event.BerryKickIn:
			fmt.Printf("berry kick x: %d\n", v.X)
			fmt.Printf("berry kick y: %d\n", v.Y)
			fmt.Printf("berry kicker: %s\n", v.Who)
		case event.BlessMaiden:
			fmt.Printf("bless maiden x: %d\n", v.X)
			fmt.Printf("bless maiden y: %d\n", v.Y)
			fmt.Printf("bless maiden: %s\n", v.Team)
		case event.CarryFood:
			fmt.Printf("carry food by: %s\n", v.Who)
		case event.GameEnd:
			fmt.Printf("game end on map: %s\n", v.Map)
			fmt.Printf("game end orientation: %s\n", v.Orientation)
			fmt.Printf("game end duration: %s\n", v.Duration)
		case event.GameStart:
			fmt.Printf("game start on map: %s\n", v.Map)
			fmt.Printf("game start orientation: %s\n", v.Orientation)
		case event.GetOffSnail:
			fmt.Printf("get off snail x: %d\n", v.X)
			fmt.Printf("get off snail y: %d\n", v.Y)
			fmt.Printf("get off snail rider: %s\n", v.Who)
		case event.GetOnSnail:
			fmt.Printf("get on snail x: %d\n", v.X)
			fmt.Printf("get on snail y: %d\n", v.Y)
			fmt.Printf("get on snail rider: %s\n", v.Who)
		case event.Glance:
			fmt.Printf("glance attacker: %s\n", v.Attacker)
			fmt.Printf("glance target: %s\n", v.Target)
		case event.PlayerKill:
			fmt.Printf("player kill x: %d\n", v.X)
			fmt.Printf("player kill y: %d\n", v.Y)
			fmt.Printf("player kill slayer: %s\n", v.Slayer)
			fmt.Printf("player kill slain: %s\n", v.Slain)
			fmt.Printf("player kill slain class: %s\n", v.SlainClass)
		case event.PlayerNames:
			fmt.Printf("player names: %s\n", strings.Join(v, ", "))
		case event.ReserveMaiden:
			fmt.Printf("reserve maiden x: %d\n", v.X)
			fmt.Printf("reserve maiden y: %d\n", v.Y)
			fmt.Printf("reserve maiden who: %s\n", v.Who)
		case event.SnailEat:
			fmt.Printf("snail eat x: %d\n", v.X)
			fmt.Printf("snail eat y: %d\n", v.Y)
			fmt.Printf("snail eat rider: %s\n", v.Rider)
			fmt.Printf("snail eat meal: %s\n", v.Meal)
		case event.SnailEscape:
			fmt.Printf("snail escape x: %d\n", v.X)
			fmt.Printf("snail escape y: %d\n", v.Y)
			fmt.Printf("snail escape who: %s\n", v.Who)
		case event.Spawn:
			fmt.Printf("spawn who: %s\n", v.Who)
			fmt.Printf("spawn AI: %t\n", v.IsAI)
		case event.UnreserveMaiden:
			fmt.Printf("unreserve maiden x: %d\n", v.X)
			fmt.Printf("unreserve maiden y: %d\n", v.Y)
			fmt.Printf("unreserve maiden who: %s\n", v.Who)
		case event.UseMaiden:
			fmt.Printf("use maiden x: %d\n", v.X)
			fmt.Printf("use maiden y: %d\n", v.Y)
			fmt.Printf("use maiden buff: %s\n", v.Buff)
			fmt.Printf("use maiden who: %s\n", v.Who)
		case event.Victory:
			fmt.Printf("victory team: %s\n", v.Team)
			fmt.Printf("victory type: %s\n", v.Type)
		}
	}
}

// mylog embeds a *log.Logger and gives it the required Logf method.
type mylog struct {
	*log.Logger
}

func newMylog(l *log.Logger) *mylog {
	return &mylog{l}
}

func (l *mylog) Logf(format string, a ...interface{}) {
	l.Printf(format, a...)
}
