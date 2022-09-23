package serializer

import "net/http"

func HttpError(code int) Response {
	return BuildResponse(code, http.StatusText(code), nil)
}

func NotFoundError() Response {
	return HttpError(http.StatusNotFound)
}

func InternalServerError() Response {
	return HttpError(http.StatusInternalServerError)
}

func BadRequestError() Response {
	return HttpError(http.StatusBadRequest)
}
