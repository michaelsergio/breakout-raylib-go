package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func ReadMaxScoreFile(path string) (SavedGames, error) {
	savedGames := SavedGames{}
	data, err := os.ReadFile(path)
	if err != nil {
		return savedGames, err
	}

	var contents = string(data[:])
	nlSplit := strings.Split(contents, "\n")
	if len(nlSplit) == 0 {
		return savedGames, errors.New("Count not read lines of file")
	}
	if strings.HasPrefix(nlSplit[0], "MAX_SCORE=") {
		var eqSplit = strings.SplitAfterN(nlSplit[0], "=", 2)
		if len(eqSplit) == 0 {
			return savedGames, errors.New("Weirdly formatted MAX_SCORE")
		}
		var scoreStr = eqSplit[1]
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			return savedGames, err
		}
		savedGames.MaxScore = score
	}
	return savedGames, nil
}

func WriteMaxScoreFile(path string, save SavedGames) {
	var content = fmt.Sprintf("MAX_SCORE=%d", save.MaxScore)
	os.WriteFile(path, []byte(content), 0644)
}

