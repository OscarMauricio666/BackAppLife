package Connection

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

func connect() string {

	// cfg := mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "example",
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3306",
	// 	DBName:               "",
	// 	AllowNativePasswords: true,
	// }
	// var err error
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Conected!!!")
	return "Conectado!!!"
}
