package router

import (
	"awesomeProject/model"
	"awesomeProject/model/response"
	"awesomeProject/serializer"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const routerKeyBcChat = "/bcChat"

type BcChatRouter struct {
}

func (router BcChatRouter) get(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcChatService service.BcChatService
	byId := bcChatService.OneById(id)
	if byId.IsEmpty() {
		response.FailWithMessage(response.MessageNoDBRecords, c)
		return
	}
	query, err := serializer.SerializeBcChat(byId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(query, response.MessageOk, c)
	return
}

func (router BcChatRouter) list(c *gin.Context) {
	var bcChatParam serializer.BcChatParam
	var bcChat *model.BcChat
	err := c.ShouldBind(&bcChatParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcChatParam.BuildDBData(bcChat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcChatService service.BcChatService
	chats, count := bcChatService.List(bcChat)
	var resultList []*serializer.BcChatParam
	for _, bcChat := range chats {
		query, err := serializer.SerializeBcChat(bcChat)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		resultList = append(resultList, query)
	}
	response.OkWithDetailed(response.ListResult[*serializer.BcChatParam]{
		resultList,
		count,
	}, response.MessageOk, c)
}

func (router BcChatRouter) page(c *gin.Context) {
	var bcChatParam serializer.BcChatParam
	var chat *model.BcChat
	err := c.ShouldBind(&bcChatParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcChatParam.BuildDBData(chat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcChatService service.BcChatService
	page := serializer.CreatePage(bcChatParam.Page, bcChatParam.PageSize)
	chats, count := bcChatService.Page(page, chat)
	var resultList []*serializer.BcChatParam
	for _, bcChat := range chats {
		query, err := serializer.SerializeBcChat(bcChat)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		resultList = append(resultList, query)
	}
	response.OkWithDetailed(response.PageResult[*serializer.BcChatParam]{
		page,
		resultList,
		count,
	}, response.MessageOk, c)
}

func (router BcChatRouter) create(c *gin.Context) {
	var bcChatParam serializer.BcChatParam
	var chat *model.BcChat
	err := c.ShouldBindJSON(&bcChatParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcChatParam.BuildDBData(chat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcChatService service.BcChatService
	err = bcChatService.Create(chat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(chat.ID, c)
}

func (router BcChatRouter) update(c *gin.Context) {
	var bcChatParam serializer.BcChatParam
	var chat *model.BcChat
	err := c.ShouldBindJSON(&bcChatParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcChatParam.BuildDBData(chat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if chat.ID == 0 {
		response.FailWithMessage("传入用户数据有误", c)
		return
	}
	var bcChatService service.BcChatService
	err = bcChatService.Modify(chat)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(chat.ID, c)
}

func (router BcChatRouter) remove(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var bcChatService service.BcChatService
	errRemove := bcChatService.Remove(id)
	if errRemove != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (router BcChatRouter) Register(engine *gin.Engine) {
	group := engine.Group(routerKeyBcChat)
	{
		group.GET("/page", router.page)
		group.GET("/", router.list)
		group.GET("/:id", router.get)
		group.POST("/", router.create)
		group.PUT("/", router.update)
		group.DELETE("/:id", router.remove)
	}
}
