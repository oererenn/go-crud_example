package rest

import (
	"comment-service/internal/service"
	"comment-service/pkg/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
)

type CommentController struct {
	service service.ICommentService
}

func NewCommentController(service service.ICommentService) *CommentController {
	return &CommentController{service: service}
}

// GetById godoc
// @Summary Get a comment by id
// @Description Get a comment by id
// @Tags comments
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommentDTO
// @Failure 500 {object} string
// @Router /comments/{id} [get]
func (controller *CommentController) GetById(c echo.Context) error {
	id := c.Param("id")
	res, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, res)
}

// Create godoc
// @Summary Create a comment
// @Description Create a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommentDTO
// @Failure 500 {object} string
// @Router /comments [post]
func (controller *CommentController) Create(c echo.Context) error {
	var request model.CreateCommentDTO
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(500, err)
	}
	err = controller.service.Create(&request)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "OK")
}

func (controller *CommentController) RegisterRoutes(e *echo.Echo) {
	e.GET("/comments/:id", controller.GetById)
	e.POST("/comments", controller.Create)
}
