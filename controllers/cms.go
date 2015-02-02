package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

// CMS API
type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("Product", c.Product)
}

// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     path    string  false        "The email for login"
// @Success 200 {string} map[string]string
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (c *CMSController) StaticBlock() {
	s := c.Ctx.Input.Param(":key")
	c.Data["json"] = map[string]interface{}{"category_id": s}
	c.ServeJson()
}

// @Title Get Product list
// @Description Get Product list by some info
// @Success 200 {string} map[string]interface{}
// @Param   category_id     query   int false       "category id"
// @Param   brand_id    query   int false       "brand id"
// @Param   query   query   string  false       "query of search"
// @Param   segment query   string  false       "segment"
// @Param   sort    query   string  false       "sort option"
// @Param   dir     query   string  false       "direction asc or desc"
// @Param   offset  query   int     false       "offset"
// @Param   limit   query   int     false       "count limit"
// @Param   price           query   float       false       "price"
// @Param   special_price   query   bool        false       "whether this is special price"
// @Param   size            query   string      false       "size filter"
// @Param   color           query   string      false       "color filter"
// @Param   format          query   bool        false       "choose return format"
// @Failure 400 not enough input
// @Failure 500 get products common error
// @router /products [get]
func (c *CMSController) Product() {
	cId := strings.TrimSpace(c.Ctx.Input.Query("category_id"))
	bId := strings.TrimSpace(c.Ctx.Input.Query("brand_id"))
	c.Data["json"] = map[string]interface{}{"category_id": cId, "brand_id": bId}
	c.ServeJson()
}
