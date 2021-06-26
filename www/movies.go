package movies

import (
	"github.com/gin-gonic/gin"
	clientapi "github.com/italomarca/qqrfilme-backend-go/utils"
)

const API_KEY string = "6af1acbb5b00250f0669d50b891c76c6"
const API_URL string = "https://api.themoviedb.org/3"

// TODO: make it return only one movie, and make it random
func GetRandomMovie(ctx *gin.Context) {
	response := clientapi.GetJson(API_URL + "/movie/popular?api_key=" + API_KEY + "&language=pt-BR&page=1")
	ctx.JSON(200, response)
}
