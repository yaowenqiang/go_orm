package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
    _ "log"
    "context"
    "time"
    _ "os"
	"gorm.io/gorm/logger"
    "github.com/davecgh/go-spew/spew"
    "errors"
)

func main() {
    /*
		newLogger := logger.New(
	  log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	  logger.Config{
		SlowThreshold: time.Second,   // 慢 SQL 阈值
		//LogLevel:      logger.Silent, // Log level
		LogLevel:      logger.Info, // Log level
		Colorful:      true,         // 禁用彩色打印
	  },
	)
    */
    newLogger := myLogger{}
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
    //db = db.Debug()
    //db.LogMode(true)
    //db.SetLogger(true)

    user := User{
        Username:"jackyao",
        FirstName: "jack",
        LastName: "yao",
        Calendar: Calendar{Name:"My Calendar",},
    }

    fmt.Println("Creating")
    db.Create(&user)

    user.Calendar.Name = "new Calendar"
    fmt.Println("Updating")
    db.Save(&user)

    fmt.Println("Deleting")


    //Scopes

    appts := []Appointment{}
    db.Scopes(LongMeetings).Find(&appts)

    spew.Dump(appts)
    db.Delete(&user)

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


func (u *User) BeforeSave(db *gorm.DB) error {
    fmt.Println("Before Save")
    return nil
    //return errors.New("We Can't create new user!")
}
func (u *User) BeforeCreate(db *gorm.DB) error {
    fmt.Println("Before Create")
    return nil
}
func (u *User) AfterSave(db *gorm.DB) error {
    fmt.Println("After Save")
    return nil
}
func (u *User) AfterCreate(db *gorm.DB) error {
    fmt.Println("After Create")
    return nil
}

func (c *Calendar) BeforeCreate(db *gorm.DB) error {
    fmt.Println("Before Create Calendar")

    //return nil

    return errors.New("can't create new calendar!")
}

func (c *Calendar) AfterCreate(db *gorm.DB) error {
    fmt.Println("After Create Calendar")

    //return nil

    return errors.New("can't create new calendar!")
}

func LongMeetings(db *gorm.DB) *gorm.DB {
    return db.Where("length > ?", 60)
}

type myLogger struct {
}

func (ml *myLogger) jogMode(level logger.LogLevel) Logger.Interface {
    newLogger := myLogger{}
    return &newLogger

}

func (ml *myLogger) Info(ctx context.Context, msg string, data ...interface{}) {
    fmt.Printf("%s\n", msg)
}
func (ml *myLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
    fmt.Printf("%s\n", msg)
}
func (ml *myLogger) Error(ctx context.Context, msg string, data ...interface{}) {
    fmt.Printf("%s\n", msg)
}
func (ml *myLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    fmt.Printf("trace\n")
}
