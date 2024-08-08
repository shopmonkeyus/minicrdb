package main

import (
	"database/sql"
	"fmt"

	"github.com/cockroachdb/cockroach-go/v2/testserver"
	"github.com/shopmonkeyus/go-common/sys"
)

const crdbVersion = "v23.1.10"

func main() {
	localities := testserver.LocalityFlagsOpt("region=gcp-us-central1", "region=gcp-us-east1", "region=gcp-us-west1")

	ts, err := testserver.NewTestServer(
		testserver.ThreeNodeOpt(),
		localities,
		testserver.AddListenAddrPortOpt(26257),
		testserver.AddListenAddrPortOpt(26258),
		testserver.AddListenAddrPortOpt(26259),
		testserver.CustomVersionOpt(crdbVersion),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ts.Stop()

	uri := ts.PGURL().String()
	fmt.Println(uri)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		fmt.Println(err)
		return
	}

	row, err := db.Query("select 1")
	if err != nil {
		fmt.Println(err)
		return
	}
	row.Close()

	<-sys.CreateShutdownChannel()

}
