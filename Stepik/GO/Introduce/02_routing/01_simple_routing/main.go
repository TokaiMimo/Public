package main

import (
	"net/http"
)

// Создание структуры Handler
type MyHandler struct {
}

//		метод MyHandler			ResponseWriter и Request лежат в пакете http, обязательно
//	                         нужно добавить имена аргументов, к примеру w,r
func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Есть Request у которого есть вызываемые поля, в зависимости от поля мы
	//можем возвращать разный результат, если обратились в корень, то напишем Home
	if r.URL.Path == "/" {
		//чтобы вернуть какой-то результат в браузер вызываем метод Write,
		// Write принимает слайс байтов, преобразует "Hello World!" в слайс байтов
		w.Write([]byte("Home"))
		return
	}
	if r.URL.Path == "/hello" {
		w.Write([]byte("Hello, user"))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Page Not Found"))
}

func main() {

	//пакет http содержит метод ListenAndServe, первый аргумент, это адрес на котором он будет слушать, если адрес не указан
	//, то будет на всех слушать, второй аргумент - функция Handler
	// Handler - это интерфейс, который реализует одну функцию
	//ServeHTTP(w http.ResponseWriter, r *http.Request) , которая принимет
	//ResponseWriter - это интерфейс с тремя методами,
	//Header() Header
	//Write([]byte) (int, error)
	//WriteHeader(statusCode int)
	// И принимает Request - это отдельная структура с полями metod,header,body и так далее

	//передадим нашу Функцию Handler
	err := http.ListenAndServe(":3000", MyHandler{})
	//http.ListenAndServe возвращает ошибку, получаем значение
	//проверяем на nil и вызовем панику
	//если не выводить ошибку, не поймем почему программа завершилась
	if err != nil {
		panic(err)
	}
}
