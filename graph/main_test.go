package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sRRRs-7/loose_style.git/cfg"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
	"github.com/sRRRs-7/loose_style.git/graph/generated"
	"github.com/sRRRs-7/loose_style.git/token"
	"github.com/sRRRs-7/loose_style.git/utils"
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

func NewServer() {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))
	h.AddTransport(transport.POST{})
}

func NewRequest(t *testing.T, q struct{ Query string }, endpoint string) (arr []string, list map[string]string, result []byte) {
	fmt.Println(q)

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, endpoint, &body)
	if err != nil {
		t.Fatal("error new request", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Cookie", resolver.config.AdminCookieKey)
	req.Header.Add("Cookie", "token")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("error request", err)
	}
	defer res.Body.Close()

	result, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("error read body", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("error request code:", res.StatusCode, string(result))
	}

	list = make(map[string]string)
	arr = utils.RegexpArray("[:,]", string(result))
	for i := 1; i < len(arr); i = i + 2 {
		key := utils.RetrieveRegexp("\".*\"", arr[i-1])
		value := utils.RetrieveRegexp("\".*\"", arr[i])
		list = utils.CreateMap(key, value, list)
	}

	return arr, list, result
}
