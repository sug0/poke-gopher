package pokemon_test

import (
    "testing"

    "github.com/sug0/poke-gopher"
    "github.com/sug0/poke-gopher/gen7"
)

func TestCreate(t *testing.T) {
    poke, err := gen7.NewPokemonFromName("Arceus", &pokemon.StatModifiers{
        Nature: pokemon.NatureTimid,
        IVs: pokemon.BaseStats{31, 31, 31, 31, 31, 31},
        EVs: pokemon.BaseStats{
            pokemon.StatSpAtk: 252,
            pokemon.StatSpDef: 4,
            pokemon.StatSpd: 252,
        },
    })
    check(t, err)
    dump(t, poke.Describe())
    dump(t, poke.Stats())
}

func check(t *testing.T, err error) {
    if err != nil {
        t.Fatalf("failed: %s\n", err)
    }
}

func dump(t *testing.T, x interface{}) {
    t.Logf("%#v\n", x)
}
