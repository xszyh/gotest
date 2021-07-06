package entrypoint
import(
	"github.com/gin-gonic/gin"
	"github.com/xszyh/gotest/app/config"
	"github.com/xszyh/gotest/app/service"
	"net/http"
	"log"
	"fmt"
)

type RestApp struct{
	Config *config.Config
}

func NewRestApp(conf *config.Config)*RestApp{
	return &RestApp{conf}
}

func (a RestApp)Run(){
	r := gin.New()

	r.GET("/ping", Pong)

	r.Run(":" + a.Config.Port)
}

func Pong(c *gin.Context) {
	ret := "pong"
	userId := c.Query("userId")
	log.Println(userId)
	service := service.NewUserService()
	user := service.GetUserInfo(userId)

	ret = fmt.Sprintf("%s %s %d", ret, user.Name, user.Age) 

	c.String(http.StatusOK, ret)
}