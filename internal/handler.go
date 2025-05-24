package internal

import (
	"database/sql"
	"net/http"
	"strconv"
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
	c.IndentedJSON(http.StatusOK, res)
}

func (h handler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	res, err := h.r.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.Error(err)
		}
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (h handler) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}

	_, err = h.r.GetById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.Error(err)
		}
	}

	h.r.DeleteById(id)

	c.Status(http.StatusOK)
}

func (h handler) Create(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	id, err := h.r.Create(newTask)
	if err != nil {
		c.Error(err)
	}

	c.IndentedJSON(http.StatusOK, map[string]int{"id": id})
}

func (h handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
	}
	_, err = h.r.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, err)
		} else {
			c.Error(err)
		}
	}

	var updateTask task

	if err := c.Bind(&updateTask); err != nil {
		return
	}

	res, err := h.r.Update(id, updateTask)

	if err != nil {
		c.Error(err)
	}

	c.IndentedJSON(http.StatusOK, res)
}
