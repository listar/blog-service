package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "qqfav-service/models"
	"qqfav-service/modules/log"
	"regexp"
	"strings"
)

type Count struct {
	Total int
}

// 列表
func ArticleList(c *gin.Context) {
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
	var article []m.Article
	m.Model.Order("id desc").Offset((pageIndex.(float64) - float64(1)) * pageSize.(float64)).Limit(pageSize).Find(&article)

	var total Count
	m.Model.Table(m.Article.TableName(m.Article{})).Select("count(1) as total").Where("deleted_at IS NULL").Find(&total)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": article,
			"total": total.Total,
		},
	})
}

// 详情
func ArticleDetail(c *gin.Context) {
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
	var article m.Article
	m.Model.Where("id = ?", ID).Find(&article)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": article,
		},
	})
}

// 操作
func ArticleAction(c *gin.Context) {
	var artice m.Article
	json := make(map[string]interface{})
	bjerr := c.BindJSON(&json)
	if bjerr != nil {
		log.Error(bjerr)
		return
	}

	Intro := trimHtml(string([]rune(json["Content"].(string))[:300]))
	Intro = strings.ReplaceAll(Intro, "#", "")
	Intro = strings.ReplaceAll(Intro, "> ", "")
	Intro = strings.ReplaceAll(Intro, "+ ", "")
	Intro = strings.ReplaceAll(Intro, "- ", "")

	articeData := m.Article{
		Title: json["Title"].(string),
		Author: json["Author"].(string),
		OuterLink: json["OuterLink"].(string),
		Tags: json["Tags"].(string),
		Intro: Intro,
		Content: json["Content"].(string)}
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
func ArticleDel(c *gin.Context) {
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
	var article m.Article
	m.Model.Where("id = ?", ID).Delete(&article)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": gin.H{
			"result": article,
		},
	})
}

func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
