package service

// создаем интрфс с методом инсерт из файлов выше
// создаем стрктр Аудит с полем репо типа интерфейса выше
// создаем его конструктор
// 
// создаем метод Insert(ctx, req) 
// в котором создаем айтем типа audit.LogItem
// и в каждом поле используем соответствующий метод Гет с доп методом стринг кроме айди и таймстемпа
// в таймстемпе используем доп метод АсТайм
// 
// возвращаем инсерт из репозитория
// 
// 
// 
// 
// 