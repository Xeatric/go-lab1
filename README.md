# Лабораторная работа №1: Знакомство с клиент-серверным взаимодействием
## Запуск приложения
1. **Клонировать репозиторий** (или скопировать файлы проекта):

   git clone <https://github.com/Xeatric/go-lab1>
   cd lab1-go
   
2. **Собрать и запустить контейнер**

   docker-compose up --build

3. **После успешного запуска вы увидите**
4. 
      Сервер запущен на http://localhost:8080
      Доступные маршруты:
      GET  /info     - JSON с количеством дней до Нового года

5. **Проверить работу сервера(Postman)**

      Метод: GET
      URL: http://localhost:8080/info
      Нажать Send

      Ответ (пример для июня 2026 года)
         {
            "days_before_new_year": 210
         }  
      Статус-код: 200 OK
      Content-Type: application/json
    
    Ошибочный запрос POST /info (метод не разрешён)
      Метод: POST
      URL: http://localhost:8080/info
      Ответ:
         {
            "error": "Метод не поддерживается",
            "status": "Method Not Allowed"
         }
      Статус-код: 405 Method Not Allowed
