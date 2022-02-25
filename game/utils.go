/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import (
	"fmt"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"pkg.nimblebun.works/clipboard"
	"pkg.nimblebun.works/wordle-cli/common"
	"pkg.nimblebun.works/wordle-cli/common/save"
	"pkg.nimblebun.works/wordle-cli/words"
)

func (m *AppModel) getLetterState(letter byte) common.LetterState {
	if state, ok := m.LetterStates[letter]; ok {
		return state
	}

	return common.LetterStateUnknown
}

func (m *AppModel) getGridLetterState(row, col int) common.LetterState {
	if row >= common.WordleMaxGuesses || col >= common.WordleWordLength {
		return common.LetterStateUnknown
	}

	entry := m.Grid[row][col]

	if entry == nil {
		return common.LetterStateUnknown
	}

	return entry.State
}

func (m *AppModel) quit() tea.Cmd {
	return tea.Quit
}

func (m *AppModel) backspace() tea.Cmd {
	if m.GameState != common.GameStateRunning {
		return nil
	}

	if m.CurrentColumn > 0 {
		m.CurrentColumn--
	}

	return nil
}

func (m *AppModel) input(r rune) tea.Cmd {
	if m.GameState != common.GameStateRunning {
		return nil
	}

	if m.CurrentColumn >= common.WordleWordLength || m.CurrentRow >= common.WordleMaxGuesses {
		return nil
	}

	r = unicode.ToUpper(r)

	if !unicode.IsUpper(r) {
		return nil
	}

	m.setGridItem(m.CurrentRow, m.CurrentColumn, byte(r), common.LetterStateUnknown)
	m.CurrentColumn++

	return nil
}

func (m *AppModel) enter() tea.Cmd {
	if m.CurrentColumn < common.WordleWordLength {
		return nil
	}

	if m.GameState != common.GameStateRunning {
		return nil
	}

	wb := strings.Builder{}

	for _, entry := range m.Grid[m.CurrentRow] {
		lowercase := strings.ToLower(string(entry.Letter))
		wb.WriteString(lowercase)
	}

	ok := false

	word := wb.String()

	for _, w := range words.WordList {
		if w == word {
			ok = true
			break
		}
	}

	if !ok {
		for _, w := range words.ValidWordList {
			if w == word {
				ok = true
				break
			}
		}
	}

	if !ok {
		return nil
	}

	word = strings.ToUpper(word)
	targetWord := strings.ToUpper(string(m.Word[:]))

	perfectGuesses := 0
	matchedIndices := make([]bool, len(targetWord))

	for i := range word {
		ok = false

		for j := range targetWord {
			if word[i] == targetWord[j] && !matchedIndices[j] {
				if i == j {
					m.LetterStates[word[i]] = common.LetterStateExactMatch
					m.setGridItem(m.CurrentRow, i, word[i], common.LetterStateExactMatch)
					ok = true
					matchedIndices[j] = true
					perfectGuesses++
					break
				} else {
					m.LetterStates[word[i]] = common.LetterStateContainedMatch
					m.setGridItem(m.CurrentRow, i, word[i], common.LetterStateContainedMatch)
					ok = true
					matchedIndices[j] = true
				}
			}
		}

		if !ok {
			m.LetterStates[word[i]] = common.LetterStateNoMatch
			m.setGridItem(m.CurrentRow, i, word[i], common.LetterStateNoMatch)
		}
	}

	m.CurrentRow++
	m.CurrentColumn = 0

	if perfectGuesses == common.WordleWordLength {
		m.GameState = common.GameStateWon
	}

	if m.CurrentRow > common.WordleMaxGuesses {
		m.GameState = common.GameStateLost
	}

	if m.GameType != common.GameTypeRandom {
		m.save()
	}

	return nil
}

func (m *AppModel) new() tea.Cmd {
	if m.GameType != common.GameTypeRandom {
		return nil
	}

	if m.GameState == common.GameStateRunning {
		return nil
	}

	word, idx := words.GetRandomWordle()

	newModel := NewGame(word, common.GameTypeRandom, idx)

	m.ID = newModel.ID
	m.Word = newModel.Word
	m.LetterStates = newModel.LetterStates
	m.Grid = newModel.Grid
	m.CurrentColumn = newModel.CurrentColumn
	m.CurrentRow = newModel.CurrentRow
	m.GameState = newModel.GameState
	m.NewGame = newModel.NewGame

	return nil
}

func (m *AppModel) setDisplayStatistics(newState bool) tea.Cmd {
	if m.GameState == common.GameStateRunning {
		return nil
	}

	if m.SaveData == nil {
		return nil
	}

	m.DisplayStatistics = newState
	return nil
}

func (m *AppModel) displayStatistics() tea.Cmd {
	return m.setDisplayStatistics(true)
}

func (m *AppModel) displayGameSummary() tea.Cmd {
	return m.setDisplayStatistics(false)
}

func (m *AppModel) getShareString() string {
	var rows []string

	if m.GameState == common.GameStateLost {
		rows = append(rows, fmt.Sprintf("Wordle %d X/6\n", m.ID))
	} else {
		rows = append(rows, fmt.Sprintf("Wordle %d %d/6\n", m.ID, m.CurrentRow))
	}

	for i := 0; i < m.CurrentRow; i++ {
		var row []string

		for _, entry := range m.Grid[i] {
			row = append(row, entry.State.String())
		}

		rows = append(rows, strings.Join(row, ""))
	}

	return strings.Join(rows, "\n")
}

func (m *AppModel) copyShareString(automatic bool) tea.Cmd {
	if m.GameState == common.GameStateRunning {
		return nil
	}

	if automatic && !m.NewGame {
		return nil
	}

	str := strings.ReplaceAll(m.getShareString(), "🔳", "⬜")
	_ = clipboard.WriteAll(str)

	return nil
}

func (m *AppModel) save() {
	m.SaveData.LastGameID = m.ID
	m.SaveData.LastGameGrid = m.Grid
	m.SaveData.LastGameStatus = m.GameState

	if m.GameState != common.GameStateRunning {
		m.SaveData.Statistics.GamesPlayed++
	}

	if m.GameState == common.GameStateWon {
		m.SaveData.Statistics.GamesWon++
		m.SaveData.Statistics.GuessDistribution[m.CurrentRow]++
	}

	_ = save.Save(m.SaveData, m.GameType.ID())
}
