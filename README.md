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

##### Структуры <br>

* Структуры, разные способы создания <br>
[structures](https://github.com/eatae/masteringGo/blob/master/ch04/structures.go)
    
* Создание Указателя на Структуру <br>
[pointerStruct](https://github.com/eatae/masteringGo/blob/master/ch04/pointerStruct.go)
    
* Кортежи, возврат нескоьких значений <br>
[tuples](https://github.com/eatae/masteringGo/blob/master/ch04/tuples.go)


##### RegExp <br>

* RegExp, чтение файла построчно, нахождение подстроки <br>
[selectColumn](https://github.com/eatae/masteringGo/blob/master/ch04/selectColumn.go)

* RegExp, сопоставление даты и времени <br>
[changeDT](https://github.com/eatae/masteringGo/blob/master/ch04/changeDT.go)
    
* RegExp, проверка IPv4-адресов <br>
[findIPv4](https://github.com/eatae/masteringGo/blob/master/ch04/findIPv4.go)


##### Строки <br>

* Создание строки, итерация <br>
[strings](https://github.com/eatae/masteringGo/blob/master/ch04/strings.go)

* Руны (строка - множество рун) <br>
[ссылка](https://github.com/eatae/masteringGo/blob/master/ch04/runes.go)

* Какие символы являются рунами (пакет unicode) <br>
[unicode](https://github.com/eatae/masteringGo/blob/master/ch04/unicode.go)
    
* Манипуляция строками (пакет string) <br>
[useStrings](https://github.com/eatae/masteringGo/blob/master/ch04/useStrings.go)
    
* switch, string + RegExp <br>
[switch](https://github.com/eatae/masteringGo/blob/master/ch04/switch.go)

*  Хранилище Key:Value. В цикле слушаем StdIn и выполняем комманды <br>
[ссылка](адрес)


##### JSON <br>

* Читаем JSON из файла и Декодируем во вложенные структуры <br>
[readJSON](https://github.com/eatae/masteringGo/blob/master/ch04/readJSON.go)

* Пишем наши структуры в виде JSON в файл <br>
[writeJSON](https://github.com/eatae/masteringGo/blob/master/ch04/writeJSON.go)
    
* функции Marshal() и Unmarshal() (удобнее чем encode/decode) <br>
[сmUJSON](https://github.com/eatae/masteringGo/blob/master/ch04/mUJSON.go)
    
* Чтение и запись JSON без структурирования (используется Map) <br>
[parsingJSON](https://github.com/eatae/masteringGo/blob/master/ch04/parsingJSON.go)


##### XML <br>

* Структуру переводим в XML <br>
[rwXML](https://github.com/eatae/masteringGo/blob/master/ch04/rwXML.go)

* Чтение XML-файла <br>
[readXML](https://github.com/eatae/masteringGo/blob/master/ch04/readXML.go)
    
* Поля структуры как XML данные <br>
[modXML](https://github.com/eatae/masteringGo/blob/master/ch04/modXML.go)




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

##### Функции <br>

* Функции которые возвращают несколько значений <br>
[functions](https://github.com/eatae/masteringGo/blob/master/ch06/functions.go)

* Функции, возвращающие именованные значения <br>
[returnNames](https://github.com/eatae/masteringGo/blob/master/ch06/returnNames.go)

* Функции, принимающие указатели <br>
[ptrFun](https://github.com/eatae/masteringGo/blob/master/ch06/ptrFun.go)

* Функции, которые возвращают указатели <br>
[returnPtr](https://github.com/eatae/masteringGo/blob/master/ch06/returnPtr.go)

* Функции, которые возвращают другие функции <br>
[returnFunction](https://github.com/eatae/masteringGo/blob/master/ch06/returnFunction.go)

* Функции, которые принимают другие функции в качестве параметров <br>
[funFun](https://github.com/eatae/masteringGo/blob/master/ch06/funFun.go)

* Функции с переменным числом параметров <br>
[variadic](https://github.com/eatae/masteringGo/blob/master/ch06/variadic.go)


##### Разработка Go-пакетов <br>

* Создание и использование пакета <br>
[aPackage](https://github.com/eatae/masteringGo/blob/master/ch06/aPackage.go) <br>
[useAPackage](https://github.com/eatae/masteringGo/blob/master/ch06/useAPackage.go) <br>
    Install:
    >mkdir ~/go/src/aPackage <br>
    cp aPackage.go ~/go/src/aPackage/ <br>
    go install aPackage <br>
    >
    Compile:
    >go tool compile aPackage.go
    >
                                                                                                                                                                                                                                                                                                                                                        
* Функция init() <br>
[a](https://github.com/eatae/masteringGo/blob/master/ch06/a.go) <br>
[b](https://github.com/eatae/masteringGo/blob/master/ch06/b.go) <br>
[manyInit](https://github.com/eatae/masteringGo/blob/master/ch06/manyInit.go) <br>
    Install:
    >mkdir ~/go/src/a <br>
     cp a.go ~/go/src/a/ <br>
     go install a <br>
     mkdir ~/go/src/b <br>
     cp b.go ~/go/src/b/ <br>
     go install b <br>

##### Модули <br>

* Читаем книгу (стр.290) <br>


##### Некоторые пакеты Go (scanner/parser/token/ast)<br>

* Пакет syscall <br>
[useSyscall](https://github.com/eatae/masteringGo/blob/master/ch06/useSyscall.go)

* Встроенный пакет go/scanner (синтаксический сканер программы на Go) <br>
[goScanner](https://github.com/eatae/masteringGo/blob/master/ch06/goScanner.go)

* Встроенный пакет go/parser (синтаксический анализ программы на Go) <br>
[goParser](https://github.com/eatae/masteringGo/blob/master/ch06/goParser.go)

* Парсим файл Go, подсчитываем кол-во яз.конструкций var <br>
[varTimes](https://github.com/eatae/masteringGo/blob/master/ch06/varTimes.go)

* Поиск имен переменных заданной длины (сложно, куча вложенных if) <br>
[varSize](https://github.com/eatae/masteringGo/blob/master/ch06/varSize.go)


##### Шаблонизатор для текста, HTML <br>

* Выод простого текста <br>
[textT](https://github.com/eatae/masteringGo/blob/master/ch06/textT.go) <br>
[text.gotext](https://github.com/eatae/masteringGo/blob/master/ch06/text.gotext)

* HTML template + DataBase <br>
[htmlT](https://github.com/eatae/masteringGo/blob/master/ch06/htmlT.go) <br>
[html.gohtml](https://github.com/eatae/masteringGo/blob/master/ch06/html.gohtml)



## Рефлексия и интерфейсы на все случае жизни (ch07)

##### Свои типы, интерфейсы и проверка типа <br>

* Создаём структуру (тип) и метод к ней <br>
[methods](https://github.com/eatae/masteringGo/blob/master/ch07/methods.go)

* Проверка типа <br>
[assertion](https://github.com/eatae/masteringGo/blob/master/ch07/assertion.go)

* Создание интерфейс / использование интерфейса <br>
[myInterface](https://github.com/eatae/masteringGo/blob/master/ch07/myInterface.go) <br>
[useInterface](https://github.com/eatae/masteringGo/blob/master/ch07/useInterface.go)

* Switch для типов данных <br>
[switch](https://github.com/eatae/masteringGo/blob/master/ch07/switch.go)


##### Рефлексия <br>

* Простой пример <br>
[reflection](https://github.com/eatae/masteringGo/blob/master/ch07/reflection.go)

* Более сложный пример <br>
[advRefl](https://github.com/eatae/masteringGo/blob/master/ch07/advRefl.go)

* Библиотека reflectwalk (рефлексия) <br>
[walkRef](https://github.com/eatae/masteringGo/blob/master/ch07/walkRef.go)

##### ООП (композиция) <br>

* Пример ООП <br>
[ooo](https://github.com/eatae/masteringGo/blob/master/ch07/ooo.go) <br>
[goCoIn](https://github.com/eatae/masteringGo/blob/master/ch07/goCoIn.go)





## Как объяснить UNIX-системе, что она должна делать (ch08)

##### Аргументы командной строки <br>

* Пакет flag <br>
[ссылка](https://github.com/eatae/masteringGo/blob/master/ch08/simpleFlag.go) <br>
[funWithFlag](https://github.com/eatae/masteringGo/blob/master/ch08/funWithFlag.go) <br>

* Пакет viper (директория) <br>
[viper](https://github.com/eatae/masteringGo/tree/master/ch08/viper)

* Пакет cobra (директория) <br>
[cobra](https://github.com/eatae/masteringGo/tree/master/ch08/cobra)


##### Интерфейсы io.Reader и io.Writer (Чтение / Запись в файл) <br>

* Посотрочное чтение текстового файла (пакет bufio) <br>
[byLine](https://github.com/eatae/masteringGo/blob/master/ch08/byLine.go)

* Чтение текстового файла по словам (пакет bufio, regexp) <br>
[byWord](https://github.com/eatae/masteringGo/blob/master/ch08/byWord.go)

* Чтение текстового файла по символам <br>
[byCharacter](https://github.com/eatae/masteringGo/blob/master/ch08/byCharacter.go)

* Чтение бинарных данных из /dev/random (пакет encoding/binary) <br>
[devRandom](https://github.com/eatae/masteringGo/blob/master/ch08/devRandom.go)

* Чтение заданного кол-ва данных <br>
[readSize](https://github.com/eatae/masteringGo/blob/master/ch08/readSize.go)

* Чтение CSV файла и создание PNG <br>
[CSVplot](https://github.com/eatae/masteringGo/blob/master/ch08/CSVplot.go)

* Запись в файл (несколько способов) <br>
[save](https://github.com/eatae/masteringGo/blob/master/ch08/save.go)

* Загрузка и сохранение данных на диске<br>
CRUD <br>
сереализация - encoding/gob <br>
[kvSaveLoad](https://github.com/eatae/masteringGo/blob/master/ch08/kvSaveLoad.go)

* Пакет strings - ввод и вывод данных в файл <br>
[str](https://github.com/eatae/masteringGo/blob/master/ch08/str.go)

* Пакет bytes <br>
сохранение в буфер bytes.Buffer <br>
работа с байтовыми срезами <br>
[bytes](https://github.com/eatae/masteringGo/blob/master/ch08/bytes.go)

* Полномочия доступа к файлам <br>
смотрим доступы к файлу <br>
[permissions](https://github.com/eatae/masteringGo/blob/master/ch08/permissions.go)

* Обработка сигналов в UNIX <br>
пакет os/signal <br>
обрабатываем сигналы передавая их  в процесс <br>
[handleTwo](https://github.com/eatae/masteringGo/blob/master/ch08/handleTwo.go)

* Обработка всех сигналов системы <br>
так же передаём сигналы в процесс <br>
[handleAll](https://github.com/eatae/masteringGo/blob/master/ch08/handleAll.go)

* Реализация утилиты cat(1) на Go <br>
вывод содержимого файла в стандартный поток вывода <br>
[cat](https://github.com/eatae/masteringGo/blob/master/ch08/cat.go)

* Структура syscall.PtraceRegs <br>
вызов внешней команды для дочерних процессов <br>
[ptraceRegs](https://github.com/eatae/masteringGo/blob/master/ch08/ptraceRegs.go)

* Отслеживание системных вызовов <br>
[traceSyscall](https://github.com/eatae/masteringGo/blob/master/ch08/traceSyscall.go)

* Идентификаторы пользователя и группы <br>
пакеты os, os/user <br>
находим id пользователя системы и групп к которой принадлежит пользователь <br>
[ids](https://github.com/eatae/masteringGo/blob/master/ch08/ids.go)




## Конкурентность в Go: горутины, каналы, и конвейеры  (ch09)

##### Горутины <br>

* Создание горутины <br>
[simple](https://github.com/eatae/masteringGo/blob/master/ch09/simple.go)

* Создание нескольких горутин <br>
[create](https://github.com/eatae/masteringGo/blob/master/ch09/create.go)

* Как дождаться завершения горутин <br>
использование sync.WaitGroup <br>
кол-во Add() и Done() должно совпадать <br>
[syncGo](https://github.com/eatae/masteringGo/blob/master/ch09/syncGo.go)

##### Каналы <br>

* Запись в канал <br>
после закрытия канала горутина остановит своё выполнение <br>
[writeCh](https://github.com/eatae/masteringGo/blob/master/ch09/writeCh.go)

* Чтение из канала <br>
вторая попытка чтения не выходит, так как канал уже закрыт <br>
[readCh](https://github.com/eatae/masteringGo/blob/master/ch09/readCh.go)

* Прием данных из закрытого канала <br>
что при чтении данных из закрытого канала возвращается нулевое значение соответствующего типа <br>
[readClose](https://github.com/eatae/masteringGo/blob/master/ch09/readClose.go)

* Каналы как аргументы функции <br>
func someFunc(out chan<- int64, in <-chan int64) <br>

##### Конвейеры <br>

* Конвейер <br>
выходные данные одной горутины становятся входными данными для другой горутины <br>
для передачи данных используются каналы <br>
три горутины передают между собой данные <br>
в аргументах можно указать на запись или чтение канал, а сам канал может быть двусторонний <br> 
[pipeline](https://github.com/eatae/masteringGo/blob/master/ch09/pipeline.go)

##### Сосотояние гонки <br>

* Выявление состояния гонки <br>
go run -race pipeline.go 1 10 <br>
здесь есть общие переменные которые используют разные горутины <br>
[pipeline](аhttps://github.com/eatae/masteringGo/blob/master/ch09/pipeline.go)





## Конкурентность в Go: горутины, расширенные возможности  (ch10)

* Переменная среды GOMAXPROCS <br>
кол-во потоков которые могут одновременно выполняться <br>
по умолчанию равно кол-ву процессоров <br>
[maxprocs](https://github.com/eatae/masteringGo/blob/master/ch10/maxprocs.go)

* Ключевое слово select <br>
select похож на switch, но только для каналов <br>
можно запустить в цикле, тогда select будет работать пока не вызовим return <br>
[select](https://github.com/eatae/masteringGo/blob/master/ch10/select.go)

* Принудительное ожидание завершения горутины 1 <br>
просто с помощью time.Sleep() или time.After() <br>
[timeOut1](https://github.com/eatae/masteringGo/blob/master/ch10/timeOut1.go)

* Принудительное ожидание завершения горутины 2 <br>
с помощью sync.WaitGroup <br>
[timeOut2](https://github.com/eatae/masteringGo/blob/master/ch10/timeOut2.go)

##### И снова о Go-каналах

* Закрытие nil-channel вызовит panic <br>
var c chan string <br>
close(c) <br>

* Сигнальные каналы <br>
создаём последовательное выполнение горутин с помошью сигнальных каналов <br>
[defineOrder](https://github.com/eatae/masteringGo/blob/master/ch10/defineOrder.go)

* Буферизованные каналы <br>
интересный пример с counter <br>
несколько значений в канале как бы в очереди на чтение <br>
[bufChannel](https://github.com/eatae/masteringGo/blob/master/ch10/bufChannel.go)

* Нулевые каналы <br>
собый вид каналов, который всегда блокирован <br>
[nilChannel](https://github.com/eatae/masteringGo/blob/master/ch10/nilChannel.go)

* Каналы каналов <br>
[chSquare](https://github.com/eatae/masteringGo/blob/master/ch10/chSquare.go)

* Выбор последовательности исполнения горутин <br>
используются сигнальные каналы <br>
[defineOrder](https://github.com/eatae/masteringGo/blob/master/ch10/defineOrder.go)

* Как не надо использовать горутины <br>
наивный способ сортировки натуральных чисел с использованием горутин <br>
[sillySort](https://github.com/eatae/masteringGo/blob/master/ch10/sillySort.go)

##### Общая память и общие переменные (Mutex)

* Тип sync.Mutex <br>
блокировка общих ресурсов для определенной горутины в момент исполлнения <br>
позволяет избежать состояние гонки (race) <br>
sync.Lock() и sync.Unlock() <br>
[mutex](https://github.com/eatae/masteringGo/blob/master/ch10/mutex.go)

* Что произойдет, если вы забудете разблокировать мьютекс <br>
Locked! <br>
Fatal <br>
[forgetMutex](https://github.com/eatae/masteringGo/blob/master/ch10/forgetMutex.go)

* Тип sync.RWMutex <br>
считается улучшенной версией Mutex <br>
разделение блокировки на чтение и запись, то есть читать могут сразу несколько горутин <br>
[rwMutex](https://github.com/eatae/masteringGo/blob/master/ch10/rwMutex.go)

* Пакет atomic <br>
атомарная операция не может быть прервана на середине <br>
atomic иногда позволяет обойтись без использования мьютекса <br>
можно создавать атомные счетчики, доступные для нескольких горутин <br>
[atom](https://github.com/eatae/masteringGo/blob/master/ch10/atom.go)

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