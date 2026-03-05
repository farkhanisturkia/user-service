package models

import "time"

type User struct {
	Id        		uint      	`json:"id" gorm:"primaryKey"`
	Name      		string    	`json:"name"`
	Username  		string    	`json:"username" gorm:"unique;not null"`
	Email     		string    	`json:"email" gorm:"unique;not null"`
	Role      		string    	`json:"role" gorm:"default:'user';not null"`
	Password  		string    	`json:"password"`

	// CreatedCourses []Course     `gorm:"foreignKey:CreatorID"`
	// Enrollments    []UserCourse `gorm:"foreignKey:ParticipantID"`

	CreatedAt      time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}
