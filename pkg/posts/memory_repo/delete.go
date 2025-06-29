package posts_memory_repo

import (
	"github.com/google/uuid"
)

func (repo *PostsMemoryRepo) Delete(
	postId uuid.UUID,
	userId uuid.UUID,
) (uuid.UUID, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	post, ok := repo.Posts[postId]
	if !ok {
		return uuid.Nil, &PostNotFoundError{PostID: postId}
	}
	if post.Author.Id != userId {
		return uuid.Nil, &UserHasNotEnoughRights{UserId: userId}
	}

	delete(repo.Posts, postId)
	return post.Id, nil
}
