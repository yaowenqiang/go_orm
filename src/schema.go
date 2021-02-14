package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    _ "fmt"
    "log"
    "time"
    "os"
	"gorm.io/gorm/logger"
    _ "github.com/davecgh/go-spew/spew"
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
    db = db.Debug()

    db.AutoMigrate(&User{}, &Calendar{}, &Appointment{}, &Attachment{})
    //db.DropTableIfExists("temptabl3")
    db.Migrator().DropTable("temptabl3")


    //db.Model(&User{}).ModifyColumn("first_name", "VARCHAR(100)")
    //Alter table
    // just works for postgres
    db.Exec(`ALTER TABLE users ALTER "first_name" TYPE VARCHAR(100)`)
    //db.Migrator().AlterColumn(&User{}, "last_name TYPE VARCHAR(100)")
    db.Migrator().DropColumn(&User{}, "tmp")

    //db.DroptableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
    //db.CreateTable(&User{}, &Calendar{}, &Appointment{})

}

type User struct {
	 gorm.Model
    Username string `gorm:"comment:用户名;size:15;not null"`
	//FirstName  string `gorm:"size:15;not null;column:FirstName"`
	FirstName  string `gorm:"size:14"`
	//LastName string `gorm:"unique;uniqueIndex;not null;column:LastName;default:smith"`
	//LastName string `gorm:"not null;column:LastName;default:smith"`
    LastName string `gorm:"not null;default:smith;size:1"`
    Calendar Calendar `gorm:"foreignKey:UserID"`
    //Calendar Calendar
    //CalendarID uint `gorm:"foreignKey:CalendarID"`
}


type Calendar struct {
    gorm.Model
    Name string
    UserID uint
    Appointments []Appointment `gorm:"polymorphic:Owner"`
}



type Appointment struct {
    gorm.Model
    Subject string
    Description string
    StartTime time.Time
    Length uint
    //CalendarID uint
    Recurring bool
    RecurrencePattern string
    OwnerID uint
    OwnerType string
    Attendees []User  `gorm:"many2many:appointment_user"`
    Attachments []Attachment `gorm:"foreignKey:AppointmentID"`
}

type TaskList struct {
    gorm.Model
    Appointments []Appointment `gorm:"polymorphic:Owner"`
}

type Attachment struct {
    gorm.Model
    Date []byte
    AppointmentID uint `gorm:"index:idx_appointment_id"`
}
