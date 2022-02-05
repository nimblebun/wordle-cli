/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package game

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit      key.Binding
	Backspace key.Binding
	Enter     key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "Quit"),
	),
	Backspace: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "Delete last letter"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Submit word"),
	),
}

func (km KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{km.Quit, km.Backspace, km.Enter},
	}
}

func (km KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		km.Quit,
		km.Backspace,
		km.Enter,
	}
}

// func (m *AppModel) renderHelp() string {
// 	return m.Help.View(Keys)
// }
