package models

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

func init() {
	orm.RegisterTable((*UserFriend)(nil))
}

type User struct {
	ID              uint
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
	ID      uint
	Pair    bool
	User1ID uint `pg:"user1_id" json:"user1_id"`
	User2ID uint `pg:"user2_id" json:"user2_id"`
}

type UserFriend struct {
	ID       uint
	UserID   uint
	FriendID uint
}
