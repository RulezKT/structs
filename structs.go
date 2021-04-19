package structs

import (
	"math"
	"strconv"

	"github.com/RulezKT/gregdate"
)

const (
	//Astronomical Unit
	AU = 0.1495978707e9 // km 149597870.7

	// Здесь мы как начальное значение ставим eps = 23°26'21,448" градуса согласно CD
	// double const RAD_TO_DEG = 5.7295779513082320877e1;
	// Obliquity of the ecliptic  = 23°26'21,448"  - на 1 января 2000 года = 23.43929111111111
	// 23.43929111111111/5.7295779513082320877e1 = 0.4090928042223289
	MED_EPS = 0.4090928042223289

	SSB       = 0
	MERCURY   = 1 // 7,01° (относительно эклиптики)
	VENUS     = 2 // 3,39458° (относительно эклиптики)
	EARTH     = 3
	MARS      = 4 // 1,85061° (относительно эклиптики)
	JUPITER   = 5 // 1,304° (относительно эклиптики)
	SATURN    = 6 // 2,485 240° (относительно эклиптики)
	URANUS    = 7 // 0,772556° (относительно эклиптики)
	NEPTUNE   = 8 // 1,767975° (относительно эклиптики)
	PLUTO     = 9 // 17°,14 (относительно эклиптики)
	SUN       = 10
	MOON      = 11 // 5,14° (относительно эклиптики)
	NORTHNODE = 12
	SOUTHNODE = 13
	HIRON     = 14

	HEAD   = 0
	AJNA   = 1
	THROAT = 2
	G      = 3
	SACRAL = 4
	ROOT   = 5
	EGO    = 6
	SPLEEN = 7
	EMO    = 8

	NUMBEROFGATES    = 65 //from 1 to 64
	NUMBEROFCHANNELS = 37 //from 1 to 36
)

// from 0 to 13, don't count Hiron yet
const NUMBEROFPLANETS = 14

// from 0 to 8
const NUMBEROFCENTERS = 9

// main type with complete calcuated information
type CdInfo struct {
	HdInfo
	FdInfo
	AstroInfo
	NumerologyInfo
}

type HdInfo struct {
	Personality HdObjects
	Design      HdObjects
	Gates       [NUMBEROFGATES]Gate // from 1 to 64
	Channels    [NUMBEROFCHANNELS]Channel
	Centers     Centers
	Phs
	Variable string
	Psychology
	Cross      Cross
	Profile    string
	Authority  string
	Definition string
	Type       string
}

func (hd *HdInfo) Init() {

	for i := 1; i < NUMBEROFGATES; i++ {
		hd.Gates[i].Number = i

		if i < NUMBEROFCHANNELS {
			hd.Channels[i].Number = i
		}

	}

	hd.Personality.Planets.Init()
	hd.Design.Init()

	hd.Centers.Init()

}

type Gate struct {
	Number  int
	Pers    int //сколько раз активированы по личности
	Des     int // сколько раз активированы по дизайну
	Defined bool
}

// 36 каналов ДЧ
/*

  1 - 64-47
  2 - 61-24
  3 - 63-4

  4 - 17-62
  5 - 43-23
  6 - 11-56

  7 - 48-16

  8 - 57-20
  9 - 34-20
  10- 10-20
  11- 57-10
  12- 57-34
  13- 34-10

  14- 7-31
  15- 1-8
  16- 13-33
  17- 21-45
  18- 22-12
  19- 36-35

  20- 5-15
  21- 14-2
  22- 29-46
  23- 51-25

  24- 44-26
  25- 27-50
  26- 59-6
  27- 37-40

  28- 54-32
  29- 38-28
  30- 58-18

  31- 53-42
  32- 60-3
  33- 52-9

  34- 19-49
  35- 39-55
  36- 41-30

*/
type Channel struct {
	Number     int
	FirstGate  Gate
	SecondGate Gate
	Defined    bool
}

type HdObjects struct {
	Planets
	Centers Centers
	TimeData
	Authority string
}

type Planet struct {
	Longitude float64 //in Radians
	Name      string
	Number    int
	HdStructure
	FdStructure
	ZodiacStructure
}

type Planets struct {
	Planet [NUMBEROFPLANETS]Planet
}

