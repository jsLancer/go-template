package controllers

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/app/ops/models"
	"go-template/internal/app/ops/services"
	"net/http"
	"strconv"
)

type DemoController struct {
	service services.DemoService
}

func NewDemoController(service services.DemoService) *DemoController {
	return &DemoController{
		service: service,
	}
}

func (d *DemoController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "parse id err:%v", err)
		return
	}
	demo, err := d.service.GetByID(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "call demoService.GetByID() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, demo)
}

func (d *DemoController) Create(c *gin.Context) {
	req := new(models.Demo)
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	err := d.service.CreateDemo(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "call demoService.Create() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (d *DemoController) Update(c *gin.Context) {
	req := new(models.Demo)
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, "bind params error:%v", err)
		return
	}
	err := d.service.Update(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "call demoService.Update() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (d *DemoController) FindAll(c *gin.Context) {
	demos, err := d.service.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "call demoService.FindAll() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, demos)
}

func (d *DemoController) DeleteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "parse id err:%v", err)
		return
	}
	err = d.service.DeleteByID(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "call demoService.DeleteByID() error:%v", err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
