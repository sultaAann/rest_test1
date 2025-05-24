package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type task struct {
	Id           int        `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Is_completed bool       `json:"is_completed"`
	Created_at   *time.Time `json:"created_at"`
	Updated_at   *time.Time `json:"updated_at"`
}

type handler struct {
	r *repository
}

func NewHandler(r *repository) *handler {
	return &handler{r: r}
}

func (h handler) GetAll(c *gin.Context) {
	res, err := h.r.GetAll()
	if err != nil {
		c.Error(err)
	}
	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}

// func Create(c *gin.Context) {
// 	c.
// }
