package event

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
	Gold Team = "GOLD"
	Blue      = "BLUE"
)

type Maiden string

const (
	Warrior Maiden = "WARRIOR"
	Speed          = "SPEED"
)

type Map string

const (
	Day   Map = "DAY"
	Night     = "NIGHT"
	Dusk      = "DUSK"
)

type Victory struct {
	Team Team
	Type WinCondition
}

type WinCondition string

const (
	Military WinCondition = "MILITARY"
	Econimic              = "ECONIMIC"
	Snail                 = "SNAIL"
)

type Axis struct {
	X int
	Y int
}

type Kill struct {
	Pos    Axis
	Slayer Bee
	Slain  Bee
}

type BlessMaiden struct {
	Pos  Axis
	Team Team
}

type ReserveMaiden struct {
	Pos  Axis
	Team Team
}

type UnreserveMaiden struct {
	Pos  Axis
	Team Team
}

type UseMaiden struct {
	Pos  Axis
	Type Maiden
	Who  Bee
}

// This is a placeholder for the rfid stuff.
type PlayerNames struct{}

type CabOrientation string

// Read keys as LeftRight
const (
	BlueOnLeft = "BLUE_ON_LEFT"
	GoldOnLeft = "GOLD_ON_LEFT"
)

type Glance struct {
	Attacker Bee
	Target   Bee
}

type CarryFood struct {
	Drone Bee
}

type GameStart struct {
	Map         Map
	Orientation CabOrientation
}

type GameEnd struct {
	Map         Map
	Orientation CabOrientation
	Duration    float64 // time in seconds? time.Duration?
}

type Spawn struct {
	Who  Bee
	IsAI bool
}

type GetOnSnail struct {
	Pos Axis
	Who Bee
}

type GetOffSnail struct {
	Pos Axis
	Who Bee
}

type SnailEat struct {
	Pos   Axis
	Rider Bee
	Meal  Bee
}

type SnailEscape struct {
	Pos Axis
	Who Bee
}

type BerryDeposit struct {
	Pos Axis
	Who Bee
}

type BerryKickIn struct {
	Pos Axis
	Who Bee
}
