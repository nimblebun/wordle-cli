/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package save

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"pkg.nimblebun.works/wordle-cli/common"
)

type Statistics struct {
	GamesPlayed       int         `json:"games_played"`
	GamesWon          int         `json:"games_won"`
	GuessDistribution map[int]int `json:"guess_distribution"`
}

type SaveFile struct {
	LastGameID     int                                                                `json:"last_game_id"`
	LastGameStatus common.GameState                                                   `json:"last_game_status"`
	LastGameGrid   [common.WordleMaxGuesses][common.WordleWordLength]*common.GridItem `json:"last_game_grid"`
	Statistics     Statistics                                                         `json:"statistics"`
}

func loadSave(savepath string) (*SaveFile, error) {
	data, err := ioutil.ReadFile(savepath)
	if err != nil {
		return nil, err
	}

	save := New()
	err = json.Unmarshal(data, save)

	if err != nil {
		return nil, err
	}

	return save, nil
}

func getSaveLocation(id string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.wordlecli_%s.save.json", dir, id), nil
}

func New() *SaveFile {
	return &SaveFile{
		LastGameID:     -1,
		LastGameStatus: common.GameStateRunning,
		LastGameGrid:   [common.WordleMaxGuesses][common.WordleWordLength]*common.GridItem{},
		Statistics: Statistics{
			GamesPlayed: 0,
			GamesWon:    0,
			GuessDistribution: map[int]int{
				1: 0,
				2: 0,
				3: 0,
				4: 0,
				5: 0,
				6: 0,
			},
		},
	}
}

func Load(id string) (*SaveFile, error) {
	savepath, err := getSaveLocation(id)
	if err != nil {
		return nil, err
	}

	return loadSave(savepath)
}

func Save(save *SaveFile, id string) error {
	savepath, err := getSaveLocation(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(save)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(savepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
