package models

type Game struct {
	Fields     []Field
	Player1    Player
	Player2    Player
	Winner     *Player
	NextTurn   *Player
	IsFinished bool
}

type Field struct {
	PositionX  int
	PositionY  int
	OccupiedBy *Player
}

func NewGame(player1 *Player, player2 *Player) Game {

	var instance = Game{
		Fields: []Field{
			{
				PositionX: 0,
				PositionY: 0,
			},
			{
				PositionX: 1,
				PositionY: 0,
			},
			{
				PositionX: 2,
				PositionY: 0,
			},
			{
				PositionX: 0,
				PositionY: 1,
			},
			{
				PositionX: 1,
				PositionY: 1,
			},
			{
				PositionX: 2,
				PositionY: 1,
			},
			{
				PositionX: 0,
				PositionY: 2,
			},
			{
				PositionX: 1,
				PositionY: 2,
			},
			{
				PositionX: 2,
				PositionY: 2,
			},
		},
		Player1:  *player1,
		Player2:  *player2,
		NextTurn: player1,
	}

	return instance
}

type gameError struct {
	message string
}

func (ge gameError) Error() string {
	return ge.message
}

func (g *Game) PlayTurn(p *Player, fieldId int) error {
	if g.IsFinished {
		return gameError{
			message: "Game already finished",
		}
	}

	f := &g.Fields[fieldId]

	if *g.NextTurn == *p {
		if g.Player1 == *p {
			g.NextTurn = &g.Player2
		}
		if g.Player2 == *p {
			g.NextTurn = &g.Player1
		}

		if f.OccupiedBy != nil {
			return gameError{
				message: "Field already occupied",
			}
		} else {
			f.OccupiedBy = p

			// check if Game is finished after each turn
			g.IsGameFinished()
			return nil
		}
	} else {
		return gameError{
			message: "Player not at turn",
		}
	}
}

func (g *Game) IsGameFinished() bool {

	// check horizontal rows
	for i := 0; i < 3; i++ {
		// check single vertical row
		if g.Fields[i].OccupiedBy != nil && g.Fields[i+1].OccupiedBy != nil && g.Fields[i+2].OccupiedBy != nil &&
			*g.Fields[i].OccupiedBy == *g.Fields[i+1].OccupiedBy && *g.Fields[i].OccupiedBy == *g.Fields[i+2].OccupiedBy {
			g.Winner = g.Fields[i].OccupiedBy
			g.IsFinished = true
			return true
		}
	}

	// check vertical columns
	for i := 0; i < 3; i++ {
		// check single vertical column
		if g.Fields[i].OccupiedBy != nil && g.Fields[i+3].OccupiedBy != nil && g.Fields[i+6].OccupiedBy != nil &&
			g.Fields[i].OccupiedBy == g.Fields[i+3].OccupiedBy && g.Fields[i].OccupiedBy == g.Fields[i+6].OccupiedBy {
			g.Winner = g.Fields[i].OccupiedBy
			g.IsFinished = true
			return true
		}
	}

	// check diagonals
	if g.Fields[0].OccupiedBy != nil && g.Fields[4].OccupiedBy != nil && g.Fields[8].OccupiedBy != nil &&
		g.Fields[0].OccupiedBy == g.Fields[4].OccupiedBy && g.Fields[0].OccupiedBy == g.Fields[8].OccupiedBy {
		g.Winner = g.Fields[4].OccupiedBy
		g.IsFinished = true
		return true
	}
	if g.Fields[2].OccupiedBy != nil && g.Fields[4].OccupiedBy != nil && g.Fields[6].OccupiedBy != nil &&
		g.Fields[2].OccupiedBy == g.Fields[4].OccupiedBy && g.Fields[2].OccupiedBy == g.Fields[6].OccupiedBy {
		g.Winner = g.Fields[4].OccupiedBy
		g.IsFinished = true
		return true
	}

	// check tie
	for i := 0; i < 9; i++ {
		if g.Fields[i].OccupiedBy != nil {
			continue
		}
		return false
	}

	return true
}
