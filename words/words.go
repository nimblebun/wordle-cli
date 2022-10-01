/* -----------------------------------------------------------------------------
 * Copyright (c) Nimble Bun Works. All rights reserved.
 * This software is licensed under the MIT license.
 * See the LICENSE file for further information.
 * -------------------------------------------------------------------------- */

package words

import (
	"math"
	"math/rand"
	"time"
)

func getWordOfTheDay(epoch time.Time) (string, int) {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()

	currentDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	idx := int(math.Round(currentDay.Sub(epoch).Hours()/24)) % len(WordList)
	return WordList[idx], idx
}

func GetOfficialWordOfTheDay() (string, int) {
	wordleDay := time.Date(2021, time.Month(6), 19, 0, 0, 0, 0, time.Local)
	return getWordOfTheDay(wordleDay)
}

func GetWordOfTheDay() (string, int) {
	wordleCliDay := time.Date(2001, time.Month(3), 7, 0, 0, 0, 0, time.Local)
	return getWordOfTheDay(wordleCliDay)
}

func GetRandomWordle() (string, int) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(WordList))
	return WordList[idx], idx
}
