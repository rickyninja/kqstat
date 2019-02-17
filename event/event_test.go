package event

import (
	"testing"
	"time"
)

func TestNewVictory(t *testing.T) {
	v := NewVictory("Blue,economic")
	if v.Team != Blue {
		t.Errorf("wrong team, got %s want %s", v.Team, Blue)
	}
	if v.Type != Economic {
		t.Errorf("wrong win condition, got %s want %s", v.Type, Economic)
	}
}

func TestNewKill(t *testing.T) {
	cases := []struct {
		value string
		want  Kill
	}{
		{value: "638,519,1,6,Worker", want: Kill{
			Pos:        Axis{X: 638, Y: 519},
			Slayer:     GoldQueen,
			Slain:      BlueAbs,
			SlainClass: Worker,
		}},
		{value: "1053,418,2,9,Soldier", want: Kill{
			Pos:        Axis{X: 1053, Y: 418},
			Slayer:     BlueQueen,
			Slain:      GoldChecks,
			SlainClass: Soldier,
		}},
		{value: "870,275,2,1,Queen", want: Kill{
			Pos:        Axis{X: 870, Y: 275},
			Slayer:     BlueQueen,
			Slain:      GoldQueen,
			SlainClass: Queen,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewKill(tc.value)
		if got.Pos.X != want.Pos.X {
			t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
		}
		if got.Pos.Y != want.Pos.Y {
			t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
		}
		if got.Slayer != want.Slayer {
			t.Errorf("wrong Slayer, got %d want %d", got.Slayer, want.Slayer)
		}
		if got.Slain != want.Slain {
			t.Errorf("wrong Slain, got %d want %d", got.Slain, want.Slain)
		}
		if got.SlainClass != want.SlainClass {
			t.Errorf("wrong SlainClass, got %s want %s", got.SlainClass, want.SlainClass)
		}
	}
}

func TestNewBlessMaiden(t *testing.T) {
	cases := []struct {
		value string
		want  BlessMaiden
	}{
		{value: "700,560,Blue", want: BlessMaiden{
			Pos:  Axis{X: 700, Y: 560},
			Team: Blue,
		}},
		{value: "960,700,Red", want: BlessMaiden{
			Pos:  Axis{X: 960, Y: 700},
			Team: Red,
		}},
		{value: "1220,260,Gold", want: BlessMaiden{
			Pos:  Axis{X: 1220, Y: 260},
			Team: Gold,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewBlessMaiden(tc.value)
		if got.Pos.X != want.Pos.X {
			t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
		}
		if got.Pos.Y != want.Pos.Y {
			t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
		}
		if got.Team != want.Team {
			t.Errorf("wrong Team, got %s want %s", got.Team, want.Team)
		}
	}
}

func TestNewReserveMaiden(t *testing.T) {
	cases := []struct {
		value string
		want  ReserveMaiden
	}{
		{value: "410,860,3", want: ReserveMaiden{
			Pos: Axis{X: 410, Y: 860},
			Who: GoldStripes,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewReserveMaiden(tc.value)
		if got.Pos.X != want.Pos.X {
			t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
		}
		if got.Pos.Y != want.Pos.Y {
			t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
		}
		if got.Who != want.Who {
			t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
		}
	}
}

func TestNewUnreserveMaiden(t *testing.T) {
	cases := []struct {
		value string
		want  UnreserveMaiden
	}{
		{value: "410,860,,3", want: UnreserveMaiden{
			Pos: Axis{X: 410, Y: 860},
			Who: GoldStripes,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewUnreserveMaiden(tc.value)
		if got.Pos.X != want.Pos.X {
			t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
		}
		if got.Pos.Y != want.Pos.Y {
			t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
		}
		if got.Who != want.Who {
			t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
		}
	}
}

func TestNewUseMaiden(t *testing.T) {
	cases := []struct {
		value string
		want  UseMaiden
	}{
		{value: "960,500,maiden_wings,3", want: UseMaiden{
			Pos:  Axis{X: 960, Y: 500},
			Buff: Wings,
			Who:  GoldStripes,
		}},
		{value: "340,140,maiden_speed,10", want: UseMaiden{
			Pos:  Axis{X: 340, Y: 140},
			Buff: Speed,
			Who:  BlueChecks,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewUseMaiden(tc.value)
		if got.Pos.X != want.Pos.X {
			t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
		}
		if got.Pos.Y != want.Pos.Y {
			t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
		}
		if got.Buff != want.Buff {
			t.Errorf("wrong Buff, got %s want %s", got.Buff, want.Buff)
		}
		if got.Who != want.Who {
			t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
		}
	}
}

func TestNewPlayerNames(t *testing.T) {
	want := []string{"Kim", "Kia", "Tyler M.", "Tyler D.", "Sam W.", "Sam G.", "Logan", "Liz", "Max", "San"}
	got := NewPlayerNames("Kim,Kia,Tyler M.,Tyler D.,Sam W.,Sam G.,Logan,Liz,Max,San")
	if len(got) != len(want) {
		t.Errorf("wrong length, got %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("wrong value at index %d, got %s want %s", i, got[i], want[i])
			t.Logf(" got [%s]", got[i])
			t.Logf("want [%s]", want[i])
		}
	}
}

func TestNewGlance(t *testing.T) {
	got := NewGlance("1,10")
	want := Glance{Attacker: GoldQueen, Target: BlueChecks}
	if got.Attacker != want.Attacker {
		t.Errorf("wrong Attacker, got %d want %d", got.Attacker, want.Attacker)
	}
	if got.Target != want.Target {
		t.Errorf("wrong Target, got %d want %d", got.Target, want.Target)
	}
}

func TestNewCarryFood(t *testing.T) {
	got := NewCarryFood("5")
	want := CarryFood{Who: GoldAbs}
	if got != want {
		t.Errorf("wrong CarryFood, got %d want %d", got.Who, want.Who)
	}
}

func TestNewGameStart(t *testing.T) {
	cases := []struct {
		value string
		want  GameStart
	}{
		{value: "map_day,False,0,False", want: GameStart{
			Map: Day,
		}},
		{value: "map_night,False,0,False", want: GameStart{
			Map: Night,
		}},
		{value: "map_dusk,False,0,False", want: GameStart{
			Map: Dusk,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewGameStart(tc.value)
		if got.Map != want.Map {
			t.Errorf("wrong Map, got %s want %s", got.Map, want.Map)
		}
	}
}

func TestNewGameEnd(t *testing.T) {
	cases := []struct {
		value string
		want  GameEnd
	}{
		{value: "map_day,False,88.59263,False", want: GameEnd{
			Map:      Day,
			Duration: time.Microsecond * 88592630,
		}},
		{value: "map_night,False,180.12487,False", want: GameEnd{
			Map:      Night,
			Duration: time.Minute*3 + time.Microsecond*124870,
		}},
		{value: "map_dusk,False,1800.00000,False", want: GameEnd{
			Map:      Dusk,
			Duration: time.Minute * 30,
		}},
	}
	for _, tc := range cases {
		want := tc.want
		got := NewGameEnd(tc.value)
		if got.Map != want.Map {
			t.Errorf("wrong Map, got %s want %s", got.Map, want.Map)
		}
		if got.Duration != want.Duration {
			t.Errorf("wrong Duration, got %s want %s", got.Duration, want.Duration)
		}
	}
}

func TestNewSpawn(t *testing.T) {
	got := NewSpawn("2,False")
	want := Spawn{Who: BlueQueen, IsAI: false}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
	if got.IsAI != want.IsAI {
		t.Errorf("wrong IsAI, got %t want %t", got.IsAI, want.IsAI)
	}
}

func TestNewGetOnSnail(t *testing.T) {
	got := NewGetOnSnail("621,11,6")
	want := GetOnSnail{Pos: Axis{X: 621, Y: 11}, Who: BlueAbs}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewGetOffSnail(t *testing.T) {
	got := NewGetOffSnail("579,11,,8")
	want := GetOffSnail{Pos: Axis{X: 579, Y: 11}, Who: BlueSkulls}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewSnailEat(t *testing.T) {
	got := NewSnailEat("163,11,7,6")
	want := SnailEat{Pos: Axis{X: 163, Y: 11}, Rider: Bee(7), Meal: Bee(6)}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Rider != want.Rider {
		t.Errorf("wrong Rider, got %d want %d", got.Rider, want.Rider)
	}
	if got.Meal != want.Meal {
		t.Errorf("wrong Meal, got %d want %d", got.Meal, want.Meal)
	}
}

func TestNewSnailEscape(t *testing.T) {
	got := NewSnailEscape("910,11,4")
	want := SnailEscape{Pos: Axis{X: 910, Y: 11}, Who: Bee(4)}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewBerryDeposit(t *testing.T) {
	got := NewBerryDeposit("1745,139,6")
	want := BerryDeposit{Pos: Axis{X: 1745, Y: 139}, Who: Bee(6)}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewBerryKickIn(t *testing.T) {
	got := NewBerryKickIn("1030,972,7")
	want := BerryKickIn{Pos: Axis{X: 1030, Y: 972}, Who: Bee(7)}
	if got.Pos.X != want.Pos.X {
		t.Errorf("wrong X axis, got %d want %d", got.Pos.X, want.Pos.X)
	}
	if got.Pos.Y != want.Pos.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Pos.Y, want.Pos.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}
