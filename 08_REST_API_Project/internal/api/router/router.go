package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	// eRouter := execsRouter()
	tRouter := teachersRouter()
	sRouter := studentsRouter()

	sRouter.Handle("/", execsRouter()) // --- OR ---
	// sRouter.Handle("/", eRouter)
	tRouter.Handle("/", sRouter)
	return tRouter

}
