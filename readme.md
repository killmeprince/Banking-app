Что тут вообще?
это REST API банковского сервиса на Go. Я пытался сделать почти всё по ТЗ, но кое-что ещё не допилил.

Кратко, что работает

Регистрация / логин через JWT  

Счёта  

Создание (`POST /api/accounts`)  
Депозит/списание (`POST /api/accounts/deposit`)  
Список (`GET /api/accounts`)  

Переводы между своими счетами в одной транзакции (`POST /api/transfer`)  

Карты 
Генерация валидного номера (алгоритм Луна)  
Хеш CVV + HMAC  
Возврат маскированного номера и срока  

Кредиты

Оформление кредита + расчёт аннуитета (`POST /api/credits`)  
Сохранение графика в таблицу `payment_schedules`  
Получение графика по ID (`GET /api/credits/{id}/schedule`)  

Аналитика  

Месяц доход/расход (`GET /api/analytics/month`)  
Кредитная нагрузка (`GET /api/analytics/credit`)  
Прогноз баланса (`GET /api/analytics/predict?days=N`)  
SOAP-интеграция с ЦБР(метод `GetKeyRate()` есть, но UI не проверял)  
Транзакции и SQL через `sqlx`, всё параметризовано, защиты от SQL-инъекций.  
Тестировал руками через `curl`.

Что не реализовано и почему

PGP-шифрование карт через `pgcrypto` в БД — лень было ковыряться с ключами.  
Cron-шедулер для автосписания платежей — Не получилось подключить `cron/v3`.  
SMTP-уведомления (gomail.v2 + MailHog) — не получилось:( 
Logrus вместо `log.Println`
Полноценный ACL (проверка прав на трансфер/карты) только частично в deposit 

Быстрый старт

docker-compose up -d

миграции

cat migrations/000_init.sql    | docker exec -i $(docker-compose ps -q db) psql -U user -d banking
cat migrations/001_accounts_cards.sql \
                              | docker exec -i $(docker-compose ps -q db) psql -U user -d banking
cat migrations/002_transactions_credits.sql \
                              | docker exec -i $(docker-compose ps -q db) psql -U user -d banking

настроить ENV

export DB_HOST=localhost DB_PORT=5432 DB_USER=user DB_PASS=pass DB_NAME=banking
export JWT_SECRET=supersecretkey

собрать и запустить

go build ./cmd/server
go run cmd/server/main.go
