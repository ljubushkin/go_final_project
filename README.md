Описание проекта
Этот проект представляет собой веб-сервер на Go, который реализует функциональность планировщика задач (TODO-листа). Он позволяет пользователям добавлять, получать, изменять и удалять задачи, а также отмечать их как выполненные. Задачи могут иметь дату дедлайна и могут повторяться по заданным правилам.

Директория `web` содержит файлы фронтенда.

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.
 - Реализована задача со * по фунционалу рассчета даты повторения, поэтому в фале settings.go значение переменной FullNextDate = true, тип повторения можно выбирать не только ежедневно и ежегодно, но также по дням недели и месяца
 - Реализована задача со * по фунционалу поиска задач по тексту седержащемся в загловке либо в комментарии, значение переменной 
 Search = true
 - Реализована задача со * по аутентификации по паролю, в переменной Token установлен куки, генерируемый при вводе пароля

Реализованы переменные окружения, они хранятся в файле .env

Для запуска приложения, необходимо склонировать репозиторий, git clone https://github.com/ljubushkin/go_final_project.git и запустить командой go run .

По адресу http://localhost:7540/ будет доступно приложения, при первом входе необходимо ввести пароль, он указан в файле .env

При запущенном приложении можно запусить все тесты командой go test ./tests

Так же можно создать Docker образ командой docker-compose up -d, приложение также доступно по адресу http://localhost:7540/








