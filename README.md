# zadanie

golang version - 1.21

Для запуска необходимо из директории проекта запустить команду:
make compose-up

## make comands:
make compose-up
Команда выполнит сборку исполняемого файла, соберет докер образ и запустит docker-compose среду из веб-сервера и БД postgres

make compose-down
Команда остановит выполнение compose среды и удалит контейнеры

Для сборки под образ alpine нужно выполнить команду:
make build-for-alpine

Для сборки под текущую ОС пользователя есть команда:
make build
 
Для запуска тестов:
make test

## swagger:
Сваггер доступен после запуска по ссылке:
http://localhost:8088/swagger/index.html