package app

import (
	"TestTask/config"
	order_product_usecase "TestTask/internal/domain/order_product/usecase"
	order_product_repository "TestTask/internal/infrastructure/repository/order_product"
	"TestTask/pkg/postgres"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type App struct {
	useCaseList
}

type useCaseList struct {
	orderProductUC *order_product_usecase.OrderProductUseCase
}

func NewApp(cfg *config.Config) App {

	psql, err := postgres.New(cfg.Pg.User, cfg.Pg.Password, cfg.Pg.Host, cfg.Pg.Port, cfg.Pg.DbName, cfg.Pg.SslMode)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация Repository
	orderProductRepo := order_product_repository.NewOrderProductRepository(psql)

	// Инициализация UseCase
	orderProductUC := order_product_usecase.NewOrderProductUseCase(orderProductRepo)

	return App{
		useCaseList: useCaseList{
			orderProductUC: orderProductUC,
		},
	}
}

func (app *App) Run() {

	args := os.Args[1:]
	var orders []int
	for {
		if len(args) == 0 {
			fmt.Println("Пожалуйста, введите номера заказов:")
		}

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		}
		input = strings.TrimSpace(input)
		args = strings.Fields(input)

		valid := true
		for _, arg := range args {
			numb, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Ошибка: '%s' не является числом.\n", arg)
				valid = false
				args = []string{}
				break
			}
			orders = append(orders, numb)
		}
		if valid {
			break
		}
	}

	orderProductMap, err := app.orderProductUC.GetOrderProductList(orders)
	if err != nil {
		fmt.Println(err)
	}

	for rack, orderProductList := range orderProductMap {
		fmt.Println("=== Стеллаж", rack)
		for _, product := range orderProductList {
			fmt.Println(product)
		}
	}

}
