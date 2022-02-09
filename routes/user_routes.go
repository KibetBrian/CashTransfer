package routes
import(
	"github.com/gin-gonic/gin"
	"github.com/KibetBrian/fisa/controllers"
)

var UserRoutes = func ( router *gin.Engine){
	router.GET("/hello", controllers.SayHello)
	router.POST("/register", controllers.RegisterUser) 
	router.POST("/login", controllers.Login)
}
