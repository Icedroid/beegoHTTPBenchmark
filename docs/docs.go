package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/cms","description":"CMS API\n"}],"info":{"title":"mobile API","description":"mobile has every tool to get any job done, so codename for the new mobile APIs.","contact":"astaxie@gmail.com"}}`
var subapi string = `{"/cms":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/cms","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"[get]","description":"","operations":[{"httpMethod":"GET","nickname":"getStaticBlock","type":"","summary":"get all the staticblock by key","parameters":[{"paramType":"path","name":"key","description":"\"The email for login\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"map[string]string","responseModel":""},{"code":400,"message":"Invalid email supplied","responseModel":""},{"code":404,"message":"User not found","responseModel":""}]}]},{"path":"/products","description":"","operations":[{"httpMethod":"GET","nickname":"Get Product list","type":"","summary":"Get Product list by some info","parameters":[{"paramType":"query","name":"category_id","description":"\"category id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"brand_id","description":"\"brand id\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"query","description":"\"query of search\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"segment","description":"\"segment\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sort","description":"\"sort option\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"dir","description":"\"direction asc or desc\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"offset\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"count limit\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"price","description":"\"price\"","dataType":"float","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"special_price","description":"\"whether this is special price\"","dataType":"bool","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"size","description":"\"size filter\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"color","description":"\"color filter\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"format","description":"\"choose return format\"","dataType":"bool","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"map[string]interface{}","responseModel":""},{"code":400,"message":"not enough input","responseModel":""},{"code":500,"message":"get products common error","responseModel":""}]}]}]}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
