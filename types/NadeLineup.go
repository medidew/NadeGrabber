package types

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
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

func NewNadeLineup(
	x 			float64,
	y 			float64,
	z 			float64,
	view_x 		float32,
	view_y 		float32,
	isCrouching bool,
	isWalking 	bool,
	grenadeType common.EquipmentType,
	username 	string,
	steamid 	uint64,
) NadeLineup {
	return NadeLineup{x, y, z, view_x, view_y, isCrouching, isWalking, grenadeType, username}
}

// Returns a human readable description of the non-numerical Lineup data
func (nl NadeLineup) String() string {
	descriptor := "<" + nl.username + "> threw a <" + nl.grenadeType.String() + ">"

	if nl.isCrouching {
		descriptor += " while crouching"
	} else if nl.isWalking {
		descriptor += " while walking"
	}

	return descriptor
}

// Returns a setpos+setang command to the player's location at the time the nade is thrown.
func (nl NadeLineup) Command() string {
	return "setpos " +
		strconv.FormatFloat(nl.x, 'f', -1, 64) + " " +
		strconv.FormatFloat(nl.y, 'f', -1, 64) + " " +
		strconv.FormatFloat(nl.z, 'f', -1, 64) + "; " +
		"setang " +
		strconv.FormatFloat(float64(nl.view_y), 'f', -1, 32) + " " +
		strconv.FormatFloat(float64(nl.view_x), 'f', -1, 32) + ";"
}

// Concatenates the results of String() and Command() on separate lines
func (nl NadeLineup) Descriptor() string {
	return nl.String() + "\n	" + nl.Command() + "\n";
}