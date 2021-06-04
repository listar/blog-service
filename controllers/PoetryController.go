package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "qqfav-service/models"
	"qqfav-service/modules/log"
)


// 列表
func PoetryList(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	bjerr := c.BindJSON(&json)
	if bjerr != nil {
		log.Error(bjerr)
		return
	}
	pageIndex := json["pageIndex"]
	pageSize := json["pageSize"]

	if pageSize.(float64) >= 100 {
		pageSize = 100
	}
	var Poetry []m.Poetry
	m.Model.Order("id desc").Offset((pageIndex.(float64) - float64(1)) * pageSize.(float64)).Limit(pageSize).Find(&Poetry)

	var total Count
	m.Model.Table(m.Poetry.TableName(m.Poetry{})).Select("count(1) as total").Where("deleted_at IS NULL").Find(&total)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": Poetry,
			"total": total.Total,
		},
	})
}

// 详情
func PoetryDetail(c *gin.Context) {
	json := make(map[string]interface{})
	bjerr := c.BindJSON(&json)
	if bjerr != nil {
		log.Error(bjerr)
		return
	}
	ID := json["ID"]
	if ID.(float64) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "id param is error",
			"data": gin.H{
			},
		})
		return
	}
	var Poetry m.Poetry
	m.Model.Where("id = ?", ID).Find(&Poetry)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": Poetry,
		},
	})
}

// 操作
func PoetryAction(c *gin.Context) {
	var artice m.Poetry
	json := make(map[string]interface{})
	bjerr := c.BindJSON(&json)
	if bjerr != nil {
		log.Error(bjerr)
		return
	}
	articeData := m.Poetry{Title: json["Title"].(string), Author: json["Author"].(string), Content: json["Content"].(string), Remark: json["Remark"].(string) }
	if json["id"].(float64) == 0 {
		m.Model.Create(&articeData)
	} else {
		// update
		m.Model.Model(&artice).Where("id = ?", json["id"].(float64)).Update(articeData)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
		},
	})
}


// 删除
func PoetryDel(c *gin.Context) {
	json := make(map[string]interface{})
	bjerr := c.BindJSON(&json)
	if bjerr != nil {
		log.Error(bjerr)
		return
	}
	ID := json["ID"]
	if ID.(float64) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 100,
			"msg":  "id param is error",
			"data": gin.H{
			},
		})
		return
	}
	var Poetry m.Poetry
	m.Model.Where("id = ?", ID).Delete(&Poetry)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": Poetry,
		},
	})
}
