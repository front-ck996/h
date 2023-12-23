package gin_middleware

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	//go:embed dist
	front embed.FS
)

type SwaggerConfig struct {
	RelativePath string
}

type swaggerService struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	SwaggerVersion string `json:"swaggerVersion"`
	Location       string `json:"location"`
}
type swaggerContent struct {
	Swagger string `json:"swagger"`

	Info struct {
		Contact     struct{} `json:"contact"`
		Description string   `json:"description"`
		Title       string   `json:"title"`
		Version     string   `json:"version"`
	} `json:"info"`
}

func init() {
	var err error
	if err != nil {
		log.Println("no swagger.json found in ./docs")
	}
}

func SwaggerHandler(config SwaggerConfig) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		switch ctx.Request.URL.Path {
		case config.RelativePath + "/docList":
			var result []swaggerService
			filepath.Walk("docs", func(path string, info fs.FileInfo, err error) error {
				if strings.HasSuffix(path, ".json") {
					parseDesc := swaggerContent{}
					file, err := os.ReadFile(path)
					if err == nil {
						json.Unmarshal(file, &parseDesc)
					}
					result = append(result, swaggerService{
						Name:           parseDesc.Info.Title,
						SwaggerVersion: parseDesc.Swagger,
						Url:            config.RelativePath + "/docJson?file=" + filepath.Base(path),
						Location:       config.RelativePath + "/docJson?file=" + filepath.Base(path),
					})
				}
				return nil
			})
			ctx.JSON(http.StatusOK, result)
			break
		case config.RelativePath + "/docJson":
			docFile := path.Base(ctx.DefaultQuery("file", ""))
			file, err := os.ReadFile("docs/" + docFile)
			if err != nil {
				return
			}
			ctx.JSON(200, string(file))
			break

		default:
			fmt.Println(strings.TrimPrefix(ctx.Request.RequestURI, config.RelativePath), http.FS(front))
			ctx.FileFromFS(strings.TrimPrefix(ctx.Request.RequestURI, config.RelativePath), http.FS(front))
		}

	}
}
