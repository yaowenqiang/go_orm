package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
    "log"
    "time"
    "os"
	"gorm.io/gorm/logger"
    "github.com/davecgh/go-spew/spew"
)

func main() {
		newLogger := logger.New(
	  log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	  logger.Config{
		SlowThreshold: time.Second,   // 慢 SQL 阈值
		LogLevel:      logger.Silent, // Log level
		Colorful:      true,         // 禁用彩色打印
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

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Appointment{})
    u := User{}
    fmt.Println(u)
    db = db.Debug()

    //db.First(&u)
    /*db.FirstOrInit(&u, &User{
        FirstName: "Trump",
    })
    */
    db.FirstOrCreate(&u, &User{
        FirstName: "Trump",
        Username: "tower",

    })

    last := User{}
    db.Last(&last)
    fmt.Println(last)


    //find all records
    users := []User{}
    //db.Find(&users,&User{Username:"jack"})
    //db.Find(&users,&User{Username:"jack"})
    //db.Find(&users,map[string]interface{}{"username":"jack"})
    db.Find(&users,"username = ? ", "jack")

    for _, r :=  range users {
        //fmt.Printf("\n%+v\n", r)
        spew.Dump(r)
    }


    fmt.Println("done")

}

type User struct {
	 gorm.Model
	FirstName  string
	LastName  string
    Salary uint
    Username string
    Appointments []Appointment `gorm:"foreignKey:UserID"`
}


type Appointment struct {
    gorm.Model
    StartTime *time.Time
    Duration uint
    UserID uint
    //Attendees []*User
    Subject string
    Description string
    Length uint
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
    println("Before Update")
    return nil
}

func (u *User) AfterUpdate(db *gorm.DB) error {
    println("After Update")
    return nil
}

