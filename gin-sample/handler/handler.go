package handler

var (
	defaultUserHandler *userHandler
)

func Init() error {
	defaultUserHandler = NewUserHandler()
	return nil
}

func GetUserHandler() *userHandler {
	return defaultUserHandler
}
