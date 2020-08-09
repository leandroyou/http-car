package dao

type Car struct {
	Id    string `bson:"_id,omitempty"`
	Make  string `bson:"Make"`
	Model string `bson:"Model"`
}
