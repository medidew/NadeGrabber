package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	ev "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

type NadeLineup struct {
	x           float64
	y           float64
	z           float64
	view_x      float32
	view_y      float32
	isCrouching bool
	isWalking   bool
	grenadeType common.EquipmentType
	username    string
}

func (nl NadeLineup) String() string {
	descriptor := "<" + nl.username + "> threw a <" + nl.grenadeType.String() + ">"

	if nl.isCrouching {
		descriptor += " while crouching"
	} else if nl.isWalking {
		descriptor += " while walking"
	}

	return descriptor + "\n" +
		"	setpos " +
		strconv.FormatFloat(nl.x, 'f', -1, 64) + " " +
		strconv.FormatFloat(nl.y, 'f', -1, 64) + " " +
		strconv.FormatFloat(nl.z, 'f', -1, 64) + "; " +
		"setang " +
		strconv.FormatFloat(float64(nl.view_y), 'f', -1, 32) + " " +
		strconv.FormatFloat(float64(nl.view_x), 'f', -1, 32) + "\n"
}

func main() {
	filename := os.Args[1] // get the demo file path as cmd-line arg

	demo, err := os.Open(filename)

	fmt.Println("File opened")
	fmt.Printf("demo: %v\n", demo.Name())

	if err != nil {
		fmt.Printf("Error occurred opening demo file: %v", err)
	}

	smokes := []NadeLineup{}
	flashes := []NadeLineup{}
	nades := []NadeLineup{}
	molotovs := []NadeLineup{}

	p := demoinfocs.NewParser(demo)
	p.RegisterEventHandler(func(event ev.GrenadeProjectileThrow) {
		thrower := event.Projectile.Thrower
		pos := thrower.Position()

		nl := NadeLineup{
			x:           pos.X,
			y:           pos.Y,
			z:           pos.Z,
			view_x:      thrower.ViewDirectionX(),
			view_y:      thrower.ViewDirectionY(),
			isCrouching: thrower.IsDucking(),
			isWalking:   thrower.IsWalking(),
			grenadeType: event.Projectile.WeaponInstance.Type,
			username:    thrower.Name,
		}

		switch event.Projectile.WeaponInstance.Type {
		case common.EqMolotov:
			molotovs = append(molotovs, nl)
		case common.EqIncendiary:
			molotovs = append(molotovs, nl)
		case common.EqFlash:
			flashes = append(flashes, nl)
		case common.EqSmoke:
			smokes = append(smokes, nl)
		case common.EqHE:
			nades = append(nades, nl)
		}
	})

	p.ParseToEnd()

	output_file, err := os.Create("lineups")

	if err != nil {
		fmt.Printf("Error creating output file: %v", err)
	}

	for _, smoke := range smokes {
		output_file.WriteString(smoke.String())
	}
}
