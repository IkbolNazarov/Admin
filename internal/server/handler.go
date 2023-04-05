package server

import (
	"admin/internal/models"
	"admin/internal/repository"
	"admin/internal/services"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine *gin.Engine					//eсли правильно понял, то речь идет о сервисах, убрал
	Repository *repository.Repository //TODO: не используется и вообще не должен даже
}

func NewHandler(engine *gin.Engine, services *services.Services) *Handler {
	return &Handler{
		Engine: engine,
	}
}

func (h *Handler) Init() {
	h.Engine.GET("/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Connected"})
	})
	h.Engine.POST("/add_user", h.AddData)
	h.Engine.POST("/add_image", h.UploadImage)
	h.Engine.GET("/get_user", h.GetData)
	h.Engine.POST("/update_user", h.UpdateData)
	h.Engine.DELETE("/delete_user", h.DeleteUserData)
}

func (h *Handler) AddData(ctx *gin.Context) {
	var UserInfo models.UserInfo
	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.AddData(&UserInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "info is added")
}

func (h *Handler) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image") 							//TODO: файл сам не грузится
	if err != nil {												//проверил, сохраняется в ./pics/
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	randomise := strconv.Itoa(rand.Intn(99999))
	file.Filename = randomise + file.Filename            //добавив в конце данные ты формат файла испортил
	imagePath := filepath.Join("./pics/", file.Filename) //DONE
	imageFile, err := os.Create(imagePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer imageFile.Close()

	ctx.JSON(http.StatusOK, gin.H{"AddUser": file.Filename})
}

func (h *Handler) GetData(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)
	var listLenght int64
	UserLists, err := h.Repository.GetData(&pagination, listLenght)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	TotalPages, err := h.TotalPageUserInfo(int64(pagination.Limit), listLenght)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = UserLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h *Handler) UpdateData(ctx *gin.Context) {
	var userInfo *models.UserInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateData(userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"UpdateData": "Done"})
}

func (h *Handler) DeleteUserData(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Query("id")) //TODO: исполльзуй возможности если используешь gin gin ctx.Query()
	if err != nil {                          //DONE
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(id)
	err = h.Repository.DeleteData(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"DeleteData": "Done"})
}
