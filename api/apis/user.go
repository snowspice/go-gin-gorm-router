package apis
import (
	"github.com/gin-gonic/gin"
	model "github.com/snowspice/go-gin-gorm-router/api/models"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin/json"
)

//列表数据
func Index(c *gin.Context){
	c.String(http.StatusOK,"ok")
}


//curl -X GET http://127.0.0.1:8000/users
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data":   result,
		"message":"成功",
	})
}

//添加数据
//curl -X POST http://127.0.0.1:8000/user -H "Content-Type:application/x-www-form-urlencoded" -d "username=hill&password=123456"
func Add(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
//curl -X PUT http://127.0.0.1:8000/user/2 -H "Content-Type:application/x-www-form-urlencoded"  -d "password=345678"
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "修改成功",
	})
}

//删除数据
func Delete(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Delete(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "删除成功",
	})
}

// 分页查询
//curl -X POST http://127.0.0.1:8000/user/page -H "Content-Type:application/json" -d "{\"page\":1,\"pageSize\":5,\"total\":0,\"pageCount\":0,\"nums\":null,\"numsCount\":0,\"username\":\"hill\"}"
//curl -X POST http://127.0.0.1:8000/user/page -H "Content-Type:application/json" -d "{\"page\":1,\"pageSize\":3,\"total\":0,\"pageCount\":0,\"nums\":null,\"numsCount\":0,\"username\":\"hill\"}"
func FindPage(c *gin.Context) {
	var user model.User
	var userArg model.UserArgs
	c.Bind(&userArg)

	if b, err := json.Marshal(userArg); err == nil {
		fmt.Println("============controller====struct 到json str==")
		fmt.Println(string(b))
	}
    result,pager,err :=user.FindPage(userArg)
    if(err != nil){
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "查询页面信息出错",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"message": "查询成功",
		"data":    result,
		"pager":pager,
	})

}
