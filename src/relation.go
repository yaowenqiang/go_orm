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
    //db.Debug().AutoMigrate(&User{})
    //db.Debug().AutoMigrate(&Calendar{})

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Calendar{})

    user := User{
        Username: "jack",
        FirstName: "jack",
        LastName: "yao",
        Calendar: Calendar{
            Name: "Improbable Events",
        },
    }

    db.Create(&user)

    u := User{}
    c := Calendar{}

    db.First(&u)

    db.Debug().Model(&u).Association("Calendar").Find(&c)

    fmt.Println(u)
    fmt.Println(c)



    fmt.Println("done")
}

type User struct {
	 gorm.Model
    Username string `gorm:"comment:用户名;size:15;not null"`
	FirstName  string `gorm:"size:15;not null;column:FirstName"`
	LastName string `gorm:"unique;uniqueIndex;not null;column:LastName;default:smith"`
    Calendar Calendar `gorm:"foreignKey:UserID"`
}


type Calendar struct {
    gorm.Model
    Name string
    UserID uint
}


