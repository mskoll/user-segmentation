# User Segmentation service

Сервис для хранения пользователей и сегментов, в которых он состоит.

## Getting Started
Перед запуском сервиса нужно заполнить .env файл

Запуск сервиса:
```
docker-compose up
```

## Progress

Создание и удаление сегмента
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>

Добавление пользователя в сегмент
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>

Получение активных сегментов пользователя
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>

Доп.1: Получение истории попадания/выбывания пользователя из сегмента
<img alt="75" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/80-percents.png"/>

Доп.2: Установка времени автоматического удаления пользователя из сегмента
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>

Доп.3: Установка процента пользователей, автоматически попадающих в сегмент
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>

Тесты
<img alt="25" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/25-percents.png"/>

Swagger
<img alt="100" width="15" height="15" vspace="5" hspace="5" align="left" src="https://img.icons8.com/doodle/48/100-percents.png"/>


## Examples
* [Создание пользователя](#Создание-пользователя)
* [Получение активных сегментов пользователя](#Получение-активных-сегментов-пользователя)
* [Добавление/удаление сегментов пользователю](#Добавлениеудаление-сегментов-пользователю)
* [История попадания/выбывания пользователя из сегмента](#История-попаданиявыбывания-пользователя-из-сегмента)
* [Создание сегмента](#Создание-сегмента)
* [Удаление сегмента](#Удаление-сегмента)

### Создание пользователя 
Пример запроса:
```
curl --location --request POST 'http://localhost:8000/user/' \
--data '{
    "username": "username3"
}'
```
Пример ответа:
```
{
    "id": 3
}
```

### Получение активных сегментов пользователя
Пример запроса:
```
curl --location --request GET 'http://localhost:8000/user/2'
```
Пример ответа 1 (сегмент добавился пользователю автоматически при [создании сегмента](#Создание-сегмента)):
```
{
    "user": {
        "id": 2,
        "username": "username2"
    },
    "segments": [
        {
            "id": 4,
            "name": "AVITO_DISCOUNT_50",
            "percent": 60
        }
    ]
}
```
Пример ответа 2 (сегменты с id 1 и 2 были добавлены, сегмент с id 4 удален в [запросе](#Добавлениеудаление-сегментов-пользователю)):
```
{
    "user": {
        "id": 2,
        "username": "username2"
    },
    "segments": [
        {
            "id": 1,
            "name": "AVITO_VOICE_MESSAGES"
        },
        {
            "id": 2,
            "name": "AVITO_PERFORMANCE_VAS"
        }
    ]
}
```

### Добавление/удаление сегментов пользователю
На вход подается 2 списка. В списке на добавление можно добавить ttl. Если ttl не указать - время жизни сегмента будет 
не ограничено

Пример запроса:
```
curl --location --request POST 'http://localhost:8000/user/segment' \
--header 'Content-Type: application/json' \
--data '{
    "user_id" : 2,
    "to_add": [
        {"name":"AVITO_VOICE_MESSAGES", "ttl":"2023-09-10T18:00:00.28Z"},
        {"name":"AVITO_PERFORMANCE_VAS"}
    ],
    "to_del":[
        {"name":"AVITO_DISCOUNT_50"}
    ]
   
}'
```
Пример ответа:
```
{
    "message": "success"
}
```

### История попадания/выбывания пользователя из сегмента
Возвращение ссылки на csv файл не реализовано. На данный момент отчет возвращается в виде json

Пример запроса:
```
curl --location --request POST 'http://localhost:8000/user/operations' \
--header 'Content-Type: application/json' \
--data '{
    "user_id":2,
    "month":8,
    "year":2023
}'
```
Пример ответа:
```
[
    {
        "user_id": 2,
        "segment_name": "AVITO_DISCOUNT_50",
        "operation": "created",
        "datetime": "2023-08-31T10:21:13.130483Z"
    },
    {
        "user_id": 2,
        "segment_name": "AVITO_DISCOUNT_50",
        "operation": "deleted",
        "datetime": "2023-08-31T10:28:15.006221Z"
    },
    {
        "user_id": 2,
        "segment_name": "AVITO_PERFORMANCE_VAS",
        "operation": "created",
        "datetime": "2023-08-31T10:28:15.012057Z"
    },
    {
        "user_id": 2,
        "segment_name": "AVITO_VOICE_MESSAGES",
        "operation": "created",
        "datetime": "2023-08-31T10:28:15.012057Z"
    }
]
```

### Создание сегмента
Процент пользователей можно не указывать, в этом случае он будет нулевым

Пример запроса:
```
curl --location --request POST 'http://localhost:8000/segment/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "AVITO_DISCOUNT_50",
    "percent": 60
}'
```
Пример ответа:
```
{
    "id": 4
}
```

### Удаление сегмента
Пример запроса:
```
curl --location --request DELETE 'http://localhost:8000/segment/' \
--header 'Content-Type: application/json' \
--data '{
    "name":"AVITO_DISCOUNT_50"
}'
```
Пример ответа:
```
{
    "message": "success"
}
```

## Solutions
В процессе разработки были приняты  решения:
1. Как хранить историю попадания/выпадания пользователя из сегмента
> Можно было бы создать отдельную таблицу для хранения истории и при добавлении/удалении пользователя из сегмента 
> обновлять ее, либо же написать триггер на insert. Я решила поступить проще и делать union select для таблицы, 
> хранящей информацию о сегментах пользователей (в том числе дату попадания/выбывания из сегмента)
2. В каком формате принимать TTL
> Можно было принимать время жизни сегмента в форматах, например количество дней/часов/минут, но было решено принимать 
> таймстамп выбывания из сегмента, чтобы при изменении внешнего вида формы (на теоретическом фронте) ничего не менялось
> в сервисе