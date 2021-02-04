package rest


// List ...
type List struct {
	ID 			int    `json:"id"`
	Title 		string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// UsersList ...
type UsersList struct {
	ID 			int
	UserID 		string
	ListID 		string
}

// Item ...
type Item struct {
	ID 			int    `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Done 		bool   `json:"done"`
}


// ListItem ...
type ListItem struct {
	ID	   int
	ListID string
	ItemID string
}
