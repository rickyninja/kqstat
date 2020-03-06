package event

import (
	"reflect"
	"testing"
	"time"
)

func TestParseKV(t *testing.T) {
	t.Parallel()
	line := "![k[alive],v[10:26:03 PM]]!\n"
	want := pair{
		Key:   "alive",
		Value: "10:26:03 PM",
	}
	got, err := parseKV(line)
	if err != nil {
		t.Fatal(err)
	}
	if got.Key != want.Key {
		t.Errorf("wrong Key: got %s want %s", got.Key, want.Key)
	}
	if got.Value != want.Value {
		t.Errorf("wrong Value: got %s want %s", got.Value, want.Value)
	}
}

func TestParseEvent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line string
		want Event
	}{
		{"![k[alive],v[1:25:40 PM]]!", Alive{"1:25:40 PM"}},
		{"![k[berryDeposit],v[763,957,3]]!", BerryDeposit{X: 763, Y: 957, Who: GoldStripes}},
		{"![k[berryKickIn],v[831,684,4]]!", BerryKickIn{X: 831, Y: 684, Who: BlueStripes}},
		{"![k[blessMaiden],v[1220,260,Blue]]!", BlessMaiden{X: 1220, Y: 260, Team: Blue}},
		{"![k[carryFood],v[3]]!", CarryFood{Who: GoldStripes}},

		{"![k[gameend],v[map_night,False,128.2457,False]]!", GameEnd{Map: Night, Orientation: BlueOnLeft, Duration: time.Duration(128245700000)}},
		{"![k[gamestart],v[map_dusk,False,0,False]]!", GameStart{Map: Dusk, Orientation: BlueOnLeft}},
		{"![k[getOffSnail: ],v[950,11,,4]]!", GetOffSnail{X: 950, Y: 11, Who: BlueStripes}},
		{"![k[getOnSnail: ],v[950,11,4]]!", GetOnSnail{X: 950, Y: 11, Who: BlueStripes}},
		{"![k[glance],v[1,2]]!", Glance{Attacker: GoldQueen, Target: BlueQueen}},
		{"![k[playerKill],v[1301,1014,1,10,Soldier]]!", PlayerKill{X: 1301, Y: 1014, Slayer: GoldQueen, Slain: BlueChecks, SlainClass: Soldier}},
		{"![k[playernames],v[one,two,three,four,five,six,seven,eight,nine,ten]]!", PlayerNames{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}},
		{"![k[reserveMaiden],v[560,260,10]]!", ReserveMaiden{X: 560, Y: 260, Who: BlueChecks}},
		{"![k[snailEat],v[976,11,3,4]]!", SnailEat{X: 976, Y: 11, Rider: GoldStripes, Meal: BlueStripes}},
		{"![k[snailEscape],v[317,11,4]]!", SnailEscape{X: 317, Y: 11, Who: BlueStripes}},
		{"![k[spawn],v[3,False]]!", Spawn{Who: GoldStripes, IsAI: false}},
		{"![k[unreserveMaiden],v[1220,260,,7]]!", UnreserveMaiden{X: 1220, Y: 260, Who: GoldSkulls}},
		{"![k[useMaiden],v[700,260,maiden_wings,6]]!", UseMaiden{X: 700, Y: 260, Buff: Wings, Who: BlueAbs}},
		{"![k[victory],v[Gold,military]]!", Victory{Team: Gold, Type: Military}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.line, func(t *testing.T) {
			t.Parallel()
			ev, err := Parse(tc.line)
			if err != nil {
				t.Fatal(err)
			}
			switch want := tc.want.(type) {
			case Alive:
				got := ev.(Alive)
				if got != want {
					t.Errorf("wrong Alive got %#v want %#v", got, want)
				}
			case BerryDeposit:
				got := ev.(BerryDeposit)
				if got != want {
					t.Errorf("wrong BerryDeposit got %#v want %#v", got, want)
				}
			case BerryKickIn:
				got := ev.(BerryKickIn)
				if got != want {
					t.Errorf("wrong BerryKickIn got %#v want %#v", got, want)
				}
			case BlessMaiden:
				got := ev.(BlessMaiden)
				if got != want {
					t.Errorf("wrong BlessMaiden got %#v want %#v", got, want)
				}
			case CarryFood:
				got := ev.(CarryFood)
				if got != want {
					t.Errorf("wrong CarryFood got %#v want %#v", got, want)
				}
			case GameEnd:
				got := ev.(GameEnd)
				if got != want {
					t.Errorf("wrong GameEnd got %#v want %#v", got, want)
				}
			case GameStart:
				got := ev.(GameStart)
				if got != want {
					t.Errorf("wrong GameStart got %#v want %#v", got, want)
				}
			case GetOffSnail:
				got := ev.(GetOffSnail)
				if got != want {
					t.Errorf("wrong GetOffSnail got %#v want %#v", got, want)
				}
			case GetOnSnail:
				got := ev.(GetOnSnail)
				if got != want {
					t.Errorf("wrong GetOnSnail got %#v want %#v", got, want)
				}
			case Glance:
				got := ev.(Glance)
				if got != want {
					t.Errorf("wrong Glance got %#v want %#v", got, want)
				}
			case PlayerKill:
				got := ev.(PlayerKill)
				if got != want {
					t.Errorf("wrong PlayerKill got %#v want %#v", got, want)
				}
			case PlayerNames:
				got := ev.(PlayerNames)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("wrong PlayerNames got %#v want %#v", got, want)
				}
			case ReserveMaiden:
				got := ev.(ReserveMaiden)
				if got != want {
					t.Errorf("wrong ReserveMaiden got %#v want %#v", got, want)
				}
			case SnailEat:
				got := ev.(SnailEat)
				if got != want {
					t.Errorf("wrong SnailEat got %#v want %#v", got, want)
				}
			case SnailEscape:
				got := ev.(SnailEscape)
				if got != want {
					t.Errorf("wrong SnailEscape got %#v want %#v", got, want)
				}
			case Spawn:
				got := ev.(Spawn)
				if got != want {
					t.Errorf("wrong Spawn got %#v want %#v", got, want)
				}
			case UnreserveMaiden:
				got := ev.(UnreserveMaiden)
				if got != want {
					t.Errorf("wrong UnreserveMaiden got %#v want %#v", got, want)
				}
			case UseMaiden:
				got := ev.(UseMaiden)
				if got != want {
					t.Errorf("wrong UseMaiden got %#v want %#v", got, want)
				}
			case Victory:
				got := ev.(Victory)
				if got != want {
					t.Errorf("wrong Victory got %#v want %#v", got, want)
				}
			}
		})
	}
}

