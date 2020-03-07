// Package event is focused on deserializing text obtained from the Killerqueen stats service into types representing those events.
package event

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Event represents an event returned from the stats service.  You'll probably want to type switch this to get details.
type Event interface{}

// Bee is a positive integer in the stats text representing one of the 10 player positions.
type Bee int

func (b Bee) String() string {
	switch b {
	case GoldQueen:
		return "gold-queen"
	case BlueQueen:
		return "blue-queen"
	case GoldStripes:
		return "gold-stripes"
	case BlueStripes:
		return "blue-stripes"
	case GoldAbs:
		return "gold-abs"
	case BlueAbs:
		return "blue-abs"
	case GoldSkulls:
		return "gold-skulls"
	case BlueSkulls:
		return "blue-skulls"
	case GoldChecks:
		return "gold-checks"
	case BlueChecks:
		return "blue-checks"
	}
	return "unknown-bee"
}

const (
	_           Bee = iota
	GoldQueen       // 1
	BlueQueen       // 2
	GoldStripes     // 3
	BlueStripes     // 4
	GoldAbs         // 5
	BlueAbs         // 6
	GoldSkulls      // 7
	BlueSkulls      // 8
	GoldChecks      // 9
	BlueChecks      // 10
)

// Team is which team an event occurred for.
type Team string

const (
	Gold Team = "Gold"
	Blue      = "Blue"
	// Red team is used in the game's demo, and isn't seen in real games.
	Red = "Red"
)

// Buff is a worker getting speed or becoming a warrior.
type Buff string

const (
	// Wings is warrior gate usage.
	Wings Buff = "maiden_wings"
	// Speed is speed gate usage.
	Speed Buff = "maiden_speed"
)

// Map is a game map or level.
type Map string

// TODO should bonus maps be listed here?
const (
	Day   Map = "map_day"
	Night     = "map_night"
	Dusk      = "map_dusk"
)

// CabOrientation indicates which side of each other gold & blue cab are positioned.
type CabOrientation string

const (
	BlueOnLeft CabOrientation = "BLUE_ON_LEFT"
	GoldOnLeft                = "GOLD_ON_LEFT"
)

// NewVictory creates a Victory type from victory event text.
func NewVictory(v string) Victory {
	vals := strings.Split(v, ",")
	if len(vals) < 2 {
		log.Printf("value should have at least 2 values: %s", v)
		return Victory{}
	}
	return Victory{
		Team: Team(vals[0]),
		Type: WinCondition(vals[1]),
	}
}

// Victory is info on the winning team.
type Victory struct {
	// Team is which team won.
	Team Team
	// Type is how the game was won.
	Type WinCondition
}

// WinCondition is how the game was won.
type WinCondition string

const (
	Military WinCondition = "military"
	Economic              = "economic"
	Snail                 = "snail"
)

// event: kqstat.Pair{Key:"playerKill", Value:"638,519,1,6,Worker"}
// event: kqstat.Pair{Key:"playerKill", Value:"1053,418,2,9,Soldier"}
// event: kqstat.Pair{Key:"playerKill", Value:"870,275,2,1,Queen"}

// Class is the role a bee is performing in the game.
type Class string

const (
	// Worker is what we refer to as drone.
	Worker Class = "Worker"
	// Soldier is what we refer to as warrior.
	Soldier = "Soldier"
	// Queen is the queen.
	Queen = "Queen"
)

// ![k[alive],v[10:26:03 PM]]!
// I don't think the value is needed; the reply being used has an empty value.

// Alive is a keepalive event.  The time string in this event doesn't appear to be used for anything,
// so it's simply stored as is, so it can be seen in log output potentially.
type Alive struct {
	Time string
}

// NewAlive creates an Alive event from alive event text.
func NewAlive(v string) Alive {
	return Alive{Time: v}
}

// NewPlayerKill creates a PlayerKill event from playerkill event text.
func NewPlayerKill(v string) PlayerKill {
	vals := strings.Split(v, ",")
	if len(vals) < 5 {
		log.Printf("value should have at least 5 values: %s", v)
		return PlayerKill{}
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[0], err)
		return PlayerKill{}
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[1], err)
		return PlayerKill{}
	}
	slayer, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return PlayerKill{}
	}
	slain, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return PlayerKill{}
	}
	class := vals[4]
	return PlayerKill{
		X:          x,
		Y:          y,
		Slayer:     Bee(slayer),
		Slain:      Bee(slain),
		SlainClass: Class(class),
	}
}

