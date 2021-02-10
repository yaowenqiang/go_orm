package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
    "log"
    "time"
    "os"
	"gorm.io/gorm/logger"
)

func main() {
		newLogger := logger.New(
	  log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	  logger.Config{
		SlowThreshold: time.Second,   // 慢 SQL 阈值
		LogLevel:      logger.Silent, // Log level
		Colorful:      false,         // 禁用彩色打印
	  },
	)
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

    if err != nil {
        panic(err.Error())
    }

    dbase, err := db.DB()
    if err != nil {
        panic(err.Error())
    }

    defer dbase.Close()

    err = dbase.Ping()

    if err != nil {
        panic(err.Error())
    }

    //println("Connection to database established")

    //db.DropTable(&People{})
    //db.Singular(true)
    db.Debug().AutoMigrate(&People{})
	//db.Model(&People{}).AddIndex("idx_first_name", "first_name")
	//db.Model(&People{}).RemoveIndex("idx_first_name")
	//db.Model(&People{}).AddUniqueIndex("idx_last_name", "last_name")

	/*
	for _, f := range db.NewScope(&People{}).Fields() {
		fmt.Println(f.Name)
	}
	*/

    user := People{
        Username: "jack",
        FirstName: "jack",
        LastName: "yao",
    }

    db.Create(&user)


    var users = []People{
        People{
            Username: "Adent",
            FirstName: "Arthur",
            LastName: "Dent",
        },
        People{
            Username: "fprefect",
            FirstName: "Ford",
            LastName: "Prefect",
        },
        People{
            Username: "tmacmillan",
            FirstName: "Tricia",
            LastName: "MacMillan",
        },
        People{
            Username: "mrobot",
            FirstName: "Marvin",
            LastName: "Robot",
        },
    }

    for _, u := range users{
        db.Create(&u)
    }

    uu  := People{}
    //db.First(&uu)
    db.Last(&uu)
    fmt.Println(uu)
    //println("done")

    //fmt.Println(user)
    updateUser := People{
        Username: "jack",
    }

    db.Where(&updateUser).First(&updateUser)

    fmt.Println(updateUser)

    updateUser.LastName = "LastMan"

    db.Save(&updateUser)


    updatedUser := People{}
    db.Where(&updateUser).First(&updatedUser)
    fmt.Println(updatedUser)


    db.Where(&People{Username:"jack"}).Delete(&People{})

    fmt.Println("done")
}

type People struct {
    //ID uint
	Model gorm.Model `gorm:"embedded"`
	UserID int  `gorm:"primaryKey"`
    Username string `gorm:"comment:用户名;size:15;not null"`
	FirstName  string `gorm:"size:15;not null;column:FirstName"`
	LastName string `gorm:"unique;uniqueIndex;not null;column:LastName;default:smith"`
    Count int `gorm:"autoIncrement"`
    TempField bool `gorm:"-"`
}

func (u People) TableName() string {
    return "stackholders"
}


