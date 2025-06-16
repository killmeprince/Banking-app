# Банковский REST API на Go

**Монолитное приложение с автосборкой через Docker Compose.**  
Реализованы: JWT-аутентификация, счета, карты с шифрованием, переводы, кредиты, аналитика, планировщик автоплатежей и почтовые уведомления через Mailhog.

---

## старт

1. **все сервисы одной командой:**
    ```bash
    docker-compose up --build -d
    ```

2. **доступность сервисов:**
    - API: [http://localhost:8080/api](http://localhost:8080/api)
    - Почта (Mailhog): [http://localhost:8025](http://localhost:8025)

3. **логи:**
    ```bash
    docker-compose logs app
    docker-compose logs migrate
    docker-compose logs db
    ```
4. **endpoints**
        *AUTH*

        POST /api/register
        — Регистрация пользователя
        body: { "Email": "...", "Password": "..." }

        POST /api/login
        — Логин, получить JWT
        body: { "Email": "...", "Password": "..." }
        response: { "token": "..." }

        *ACCOUNTS*

        POST /api/accounts
        — Создать новый счет для пользователя

        GET /api/accounts
        — Получить список всех своих счетов

        POST /api/accounts/deposit
        — Внести или снять деньги со счета
        body: { "account_id": ..., "amount": ... }
        (отрицательное amount — это списание)

        POST /api/transfer
        — Перевод между своими счетами
        body: { "from_id": ..., "to_id": ..., "amount": ... }

        *CARDS*

        POST /api/cards
        — Создать карту к своему счету
        body: { "account_id": ..., "number": "...", "exp": "...", "cvv": "..." }
        (number, exp, cvv — опционально, если пустые — сгенерятся автоматически)

        GET /api/cards
        — Получить список всех своих карт

        *CREDITS*

        POST /api/credits
        — Оформить кредит
        body:
        {
        "account_id": ...,
        "principal": ...,
        "rate": ...,
        "term_months": ...,
        "margin": ...
        }
        GET /api/credits/{id}/schedule
        — Получить график платежей по кредиту

        *TRANSACTIONS*
        
        GET /api/transactions
        — Получить список всех своих транзакций

        ANALYTICS
        GET /api/analytics/month
        — Доход/расход по месяцам (суммы)

        GET /api/analytics/credit
        — Текущая кредитная нагрузка

        GET /api/analytics/predict?days=N
        — Прогноз баланса на N дней вперед

       