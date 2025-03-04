# test_galileosky
Тестовое задание
---
1\. Задача
 Необходимо разработать REST-сервис, который будет реализовывать следующие функции: - Добавление записи какой-либо сущности (например «задачи»). - Получение списка записей с возможностью сортировки по наименованию и фильтрацией по дате. - Произвольное изменение порядка следования записей. - Экспорт списка записей в формат XLSX.

---
2\. Объём данных
 Сервис должен обрабатывать 100 миллионов записей.

---
3\. Инфраструктура
 - Предоставить файл docker-compose для развертывания необходимой базы данных.
 - Предоставить миграцию, которая заполнит базу данными (сгенерирует записи) до требуемого объёма в 100 млн.

---
4\. Нагрузочное тестирование
 - Разработать и приложить нагрузочный тест, имитирующий одновременную работу 2000 пользователей с реализованными endpoint’ами.

---
5\. Технологии
 - Выберите любой удобный для вас язык программирования и необходимые библиотеки/фреймворки.

---
6\. Результат
 - Разместите программный код в любом доступном репозитории (например, GitHub, GitLab, Bitbucket) и пришлите ссылку на него.
