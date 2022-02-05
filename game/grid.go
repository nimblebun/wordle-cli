/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import (
	"github.com/charmbracelet/lipgloss"
	"pkg.nimblebun.works/wordle-cli/common"
)

func (m *AppModel) setGridItem(i, j int, letter byte, state common.LetterState) {
	m.Grid[i][j] = &common.GridItem{
		Letter: letter,
		State:  state,
	}
}

func (m *AppModel) getLetterForIndex(row int, col int) (byte, bool) {
	if row == m.CurrentRow && col == m.CurrentColumn {
		return '_', true
	}

	if (row >= m.CurrentRow && col >= m.CurrentColumn) || row > m.CurrentRow {
		return ' ', true
	}

	if m.Grid[row][col] == nil {
		return ' ', true
	}

	return m.Grid[row][col].Letter, row == m.CurrentRow
}

func (m *AppModel) renderGridRow(rowIndex int, row [common.WordleWordLength]*common.GridItem) string {
	output := make([]string, len(row))

	for colIndex := range row {
		letter, active := m.getLetterForIndex(rowIndex, colIndex)

		if letter == '_' {
			output = append(output, m.renderTile(letter, lipgloss.Color("#ffffff")))
			continue
		}

		if active {
			output = append(output, m.renderTile(letter, lipgloss.Color(common.WordleColorUnknown.Hex())))
			continue
		}

		state := m.getGridLetterState(rowIndex, colIndex)
		color := state.ToLipglossColor()

		output = append(output, m.renderTile(letter, color))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, output...)
}

func (m *AppModel) renderGrid() string {
	var output []string

	for rowIndex, row := range m.Grid {
		renderedRow := m.renderGridRow(rowIndex, row)
		output = append(output, lipgloss.NewStyle().Padding(0, 1).Render(renderedRow))
	}

	return lipgloss.JoinVertical(lipgloss.Top, output...)
}
