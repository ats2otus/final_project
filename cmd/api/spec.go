package main

// Result - структура для ответов по bruteforce
type Result struct {
	Ok bool `json:"ok"`
}

// NoContent - пустая структура для ответов
type NoContent struct{}

// Error - для ответов с ошибками
type Error struct {
	Date  string `json:"date" format:"dateTime" example:"2006-01-02T15:04:05Z07:00"`
	Error string `json:"errors" example:"error description"`
}

// CheckItem - данные для проверки на bruteforce
type CheckItem struct {
	IP       string `json:"ip" example:"127.0.0.1"`
	Login    string `json:"login" example:"freak192"`
	Password string `json:"password" example:"pa$$W0rD"`
}

// ResetItem - данные для сброса бакетов
type ResetItem struct {
	IP    string `json:"ip" example:"127.0.0.1"`
	Login string `json:"login" example:"freak192"`
}

// ListItem - данные для работы с white/black листами
type ListItem struct {
	Subnet string `json:"subnet" example:"192.1.1.0/25"`
}
