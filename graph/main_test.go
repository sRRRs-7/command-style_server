package graph

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sRRRs-7/loose_style.git/cfg"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/dataloaders"
	"github.com/sRRRs-7/loose_style.git/token"
)

var testDB *pgxpool.Pool
var resolver *Resolver
var tokenMaker token.Maker

func TestMain(m *testing.M) {
	config, err := cfg.LoadConfig("../")
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

	store := db.NewStore(testDB)
	// all instance connect
	resolver, tokenMaker, err = NewResolver(config, store, testDB)
	if err != nil {
		log.Fatalf("new resolver error: %v", err)
	}

	os.Exit(m.Run())
}

func GinTestRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(GinContextToContextCookie(tokenMaker))
	r.Use(dataloaders.DataLoaderMiddleware(resolver.store))
	r.POST("/query", graphqlHandler(resolver))
	r.POST("/admin/query", graphqlHandler(resolver))
	return r
}
