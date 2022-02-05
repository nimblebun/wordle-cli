/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package common

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type LetterState int

const (
	LetterStateUnknown LetterState = iota
	LetterStateExactMatch
	LetterStateContainedMatch
	LetterStateNoMatch
)

func (ls LetterState) ToLipglossColor() lipgloss.Color {
	switch ls {
	case LetterStateUnknown:
		return lipgloss.Color(WordleColorUnknown.Hex())
	case LetterStateExactMatch:
		return lipgloss.Color(WordleColorExactMatch.Hex())
	case LetterStateContainedMatch:
		return lipgloss.Color(WordleColorContainedMatch.Hex())
	case LetterStateNoMatch:
		return lipgloss.Color(WordleColorNoMatch.Hex())
	default:
		panic(fmt.Sprintf("Unknown letter state: %d", ls))
	}
}

func (ls LetterState) String() string {
	switch ls {
	case LetterStateExactMatch:
		return "🟩"
	case LetterStateContainedMatch:
		return "🟨"
	case LetterStateNoMatch:
		return "⬜"
	default:
		return ""
	}
}

type GridItem struct {
	Letter byte        `json:"letter"`
	State  LetterState `json:"state"`
}

type GameState int

const (
	GameStateRunning GameState = iota
	GameStateWon
	GameStateLost
)

func (gs GameState) GetMessage(attempts int) string {
	switch gs {
	case GameStateWon:
		switch attempts {
		case 1:
			return "Genius! 😱"
		case 2:
			return "Magnificent! 😲"
		case 3:
			return "Impressive! 🤩"
		case 4:
			return "Splendid! 👏"
		case 5:
			return "Great! 😊"
		case 6:
			return "Phew! 🎉"
		default:
			return "You won! 🎉"
		}
	case GameStateLost:
		return "You lost! 😔"
	default:
		return ""
	}
}

type GameType int

const (
	GameTypeOfficial GameType = iota
	GameTypeDaily
	GameTypeRandom
)

func (gt GameType) String() string {
	switch gt {
	case GameTypeOfficial:
		return "Official word of the day"
	case GameTypeDaily:
		return "Wordle CLI word of the day"
	case GameTypeRandom:
		return "Random wordle"
	default:
		panic(fmt.Sprintf("Unknown game type: %d", gt))
	}
}

func (gt GameType) ID() string {
	switch gt {
	case GameTypeOfficial:
		return "official"
	case GameTypeDaily:
		return "daily"
	case GameTypeRandom:
		return "random"
	default:
		panic(fmt.Sprintf("Unknown game type: %d", gt))
	}
}
