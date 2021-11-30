### Сборка приложения (Build Info)

При сборке приложения будут автоматически определяться (вшиты в исходник) ряд переменных:
+ Репозиторий (ссылка) - переменная `Repo`
+ Имя ветки с исходным кодом - переменная `Branch`
+ Commit из которого происходит сборка - переменная `Commit`
+ Имя компонента - переменная `Name`
+ Время сборки компонента - переменная `Time`

В коде такие переменные объявлять так (**пакет main**):
```golang
var (
	Name   string = "API"
	Repo   string
	Branch string
	Commit string
	Time   string
)
```
Для такой сборки см. файл `Makefile` и весь исходным код с примером в данном репозитории.
Данная сборка будет запускать автоматически в Jenkins.

### Конфигурация приложения
Предлагаю обозначать как в примере у меня в репозитории (main.go)

### Запуск приложения
Весь запуск можно разбить на 4 этапа:
1. При запуске в самом начале распечатываем **Build Info** в stdout;
2. Загружаем весь **Config** и распечатываем его в stdout;
3. Проверяем все зависимости от других компонентов и распечатываем каждый статус в stdout;
4. Основной запуск приложения и далее применимаем весь трафик.
