package http

import (
	"blogging-service/internal/app/config"
	"blogging-service/internal/service/post"
	"blogging-service/pkg/reqresp"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	srv     post.Service
	timeout time.Duration
}

func NewHandler(cfg config.HttpConfig, srv post.Service) *Handler {
	return &Handler{
		srv:     srv,
		timeout: time.Duration(cfg.RequestTimeoutSeconds) * time.Second,
	}
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (h *Handler) Create(c echo.Context) error {
	ctx, cancel := h.context(c)

	defer cancel()

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	idFloat := jsonBody["author_id"].(float64)

	id := int(idFloat)

	resp, err := h.srv.CreatePost(ctx, reqresp.PostRequest{Title: jsonBody["title"].(string), Content: jsonBody["content"].(string), AuthorId: id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)

}

func (h *Handler) Read(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	resp, err := h.srv.GetPosts(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Update(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	authidFloat := jsonBody["author_id"].(float64)
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	authid := int(authidFloat)

	resp, err := h.srv.UpdatePost(ctx, reqresp.PostRequest{Id: id, Title: jsonBody["title"].(string), Content: jsonBody["content"].(string), AuthorId: authid})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Delete(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	resp, err := h.srv.DeletePost(ctx, reqresp.PostRequest{Id: id})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "request_id", c.Response().Header().Get(echo.HeaderXRequestID))

	return context.WithTimeout(ctx, h.timeout)
}

//func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	resp := &Response{}
//	defer json.NewEncoder(w).Encode(resp)
//
//	repo := repository.PostRepository{Collection: h.Repo.Collection}
//	res, err := repo.GetAll()
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: ", err)
//	}
//
//	resp.Data = res
//	w.WriteHeader(http.StatusOK)
//}
//
//func (h *PostHandler) GetPostById(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	resp := &Response{}
//	defer json.NewEncoder(w).Encode(resp)
//
//	repo := repository.PostRepository{Collection: h.Repo.Collection}
//	postId := mux.Vars(r)["id"]
//
//	objId, err := primitive.ObjectIDFromHex(postId)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		log.Println("error: invalid object id format")
//		resp.Error = "invalid object id format"
//		return
//	}
//
//	res, err := repo.GetByID(objId.Hex())
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: ", err)
//		return
//	}
//
//	resp.Data = res
//	w.WriteHeader(http.StatusOK)
//
//}
//
//func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var post models.Post
//	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	post.ID = primitive.NewObjectID()
//	post.CreatedAt = time.Now()
//
//	res, err := h.Repo.Create(&post)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		log.Println("error: ", err)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(res)
//
//}
//
//func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	resp := &Response{}
//	defer json.NewEncoder(w).Encode(resp)
//
//	postID := mux.Vars(r)["id"]
//
//	if postID == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = "invalid post id"
//		return
//	}
//
//	objId, err := primitive.ObjectIDFromHex(postID)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: invalid object id format")
//		return
//	}
//
//	var post models.Post
//
//	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: invalid post body")
//		return
//	}
//
//	post.UpdatedAt = time.Now()
//	post.ID = objId
//
//	repo := repository.PostRepository{Collection: h.Repo.Collection}
//	count, err := repo.Update(post.ID, &post)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: ", err)
//		return
//	}
//
//	resp.Data = count
//	w.WriteHeader(http.StatusOK)
//}
//
//func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	resp := &Response{}
//	defer json.NewEncoder(w).Encode(resp)
//
//	postID := mux.Vars(r)["id"]
//
//	repo := repository.PostRepository{Collection: h.Repo.Collection}
//	count, err := repo.Delete(postID)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Println("error: ", err)
//		return
//	}
//
//	resp.Data = count
//	w.WriteHeader(http.StatusOK)
//}
