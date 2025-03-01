# Этап сборки
FROM golang:1.23.3 AS build

WORKDIR /usr/local/src

# Копируем файлы зависимостей и загружаем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Проверяем версию Go (опционально, для диагностики)
RUN go version

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o /todolist-drpetproject ./cmd/todolist

# Этап финального образа
FROM alpine:3.18 AS runner

# Устанавливаем bash, если необходимо
RUN apk --no-cache add bash

WORKDIR /usr/local/src

# Копируем скомпилированный бинарник и конфигурационный файл
COPY --from=build /todolist-drpetproject /usr/local/src/todolist-drpetproject
COPY ./config.yml ./

# Определяем точку входа
ENTRYPOINT ["/usr/local/src/todolist-drpetproject"]
