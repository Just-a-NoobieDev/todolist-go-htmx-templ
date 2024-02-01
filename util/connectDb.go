package util


import (
	"github.com/eaigner/jet"
  "github.com/lib/pq"
	"github.com/labstack/gommon/log"
)

// Log error
func logFatal(err error) {
	if err != nil {
			log.Fatal(err)
	}
}

func ConnectToDB() *jet.Db {
	pgUrl, err := pq.ParseURL("postgres://iemsowpq:X1G2a23gqqPkByfyXsR1bZj4wJX-Cvq2@rosie.db.elephantsql.com/iemsowpq")
	logFatal(err)

	db, err := jet.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}