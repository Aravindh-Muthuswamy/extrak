package configuration
import(
  "os"
  "fmt"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

func ConfigureDatabase(){
  fmt.Println("Configuring database")
  if _, err := os.Stat("./data/extrak.db"); err != nil {
    os.MkdirAll("./data", 0755)
    os.Create("./data/extrak.db")
    fmt.Println("database created")
    db, err := sql.Open("sqlite3", "./data/extrak.db")
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    _, err = db.Exec("CREATE TABLE `app_user` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `user_name` VARCHAR(64) NULL, `additional_details` VARCHAR(255) NULL")
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    db.Close()
  }else{
    fmt.Println("file already exists")
  }
}

func OpenDatabaseConnection() (*sql.DB){
  db, err := sql.Open("sqlite3", "./data/extrak.db")
  if(err != nil){
    fmt.Println(err)
  }
  return db
}
