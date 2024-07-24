// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	orderHanler "applicationDesignTest/internal/handler/order"
	orderRepository "applicationDesignTest/internal/repository/order"
	orderService "applicationDesignTest/internal/service/order"
	"applicationDesignTest/pkg/logger"
)

func main() {
	mux := http.NewServeMux()

	l := logger.NewLogger(log.Default())

	orderRepo := orderRepository.NewRepo(orderRepository.NewStorage())
	orderService := orderService.NewService(orderRepo)
	orderHandler := orderHanler.NewHandler(orderService, l)

	mux.HandleFunc("/orders", orderHandler.CreateOrder)

	l.LogInfo("Server listening on localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		l.LogInfo("Server closed")
	} else if err != nil {
		l.LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}
