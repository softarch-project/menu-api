package util

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/config"
)

func NewConnectionUrlBuilder(stuff string, db config.Database) string {
	var url string
	switch stuff {
	case "mysql":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			db.Username,
			db.Password,
			db.Hostname,
			db.Port,
			db.Database,
		)
	case "dns":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?tls=true&parseTime=true",
			db.Username,
			db.Password,
			db.Hostname,
			db.Database,
		)
	default:
		log.Errorf("error, unknow the stuff: %s", stuff)
	}

	return url
}
