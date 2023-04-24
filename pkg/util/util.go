package util

import (
	"fmt"

	"github.com/softarch-project/menu-api/config"
)

func NewConnectionUrlBuilder(db config.Database) string {
	var url string
	url = fmt.Sprintf(
		"mongodb+srv://%s:%s@menuappcluster.kbempqg.mongodb.net/?retryWrites=true&w=majority",
		db.Username,
		db.Password,
	)

	return url
}
