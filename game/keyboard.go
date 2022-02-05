/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *AppModel) renderRow(keys []string) string {
	output := make([]string, len(keys))

	for _, key := range keys {
		state := m.getLetterState(key[0])
		color := state.ToLipglossColor()
		output = append(output, m.renderTile(key[0], color))
	}

	return lipgloss.JoinHorizontal(lipgloss.Bottom, output...)
}

func (m *AppModel) renderKeyboard() string {
	rows := [][]string{
		strings.Split("QWERTYUIOP", ""),
		strings.Split("ASDFGHJKL", ""),
		strings.Split("ZXCVBNM", ""),
	}

	keys := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().Padding(1, 0).Render(m.renderLeadingBlock()),
		lipgloss.NewStyle().Padding(0, 1).Render(m.renderRow(rows[0])),
		lipgloss.NewStyle().Padding(0, 1).Render(m.renderRow(rows[1])),
		m.renderRow(rows[2]),
	)

	return keys
}
