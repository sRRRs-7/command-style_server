package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sRRRs-7/loose_style.git/cfg"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := cfg.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	cfg, err := pgxpool.ParseConfig(config.DBsource)
	if err != nil {
		log.Fatalf("pgx configuration error: %v", err)
	}

	testDB, err = pgxpool.ConnectConfig(context.TODO(), cfg)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
