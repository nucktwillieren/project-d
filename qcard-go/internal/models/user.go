package models

import (
	"time"
)

func init() {
}

type User struct {
	ID              uint      `json:"id"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Gender          string    `json:"gender"`
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
	Friends         []*Friend `pg:"rel:has-many, join_fk:user_one_id" json:"friends"`
	PairingID       uint      `json:"pairingId"`
	Pairing         *Friend   `pg:"rel:has-one" json:"pairing"`
}

type Friend struct {
	ID        uint      `json:"id"`
	Pair      bool      `pg:",use_zero" json:"pair"`
	UserOneID uint      `pg:"user_one_id" json:"userOneId"`
	UserTwoID uint      `pg:"user_two_id" json:"userTwoId"`
	CreatedAt time.Time `json:"createdAt"`
}
