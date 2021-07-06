package main
import(
	"log"
	"github.com/xszyh/gotest/config"
)

func main(){
	conf := config.GetConfig()
	app := NewRestApp(conf)
	app.Run()
}