package fbs

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"net/url"
	"github.com/kataras/iris/core/host"
	"os"
)

func FBS() {
	fmt.Println("yea!")
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("dist/main.js", false)
	})
	setupWebsocket(app)
}

func indexHandler(ctx iris.Context) {
	ctx.ServeFile("index.html", true)
}

func setupWebsocket(app *iris.Application) {

	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})

	ws.OnConnection(handleConnection)

	// return index for base path.
	app.Get("/index", indexHandler)

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/echo", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})

	target, err := url.Parse("http://localhost:22")

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting Server")

	go host.NewProxy("localhost:2733", target).ListenAndServe()
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	app.Run(iris.TLS("localhost:2733", string(wd + `/../fbs/pems/domain.crt`), string(wd + "/../fbs/pems/domain.key")))
}

func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner with:
		// c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}

func FBScrapperHandler() {

}