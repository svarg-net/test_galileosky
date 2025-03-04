package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host              = "localhost"
	port              = 5432
	user              = "user"
	password          = "password"
	dbname            = "tasks_db"
	maxGoroutines     = 20    // Максимальное количество Goroutines
	batchSize         = 63000 // Размер пакета для каждой Goroutine
	maxParamsPerQuery = 63000 // Максимум 6000 параметров на один запрос
)

func main() {
	if false {
		// Начало отсчета времени выполнения
		startTime := time.Now()

		// Имя выходного файла
		filename := "./migration/data/tasks.csv"

		// Открываем файл для записи
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
			return
		}
		defer file.Close()

		// Создаем буферизованный writer для оптимизации записи
		writer := bufio.NewWriter(file)

		// Записываем заголовок CSV
		if _, err := writer.WriteString("name,description,created_at,order_index\n"); err != nil {
			fmt.Println("Ошибка при записи заголовка:", err)
			return
		}

		// Генерация 100 миллионов записей
		totalRecords := 100_000_000
		for i := 1; i <= totalRecords; i++ {
			name := fmt.Sprintf("Task_%d", i)
			description := fmt.Sprintf("Description for task %d", i)
			createdAt := time.Now().Format(time.RFC3339) // Текущее время в формате ISO 8601
			orderIndex := strconv.Itoa(i)

			// Формируем строку CSV
			line := fmt.Sprintf("%s,%s,%s,%s\n", name, description, createdAt, orderIndex)

			// Записываем строку в файл
			if _, err := writer.WriteString(line); err != nil {
				fmt.Println("Ошибка при записи строки:", err)
				return
			}

			// Периодическая очистка буфера для освобождения памяти
			if i%1000_000 == 0 {
				writer.Flush()
				fmt.Printf("Записано %d записей...\n", i)
			}
		}

		// Очистка буфера перед закрытием файла
		writer.Flush()

		// Время завершения
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		fmt.Printf("Файл '%s' успешно создан. Время выполнения: %v\n", filename, duration)
	}
	if true {
		startTime := time.Now()
		db()
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		fmt.Printf("Время выполнения: %v\n", duration)
	}

}

func db() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось проверить соединение с базой данных: %v", err)
	}

	// Путь к файлу CSV
	filePath := "/docker-entrypoint-initdb.d/data/tasks.csv"

	// Вызов функции для импорта данных
	err = copyDataToPostgres(db, filePath)
	if err != nil {
		log.Fatalf("Ошибка при импорте данных: %v", err)
	}
}
func copyDataToPostgres(db *sql.DB, filePath string) error {
	// SQL-команда COPY
	query := fmt.Sprintf(`
BEGIN;
SET LOCAL synchronous_commit TO OFF;
    COPY tasks (name, description, created_at, order_index) 
    FROM '%s' 
    WITH (FORMAT CSV, HEADER TRUE, DELIMITER ',');
COMMIT;
    `, filePath)

	// Выполняем команду COPY
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении COPY: %w", err)
	}

	fmt.Println("Данные успешно импортированы в таблицу tasks.")
	return nil
}
