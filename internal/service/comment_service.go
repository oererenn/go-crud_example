package service

import (
	"comment-service/internal/repository"
	"comment-service/pkg/model"
	"github.com/google/uuid"
	"time"
)

type ICommentService interface {
	GetById(id string) (*model.CommentDTO, error)
	Create(request *model.CreateCommentDTO) error
}

type CommentService struct {
	repository repository.IMongoDBRepository
}

func (c CommentService) Create(request *model.CreateCommentDTO) error {
	comment := model.Comment{
		ID:        uuid.New().String(),
		UserId:    request.UserId,
		Content:   request.Content,
		UserName:  request.UserName,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
	}
	return c.repository.Create(&comment)
}

func (c CommentService) GetById(id string) (*model.CommentDTO, error) {
	comment, err := c.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return convertToCommentDTO(comment), nil
}

func NewCommentService(repository repository.IMongoDBRepository) ICommentService {
	return &CommentService{repository: repository}
}

func convertToCommentDTO(comment *model.Comment) *model.CommentDTO {
	return &model.CommentDTO{
		ID:      comment.ID,
		UserId:  comment.UserId,
		Content: comment.Content,
	}
}

func convertToComment(comment *model.CommentDTO) *model.Comment {
	return &model.Comment{
		ID:      comment.ID,
		UserId:  comment.UserId,
		Content: comment.Content,
	}
}
