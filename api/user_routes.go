package api
import(
	"github.com/gin-gonic/gin"
	"github.com/KibetBrian/fisa/handlers"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", handlers.SayHello)
	router.POST("/user/register", handlers.RegisterUser) 
	router.POST("/user/login", handlers.Login)
}
