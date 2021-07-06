package entrypoint
import(
	"github.com/gin-gonic/gin"
	"github.com/xszyh/gotest/app/config"
	"net/http"
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
	c.String(http.StatusOK, "pong")
}