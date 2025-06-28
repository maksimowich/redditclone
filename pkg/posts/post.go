package posts

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/comments"
	"github.com/maksimowich/redditclone/pkg/users"
)

type Vote struct {
	User *users.User
	Vote int8
}

type Post struct {
	Id uuid.UUID

	Title    string
	Type     string
	Category string
	Text     string
	Url      string
	Author   *users.User

	Score            uint32
	Views            uint32
	Votes            []*Vote
	UpvotePercentage float64

	Comments []*comments.Comment

	Created string
}

func (p *Post) String() string {
	return fmt.Sprintf(
		"Post{id=%s, category=%s, title=%s}",
		p.Id, p.Category, p.Title,
	)
}

type PostAdd struct {
	Title    string
	Type     string
	Category string
	Text     string
	Url      string
	AuthorId uuid.UUID
}

type VoteResponse struct {
	UserId uuid.UUID `json:"user"`
	Vote   int8      `json:"vote"`
}

type PostResponse struct {
	Id               uuid.UUID           `json:"id"`
	Title            string              `json:"title"`
	Type             string              `json:"type"`
	Category         string              `json:"category"`
	Text             string              `json:"text,omitempty"`
	Url              string              `json:"url,omitempty"`
	Author           *users.User         `json:"author"`
	Score            uint32              `json:"score"`
	Views            uint32              `json:"views"`
	Votes            []*VoteResponse     `json:"votes"`
	UpvotePercentage float64             `json:"upvotePercentage"`
	Comments         []*comments.Comment `json:"comments"`
	Created          string              `json:"created"`
}

func (p *Post) ToResponse() *PostResponse {
	votes := make([]*VoteResponse, 0, len(p.Votes))
	for _, vote := range p.Votes {
		votes = append(votes, &VoteResponse{
			UserId: vote.User.Id,
			Vote:   vote.Vote,
		})
	}

	return &PostResponse{
		Id:               p.Id,
		Title:            p.Title,
		Type:             p.Type,
		Category:         p.Category,
		Text:             p.Text,
		Url:              p.Url,
		Author:           p.Author,
		Score:            p.Score,
		Views:            p.Views,
		Votes:            votes,
		UpvotePercentage: p.UpvotePercentage,
		Comments:         p.Comments,
		Created:          p.Created,
	}
}

type PostsRepoInterface interface {
	// Create
	Add(postAdd *PostAdd) (*Post, error)

	// Read
	GetAll() ([]*Post, error)
	GetByCategory(category string) ([]*Post, error)
	GetByUser(user *users.User) ([]*Post, error)
	GetById(id uuid.UUID) (*Post, error)

	// Update
	AddComment(id uuid.UUID, comment string, commentAuthor *users.User) (*Post, error)
	DeleteComment(id uuid.UUID, commentId uuid.UUID, deletingUser *users.User) (*Post, error)
	Upvote(id uuid.UUID, user *users.User) (*Post, error)
	Downvote(id uuid.UUID, user *users.User) (*Post, error)

	// Delete
	Delete(id uuid.UUID) error
}
