# CT Stickers Bot

Бот для текстового поиска стикеров

## Сборка

```
go get .
go build -o CTStickersBot
```

## Запуск

Поднять локальную базу (можно не поднимать, если есть внешняя).

```
docker-compose -p stickers_bot -f postgresql/docker-compose.yml build
docker-compose -p stickers_bot -f postgresql/docker-compose.yml up -d
```

Задать переменные окружения:

* BOT_AUTH_TOKEN - telegram токен бота;
* OCRAPI_TOKEN - токен для парсинга текста.

```
./CTStickersBot
```

## Характеристики

### Telegram Api

* Используется https://github.com/go-telegram-bot-api/telegram-bot-api.
* Реализована Long-polling схема.
* Поиск стикеров работает через inline сообщения. Не забудьте подключить inline через Bot Father.

### Логирование

* Используется struct логирование через http://go.uber.org/zap.

### Хранилище

* __PostgreSQL__.
* SQL для наливки базы в `postgresql/initial.sql`, в нем уже есть несколько стикеров для примера.
* Можно запустить как Docker контейнер. Образ лежит в `postgresql/docker-compose.yml`.

### Парсинг текста стикеров

* Используется api https://api.ocr.space/.
* Для использования нужно получить бесплатный токен.

### Конфиги

* Все настраивается в `config_dev.yaml`.
* Почти все настройки можно переопределять через соответствующие переменные окружения. Соответствие настроек и
  переменных можно посмотреть в коде.

### Web-Сервер

* Запускается при старте бота.
* Умеет отвечать на две ручки:
  * `/`  - некоторая статистика из базы;
  * `/ping` - должна отвечать `pong`.

### Ticker

* Используется для запуска бота в heroku. Раз в некоторое время бот дергает сам у себя _ping_ через внешний хост, чтобы
  heroku думал, что сайтом кто-то пользуется.