// PlayerKill is a playerkill event.
type PlayerKill struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Slayer is who did the killing.
	Slayer Bee
	// Slain is the target of the killing.
	Slain Bee
	// SlainClass is the class or role of the bee that was slain.
	SlainClass Class
}

// NewBlessMaiden creates a BlessMaiden event from blessMaiden event text.
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
		X:    x,
		Y:    y,
		Team: t,
	}
}

// BlessMaiden represents a queen tagging a gate.
type BlessMaiden struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Team is blue or gold team.
	Team Team
}

// NewReserveMaiden creates a ReserveMaiden type from reserveMaiden event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

// ReserveMaiden represents a worker using a gate to obtain a Buff.
type ReserveMaiden struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is which player is using the gate.
	Who Bee
}

// NewUnreserveMaiden create a UnreserveMaiden type from unreserveMaiden event text.
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
	// vals[2] seems to always be empty.
	w, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return UnreserveMaiden{}
	}
	return UnreserveMaiden{
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

// UnreserveMaiden represents a worker exiting a gate before the Buff is received.
type UnreserveMaiden struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is which player is using the gate.
	Who Bee
}

// NewUseMaiden creates a UseMaiden event from useMaiden event text.
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
		X:    x,
		Y:    y,
		Buff: Buff(vals[2]),
		Who:  Bee(w),
	}
}

// UseMaiden represents a worker using a gate to obtain a Buff.
type UseMaiden struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Buff is the buff the gate provides.
	Buff Buff
	// Who is which player is using the gate.
	Who Bee
}

// NewPlayerNames creates a PlayerNames event from playernames event text.
// This is a placeholder for the rfid stuff.
func NewPlayerNames(v string) PlayerNames {
	vals := strings.Split(v, ",")
	if len(vals) != 10 {
		log.Printf("value should have 10 values: %s", v)
		return nil
	}
	return PlayerNames(vals)
}

// PlayerNames is a list of players in order of their positions on the cabs.
type PlayerNames []string

// NewGlance creates a Glance event from glance event text.
func NewGlance(v string) Glance {
	vals := strings.Split(v, ",")
	if len(vals) < 2 {
		log.Printf("value should have at least 2 values: %s", v)
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

// Glance is an event representing when an attacker bounces off their target instead of delivering a killing blow.
type Glance struct {
	// Attacker is intended killer.
	Attacker Bee
	// Target is the intended victim.
	Target Bee
}

// NewCarryFood creates a CarryFood event from carryFood event text.
func NewCarryFood(v string) CarryFood {
	w, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return CarryFood{Who: Bee(w)}
}

// CarryFood is an event representing a worker picking up a berry.
type CarryFood struct {
	// Who is the worker picking up a berry.
	Who Bee
}

// NewGameStart creates a GameStart event from gamestart event text.
func NewGameStart(v string) GameStart {
	vals := strings.Split(v, ",")
	if len(vals) < 4 {
		log.Printf("value should have at least 4 values: %s", v)
		return GameStart{}
	}
	or := parseOrientation(vals[1])
	return GameStart{
		Map:         Map(vals[0]),
		Orientation: or,
	}
}

// GameStart is an event indicating a new game is beginning.
type GameStart struct {
	// Map is which map the game will be played on.
	Map Map
	// Orientation is how the cabs are positioned next to each other.
	Orientation CabOrientation
}

// NewGameEnd creates a GameEnd event from gameend event text.
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
	or := parseOrientation(vals[1])
	return GameEnd{
		Map:         Map(vals[0]),
		Orientation: or,
		Duration:    dur,
	}
}

// GameEnd is an event representing the end of a game.
type GameEnd struct {
	// Map is which map the game was played on.
	Map Map
	// Orientation is how the cabs are positioned next to each other.
	Orientation CabOrientation
	// Duration is how long the game lasted.
	Duration time.Duration
}

// NewSpawn creates a Spawn event from spawn event text.
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

// Spawn is an event that represents a player spawning.
type Spawn struct {
	// Who is which player is spawning by their position on the sticks.
	Who Bee
	// IsAI indicates if this is a player or a robot.
	IsAI bool
}

// NewGetOnSnail creates a GetOnSnail event from getOnSnail event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

// GetOnSnail is an event representing a worker beginning to ride on the snail.
type GetOnSnail struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is which worker is mounting the snail.
	Who Bee
}

