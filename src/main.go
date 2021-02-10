package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
)

func main() {
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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

    db.AutoMigrate(&People{})

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
    ID uint
    Username string
    FirstName  string
    LastName string
}
