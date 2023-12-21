# Постановка задачи

Написать json REST сервер с двумя ручками:
- `POST /edit/:Id` - изменение новости по Id
- `GET /list` - список новостей

БД - mysql (можно и postgres, мы перешли на него).

В качестве сервера использовать fiber. Для работы с базой reform.

Соединение с базой должно использовать connection pool. Все настройки через переменные окружения и/или viper.

БД:
```SQL


--
-- Структура таблицы `News`
--

CREATE TABLE `News` (
  `Id` bigint NOT NULL,
  `Title` tinytext NOT NULL,
  `Content` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `NewsCategories`
--

CREATE TABLE `NewsCategories` (
  `NewsId` bigint NOT NULL,
  `CategoryId` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `News`
--
ALTER TABLE `News`
  ADD PRIMARY KEY (`Id`);

--
-- Индексы таблицы `NewsCategories`
--
ALTER TABLE `NewsCategories`
  ADD PRIMARY KEY (`NewsId`,`CategoryId`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `News`
--
ALTER TABLE `News`
  MODIFY `Id` bigint NOT NULL AUTO_INCREMENT;
```

Т.е. легко видеть, что связка новостей и категорий идёт в отдельной таблице.

Формат входных данных для первой ручки:

```json
{
  "Id": 64,
  "Title": "Lorem ipsum",
  "Content": "Dolor sit amet <b>foo</b>",
  "Categories": [1,2,3]
}
```

При этом, если какое-то из полей не задано - это поле не должно быть обновлено.

Формат данных на выходе list:

```json
{
    "Success": true,
    "News": [
      {
        "Id": 64,
        "Title": "Lorem ipsum",
        "Content": "Dolor sit amet <b>foo</b>",
        "Categories": [1,2,3]
      },
      {
        "Id": 1,
        "Title": "first",
        "Content": "tratata",
        "Categories": [1]
      }
    ]
}
```

## Требования и пожелания

1. Срок выполнения - не больше трёх дней. Мы высылаем тестовые задания пачками сразу десяткам людей, поэтому вас просто могут опередить другие кандидаты.
2. Если знакомы с docker - хотелось бы посмотреть упаковку сервиса в контейнер.
3. Дополнительно хотелось бы увидеть(плюс в карму по сравнению с другими кандидатами):
- авторизацию через Authorization заголовок и грамотную структуризацию кода и ручек по группам/папкам;
- валидацию полей при редактировании;
- пагинацию на ручке списка;
- грамотное логгирование с использованием любого популярного логгера(напр. logrus);
- грамотную обработку ошибок.