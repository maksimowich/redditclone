package posts_memory_repo

import (
	"time"

	"github.com/google/uuid"
)

func (repo *PostsMemoryRepo) Add(
	postAdd *PostAdd,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	author, err := repo.UsersRepo.GetById(postAdd.AuthorId)
	if err != nil {
		return nil, err
	}

	postId := uuid.New()
	newPost := &Post{
		Id:               postId,
		Title:            postAdd.Title,
		Type:             postAdd.Type,
		Category:         postAdd.Category,
		Text:             postAdd.Text,
		Url:              postAdd.Url,
		Author:           author,
		Score:            0,
		Views:            0,
		Votes:            []*Vote{},
		UpvotePercentage: 100.0,
		Comments:         []*Comment{},
		Created:          time.Now().String(),
	}

	repo.Posts[postId] = newPost

	return newPost, nil
}
