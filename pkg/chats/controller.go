package chats

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(app *fiber.App){
	routes := app.Group("/chats")
	routes.Get("/add", h.addChat)
}

func(h* handler) createChat(person1 String){

}
