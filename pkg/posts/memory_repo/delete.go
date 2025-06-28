package posts_memory_repo

import (
	"github.com/google/uuid"
)

type ()

func (repo *PostsMemoryRepo) Delete(
	id uuid.UUID,
) error {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	_, ok := repo.Posts[id]
	if !ok {
		return &PostNotFoundError{PostID: id}
	}

	delete(repo.Posts, id)
	return nil
}
