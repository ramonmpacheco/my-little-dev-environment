package model

type ProductModel struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
