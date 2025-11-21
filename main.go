package main

import (
	"fmt"
	"os"

	"github.com/medidew/NadeGrabber/types"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	ev "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./NadeGrabber <path_to_demo_file>")
		return
	}
	filename := os.Args[1] // get the demo file path as cmd-line arg

	demo, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error occurred opening demo file: %v", err)
	}
	fmt.Printf("File '%v' opened\n", demo.Name())

	nades, err := parse(demo)
	if err != nil {
		fmt.Printf("Error occurred during parsing: %v", err)
	}

	output_file, err := os.Create("lineups")
	if err != nil {
		fmt.Printf("Error creating output file: %v", err)
	}

	for _, nade := range nades {
		for _, nl := range nade {
			output_file.WriteString(nl.Descriptor())
		}
	}
}

func parse(demo *os.File) (map[uint64][]types.NadeLineup, error) {
	p := demoinfocs.NewParser(demo)

	nades := map[uint64][]types.NadeLineup{}

	p.RegisterEventHandler(func(event ev.GrenadeProjectileThrow) {
		thrower := event.Projectile.Thrower
		pos := thrower.Position()

		nl := types.NewNadeLineup(
			pos.X,
			pos.Y,
			pos.Z,
			thrower.ViewDirectionX(),
			thrower.ViewDirectionY(),
			thrower.IsDucking(),
			thrower.IsWalking(),
			event.Projectile.WeaponInstance.Type,
			thrower.Name,
			thrower.SteamID64,
		)

		nades[thrower.SteamID64] = append(nades[thrower.SteamID64], nl)
	})

	err := p.ParseToEnd()
	if err != nil {
		return nil, err
	}

	return nades, nil
}