// NewGetOffSnail creates a GetOffSnail event from getOffSnail event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

// GetOffSnail is an event representing a worker ending a ride on the snail.
type GetOffSnail struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is which worker is dismounting the snail.
	Who Bee
}

// NewSnailEat creates a SnailEat event from snailEat event text.
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
	r, err := strconv.Atoi(vals[2])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[2], err)
		return gos
	}
	m, err := strconv.Atoi(vals[3])
	if err != nil {
		log.Printf("failed Atoi for %s: %s", vals[3], err)
		return gos
	}
	return SnailEat{
		X:     x,
		Y:     y,
		Rider: Bee(r),
		Meal:  Bee(m),
	}
}

// SnailEat is an event representing the snail beginning to eat a worker.
type SnailEat struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Rider is which worker is mounted on the snail.
	Rider Bee
	// Meal is which worker is being eaten by the snail.
	Meal Bee
}

// NewSnailEscape creates a SnailEscape event from snailEscape event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

type SnailEscape struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is the worker that escaped the mouth of the snail.
	Who Bee
}

// NewBerryDeposit creates a BerryDeposit event from berryDeposit event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

// BerryDeposit is an event representing a worker putting a berry in their hive.
type BerryDeposit struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is which worker deposited the berry.
	Who Bee
}

// NewBerryKickIn creates a BerryKickIn event from berryKickIn event text.
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
		X:   x,
		Y:   y,
		Who: Bee(w),
	}
}

type BerryKickIn struct {
	// X & Y are coordinates where the event occurred.
	X, Y int
	// Who is who kicked in the berry.
	Who Bee
}

// Parse parses a line of event text from the stats service, and returns it as an event.
func Parse(line string) (Event, error) {
	p, err := parseKV(line)
	if err != nil {
		return nil, err
	}
	switch p.Key {
	case "alive":
		return NewAlive(p.Value), nil
	case "berryDeposit":
		return NewBerryDeposit(p.Value), nil
	case "berryKickIn":
		return NewBerryKickIn(p.Value), nil
	case "blessMaiden":
		return NewBlessMaiden(p.Value), nil
	case "carryFood":
		return NewCarryFood(p.Value), nil
	case "gameend":
		return NewGameEnd(p.Value), nil
	case "gamestart":
		return NewGameStart(p.Value), nil
	case "getOffSnail: ":
		return NewGetOffSnail(p.Value), nil
	case "getOnSnail: ":
		return NewGetOnSnail(p.Value), nil
	case "glance":
		return NewGlance(p.Value), nil
	case "playerKill":
		return NewPlayerKill(p.Value), nil
	case "playernames":
		return NewPlayerNames(p.Value), nil
	case "reserveMaiden":
		return NewReserveMaiden(p.Value), nil
	case "snailEat":
		return NewSnailEat(p.Value), nil
	case "snailEscape":
		return NewSnailEscape(p.Value), nil
	case "spawn":
		return NewSpawn(p.Value), nil
	case "unreserveMaiden":
		return NewUnreserveMaiden(p.Value), nil
	case "useMaiden":
		return NewUseMaiden(p.Value), nil
	case "victory":
		return NewVictory(p.Value), nil
	}
	return nil, fmt.Errorf("unknown event: %v", p)
}

// pair is an event after parsing into a key and value.
type pair struct {
	Key   string
	Value string
}

// parseKV parses an event line into a pair.
func parseKV(line string) (pair, error) {
	p := pair{}
	v := strings.Split(line, "],v[")
	if len(v) < 2 {
		return p, fmt.Errorf("Failed to parse line: %s", line)
	}
	p.Key = strings.TrimPrefix(v[0], "![k[")
	p.Value = strings.TrimSuffix(v[1], "\n")
	p.Value = strings.TrimSuffix(p.Value, "]]!")
	return p, nil
}

// parseOrientation sets the CabOrientation based on a boolean value in the event text.
func parseOrientation(v string) CabOrientation {
	var or CabOrientation
	switch v {
	case "True":
		or = GoldOnLeft
	case "False":
		or = BlueOnLeft
	}
	return or
}
