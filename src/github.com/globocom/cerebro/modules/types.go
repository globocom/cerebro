package modules

// Version route response
type Version struct {
	Version string `json:"version"`
}

// Healthcheck route response
type Healthcheck struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

// User struct
type User struct {
	Segments []string `json:"segments" bson:"segments"`
}

func (u User) addSegment(segment string) *User {
	u.Segments = append(u.Segments, segment)
	return &u
}
