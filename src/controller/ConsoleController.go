package controller

import (
	"fmt"
	"sync"
	"tor_search/src/service"
)

type ConsoleController struct {
	service *service.Service
	wg      *sync.WaitGroup
}

func NewConsoleController(service *service.Service) *ConsoleController {
	return &ConsoleController{
		service: service,
		wg:      &sync.WaitGroup{}}
}

func (c *ConsoleController) Run() {
	word, lim := "", 0
	fmt.Println("\nИспользование:\n[слово] [максимальное кол-во страниц результата]")

	for {
		c.wg.Wait()
		fmt.Println("\nВвод:")

		_, err := fmt.Scanln(&word, &lim)
		if err != nil {
			fmt.Println("Проверьте правильность ввода!")
			continue
		}

		c.wg.Add(1)
		go c.searchAndPrint(word, lim)
	}
}

func (c *ConsoleController) searchAndPrint(word string, limit int) {
	defer c.wg.Done()

	res, err := c.service.Search.Search(word, limit)
	if err != nil {
		c.service.Logger.Log(err)
		return
	}

	for _, v := range res {
		fmt.Println(fmt.Sprintf(
			"--------------------\n%s\n(%s, %d)",
			v.Url, v.Word, v.Count))
	}
}
