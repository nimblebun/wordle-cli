/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package main

import (
	"log"
	"os"

	"pkg.nimblebun.works/wordle-cli/common"
	"pkg.nimblebun.works/wordle-cli/game"
	"pkg.nimblebun.works/wordle-cli/words"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func startGame(word string, gameType common.GameType, id int) error {
	model := game.NewGame(word, gameType, id)

	program := tea.NewProgram(model, tea.WithAltScreen())
	return program.Start()
}

func startOfficial(c *cli.Context) error {
	word, id := words.GetOfficialWordOfTheDay()
	return startGame(word, common.GameTypeOfficial, id)
}

func startDaily(c *cli.Context) error {
	word, id := words.GetWordOfTheDay()
	return startGame(word, common.GameTypeDaily, id)
}

func startRandom(c *cli.Context) error {
	word, id := words.GetRandomWordle()
	return startGame(word, common.GameTypeRandom, id)
}

func main() {
	app := &cli.App{
		Name:    "wordle-cli",
		Usage:   "play wordle in your terminal",
		Action:  startOfficial,
		Version: "1.0.9",
		Commands: []*cli.Command{
			{
				Name:   "official",
				Usage:  "play official wordle of the day",
				Action: startOfficial,
			},
			{
				Name:   "daily",
				Usage:  "play the CLI's wordle of the day",
				Action: startDaily,
			},
			{
				Name:   "random",
				Usage:  "play a random wordle",
				Action: startRandom,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
