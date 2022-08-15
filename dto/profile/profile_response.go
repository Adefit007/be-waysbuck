package profiledto

import "waysbuck/models"

type ProfileResponse struct {
		ID 			int							`json:"id" gorm:"primary_key:auto_increment"`
		Fullname	string						`json:"fullname" gorm:"type: varchar(255)"`
		Image		string						`json:"image" gorm:"type: varchar(255)"`
		UserID		int							`json:"user_id"`
		User		models.UserProfileResponse 	`json:"user"`
}