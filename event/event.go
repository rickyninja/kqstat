package event

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Bee int

const (
	_ Bee = iota
	GoldQueen
	BlueQueen
	GoldStripes
	BlueStripes
	GoldAbs
	BlueAbs
	GoldSkulls
	BlueSkulls
	GoldChecks
	BlueChecks
)

type Team string

const (
	Gold Team = "Gold"
	Blue      = "Blue"
	Red       = "Red"
)

type Buff string

const (
	Wings Buff = "maiden_wings"
	Speed Buff = "maiden_speed"
)

type Map string

const (
	Day   Map = "map_day"
	Night     = "map_night"
	Dusk      = "map_dusk"
)

// event: kqstat.Pair{Key:"victory", Value:"Blue,economic"}

func NewVictory(v string) Victory {
	vals := strings.Split(v, ",")
	if len(vals) < 2 {
		log.Printf("value should have at least 2 values: %s", v)
		return Victory{}
	}
	t, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return Victory{}
	}
	wc, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return Victory{}
	}
	return Victory{
		Team: Team(t),
		Type: WinCondition(wc),
	}
}

type Victory struct {
	Team Team
	Type WinCondition
}

type WinCondition string

const (
	Military WinCondition = "military"
	Econimic              = "econimic"
	Snail                 = "snail"
)

type Axis struct {
	X int
	Y int
}

// event: kqstat.Pair{Key:"playerKill", Value:"638,519,1,6,Worker"}
// event: kqstat.Pair{Key:"playerKill", Value:"1053,418,2,9,Soldier"}
// event: kqstat.Pair{Key:"playerKill", Value:"870,275,2,1,Queen"}

type Class string

const (
	Worker  Class = "Worker"
	Soldier       = "Soldier"
	Queen         = "Queen"
)

func NewKill(v string) Kill {
	vals := strings.Split(v, ",")
	if len(vals) < 5 {
		log.Printf("value should have at least 5 values: %s", v)
		return Kill{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return Kill{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return Kill{}
	}
	slayer, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return Kill{}
	}
	slain, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return Kill{}
	}
	class := vals[4]
	return Kill{
		Pos:        Axis{X: x, Y: y},
		Slayer:     Bee(slayer),
		Slain:      Bee(slain),
		SlainClass: Class(class),
	}
}

type Kill struct {
	Pos        Axis
	Slayer     Bee
	Slain      Bee
	SlainClass Class
}

func NewBlessMaiden(v string) BlessMaiden {
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return BlessMaiden{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return BlessMaiden{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return BlessMaiden{}
	}
	var t Team
	switch strings.ToLower(vals[2]) {
	case "blue":
		t = Blue
	case "gold":
		t = Gold
	case "red":
		t = Red
	}
	return BlessMaiden{
		Pos: Axis{
			X: x,
			Y: y,
		},
		Team: t,
	}
}

type BlessMaiden struct {
	Pos  Axis
	Team Team
}

// event: kqstat.Pair{Key:"reserveMaiden", Value:"410,860,3"}

func NewReserveMaiden(v string) ReserveMaiden {
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return ReserveMaiden{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return ReserveMaiden{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return ReserveMaiden{}
	}
	w, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return ReserveMaiden{}
	}
	return ReserveMaiden{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type ReserveMaiden struct {
	Pos Axis
	Who Bee
}

// event: kqstat.Pair{Key:"unreserveMaiden", Value:"410,860,,3"}
func NewUnreserveMaiden(v string) UnreserveMaiden {
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return UnreserveMaiden{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return UnreserveMaiden{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return UnreserveMaiden{}
	}
	w, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return UnreserveMaiden{}
	}
	return UnreserveMaiden{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type UnreserveMaiden struct {
	Pos Axis
	Who Bee
}

// event: kqstat.Pair{Key:"useMaiden", Value:"960,500,maiden_wings,3"}
// event: kqstat.Pair{Key:"useMaiden", Value:"340,140,maiden_speed,10"}
func NewUseMaiden(v string) UseMaiden {
	um := UseMaiden{}
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return um
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return um
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return um
	}
	w, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return um
	}
	return UseMaiden{
		Pos:  Axis{X: x, Y: y},
		Buff: Buff(vals[2]),
		Who:  Bee(w),
	}
}

type UseMaiden struct {
	Pos  Axis
	Buff Buff
	Who  Bee
}

// This is a placeholder for the rfid stuff.
// event: kqstat.Pair{Key:"playernames", Value:",,,,,,,,,"}
func NewPlayerNames(v string) PlayerNames {
	vals := strings.Split(v, ",")
	if len(vals) != 10 {
		log.Printf("value should have 10 values: %s", v)
		return nil
	}
	return PlayerNames(vals)
}

type PlayerNames []string

type CabOrientation string

// Read keys as LeftRight
const (
	BlueOnLeft = "BLUE_ON_LEFT"
	GoldOnLeft = "GOLD_ON_LEFT"
)

// event: kqstat.Pair{Key:"glance", Value:"1,10"}

func NewGlance(v string) Glance {
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return Glance{}
	}
	att, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return Glance{}
	}
	tar, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return Glance{}
	}
	return Glance{
		Attacker: Bee(att),
		Target:   Bee(tar),
	}
}

