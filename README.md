# random-numbers-generator

Программа представляет собой многопоточный генератор уникальных случайных чисел. Флаги: bound - верхняя граница генерации, flows - количество потоков. 
Срез уникальных случайных чисел будет сгенерирован в диапазоне [1; bound], произведена сортировка данных, подсчет времени генерации среза и вывод данных на экран.
В последней версии добавлен http-сервер, адрес - localhost:3000. 
При get запросе возвращается JSON файл, содержащий информацию об ошибке или успешно сгенерированные данные.
