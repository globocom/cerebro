package modules

// Version route response
type Version struct {
	Version string `json:"version"`
}

// Status struct
type Status struct {
	Status string `json:"status"`
	Err    string `json:"err,omitempty"`
}

// Attribute struct
type Attribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// User struct
type User struct {
	Segments []string `json:"segments" bson:"segments"`
}

func (u User) addSegment(segment string) *User {
	u.Segments = append(u.Segments, segment)
	return &u
}
