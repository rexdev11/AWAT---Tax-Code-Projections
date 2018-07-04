package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func main() {
	fmt.Println("10")
}

//Todo: Get user input and compute projections to return a data visualization.

type State struct {

}

type Entity struct {
	id string
	fullName string
	history map[string] State
	state State
}

type TaxCodeConfigSpec struct {
	currentProcess int
}

type TaxCodeConfigs struct {
	entity Entity
	inProgress bool
	pristine bool
	old bool
	specs TaxCodeConfigSpec
}

type Store struct {
	value func(key string)
	next func()
	dispatch func()
	_source map[string] string
}

func StoresFactory(configs TaxCodeConfigs) {
	// Todo return a stores obj
}

// TODO: Main UX Flow ( Init UI -> Consumes Profile Data -> Sets Session State -> Retrieves Process Library -> Assembles Processes -> Project )

func initiallizeUI() {

}

type CrawlerOptions struct {
	urls [0:]string
}

func (options )initializeCrawler(app iris.Application) {
	app.Any("/dist/crawler.js", websocket.ClientHandler())
}

func crawlerHandler(ctx iris.Context)(routes *Routes[]) {

}

// TODO: Homeostasis do while ( Resilience is > critical , no Sig Term)

// TODO: Homeostatic functions ( Health check -> audit servers -> forEach() { compare with manifest IF not normal send report/alert } )

// TODO: Internal functions ( Init Data Update -> get list of document URLs -> { init crawler -> retrieve docs and format -> parse -> store processes } -> log histories -> ... )

// TODO: Internal checks -> IF not normal send report/alert

// TODO: Alerts/Reports ->

// TODO: API ( Processes Library, Crawler, Reflection, Session, Store )