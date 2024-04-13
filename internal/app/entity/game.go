package entity

type Game struct {
	ID      int
	kills   []KillEvent
	players map[string]Player
}

func NewGame(id int) *Game {
	return &Game{
		ID:      id,
		kills:   make([]KillEvent, 0),
		players: make(map[string]Player, 0),
	}
}

func (g *Game) AddPlayer(player Player) {
	if _, ok := g.players[player.id]; !ok {
		g.players[player.id] = player
	}
}

func (g *Game) ListPlayers() []Player {
	players := make([]Player, 0)

	for _, player := range g.players {
		players = append(players, player)
	}

	return players
}

func (g *Game) AddKill(kill KillEvent) {
	g.kills = append(g.kills, kill)
}

func (g *Game) ListKills() []KillEvent {
	kills := append([]KillEvent(nil), g.kills...)
	return kills
}
