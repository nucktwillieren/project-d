package db

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/go-pg/pg/v10"
	"gopkg.in/yaml.v3"
)

type YamlPGOptions struct {
	DB map[string]*pg.Options `yaml:"db"`
}

type PGLogger struct{}

type FailProcessor func(error)

func PGDefaultFailProcessor(err error) {
	log.Println(err.Error())
}

func (p PGLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (p PGLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	output, err := q.FormattedQuery()
	log.Println(string(output))
	return err
}

// CreateConnection is
func PGCreateConnection(options *pg.Options, failed FailProcessor) (newDB *pg.DB) {

	newDB = pg.Connect(options)

	pgLoggerImp := pg.QueryHook(PGLogger{})
	newDB.AddQueryHook(pgLoggerImp)

	if err := newDB.Ping(context.Background()); err != nil {
		log.Fatalln("PostgreSQL is down", err)
	}

	log.Println("Successfully created connection to database")

	return newDB
}

func InitWithPGOptionMap(schema map[string]*pg.Options, failed FailProcessor) map[string]*pg.DB {
	db := make(map[string]*pg.DB)
	for k, v := range schema {
		db[k] = PGCreateConnection(
			v,
			failed,
		)
	}
	return db
}

func YamlToPGOptions(path string) map[string]*pg.DB {
	var a YamlPGOptions
	s, err := ioutil.ReadFile(path)
	yaml.Unmarshal(s, &a)

	log.Println(a, err)
	return InitWithPGOptionMap(a.DB, PGDefaultFailProcessor)
}
