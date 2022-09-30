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

const routerKeyBcWear = "/bcWear"

type BcWearRouter struct {
}

func (router BcWearRouter) get(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcWearService service.BcWearService
	byId := bcWearService.OneById(id)
	if byId.IsEmpty() {
		response.FailWithMessage(response.MessageNoDBRecords, c)
		return
	}
	query, err := serializer.SerializeBcWear(byId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(query, response.MessageOk, c)
	return
}

func (router BcWearRouter) list(c *gin.Context) {
	var bcWearParam serializer.BcWearParam
	var bcWear *model.BcWear
	err := c.ShouldBind(&bcWearParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcWearParam.BuildDBData(bcWear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcWearService service.BcWearService
	wears, count := bcWearService.List(bcWear)
	var resultList []*serializer.BcWearParam
	for _, bcWear := range wears {
		query, err := serializer.SerializeBcWear(bcWear)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		resultList = append(resultList, query)
	}
	response.OkWithDetailed(response.ListResult[*serializer.BcWearParam]{
		resultList,
		count,
	}, response.MessageOk, c)
}

func (router BcWearRouter) page(c *gin.Context) {
	var bcWearParam serializer.BcWearParam
	var wear *model.BcWear
	err := c.ShouldBind(&bcWearParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcWearParam.BuildDBData(wear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcWearService service.BcWearService
	page := serializer.CreatePage(bcWearParam.Page, bcWearParam.PageSize)
	wears, count := bcWearService.Page(page, wear)
	var resultList []*serializer.BcWearParam
	for _, bcWear := range wears {
		query, err := serializer.SerializeBcWear(bcWear)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		resultList = append(resultList, query)
	}
	response.OkWithDetailed(response.PageResult[*serializer.BcWearParam]{
		page,
		resultList,
		count,
	}, response.MessageOk, c)
}

func (router BcWearRouter) create(c *gin.Context) {
	var bcWearParam serializer.BcWearParam
	var wear *model.BcWear
	err := c.ShouldBindJSON(&bcWearParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcWearParam.BuildDBData(wear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var bcWearService service.BcWearService
	err = bcWearService.Create(wear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(wear.ID, c)
}

func (router BcWearRouter) update(c *gin.Context) {
	var bcWearParam serializer.BcWearParam
	var wear *model.BcWear
	err := c.ShouldBindJSON(&bcWearParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bcWearParam.BuildDBData(wear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if wear.ID == 0 {
		response.FailWithMessage("传入用户数据有误", c)
		return
	}
	var bcWearService service.BcWearService
	err = bcWearService.Modify(wear)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(wear.ID, c)
}

func (router BcWearRouter) remove(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var bcWearService service.BcWearService
	errRemove := bcWearService.Remove(id)
	if errRemove != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (router BcWearRouter) Register(engine *gin.Engine) {
	group := engine.Group(routerKeyBcWear)
	{
		group.GET("/page", router.page)
		group.GET("/", router.list)
		group.GET("/:id", router.get)
		group.POST("/", router.create)
		group.PUT("/", router.update)
		group.DELETE("/:id", router.remove)
	}
}
