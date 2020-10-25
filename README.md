# wb-test

Решение, случайно нарытого в интернете, тестового задания Wildberries по Golang.

## Задание

Программа читает из stdin строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех переданных URL, например:

```
$ cat urls | go run ./cmd/wb-test/main.go
https://golang.org/: 9
https://golang.org/doc/: 64
https://golang.org/compress/: 6
...
...
Total: 253
```

Каждый URL должен начать обрабатываться сразу после вычитывания и параллельно с вычитыванием следующего. URL должны обрабатываться параллельно, но не более k=5 одновременно. Обработчики URL не должны порождать лишних горутин, т.е. если k=5, а обрабатываемых URL-ов всего 2, не должно создаваться 5 горутин.

Нужно обойтись без глобальных переменных и использовать только стандартную библиотеку.
