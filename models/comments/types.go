package comments

import "time"

type Comment struct {
	ID         uint      `json:"id" gorm:"primary_key;autoIncrement:true"`
	ParentID   uint      `json:"parent_id" binding:"required"`
	ParentType string    `json:"parent_type" binding:"required"`
	Nickname   string    `json:"nickname" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp()"`
}
