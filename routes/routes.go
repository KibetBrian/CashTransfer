package routes
import (
	"github.com/gin-gonic/gin"
)

var Routes = func (){
	router :=gin.Default();
	UserRoutes(router)
		
	router.Run("localhost:8080")
}


