package api

import (
	"fmt"
	"hexagonal/infraestructure/entrypoints/api/apiRoutes"
	"sync"

	"github.com/labstack/echo"
)

func Run(wg *sync.WaitGroup) {
	fmt.Println(`

	  ___  _   _   _____   ___  ____________ _____ _     
	 / _ \| | | | /  __ \ / _ \ | ___ \ ___ \_   _| |    
	/ /_\ \ | | | | /  \// /_\ \| |_/ / |_/ / | | | |    
	|  _  | | | | | |    |  _  ||    /|    /  | | | |    
	| | | | |_| | | \__/\| | | || |\ \| |\ \ _| |_| |____
	\_| |_/\___/   \____/\_| |_/\_| \_\_| \_|\___/\_____/
														
																								 

	`)
	e := echo.New()
	//GetRoutesSaludo(e)
	//GetRoutesUsuarios(e)
	apiRoutes.GetRoutesSaludo(e)
	apiRoutes.GetRoutesUsuarios(e)
	e.Logger.Fatal(e.Start(":1323"))
}
