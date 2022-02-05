/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import "github.com/charmbracelet/lipgloss"

func (m *AppModel) renderTile(ch byte, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(color).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Render(string(ch))
}
