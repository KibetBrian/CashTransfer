package api
import(
	"github.com/gin-gonic/gin"
	"github.com/KibetBrian/fisa/controllers"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", controllers.SayHello)
	router.POST("/user/register", controllers.RegisterUser) 
	router.POST("/user/login", controllers.Login)
}
