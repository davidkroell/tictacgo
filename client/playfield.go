package client

import (
	"github.com/davidkroell/tictacgo/models"
	"os"
	"text/template"
)

const playFieldTemplate = `+-------------+
|  {{ (index .Fields 0) | toString }} | {{ (index .Fields 1) | toString }} | {{ (index .Fields 2) | toString }}  |
| ---+---+--- |
|  {{ (index .Fields 3) | toString }} | {{ (index .Fields 4) | toString }} | {{ (index .Fields 5) | toString }}  |
| ---+---+--- |
|  {{ (index .Fields 6) | toString }} | {{ (index .Fields 7) | toString }} | {{ (index .Fields 8) | toString }}  |
+-------------+
`

func (c *Client) RenderPlayField(game models.Game, file *os.File) error {

	parsePlayerXorO := func(field models.Field) string {
		if field.OccupiedBy == nil {
			return " "
		}

		if field.OccupiedBy.Name == c.Username {
			return "X"
		} else {
			return "O"
		}
	}

	playfield, err := template.New("playfield").
		Funcs(template.FuncMap{"toString": parsePlayerXorO}).
		Parse(playFieldTemplate)

	if err != nil {
		return err
	}

	playfield.Execute(file, game) // write to pipe, for future reads
	return nil
}
