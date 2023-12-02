package main

import "testing"

func TestExampleGames(t *testing.T) {

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	games := []Game{}
	for _, line := range input {
		games = append(games, ParseGame(line))
	}

	idSumValidGames := 0
	for _, game := range games {
		if CheckValidGame(game) {
			idSumValidGames += game.id
		}
	}

	expected := 8

	if idSumValidGames != expected {
		t.Errorf("Unexpected result: %d", idSumValidGames)
	}
}

func TestExamplePowerSet(t *testing.T) {

	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	games := []Game{}
	for _, line := range input {
		games = append(games, ParseGame(line))
	}

	powerSum := 0

	for _, game := range games {
		powerSum = powerSum + CalcPowerSet(game)
	}

	expected := 2286

	if powerSum != expected {
		t.Errorf("Unexpected result: %d", powerSum)
	}
}
