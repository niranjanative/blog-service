package articles

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primary_key;autoIncrement:true"`
	Nickname  string    `json:"nickname" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp()"`
	Title     string    `json:"title" binding:"required"`
}
