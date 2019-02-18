package pkg

// Talk  ...
type Talk struct {
	Name    string `json:"name" bson:"name"`
	Speaker string `json:"speaker" bson:"speaker"`
	Place   string `json:"place" bson:"place"`
}