func TestNewVictory(t *testing.T) {
	t.Parallel()
	v := NewVictory("Blue,economic")
	if v.Team != Blue {
		t.Errorf("wrong team, got %s want %s", v.Team, Blue)
	}
	if v.Type != Economic {
		t.Errorf("wrong win condition, got %s want %s", v.Type, Economic)
	}
}

func TestNewAlive(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  Alive
	}{
		{"10:26:03 PM", Alive{Time: "10:26:03 PM"}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewAlive(tc.value)
			if got.Time != want.Time {
				t.Errorf("wrong Time, got %s want %s", got.Time, want.Time)
			}
		})
	}
}

func TestNewKill(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  PlayerKill
	}{
		{value: "638,519,1,6,Worker", want: PlayerKill{
			X: 638, Y: 519,
			Slayer:     GoldQueen,
			Slain:      BlueAbs,
			SlainClass: Worker,
		}},
		{value: "1053,418,2,9,Soldier", want: PlayerKill{
			X: 1053, Y: 418,
			Slayer:     BlueQueen,
			Slain:      GoldChecks,
			SlainClass: Soldier,
		}},
		{value: "870,275,2,1,Queen", want: PlayerKill{
			X: 870, Y: 275,
			Slayer:     BlueQueen,
			Slain:      GoldQueen,
			SlainClass: Queen,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewPlayerKill(tc.value)
			if got.X != want.X {
				t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
			}
			if got.Y != want.Y {
				t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
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
		})
	}
}

func TestNewBlessMaiden(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  BlessMaiden
	}{
		{value: "700,560,Blue", want: BlessMaiden{
			X: 700, Y: 560,
			Team: Blue,
		}},
		{value: "960,700,Red", want: BlessMaiden{
			X: 960, Y: 700,
			Team: Red,
		}},
		{value: "1220,260,Gold", want: BlessMaiden{
			X: 1220, Y: 260,
			Team: Gold,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewBlessMaiden(tc.value)
			if got.X != want.X {
				t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
			}
			if got.Y != want.Y {
				t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
			}
			if got.Team != want.Team {
				t.Errorf("wrong Team, got %s want %s", got.Team, want.Team)
			}
		})
	}
}

