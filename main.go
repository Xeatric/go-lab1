package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//ВСПОМОГАТЕЛЬНЫЕ ФУНКЦИИ

// writeJSON отправляет JSON-ответ
func writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}

// writeErr отправляет JSON-ответ с ошибкой
func writeErr(w http.ResponseWriter, message string, statusCode int) {
	errResponse := map[string]string{
		"error":  message,
		"status": http.StatusText(statusCode),
	}
	writeJSON(w, errResponse, statusCode)
}

// logRequest - middleware для логирования входящих запросов
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем входящий запрос
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Вызываем следующий обработчик
		next(w, r)

		// Логируем время выполнения
		log.Printf("Запрос обработан за %v", time.Since(start))
	}
}

// daysBeforeNewYear вычисляет количество дней до следующего Нового года
func daysBeforeNewYear() int {
	now := time.Now()

	// Определяем дату следующего Нового года (1 января следующего года)
	year := now.Year()
	newYear := time.Date(year+1, time.January, 1, 0, 0, 0, 0, now.Location())

	// Вычисляем разницу в днях
	days := int(newYear.Sub(now).Hours() / 24)

	return days
}

//ОБРАБОТЧИКИ (HANDLERS)

// homeHandler обрабатывает GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Добро пожаловать! Ты открыл страницу: %s", r.URL.Path)
}

// infoHandler обрабатывает GET /info
func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErr(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Вычисляем количество дней до Нового года
	days := daysBeforeNewYear()

	// Создаем структуру с одним полем, как в примере
	response := struct {
		DaysBeforeNewYear int `json:"days_before_new_year"`
	}{
		DaysBeforeNewYear: days,
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Ошибка при формировании JSON", http.StatusInternalServerError)
	}
	log.Printf("Отправляем JSON с днями: %d", days)
}

//НАСТРОЙКА МАРШРУТОВ

// setupRoutes регистрирует все обработчики с middleware
func setupRoutes() {
	// Сначала регистрируем КОНКРЕТНЫЕ маршруты
	http.HandleFunc("/info", logRequest(infoHandler))
	// ПОТОМ регистрируем общий маршрут "/"
	http.HandleFunc("/", logRequest(homeHandler))
}

//ТОЧКА ВХОДА

func main() {
	// Настраиваем маршруты
	setupRoutes()

	// Выводим информацию о запуске
	fmt.Println("========================================")
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("========================================")
	fmt.Println("Доступные маршруты:")
	fmt.Println("  GET  /info     - JSON с количеством дней до Нового года")
	// Запускаем сервер
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
