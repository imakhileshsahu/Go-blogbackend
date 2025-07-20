// Blog represents a blog post with associated metadata.
// It includes fields for the blog's ID, title, description, image URL, and the ID of the user who created it.
// The User field establishes a relationship to the User model via the UserID foreign key.
package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID string `json:"userid"`
	User   User   `json:"user" ;gorm:"foreignkey:UserID"`
}
