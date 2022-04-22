/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import (
	"fmt"
	"math"

	"github.com/charmbracelet/lipgloss"
	"pkg.nimblebun.works/wordle-cli/common"
)

func (m *AppModel) renderTitle() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.renderTile('W', common.LetterStateNoMatch.ToLipglossColor()),
		m.renderTile('O', common.LetterStateExactMatch.ToLipglossColor()),
		m.renderTile('R', common.LetterStateContainedMatch.ToLipglossColor()),
		m.renderTile('D', common.LetterStateNoMatch.ToLipglossColor()),
		m.renderTile('L', common.LetterStateContainedMatch.ToLipglossColor()),
		m.renderTile('E', common.LetterStateNoMatch.ToLipglossColor()),
	)
}

func (m *AppModel) renderLeadingBlock() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		m.renderTitle(),
		fmt.Sprintf("%s (ID #%d)", m.GameType, m.ID),
	)
}

func (m *AppModel) renderTrailingBlock() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.NewStyle().Padding(0, 2).Render("ctrl+c - quit"),
		lipgloss.NewStyle().Padding(0, 2).Render("enter - submit word"),
		lipgloss.NewStyle().Padding(0, 2).Render("backspace - remove last"),
	)
}

func (m *AppModel) getMaxGuessDistribution() int {
	max := 0

	for _, v := range m.SaveData.Statistics.GuessDistribution {
		if v > max {
			max = v
		}
	}

	return max
}

func (m *AppModel) getDistributionProgressBar(max, idx int) string {
	count := m.SaveData.Statistics.GuessDistribution[idx]
	distribution := float64(count) / float64(max)
	bar := fmt.Sprintf("%d: ", idx)

	for i := 0; i < int(distribution*40); i++ {
		bar += "█"
	}

	bar += fmt.Sprintf(" %d", count)

	if count == max {
		return lipgloss.NewStyle().Foreground(common.LetterStateExactMatch.ToLipglossColor()).SetString(bar).String()
	}

	return bar
}

func (m *AppModel) renderStatisticsBlock() string {
	var output []string
	var quickStats []string

	total := m.SaveData.Statistics.GamesPlayed

	quickStats = append(
		quickStats,
		fmt.Sprintf("%s: %d", lipgloss.NewStyle().Bold(true).SetString("Total"), total),
	)

	wonPercentage := float64(m.SaveData.Statistics.GamesWon) / float64(total) * 100
	wonShortPercentage := math.Round(wonPercentage*100) / 100
	wonDisplayPercentage := fmt.Sprintf("%.2f%%", wonShortPercentage)
	if wonPercentage == wonShortPercentage {
		wonDisplayPercentage = fmt.Sprintf("%d%%", int(wonPercentage))
	}

	quickStats = append(
		quickStats,
		lipgloss.NewStyle().MarginLeft(5).Render(
			fmt.Sprintf("%s: %s", lipgloss.NewStyle().Bold(true).SetString("Won %"), wonDisplayPercentage),
		),
	)

	output = append(output, lipgloss.JoinHorizontal(lipgloss.Center, quickStats...))
	output = append(output, lipgloss.NewStyle().Bold(true).Render("\nGuess Distribution\n"))

	maxDistribution := m.getMaxGuessDistribution()

	for i := 1; i <= 6; i++ {
		output = append(output, m.getDistributionProgressBar(maxDistribution, i))
	}

	output = append(output, "\n\nPress ← to view game summary.\nPress Ctrl+S to copy the share string.\nPress Ctrl+C to exit.")

	statistics := lipgloss.JoinVertical(lipgloss.Left, output...)

	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		Render(statistics)
}

func (m *AppModel) renderFinalMessageBlock() string {
	message := fmt.Sprintf(
		"%s\n\n%s\n\n",
		m.GameState.GetMessage(m.CurrentRow, string(m.Word[:])),
		m.getShareString(),
	)

	if m.SaveData != nil {
		message += "Press → to view statistics.\n"
	}

	if m.GameType == common.GameTypeRandom {
		message += "Press Ctrl+N to start a new game.\n"
	}

	message += "Press Ctrl+S to copy the share string.\nPress Ctrl+C to exit."

	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		Render(message)
}