func TestNewReserveMaiden(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  ReserveMaiden
	}{
		{value: "410,860,3", want: ReserveMaiden{
			X: 410, Y: 860,
			Who: GoldStripes,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewReserveMaiden(tc.value)
			if got.X != want.X {
				t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
			}
			// ndf.@j
			if got.Y != want.Y {
				t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
			}
			if got.Who != want.Who {
				t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
			}
		})
	}
}

func TestNewUnreserveMaiden(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  UnreserveMaiden
	}{
		{value: "410,860,,3", want: UnreserveMaiden{
			X: 410, Y: 860,
			Who: GoldStripes,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewUnreserveMaiden(tc.value)
			if got.X != want.X {
				t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
			}
			if got.Y != want.Y {
				t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
			}
			if got.Who != want.Who {
				t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
			}
		})
	}
}

func TestNewUseMaiden(t *testing.T) {
	t.Parallel()
	cases := []struct {
		value string
		want  UseMaiden
	}{
		{value: "960,500,maiden_wings,3", want: UseMaiden{
			X: 960, Y: 500,
			Buff: Wings,
			Who:  GoldStripes,
		}},
		{value: "340,140,maiden_speed,10", want: UseMaiden{
			X: 340, Y: 140,
			Buff: Speed,
			Who:  BlueChecks,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewUseMaiden(tc.value)
			if got.X != want.X {
				t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
			}
			if got.Y != want.Y {
				t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
			}
			if got.Buff != want.Buff {
				t.Errorf("wrong Buff, got %s want %s", got.Buff, want.Buff)
			}
			if got.Who != want.Who {
				t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
			}
		})
	}
}

func TestNewPlayerNames(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	got := NewCarryFood("5")
	want := CarryFood{Who: GoldAbs}
	if got != want {
		t.Errorf("wrong CarryFood, got %d want %d", got.Who, want.Who)
	}
}

func TestNewGameStart(t *testing.T) {
	t.Parallel()
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
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewGameStart(tc.value)
			if got.Map != want.Map {
				t.Errorf("wrong Map, got %s want %s", got.Map, want.Map)
			}
		})
	}
}

func TestNewGameEnd(t *testing.T) {
	t.Parallel()
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
		tc := tc
		t.Run(tc.value, func(t *testing.T) {
			t.Parallel()
			want := tc.want
			got := NewGameEnd(tc.value)
			if got.Map != want.Map {
				t.Errorf("wrong Map, got %s want %s", got.Map, want.Map)
			}
			if got.Duration != want.Duration {
				t.Errorf("wrong Duration, got %s want %s", got.Duration, want.Duration)
			}
		})
	}
}

func TestNewSpawn(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	got := NewGetOnSnail("621,11,6")
	want := GetOnSnail{X: 621, Y: 11, Who: BlueAbs}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewGetOffSnail(t *testing.T) {
	t.Parallel()
	got := NewGetOffSnail("579,11,,8")
	want := GetOffSnail{X: 579, Y: 11, Who: BlueSkulls}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewSnailEat(t *testing.T) {
	t.Parallel()
	got := NewSnailEat("163,11,7,6")
	want := SnailEat{X: 163, Y: 11, Rider: Bee(7), Meal: Bee(6)}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Rider != want.Rider {
		t.Errorf("wrong Rider, got %d want %d", got.Rider, want.Rider)
	}
	if got.Meal != want.Meal {
		t.Errorf("wrong Meal, got %d want %d", got.Meal, want.Meal)
	}
}

func TestNewSnailEscape(t *testing.T) {
	t.Parallel()
	got := NewSnailEscape("910,11,4")
	want := SnailEscape{X: 910, Y: 11, Who: Bee(4)}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewBerryDeposit(t *testing.T) {
	t.Parallel()
	got := NewBerryDeposit("1745,139,6")
	want := BerryDeposit{X: 1745, Y: 139, Who: Bee(6)}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}

func TestNewBerryKickIn(t *testing.T) {
	t.Parallel()
	got := NewBerryKickIn("1030,972,7")
	want := BerryKickIn{X: 1030, Y: 972, Who: Bee(7)}
	if got.X != want.X {
		t.Errorf("wrong X axis, got %d want %d", got.X, want.X)
	}
	if got.Y != want.Y {
		t.Errorf("wrong Y axis, got %d want %d", got.Y, want.Y)
	}
	if got.Who != want.Who {
		t.Errorf("wrong Who, got %d want %d", got.Who, want.Who)
	}
}
