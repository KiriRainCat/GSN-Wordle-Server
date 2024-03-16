package controller

import (
	"gsn-wordle/internal/pkg/errs"
	"gsn-wordle/internal/pkg/util"
	"gsn-wordle/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Word = &WordController{s: service.Word}

type WordController struct {
	s *service.WordService
}

func (c *WordController) GetList(ctx *gin.Context) {
	list, err := c.s.GetList()
	if err != nil {
		util.InternalErrResponse(ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": list,
	})
}

func (c *WordController) GetById(ctx *gin.Context) {
	// REST 参数校验
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.ParamErrResponse(ctx)
		return
	}

	// 获取单词
	word, err := c.s.GetById(id)
	if err != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": word,
	})
}

func (c *WordController) GetWordOfTheDay(ctx *gin.Context) {
	word, err := c.s.GetWordOfTheDay()
	if err != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": word,
	})
}

func (c *WordController) GetRandomWord(ctx *gin.Context) {
	word, err := c.s.GetRandomWord()
	if err != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": word,
	})
}

func (c *WordController) Create(ctx *gin.Context) {
	// 参数模型
	type json struct {
		Word       string `json:"word" binding:"required"`
		Definition string `json:"definition" binding:"required"`
	}

	// 参数校验
	var data json
	if err := ctx.ShouldBindJSON(&data); err != nil {
		util.ParamErrResponse(ctx)
		return
	}

	// 创建单词
	if err := c.s.Create(data.Word, data.Definition); err != nil {
		if err == errs.ErrServer {
			util.InternalErrResponse(ctx)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": nil,
	})
}

func (c *WordController) Update(ctx *gin.Context) {
	// 参数模型
	type json struct {
		Word       string `json:"word" binding:"required"`
		Definition string `json:"definition" binding:"required"`
	}

	// REST 参数校验
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 参数校验
	var data json
	if err := ctx.ShouldBindJSON(&data); err != nil || id == 0 {
		util.ParamErrResponse(ctx)
		return
	}

	// 更新单词信息
	if err := c.s.Update(id, data.Word, data.Definition); err != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": nil,
	})
}

func (c *WordController) Delete(ctx *gin.Context) {
	// REST 参数校验
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.ParamErrResponse(ctx)
		return
	}

	// 删除单词
	if c.s.Delete(id) != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": nil,
	})
}
