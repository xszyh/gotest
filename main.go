package main
import(
	// "log"
	"github.com/xszyh/gotest/app/config"
	"github.com/xszyh/gotest/app/entrypoint"
)

func main(){
	conf := config.GetConfig()
	app := entrypoint.NewRestApp(conf)
	app.Run()
}