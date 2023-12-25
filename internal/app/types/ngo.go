package types

type NGO struct {
	Name        string `bson:"name" json:"name"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Description string `bson:"description" json:"description"`
	Location    string `bson:"location" json:"location"`
	ImageUrl    string `bson:"image_url" json:"image_url"`
}

type NGOLogin struct {
	Email string `bson:"email" json:"email"`
}
