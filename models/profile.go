package models

import "time"

type Profile struct {
	ID 			int						`json:"id" gorm:"primary_key:auto_increment"`
	Fullname 	string					`json:"fullname" gorm:"type: varchar(255)"`
	Image		string					`json:"image" gorm:"type: varchar(255)"`
	UserID		int						`json:"user_id"`
	User		UserProfileResponse 	`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt	time.Time				`json:"-"`
	UpdatedAt 	time.Time 				`json:"-"`
	
}

type ProfileResponse struct {
	Fullname	string	`json:"fullname"`
	Image		string	`json:"image"`
	UserID		int		`json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}