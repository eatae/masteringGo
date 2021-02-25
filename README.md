# Mastering-Go

## Go и операционная система (ch01)

* Читаем float из STDIN<br>
[cla](https://github.com/eatae/masteringGo/blob/master/ch01/cla.go) <br>
     
* Пишем в файл (логи) <br>
[customLog](https://github.com/eatae/masteringGo/blob/master/ch01/customLog.go)

* Использование внешнего пакета <br>
[getPackage](https://github.com/eatae/masteringGo/blob/master/ch01/getPackage.go)
    
* Пишем в системный LogFatal <br>
[logFatal](https://github.com/eatae/masteringGo/blob/master/ch01/logFatal.go)
    
* Пишем в системный LogFile <br>
[logFiles](https://github.com/eatae/masteringGo/blob/master/ch01/logFiles.go)
    
* Пишем в системный LogPanic <br>
[logPanic](https://github.com/eatae/masteringGo/blob/master/ch01/logPanic.go)
    
* Эмулируем свой Error <br>
[newError](https://github.com/eatae/masteringGo/blob/master/ch01/newError.go)
    
* Вывод ошибок в STDERR (поток ошибок) <br>
[stdERR](https://github.com/eatae/masteringGo/blob/master/ch01/stdERR.go)
    
* Получение данных от пользователя (читаем STDIN)<br>
[stdIN](https://github.com/eatae/masteringGo/blob/master/ch01/stdIN.go)
    
* Пишем в STDOUT <br>
[stdOUT](https://github.com/eatae/masteringGo/blob/master/ch01/stdOUT.go)
    
    
## Go изнутри (ch02)

* defer (стек наглядно)<br>
[defer](https://github.com/eatae/masteringGo/blob/master/ch02/defer.go)
    
* Смотрим затраты памяти <br>
[gColl](https://github.com/eatae/masteringGo/blob/master/ch02/gColl.go)
    
* Некоторые переменные окружения (системные) <br>
[goEnv](https://github.com/eatae/masteringGo/blob/master/ch02/goEnv.go)
    
* Простая Panic <br>
[justPanic](https://github.com/eatae/masteringGo/blob/master/ch02/justPanic.go)

* Интересная разбивка **Map** через % остаток<br>
[mapSplit](https://github.com/eatae/masteringGo/blob/master/ch02/mapSplit.go)

* **Пакет Unsafe** - получаем адрес Pointer и размер <br>
[moreUnsafe](https://github.com/eatae/masteringGo/blob/master/ch02/moreUnsafe.go)

* panic()  recover() <br>
[panicRecover](https://github.com/eatae/masteringGo/blob/master/ch02/panicRecover.go)

* Получаем число из строки (Atoi) <br>
[requiredVersion](https://github.com/eatae/masteringGo/blob/master/ch02/requiredVersion.go)

* Инфрмация о текущей среде Go (пакет runtime) <br>
[goEnv](https://github.com/eatae/masteringGo/blob/master/ch02/goEnv.go) <br>
[requiredVersion](https://github.com/eatae/masteringGo/blob/master/ch02/requiredVersion.go)



## Работа с основными типами данных (ch03)

* Примеры цикла for <br>
[loops](https://github.com/eatae/masteringGo/blob/master/ch03/loops.go)

* Массивы (многомерные) и итерация по ним <br>
[usingArrays](https://github.com/eatae/masteringGo/blob/master/ch03/usingArrays.go)

* Слайсы, копирование, make() <br>
[slices](https://github.com/eatae/masteringGo/blob/master/ch03/slices.go) <br>
[reslice](https://github.com/eatae/masteringGo/blob/master/ch03/reslice.go)

* Длина и вместимость Слайса <br>
[lenCap](https://github.com/eatae/masteringGo/blob/master/ch03/lenCap.go)

* Функция copy() для Слайсов <br>
[copySlice](https://github.com/eatae/masteringGo/blob/master/ch03/copySlice.go)

* Сортировка Слайса sort.Slice() <br>
[sortSlice](https://github.com/eatae/masteringGo/blob/master/ch03/sortSlice.go)

* Добавление существующего Массива к существующему Слайсу <br>
[appendArrayToSlice](https://github.com/eatae/masteringGo/blob/master/ch03/appendArrayToSlice.go)

* Хэш таблицы (map), создание, удаление, итерация <br>
[usingMaps](https://github.com/eatae/masteringGo/blob/master/ch03/usingMaps.go)

* Запись в хеш таблицу со значением nil невозможно <br>
[failMap](https://github.com/eatae/masteringGo/blob/master/ch03/failMap.go)

* Константы, генератор констант Iota <br>
[constants](https://github.com/eatae/masteringGo/blob/master/ch03/constants.go)

* Указатели, создание, разыменование <br>
[pointers](https://github.com/eatae/masteringGo/blob/master/ch03/pointers.go)

* Время и дата, примеры использования <br>
[usingMaps](https://github.com/eatae/masteringGo/blob/master/ch03/usingMaps.go)

* Парсинг времени <br>
[parseTime](https://github.com/eatae/masteringGo/blob/master/ch03/parseTime.go)

* Парсинг даты <br>
[parseDate](https://github.com/eatae/masteringGo/blob/master/ch03/parseDate.go)

* Изменение формата даты и времени <br>
[timeDate](https://github.com/eatae/masteringGo/blob/master/ch03/timeDate.go)

* Замер времени выполнения программы <br>
[execTime](https://github.com/eatae/masteringGo/blob/master/ch03/execTime.go)



## Использованеи составных типов данных (ch04)

* Структуры, разные способы создания <br>
[structures](https://github.com/eatae/masteringGo/blob/master/ch04/structures.go)
    
* Создание Указателя на Структуру <br>
[pointerStruct](https://github.com/eatae/masteringGo/blob/master/ch04/pointerStruct.go)
    
* Кортежи, возврат нескоьких значений <br>
[tuples](https://github.com/eatae/masteringGo/blob/master/ch04/tuples.go)
    
* RegExp, чтение файла построчно, нахождение подстроки <br>
[selectColumn](https://github.com/eatae/masteringGo/blob/master/ch04/selectColumn.go)

* RegExp, сопоставление даты и времени <br>
[changeDT](https://github.com/eatae/masteringGo/blob/master/ch04/changeDT.go)
    
* RegExp, проверка IPv4-адресов <br>
[findIPv4](https://github.com/eatae/masteringGo/blob/master/ch04/findIPv4.go)
    
* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)
    
* Название <br>
[ссылка](адрес)
    
* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)
    
* Название <br>
[ссылка](адрес)
    
* Название <br>
[ссылка](адрес)



## Структуры данных (ch05)

* Двоичное дерево <br>
[binTree](https://github.com/eatae/masteringGo/blob/master/ch05/binTree.go)
    
* Хэш-таблица <br>
[hashTable](https://github.com/eatae/masteringGo/blob/master/ch05/hashTable.go)
    
* Поиск по Хэш-таблице <br>
[hashTableLookup](https://github.com/eatae/masteringGo/blob/master/ch05/hashTableLookup.go)
    
* Связные списки <br>
[linkedList](https://github.com/eatae/masteringGo/blob/master/ch05/linkedList.go)

* Двусвязные списки <br>
[doublyLList](https://github.com/eatae/masteringGo/blob/master/ch05/doublyLList.go)
    
* Очереди (queue) <br>
[queue](https://github.com/eatae/masteringGo/blob/master/ch05/queue.go)
    
* Стеки (stack) <br>
[stack](https://github.com/eatae/masteringGo/blob/master/ch05/stack.go)

##### Пакет Container <br>
 
* Сontainer/heap (реализация интерфейса кучи) <br>
[conHeap](https://github.com/eatae/masteringGo/blob/master/ch05/conHeap.go)

* Container/list (список) <br>
[conList](https://github.com/eatae/masteringGo/blob/master/ch05/conList.go)
    
* Container/ring (Кольцо) <br>
[conRing](https://github.com/eatae/masteringGo/blob/master/ch05/conRing.go)

##### Генерация чисел и строк <br>

* Генерация случайных чисел <br>
[randomNumbers](https://github.com/eatae/masteringGo/blob/master/ch05/randomNumbers.go)
    
* Генерация случайных строк <br>
[generatePassword](https://github.com/eatae/masteringGo/blob/master/ch05/generatePassword.go)
    
* Генерация безопасной последовательности случайных чисел <br>
[cryptoRand](https://github.com/eatae/masteringGo/blob/master/ch05/cryptoRand.go)

##### Выполнение матричных вычислений <br>

* Сложение, вычитание матриц <br>
[addMat](https://github.com/eatae/masteringGo/blob/master/ch05/addMat.go)

* Умножение матриц <br>
[mulMat](https://github.com/eatae/masteringGo/blob/master/ch05/mulMat.go)

* Деление матриц <br>
[divMat](https://github.com/eatae/masteringGo/blob/master/ch05/divMat.go)

* Как определить размер массива <br>
[dimensions](https://github.com/eatae/masteringGo/blob/master/ch05/dimensions.go)

* Разгадывание Судоку <br>
[sudoku](https://github.com/eatae/masteringGo/blob/master/ch05/sudoku.go)



## Неочевидные знания о пакетах и функциях Go (ch06)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)

* Название <br>
[ссылка](адрес)