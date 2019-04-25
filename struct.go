package pokemon

// Represents a Pokémon nature.
type Nature uint8

// All the different Pokémon natures.
const (
    NatureHardy Nature = iota
    NatureLonely
    NatureBrave
    NatureAdamant
    NatureNaughty
    NatureBold
    NatureDocile
    NatureRelaxed
    NatureImpish
    NatureLax
    NatureTimid
    NatureHasty
    NatureSerious
    NatureJolly
    NatureNaive
    NatureModest
    NatureMild
    NatureQuiet
    NatureBashful
    NatureRash
    NatureCalm
    NatureGentle
    NatureSassy
    NatureCareful
    NatureQuirky
)

// Represents a Pokémon base stat.
type BaseStat uint8

// Represents a Pokémon stat.
type Stat uint32

// All the different Pokémon stats.
const (
    StatHP = iota
    StatAtk
    StatDef
    StatSpAtk
    StatSpDef
    StatSpd
)

// Represents a Pokémon stat modifier, such as EV and IVs.
type StatModifiers struct {
    Nature Nature
    EVs    [6]BaseStat
    IVs    [6]BaseStat
}

// Represents a Pokémon type.
type Type uint8

// All the different Pokémon types.
const (
    TypeNone Type = iota
    TypeNormal
    TypeFire
    TypeFighting
    TypeWater
    TypeFlying
    TypeGrass
    TypePoison
    TypeElectric
    TypeGround
    TypePsychic
    TypeRock
    TypeIce
    TypeBug
    TypeDragon
    TypeGhost
    TypeDark
    TypeSteel
    TypeFairy
)

// Represents a Pokémon description.
type Description struct {
    Name  string
    Types [2]Type
    Stats [6]BaseStat
}

// Represents a Pokémon.
type Pokemon interface {
    // The current HP of the Pokémon.
    Health() int
    // The stats of the Pokémon, taking into account EVs, IVs and Nature.
    // The returned slice should not be modified.
    Stats() []Stat
    // Description of the Pokémon. The returned value should not be modified.
    Describe() *Description
}

// Aliases for frequently used types.
type (
    Stats     = [6]Stat
    BaseStats = [6]BaseStat
)
