package external

type ApiResponse struct {
	statusCode int
	payload    interface{}
}

func (a *ApiResponse) IsValid() bool {
	return a.statusCode > 299
}
