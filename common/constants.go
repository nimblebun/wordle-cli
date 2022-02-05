/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package common

import "fmt"

const (
	WordleMaxGuesses = 6
	WordleWordLength = 5
)

type WordleColor int

const (
	WordleColorUnknown WordleColor = iota
	WordleColorExactMatch
	WordleColorContainedMatch
	WordleColorNoMatch
)

func (c WordleColor) Hex() string {
	switch c {
	case WordleColorUnknown:
		return "#d3d6da"
	case WordleColorExactMatch:
		return "#6aaa64"
	case WordleColorContainedMatch:
		return "#c9b458"
	case WordleColorNoMatch:
		return "#787c7e"
	default:
		panic(fmt.Sprintf("Unknown color: %d", c))
	}
}
