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
    //db.Find(&users,"username = ? ", "jack")
    //db.Where("username = ? ", "jack").Find(&users)
    //db.Where(&User{Username:"jack"}).Find(&users)
    //db.Where(map[string]interface{}{"username":"jack"}).Find(&users)
    //db.Where("username in (?)",[]string{"jack"}).Find(&users)
    //db.Where("username like ?", "%jack%").Find(&users)
    //db.Where("username like ? and first_name = ?", "%jack%", "jason").Find(&users)
    //db.Where("created_at < ?", time.Now()).Find(&users)
   // db.Where("created_at BETWEEN  ? AND ?", time.Now().Add(-30*24 * time.Hour), time.Now()).Find(&users)
   // db.Not("created_at BETWEEN  ? AND ?", time.Now().Add(-30*24 * time.Hour), time.Now()).Find(&users)
   // db.Not("created_at BETWEEN  ? AND ?", time.Now().Add(-30*24 * time.Hour), time.Now()).Or("Username = ?","jimmy").Find(&users)

   //Preload
    //db.Preload("Calendar.Appointments").Find(&users)
   // db.Limit(2).Order("first_name desc").Find(&users)
    //db.Limit(2).Offset(2).Order("first_name desc").Find(&users)
    //db.Preload("Calendar").Find(&users)

    //Select specific fields
    //db.Select([]string{"first_name", "last_name"}).Limit(2).Offset(2).Order("first_name desc").Find(&users)

    //Pluck
    usernames := []string{}

    db.Model(&User{}).Pluck("username", &usernames)

    userVMs := []UserViewModel{}
    db.Model(&User{}).Select([]string{"first_name", "last_name"}).Scan(&userVMs)


    for _, m :=  range userVMs {
        //fmt.Printf("\n%+v\n", r)
        //spew.Dump(r.Calendar)
        spew.Dump(m)
    }




    db.Find(&users)
    for _, r :=  range users {
        //fmt.Printf("\n%+v\n", r)
        //spew.Dump(r.Calendar)
        spew.Dump(r)
    }

    for _, u :=  range usernames {
        //fmt.Printf("\n%+v\n", r)
        //spew.Dump(r.Calendar)
        spew.Dump(u)
    }



    //Count

    var count int64
    db.Model(&User{}).Count(&count)
    fmt.Println(count)

    fmt.Println("done")

}

type User struct {
	 gorm.Model
	FirstName  string
	LastName  string
    Salary uint
    Username string
    Calendar Calendar
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
    OwnerID uint
    OwnerType string
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

type Calendar struct {
    gorm.Model
    Name string
    UserID uint
    Appointments []Appointment `gorm:"polymorphic:Owner"`
}


type  UserViewModel struct {
    FirstName string
    LastName string
}
