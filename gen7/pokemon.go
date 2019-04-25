package gen7

import pkm "github.com/sug0/poke-gopher"

type pokemon struct {
    dex    int
    health int
    stats  []pkm.Stat
}

// TODO validate EVs, IVs and Nature

// Create a new Pokémon from its dex entry number.
func NewPokemonFromDex(dex int, mods *pkm.StatModifiers) (pkm.Pokemon, error) {
    if dex <= 0 || dex >= len(allPokes) {
        return nil, ErrNonexistantPokemon
    }

    base := allPokes[dex].Stats[:]
    stats := make([]pkm.Stat, 6)

    // no modifiers to apply...
    if mods == nil {
        // calculate hp
        stats[0] = calcStatHP(base[0], 0, 0)

        // calculate other stats
        for i := 1; i < 6; i++ {
            stats[i] = calcStat(i, pkm.NatureHardy, base[i], 0, 0)
        }

        return &pokemon{
            dex: dex,
            health: int(stats[0]),
            stats: stats,
        }, nil
    }
    // apply modifiers!

    // calculate hp
    stats[0] = calcStatHP(base[0], mods.EVs[0], mods.IVs[0])

    // calculate other stats
    for i := 1; i < 6; i++ {
        stats[i] = calcStat(i, mods.Nature, base[i], mods.EVs[i], mods.IVs[i])
    }

    return &pokemon{
        dex: dex,
        health: int(stats[0]),
        stats: stats,
    }, nil
}

// Create a new Pokémon from its name.
func NewPokemonFromName(name string, mods *pkm.StatModifiers) (pkm.Pokemon, error) {
    return NewPokemonFromDex(nameToDexEntry[name], mods)
}

func (p *pokemon) Health() int {
    return p.health
}

func (p *pokemon) Stats() []pkm.Stat {
    return p.stats
}

func (p *pokemon) Describe() *pkm.Description {
    return &allPokes[p.dex]
}

func calcStatHP(base, ev, iv pkm.BaseStat) pkm.Stat {
    return (pkm.Stat(base) << 1) + pkm.Stat(iv) + (pkm.Stat(ev) >> 2) + 110
}

func calcStat(which int, nature pkm.Nature, base, ev, iv pkm.BaseStat) pkm.Stat {
    mod, ok := natMultipliers[natToUint16(nature, which)]
    stat := (pkm.Stat(base) << 1) + pkm.Stat(iv) + (pkm.Stat(ev) >> 2) + 5
    if !ok {
        return stat
    }
    return pkm.Stat(mod * float32(stat))
}

func natToUint16(nature pkm.Nature, stat int) uint16 {
    return uint16(nature) | (uint16(stat) << 8)
}
