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

    u := User{
        FirstName: "jack",
        LastName: "yao",
    }

    appointments := []Appointment{
        Appointment{
            Subject: "first",
        },
        Appointment{
            Subject: "Second",
            //Attendees: []*User{&u},
        },
        Appointment{
            Subject: "Third",
        },
    }

    u.Appointments = appointments

    fmt.Println(u)

    db = db.Debug()
    //fmt.Println(db.NewRecord(&u))

    db.Create(&u)
    //db.Create(&u)

    u.LastName = "john"

    fmt.Println(u)
    //db.Debug().Save(&u)
    db.Model(&u).Update("first_name","joy")
    db.Model(&u).UpdateColumn("LastName","gates") // won't trigger the callback
    db.Model(&u).UpdateColumns(map[string]interface{}{
        "LastName": "bush",
        "FirstName":"fack",
    }) // won't trigger the callback
    db.Model(&u).Updates(map[string]interface{}{
        //"first_name": "jason",
        //"last_name": "bush",
        "LastName": "bush",
        "FirstName":"fack",
    })
    fmt.Println(u)



    db.Create(&User{
        FirstName:"First",
        LastName:"Name",
        Salary: 50000,
    })
    db.Create(&User{
        FirstName:"Second",
        LastName:"Name",
        Salary: 80000,
    })

    //Batch update
    //db.Debug().Table("users").Where("last_name = ?", "Name").
    db.Table("users").Where("salary > ?", 1000).
    Update("salary", gorm.Expr("salary + 100"))

    //Delete

    //db.Delete(&u)
    db.Where("last_name like ?", "Jimmy").Delete(&User{})


    //Transaction

    tx :=  db.Begin()
    newUser := User{
        FirstName:"jerry",
    }

    if err  = tx.Create(&newUser).Error; err != nil {
        tx.Rollback()
    }

    newUser.LastName = "marvin"

    if err = tx.Save(&newUser).Error; err != nil {
        tx.Rollback()
    }

    tx.Commit()

    fmt.Println("done")

}

type User struct {
	 gorm.Model
	FirstName  string
	LastName  string
    Salary uint
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
