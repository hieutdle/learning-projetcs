package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Player struct {
	Name     string
	Strength int
	Gender   string
}

type Team struct {
	Player1 Player
	Player2 Player
	Sum     int
}

func main() {
	players, err := readPlayers("players.csv")
	if err != nil {
		log.Fatal(err)
	}

	teams := createBalancedTeams(players)
	printFirstRound(teams)
}

func readPlayers(filename string) ([]Player, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var players []Player
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		strength, _ := strconv.Atoi(record[1])
		players = append(players, Player{Name: record[0], Strength: strength, Gender: record[2]})
	}

	return players, nil
}

func createBalancedTeams(players []Player) []Team {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Strength < players[j].Strength
	})

	var teams []Team
	n := len(players)
	for i := 0; i < n/2; i++ {
		player1 := players[i]
		player2 := players[n-i-1]

		team := Team{
			Player1: player1,
			Player2: player2,
			Sum:     player1.Strength + player2.Strength,
		}

		teams = append(teams, team)
	}

	return teams
}

func printFirstRound(teams []Team) {
	fmt.Println("First Round")
	for i := 0; i < len(teams); i += 2 {
		team1 := teams[i]
		if i+1 < len(teams) {
			team2 := teams[i+1]
			fmt.Printf("Match %d: [%s and %s (Strength %d)] vs [%s and %s (Strength %d)]\n", i/2+1,
				team1.Player1.Name, team1.Player2.Name, team1.Sum,
				team2.Player1.Name, team2.Player2.Name, team2.Sum)
		} else {
			fmt.Printf("Match %d: [%s and %s (Strength %d)] has a bye\n", i/2+1,
				team1.Player1.Name, team1.Player2.Name, team1.Sum)
		}
	}
}
