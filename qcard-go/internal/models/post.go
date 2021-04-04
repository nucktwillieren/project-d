package models

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	orm.RegisterTable((*PostLike)(nil))
	orm.RegisterTable((*ReplyLike)(nil))
}

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	Rule      string    `json:"rule"`
	CreatedAt time.Time `json:"createdAt"`
}

type Post struct {
	ID          uint      `json:"id"`
	Creator     User      `pg:"rel:has-one" json:"creator"`
	Category    Category  `pg:"rel:has-one" json:"category"`
	Description string    `json:"description"`
	Like        []User    `pg:"many2many:post_likes,fk:post_id,join_fk:user_id" json:"like"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Reply struct {
	ID          uint      `json:"id"`
	CurrentPost Post      `json:"current_post"`
	Creator     User      `pg:"rel:has-one" json:"creator"`
	Description string    `json:"description"`
	Like        []User    `pg:"many2many:reply_likes,fk:reply_id,join_fk:user_id" json:"like"`
	CreatedAt   time.Time `json:"createdAt"`
}

type PostLike struct {
	ID     uint `json:"id"`
	PostID uint `json:"postId"`
	UserID uint `json:"userId"`
}

type ReplyLike struct {
	ID      uint `json:"id"`
	ReplyID uint `json:"replyId"`
	UserID  uint `json:"userId"`
}
