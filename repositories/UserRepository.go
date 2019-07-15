package repositories

import (
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" `
	Email    string        `bson:"email" json:"email"`
	Username string        `bson:"username" json:"username"`
	Password string        `bson:"password" json:"password"`
	Token    string        `bson:"token" json:"token"`
}

//Login ...
func (r User) Login() (*User, error) {
	err := db.C("User").Find(bson.M{"$or": []bson.M{{"username": r.Username, "password": r.Password}, {"email": r.Username, "password": r.Password}}}).One(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
