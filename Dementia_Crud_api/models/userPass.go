package models
import(	 
	 "time"
)

type User struct{
	Id       uint    //gorm primary key
	Email    string
	Password string
	Name     string
	CreatedAt    time.Time
    UpdatedAt    time.Time
}