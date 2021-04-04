package models

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	orm.RegisterTable((*UserFriend)(nil))
	orm.RegisterTable((*PostLike)(nil))
	orm.RegisterTable((*ReplyLike)(nil))
}

type User struct {
	ID              uint      `json:"id"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Photo           string    `json:"photo"`
	Birthday        time.Time `json:"birthday"`
	Relationship    string    `json:"relationship"`
	Interest        string    `json:"interest"`
	Club            string    `json:"club"`
	FavoriteCourse  string    `json:"favorite_course"`
	FavoriteCountry string    `json:"favorite_country"`
	Trouble         string    `json:"trouble"`
	Exchange        string    `json:"exchange"`
	Trying          string    `json:"trying"`
	Friends         []Friend  `pg:"many2many:user_friends,fk:user_id,join_fk:friend_id" json:"friends"`
	Pairing         Friend    `pg:"rel:has-one" json:"pairing"`
}

type Friend struct {
	ID      uint `json:"id"`
	Pair    bool
	User1ID uint `pg:"user1_id" json:"user1_id"`
	User2ID uint `pg:"user2_id" json:"user2_id"`
}

type UserFriend struct {
	ID       uint `json:"id"`
	UserID   uint `json:"user_id"`
	FriendID uint `json:"friend_id"`
}

type Category struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Rule  string `json:"rule"`
}

type Post struct {
	ID          uint     `json:"id"`
	Creator     User     `pg:"rel:has-one" json:"creator"`
	Category    Category `pg:"rel:has-one" json:"category"`
	Description string   `json:"description"`
	Like        []User   `pg:"many2many:post_likes,fk:post_id,join_fk:user_id" json:"like"`
}

type Reply struct {
	ID          uint   `json:"id"`
	CurrentPost Post   `json:"current_post"`
	Creator     User   `pg:"rel:has-one" json:"creator"`
	Description string `json:"description"`
	Like        []User `pg:"many2many:reply_likes,fk:reply_id,join_fk:user_id" json:"like"`
}

type PostLike struct {
	ID     uint `json:"id"`
	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}

type ReplyLike struct {
	ID      uint `json:"id"`
	ReplyID uint `json:"reply_id"`
	UserID  uint `json:"user_id"`
}
