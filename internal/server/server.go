package server

// создаем структуру сервер с двумя полями:
//	*gRPC сервер
//	audit.AuditServiceServer
// создаем его конструктор, в которой дается второй параметр
// 
// создаем метод ListenAndServe в который дается порт и возвращается ошибка
// создаем аддр в который спринтуем порт
// отдаем этот порт в метод Листен библы net вторым параметром, а первым указываем протокол tcp
// вызываем у библы audit метод РегистрацияАудитСервисаСервер, в который передаем два сервера
// вызываем у grps сервера метод Serve и передаем переменную предпоследнюю