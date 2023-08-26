package main

import (
	"github.com/modaniru/avito/internal/app"
)

// В качестве тестирующей базы данных была выбрана sqlite
// При удалении сегмента, мы отписываем всех пользователей от него и заносим в историю
func main() {
	app.App()
}
