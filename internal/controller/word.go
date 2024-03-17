package controller

import (
	"gsn-wordle/internal/pkg/config"
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
		Subject    string `json:"subject"`
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
	id, err := c.s.Create(data.Subject, data.Word, data.Definition)
	if err != nil {
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

	// 如果为管理员添加单词，直接设置为激活状态
	if ctx.Request.Header.Get("Admin-Auth") == config.Config.Server.AdminPassword {
		if err := c.s.UpdateActiveState(id, true); err != nil {
			util.InternalErrResponse(ctx)
			return
		}
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
		Subject    string `json:"subject"`
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

	// 判断是否为管理员操作
	if ctx.Request.Header.Get("Admin-Auth") == config.Config.Server.AdminPassword {
		// 更新单词信息
		if err := c.s.Update(id, data.Subject, data.Word, data.Definition); err != nil {
			util.InternalErrResponse(ctx)
			return
		}
	} else {
		// 提交单词更新申请
		if err := c.s.CommitUpdate(id, data.Subject, data.Word, data.Definition); err != nil {
			util.InternalErrResponse(ctx)
			return
		}
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

//* ------------------------------ Admin APIs ------------------------------ *//

func (c *WordController) SetActiveState(ctx *gin.Context) {
	// REST 参数校验
	id, _ := strconv.Atoi(ctx.Param("id"))
	active, err := strconv.ParseBool(ctx.Param("active"))
	if err != nil || id == 0 {
		util.ParamErrResponse(ctx)
		return
	}

	if c.s.UpdateActiveState(id, active) != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": nil,
	})
}

func (c *WordController) GetCommits(ctx *gin.Context) {
	commits, err := c.s.GetCommits()
	if err != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": commits,
	})
}

func (c *WordController) ApproveCommit(ctx *gin.Context) {
	// REST 参数校验
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		util.ParamErrResponse(ctx)
		return
	}

	// 审核提交
	if c.s.ApproveCommit(id) != nil {
		util.InternalErrResponse(ctx)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": nil,
	})
}