func (pl *Planets) Init() {
	pl.Planet[0] = Planet{Name: "SSB", Number: 0}
	pl.Planet[1] = Planet{Name: "Mercury", Number: 1}
	pl.Planet[2] = Planet{Name: "Venus", Number: 2}
	pl.Planet[3] = Planet{Name: "Earth", Number: 3}
	pl.Planet[4] = Planet{Name: "Mars", Number: 4}
	pl.Planet[5] = Planet{Name: "Jupiter", Number: 5}
	pl.Planet[6] = Planet{Name: "Saturn", Number: 6}
	pl.Planet[7] = Planet{Name: "Uranus", Number: 7}
	pl.Planet[8] = Planet{Name: "Neptune", Number: 8}
	pl.Planet[9] = Planet{Name: "Pluto", Number: 9}
	pl.Planet[10] = Planet{Name: "Sun", Number: 10}
	pl.Planet[11] = Planet{Name: "Moon", Number: 11}
	pl.Planet[12] = Planet{Name: "NorthNode", Number: 12}
	pl.Planet[13] = Planet{Name: "SouthNde", Number: 13}
	//{Name: "Hiron", Number: 14},

}

/*
func InitPlanets() Planets {

	return Planets{[NUMBEROFPLANETS]Planet{
		{Name: "SSB", Number: 0},
		{Name: "Mercury", Number: 1},
		{Name: "Venus", Number: 2},
		{Name: "Earth", Number: 3},
		{Name: "Mars", Number: 4},
		{Name: "Jupiter", Number: 5},
		{Name: "Saturn", Number: 6},
		{Name: "Uranus", Number: 7},
		{Name: "Neptune", Number: 8},
		{Name: "Pluto", Number: 9},
		{Name: "Sun", Number: 10},
		{Name: "Moon", Number: 11},
		{Name: "NorthNode", Number: 12},
		{Name: "SouthNde", Number: 13},
		//{Name: "Hiron", Number: 14},
	}}

}
*/

type TimeData struct {
	LocalTime     gregdate.GregDate //для design всегда 0
	UtcTime       gregdate.GregDate
	TypeOfTyme    int    //Изначальный источник данных 2 - local time, 1- UTC Time,  0 - Ephemeries time
	Offset        int    //смещение локального времени от UTC в секундах
	SecFromJd2000 int64  // Ephemeries time
	Place         string // не пустой, только если время изначально Local, для design всегда пустой
}

type HdStructure struct {

	//номера округляются вверх

	Hex int // соответствует воротам

	Line float64

	Color float64

	Tone float64

	Base float64

	NumberOfPassedDegrees float64 //сколько пройдено в градусах от начала гексаграммы
}

func (hd HdStructure) String() string {

	return strconv.Itoa(hd.Hex) + "." + strconv.Itoa(int(math.Ceil(hd.Line))) + "." + strconv.Itoa(int(math.Ceil(hd.Color))) + "." + strconv.Itoa(int(math.Ceil(hd.Tone))) + "." + strconv.Itoa(int(math.Ceil(hd.Base)))
}

type FdInfo struct {
}

type FdStructure struct {
	Power     int
	Direction string // D - директное  R - ретроградное  S - стационарное
}

type AstroInfo struct {
}

type ZodiacStructure struct {
	Degrees int
	Minutes int
	Seconds int
	Zodiac  string
}

type Phs struct {
	Theme     string
	NutrType  string
	Cognition string
}

type Psychology struct {
	Motivation string
	Mind       string
}

type Cross struct {
	First  int
	Second int
	Third  int
	Forth  int
}

type NumerologyInfo struct {
}

type Centers struct {
	Center map[string]bool
}

func (cent *Centers) Init() {
	/*
		cent.Centers[0] = Center{Name: "Head"}
		cent.Centers[1] = Center{Name: "Ajna"}
		cent.Centers[2] = Center{Name: "Throat"}
		cent.Centers[3] = Center{Name: "G"}
		cent.Centers[4] = Center{Name: "Sacral"}
		cent.Centers[5] = Center{Name: "Root"}
		cent.Centers[6] = Center{Name: "Ego"}
		cent.Centers[7] = Center{Name: "Spleen"}
		cent.Centers[8] = Center{Name: "Emo"}

	*/
	cent.Center = make(map[string]bool, 9)
	cent.Center["Head"] = false
	cent.Center["Ajna"] = false
	cent.Center["Throat"] = false
	cent.Center["G"] = false
	cent.Center["Sacral"] = false
	cent.Center["Root"] = false
	cent.Center["Ego"] = false
	cent.Center["Spleen"] = false
	cent.Center["Emo"] = false

}

/*
http://astro.ukho.gov.uk/nao/miscellanea/DeltaT/
https://ru.wikipedia.org/wiki/%D0%94%D0%B5%D0%BB%D1%8C%D1%82%D0%B0_T
https://eclipse.gsfc.nasa.gov/SEhelp/deltatpoly2004.html
https://en.wikipedia.org/wiki/%CE%94T
*/
type DeltaTTableStructure struct {
	Year    int
	Seconds float64
}

//первый и последнй года таблицы значений Дельта Т для быстрого доступа и сама таблица
type DeltaTTable struct {
	FirstYear int
	LastYear  int
	Table     []DeltaTTableStructure
}
