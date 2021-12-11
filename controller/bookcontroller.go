package controller

import (
	"books/domain"
	"books/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	gorilla "github.com/gorilla/websocket"
)

type BookController struct {
	bookService service.BookService
	upgrader    gorilla.Upgrader
	listeners   []gorilla.Conn
}

func upgrader() gorilla.Upgrader {
	return gorilla.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {return true},
	}
}

func New(s service.BookService) BookController {
	return BookController{
		bookService: s, 
		upgrader: upgrader(), 
		listeners: make([]gorilla.Conn, 0)}
}

func (c *BookController) notifyAll(b domain.Book) {
	for _, ws := range c.listeners {
		err := ws.WriteJSON(b)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (c *BookController) ListenForNotifications(ctx *gin.Context){
	responseWriter, request := ctx.Writer, ctx.Request

	wsConnection, err := c.upgrader.Upgrade(responseWriter, request, nil)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Could not create websocket connection")
		return
	}
	c.listeners = append(c.listeners, *wsConnection)
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.bookService.FindBooks())
}

func (c *BookController) SaveBook(ctx *gin.Context) {
	book := domain.Book{}
	err := ctx.BindJSON(&book)
	fmt.Println(book)


	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err = c.bookService.AddBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go c.notifyAll(book)
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	book := domain.Book{}
	err := ctx.BindJSON(&book)

	fmt.Println(ctx.Request.Body)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err = c.bookService.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go c.notifyAll(book)
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) GetBook(ctx *gin.Context) {
	isbn := ctx.Params.ByName("isbn")

	if isbn == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "isbn not found"})
		return
	}
	book, err := c.bookService.FindBook(isbn)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)

}
