package db
import(
	"database/sql"
	"fmt"
	"github.com/miraikeitai2020/backend-summer-vacation/pkg/config"
	"log"
	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)
var Conn *sql.DB

// Driver名
const driverName = "mysql"

func init(){
	CreateConnection()
}

func CreateConnection(){
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	var err error
	Conn, err = sql.Open(driverName,
		config.GetConnectionToken())
	if err != nil {
		log.Fatal(err)
		fmt.Println("失敗")
	}
}