type Glance struct {
	Attacker Bee
	Target   Bee
}

func NewCarryFood(v string) CarryFood {
	w, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return CarryFood{Who: Bee(w)}
}

type CarryFood struct {
	Who Bee
}

// event: kqstat.Pair{Key:"gamestart", Value:"map_day,False,0,False"}
func NewGameStart(v string) GameStart {
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return GameStart{}
	}
	return GameStart{Map: Map(vals[0])}
}

type GameStart struct {
	Map         Map
	Orientation CabOrientation
}

// event: kqstat.Pair{Key:"gameend", Value:"map_day,False,88.59263,False"}

func NewGameEnd(v string) GameEnd {
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return GameEnd{}
	}
	dur, err := time.ParseDuration(vals[2] + "s")
	if err != nil {
		log.Printf("failed ParseDuration on %s: %s", vals[2], err)
	}
	return GameEnd{
		Map:      Map(vals[0]),
		Duration: dur,
	}
}

type GameEnd struct {
	Map         Map
	Orientation CabOrientation
	Duration    time.Duration
}

// event: kqstat.Pair{Key:"spawn", Value:"2,False"}

func NewSpawn(v string) Spawn {
	vals := strings.Split(v, ",")
	if len(vals) < 2 {
		log.Printf("value should have at least 2 values: %s", v)
		return Spawn{}
	}
	w, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return Spawn{}
	}
	ai, err := strconv.ParseBool(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return Spawn{}
	}
	return Spawn{
		Who:  Bee(w),
		IsAI: ai,
	}
}

type Spawn struct {
	Who  Bee
	IsAI bool
}

// event: kqstat.Pair{Key:"getOnSnail: ", Value:"621,11,6"}

func NewGetOnSnail(v string) GetOnSnail {
	gos := GetOnSnail{}
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return gos
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return gos
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return gos
	}
	w, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return gos
	}
	return GetOnSnail{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type GetOnSnail struct {
	Pos Axis
	Who Bee
}

// event: kqstat.Pair{Key:"getOffSnail: ", Value:"579,11,,8"}
func NewGetOffSnail(v string) GetOffSnail {
	gos := GetOffSnail{}
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return gos
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return gos
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return gos
	}
	w, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return gos
	}
	return GetOffSnail{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type GetOffSnail struct {
	Pos Axis
	Who Bee
}

// event: kqstat.Pair{Key:"snailEat", Value:"163,11,7,6"}

func NewSnailEat(v string) SnailEat {
	gos := SnailEat{}
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return gos
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return gos
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return gos
	}
	r, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return gos
	}
	m, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return gos
	}
	return SnailEat{
		Pos:   Axis{X: x, Y: y},
		Rider: Bee(r),
		Meal:  Bee(m),
	}
}

type SnailEat struct {
	Pos   Axis
	Rider Bee
	Meal  Bee
}

// 1535414043963 = ![k[snailEscape],v[910,11,4]]!
func NewSnailEscape(v string) SnailEscape {
	se := SnailEscape{}
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return se
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return se
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return se
	}
	w, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return se
	}
	return SnailEscape{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type SnailEscape struct {
	Pos Axis
	Who Bee
}

func NewBerryDeposit(v string) BerryDeposit {
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return BerryDeposit{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return BerryDeposit{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return BerryDeposit{}
	}
	w, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return BerryDeposit{}
	}
	return BerryDeposit{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type BerryDeposit struct {
	Pos Axis
	Who Bee
}

// event: kqstat.Pair{Key:"berryKickIn", Value:"1030,972,7"}
func NewBerryKickIn(v string) BerryKickIn {
	bki := BerryKickIn{}
	vals := strings.Split(v, ",")
	if len(vals) < 3 {
		log.Printf("value should have at least 3 values: %s", v)
		return bki
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return bki
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return bki
	}
	w, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return bki
	}
	return BerryKickIn{
		Pos: Axis{X: x, Y: y},
		Who: Bee(w),
	}
}

type BerryKickIn struct {
	Pos Axis
	Who Bee
}
