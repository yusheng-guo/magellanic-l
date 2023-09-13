package initialize

import (
	"crypto/tls"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
)

// InitTiDB initialize TiDB
func InitTiDB() {
	err := mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-west-2.prod.aws.tidbcloud.com",
	})
	if err != nil {
		log.Fatalln("register configs, err:", err)
	}

	db, err := sql.Open("mysql", "3udZ3FZAjbZEtk5.root:5D5y1IaUVQyht6dh@tcp(gateway01.us-west-2.prod.aws.tidbcloud.com:4000)/test?tls=tidb")
	if err != nil {
		log.Fatalln("connect database, err:", err)
	}

	task := global.NewDeferTask(func(a ...any) {
		var err error
		err = a[0].(*sql.DB).Close()
		if err != nil {
			log.Println("close TiDB client, err:", err)
		}
	}, db)
	global.DeferTaskQueue = append(global.DeferTaskQueue, task)

	var dbName string
	err = db.QueryRow("SELECT DATABASE();").Scan(&dbName)
	if err != nil {
		log.Fatalln("execute query, err:", err)
	}

	log.Println("You successfully connected to TiDB!")

	global.App.TiDB = db
}
