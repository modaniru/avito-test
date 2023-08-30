# Авито тестовое задание
## Запуск
Для того, чтобы локально запустить приложение нужно
создать в директории приложения .env файл.
### Переменные окружения
Структура .env файла представлена ниже или в .env.example файле.
~~~bash
PORT=80 # Порт на котором будет запущено приложение
DB_USER=postgres # Имя пользователя бд
DB_PASSWORD=postgres # Пароль пользователя бд
DB_HOST=localhost:1111 # Хост базы данных, для запуска через docker-compose это поле игнорируется
DB_NAME=postgres # Имя базы данных
LOG_LEVEL=debug # Уровень логирования приложения (info, warn, error). debug установлен по умолчанию.
YANDEX_TOKEN=token # Yandex access token, его получение описано нижу
~~~
### Получение access token
Токен Яндекса можно получить тут: https://yandex.ru/dev/disk/poligon/

Если оставить поле YANDEX_TOKEN пустым, при запросе /history будет возвращаться статус 503 (Service Unavailable)
### Непосредственно запуск
После того, как были выполнены предыдущие шаги, можно приступать непосредственно к запуску
~~~bash
docker compose run
~~~
Приложение будет прослушиваться на 80 порту
## Описание эндпоинтов
### Postman
Описание эндпоинтов можно экспортировать в Postman, используя файл **avito.postman_collection.json**.
### Swagger
Swagger документацию можно получить по **localhost:80/swagger/**
### curl
#### Cохранение пользователя
Входные параметры:
- id - Идентификатор пользователя

Выходные параметры:
- Статус 201,  если пользователь сохранился
~~~bash
curl --location 'localhost:80/user' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1
}'
~~~
---
#### Удаление пользователя
Описание: При удалении пользователя, удаляются доступы к сегментам.
Удаленные сегменты у пользователя заносятся в историю.

Входные параметры:
- id - Идентификатор пользователя

Выходные параметры:
- Статус 204, если пользователь был удален
~~~bash
curl --location --request DELETE 'localhost:80/user/' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1
}'
~~~
---
#### Получение всех пользователей
Выходные параметры:
- Список идентификаторов
~~~bash
curl --location 'localhost:80/user/all' \
--data ''
~~~
---
#### Cохранение сегмента
Входные параметры:
- name - название сегмента

Выходные параметры:
- Идентификатор сегмента
- Статус 201, если сегмент был успешно сохранен
~~~bash
curl --location 'localhost:80/segment' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AVITO_VOICE_MESSAGES"
}'
~~~
---
#### Удаление сегмента
Описание: При удалении сегмента, отписывает всех его пользователей.
Удаленный сегмент у пользователей заносится в историю.

Входные параметры:
- name - название сегмента

Выходные параметры:
- Статус 204, если сегмент был успешно удален
~~~bash
curl --location --request DELETE 'localhost:80/segment' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AVITO_DISCOUNT_30"
}'
~~~
---
#### Получение сегментов
Выходные параметры:
- Список доступных сегментов
~~~bash
curl --location 'localhost:80/segment/all'
~~~
---
#### Добавление сегментов для пользователя
Входные параметры:
- user_id - идентификатор пользователя, если до этого пользователь не был создан, создает его
- []segments - список сегментов на добавление
- expire(optional) - добавляет время удаления сегмента у пользователя. Формат ГГГГ-ММ-ДД

Выходные параметры:
- Статус 201, если сегмент был успешно добавлен пользователю
~~~bash
curl --location 'localhost:80/user/segment/' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "segments": ["AVITO_VOICE_MESSAGES","AVITO_DISCOUNT_50","AVITO_PERFORMANCE_VAS","AVITO_DISCOUNT_30"],
    "expire": "2023-08-28"
}'
~~~
---
#### Добавление сегмента определенному проценту пользователей без него
Входные параметры:
- name - название сегмента
- percent - процент пользователей, у которых еще нет этого сегмента

Выходные параметры:
- Статус 201, если сегмент был успешно добавлен пользователям
- rows_affected - количество затронутых пользователей
~~~bash
curl --location 'localhost:80/user/segment/auto' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AVITO_IMAGES_213",
    "percent": 0.5
}'
~~~
---
#### Удаление сегментов у пользователя
Входные параметры:
- user_id - идентификатор пользователя
- []segments - список сегментов на удаление

Выходные параметры:
- Статус 204, если сегменты были успешно удалены у пользователя
~~~bash
curl --location --request DELETE 'localhost:80/user/segment/' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 1,
    "segments": ["AVITO_DISCOUNT_30"]
}'
~~~
---
#### Получение сегментов, доступных пользователю
Входные параметры:
- id - идентификатор пользователя

Выходные параметры:
- Список сегментов пользователя
~~~bash
curl --location --request GET 'localhost:80/user/segment/' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1
}'
~~~
---
#### Получение истории добавления/удаления сегментов у пользователей
Входные параметры:
- date - История за определенный промежуток времени (Формат ГГГГ-ММ-ДД или ГГГГ-ММ)

Выходные параметры:
- Ссылка на скачивание csv файла
~~~bash
curl --location 'localhost:80/history/?date=2023-08-31'
~~~
---
## Архитектура и возникшие вопросы
### архитектура базы данных
![Схема бд](files/Untitled.svg)
## Screenshots  
![App Screenshot](https://lanecdr.org/wp-content/uploads/2019/08/placeholder.png)  
