package initialize

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
	"google.golang.org/grpc"
	"log"
)

func InitDGraph() {
	// Create a connection to the Dgraph server
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial :9080, err:", err)
	}
	defer conn.Close()

	// Create a new Dgraph client
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// Create a new transaction
	txn := dgraphClient.NewTxn()
	defer txn.Discard(context.Background())

	// Run a query
	query := `{
        allPersons(func: has(name)) {
            name
        }
    }`
	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Fatalln("query, err:", err)
	}

	// Print the query result
	fmt.Println(string(resp.Json))
}
