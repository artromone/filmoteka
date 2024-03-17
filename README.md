# OpenAPI:

```
openapi: 3.0.0
info:
  title: Фильмотека API
  version: 1.0.0
servers:
  - url: http://localhost:8080
paths:
  /actors:
    post:
      summary: Добавить информацию об актёре
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                gender:
                  type: string
                  enum: [male, female, other]
                birth_date:
                  type: string
                  format: date
      responses:
        '200':
          description: OK
    get:
      summary: Получить список актёров
      responses:
        '200':
          description: OK
  /actors/{actor_id}:
    put:
      summary: Изменить информацию об актёре
      parameters:
        - in: path
          name: actor_id
          required: true
          schema:
            type: integer
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                name:
                  type: string
                gender:
                  type: string
                  enum: [male, female, other]
                birth_date:
                  type: string
                  format: date
      responses:
        '200':
          description: OK
        '404':
          description: Актёр не найден
    delete:
      summary: Удалить информацию об актёре
      parameters:
        - in: path
          name: actor_id
          required: true
          schema:
            type: integer
      responses:
        '204':
          description: No Content
        '404':
          description: Актёр не найден
  /movies:
    post:
      summary: Добавить информацию о фильме
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                title:
                  type: string
                description:
                  type: string
                release_date:
                  type: string
                  format: date
                rating:
                  type: number
                  minimum: 0
                  maximum: 10
                actors:
                  type: array
                  items:
                    type: integer
      responses:
        '200':
          description: OK
    get:
      summary: Получить список фильмов
      parameters:
        - in: query
          name: sort_by
          schema:
            type: string
            enum: [title, rating, release_date]
          description: Поле для сортировки
      responses:
        '200':
          description: OK
  /movies/search:
    get:
      summary: Поиск фильма
      parameters:
        - in: query
          name: query
          schema:
            type: string
          description: Поисковый запрос
      responses:
        '200':
          description: OK
  /actors/{actor_id}/movies:
    get:
      summary: Получить список фильмов актёра
      parameters:
        - in: path
          name: actor_id
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: OK
security:
  - basicAuth: []
```
