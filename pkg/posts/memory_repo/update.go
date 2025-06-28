package posts_memory_repo

import (
	"time"

	"github.com/google/uuid"
	"github.com/maksimowich/redditclone/pkg/comments"
	"github.com/maksimowich/redditclone/pkg/users"
)

func (repo *PostsMemoryRepo) AddComment(
	id uuid.UUID,
	comment string,
	commentAuthor *users.User,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	post, ok := repo.Posts[id]
	if !ok {
		return nil, &PostNotFoundError{PostID: id}
	}

	newComment := &comments.Comment{
		Id:      uuid.New(),
		Body:    comment,
		Author:  commentAuthor,
		Created: time.Now().String(),
	}
	post.Comments = append(post.Comments, newComment)

	return post, nil
}

func (repo *PostsMemoryRepo) DeleteComment(
	id uuid.UUID,
	commentId uuid.UUID,
	deletingUser *users.User,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	post, ok := repo.Posts[id]
	if !ok {
		return nil, &PostNotFoundError{PostID: id}
	}

	for i, comment := range post.Comments {
		if comment.Id == commentId {
			if comment.Author == deletingUser {
				post.Comments = append(post.Comments[:i], post.Comments[i+1:]...)
				return post, nil
			} else {
				return nil, &comments.NotEnoughRightsToDeleteComment{
					CommentId: commentId,
					UserId:    deletingUser.Id,
				}
			}
		}
	}

	return nil, &comments.CommentNotFoundError{
		CommentId: commentId,
	}
}

func calculateUpvotePercentage(post *Post) float64 {
	if len(post.Votes) == 0 {
		return 100.0
	}
	if post.Score <= 0 {
		return 0.0
	}
	return float64(post.Score) / float64(len(post.Votes)) * 100.0
}

func getVoteByUser(
	post *Post,
	user *users.User,
) (*Vote, bool) {
	for _, vote := range post.Votes {
		if vote.User.Id == user.Id {
			return vote, true
		}
	}
	return nil, false
}

func deleteUserVoteIfExists(
	post *Post,
	user *users.User,
) {
	var indexToDelete = -1
	for index, vote := range post.Votes {
		if vote.User.Id == user.Id {
			indexToDelete = index
			break
		}
	}
	if indexToDelete != -1 {
		post.Votes = append(post.Votes[:indexToDelete], post.Votes[indexToDelete+1:]...)
	}
}

func (repo *PostsMemoryRepo) Upvote(
	id uuid.UUID,
	user *users.User,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()
	var ok bool

	post, ok := repo.Posts[id]
	if !ok {
		return nil, &PostNotFoundError{PostID: id}
	}

	existingVote, ok := getVoteByUser(post, user)
	if !ok {
		newVote := &Vote{
			User: user,
			Vote: 1,
		}
		post.Votes = append(post.Votes, newVote)
		post.Score += 1
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	} else if existingVote.Vote == 1 {
		deleteUserVoteIfExists(post, user)
		post.Score -= 1
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	} else { // if existingVote.Vote == -1
		newVote := &Vote{
			User: user,
			Vote: 1,
		}
		deleteUserVoteIfExists(post, user)
		post.Votes = append(post.Votes, newVote)
		post.Score += 2
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	}
}

func (repo *PostsMemoryRepo) Downvote(
	id uuid.UUID,
	user *users.User,
) (*Post, error) {
	repo.Mu.Lock()
	defer repo.Mu.Unlock()

	post, ok := repo.Posts[id]
	if !ok {
		return nil, &PostNotFoundError{PostID: id}
	}

	existingVote, ok := getVoteByUser(post, user)
	if !ok {
		newVote := &Vote{
			User: user,
			Vote: -1,
		}
		post.Votes = append(post.Votes, newVote)
		post.Score -= 1
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	} else if existingVote.Vote == -1 {
		deleteUserVoteIfExists(post, user)
		post.Score += 1
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	} else { // if existingVote.Vote == 1
		newVote := &Vote{
			User: user,
			Vote: -1,
		}
		deleteUserVoteIfExists(post, user)
		post.Votes = append(post.Votes, newVote)
		post.Score -= 2
		post.UpvotePercentage = calculateUpvotePercentage(post)
		return post, nil
	}
}
