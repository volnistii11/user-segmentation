# user-segmentation

Здравствуйте, прежде всего, спасибо за внимание=)

## Описание

1. В директории `user-segmentation/cmd/segmenter` сделать `go build`, затем запустить бинарник `./segmenerter`;
2. SQL файл находится по пути `user-segmentation/config/storage.sql`.

## API

| Метод |         Путь          |            Описание            |
|:-----:|:---------------------:|:------------------------------:|
| POST  | [/api/segment/create] |     Создать новый сегмент      |
| POST  | [/api/segment/delete] |        Удалить сегмент         |
| POST  |  [/api/segment/user]  | Обновить сегменты пользователя |
|  GET  |  [/api/user/segment]  |  Узнать сегменты пользователя  |

### Создание нового сегмента

Формат запроса

```
POST /api/segment/create
Content-Type: application/json
...

{
    "slug": "AVITO_SOMETHING"
}
```

### Удаление сегмента

Формат запроса

```
POST /api/segment/delete
Content-Type: application/json
...

{
    "slug": "AVITO_SOMETHING"
}
```

### Обновление сегментов у пользователя

Формат запроса

```
POST /api/segment/user
Content-Type: application/json
...

{
    "user_id": "3",
    "add_segments": [
        "AVITO_1",
        "AVITO_6"
    ],
    "delete_segments": [
         "AVITO_1",
         "AVITO_2"
    ]
}
```

### Получение сегментов пользователя

Формат запроса

```
GET /api/user/segment
Content-Type: application/json
...

{
    "user_id": "3"
}
```

## Возникшие вопросы

1. Было не очень понятно, что делать с дублированием сегмента у пользователя. Я сделал, что возвращается ошибка при дублировании, но обработать ее как конфликт не успел.
2. Похожий вопрос с удалением, несуществующего сегмента у пользователя. Как ошибку я это расценивать не стал.

# =)