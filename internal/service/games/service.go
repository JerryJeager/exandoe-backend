package games

import (
	"fmt"

	"github.com/JerryJeager/exandoe-backend/config"
	"github.com/JerryJeager/exandoe-backend/internal/models"
)

type GameSv interface {
	Play(mv *models.GameMove)
}

type GameServ struct {
	repo GameStore
}

func NewGameService(repo GameStore) *GameServ {
	return &GameServ{repo: repo}
}

func (s *GameServ) Play(mv *models.GameMove) {
	if mv.Index < 0 || mv.Index >= len(mv.Board1D) {
		return
	}
	updateBoard(mv)

	winner := checkWinner(mv)
	if winner != "" {
		mv.Status = fmt.Sprintf("%s wins", winner)
	} else if isBoardFull(mv) {
		mv.Status = "draw"
	}

	if mv.Turn == "x" {
		mv.Turn = "o"
	} else {
		mv.Turn = "x"
	}

	for i := range config.Games {
		if config.Games[i].Room == mv.Room {
			config.Games[i] = *mv
			break
		}
	}
}

func isBoardFull(mv *models.GameMove) bool {
	var count int = 0
	for _, v := range mv.Board1D {
		if v != "" {
			count++
		}
	}
	return count == 9
}

func updateBoard(mv *models.GameMove) {
	symbol := mv.Turn
	mv.Board1D[mv.Index] = symbol

	count := 1
	for i := 0; i < len(mv.Board3D); i++ {
		for j := 0; j < len(mv.Board3D); j++ {
			if count == mv.Index && mv.Board3D[i][j] == "" {
				mv.Board3D[i][j] = symbol
			}
			count++
		}
	}

}

func checkWinner(mv *models.GameMove) string {
	var winPatterns = [][]int{
		{0, 1, 2}, // top row
		{3, 4, 5}, // middle row
		{6, 7, 8}, // bottom row
		{0, 3, 6}, // left col
		{1, 4, 7}, // middle col
		{2, 5, 8}, // right col
		{0, 4, 8}, // diagonal TL-BR
		{2, 4, 6}, // diagonal TR-BL
	}

	for _, pattern := range winPatterns {
		a, b, c := pattern[0], pattern[1], pattern[2]
		if mv.Board1D[a] != "" && mv.Board1D[a] == mv.Board1D[b] && mv.Board1D[a] == mv.Board1D[c] {
			return mv.Board1D[a] // "x" or "o"
		}
	}

	return ""
}
