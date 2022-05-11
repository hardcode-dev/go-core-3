package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	game := New()
	game.Start()
}

const W int = 11 // Количество очков требуемое для победы в партии.

// Game - Партия.
type Game struct {
	ch     chan string
	wg     *sync.WaitGroup
	scores map[string]int
}

// New - Создает новую партию.
func New() *Game {
	return &Game{
		ch:     make(chan string),
		scores: make(map[string]int),
		wg:     new(sync.WaitGroup),
	}
}

// Start - запуск партии.
func (g *Game) Start() {
	fmt.Println("Партия началась")
	g.wg.Add(2)
	go g.player("Миша")
	go g.player("Игорь")
	g.ch <- "begin"
	g.wg.Wait()
	fmt.Println("Партия окончена! Счет:", g.scores)
}

func (g *Game) player(name string) {
	defer g.wg.Done()
	for val := range g.ch {
		if g.isWin() {
			close(g.ch)
			return
		}
		var turn string
		switch val {
		case "begin", "stop", "pong":
			turn = "ping"
		case "ping":
			turn = "pong"
		}
		fmt.Println(name, turn)
		if isGoal() {
			g.scores[name]++
			fmt.Println(name, "получил очко! Счет: ", g.scores)
			g.ch <- "stop"
		} else {
			g.ch <- turn
		}
	}
}

func (g *Game) isWin() bool {
	for _, score := range g.scores {
		if score == W {
			return true
		}
	}
	return false
}

func isGoal() bool {
	return rand.Intn(5) == 4 // 20% вероятность
}
