package posts_memory_repo

import (
	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/users"
)

func (repo *PostsMemoryRepo) GetAll() ([]*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	posts := make([]*Post, 0, len(repo.Posts))
	for _, post := range repo.Posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostsMemoryRepo) GetByCategory(
	category string,
) ([]*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	res := []*Post{}
	for _, post := range repo.Posts {
		if post.Category == category {
			res = append(res, post)
		}
	}

	return res, nil
}

func (repo *PostsMemoryRepo) GetByUser(
	user *users.User,
) ([]*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	res := []*Post{}
	for _, post := range repo.Posts {
		if post.Author == user {
			res = append(res, post)
		}
	}

	return res, nil
}

func (repo *PostsMemoryRepo) GetById(
	postId uuid.UUID,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	post, ok := repo.Posts[postId]
	if !ok {
		return nil, &PostNotFoundError{PostID: postId}
	}
	post.Views += 1

	return post, nil
}
