package controller

import (
	"PostgreService/entity"
	"PostgreService/service"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Summary 按照lastappeared_id查询轨迹
// @Tags Trajectory_轨迹属性控制器
// @Value 轨迹控制器
// @Accept application/json
// @Param lastappeared_id path string true "轨迹Id"
// @Router /trajectory/objecttrajactory/{lastappeared_id} [get]
func ObjectTrajactoryApi(ctx *gin.Context) {
	id := ctx.Param("lastappeared_id")
	res := service.ObjectTrajactoryApi(id)
	ctx.JSON(200, res)
}

// @Summary 按照lastappeared_id,machinetype 查询机型
// @Tags Trajectory_轨迹属性控制器
// @Value 轨迹控制器
// @Accept application/json
// @Param lastappeared_ids query []string false "轨迹Id"
// @Param machinetype query string false "机型"
// @Router /trajectory/machineTypeSearch [get]
func MachinetypeSearch(ctx *gin.Context) {
	data := ctx.QueryArray("lastappeared_ids")
	if data != nil {
		lastappearedId := strings.Split(data[0], ",")
		machinetype := ctx.Query("machinetype")
		res := service.MachinetypeSearch(lastappearedId, machinetype)
		ctx.JSON(200, res)
	} else {
		machinetype := ctx.Query("machinetype")
		res := service.MachinetypeSearch(nil, machinetype)
		ctx.JSON(200, res)
	}
}

// @Summary 更新机型
// @Tags Trajectory_轨迹属性控制器
// @Value 轨迹控制器
// @Accept application/json
// @Param request body entity.Machinetype true "机型请求体"
// @Router /trajectory/machineTypeUpdate [post]
func MachinetypeUpdate(ctx *gin.Context) {
	dto := &entity.Machinetype{}
	err := ctx.ShouldBind(&dto)
	if err == nil {
		service.MachinetypeUpdate(dto)
		ctx.JSON(200, "更新成功")
	} else {
		ctx.JSON(500, "更新失败")
	}
}

// @Summary 添加轨迹绑定机型
// @Tags Trajectory_轨迹属性控制器
// @Value 轨迹控制器
// @Accept application/json
// @Param request body entity.Machinetype true "机型请求体"
// @Router /trajectory/machineTypeAdd [post]
func MachinetypeAdd(ctx *gin.Context) {
	dto := &entity.Machinetype{}
	err := ctx.ShouldBind(&dto)
	if err == nil {
		res := service.MachinetypeAdd(dto)
		if res == "" {
			ctx.JSON(200, "添加成功")
		} else {
			ctx.JSON(500, res)
		}
	} else {
		ctx.JSON(500, err.Error())
	}
}
