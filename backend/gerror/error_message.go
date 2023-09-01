package gerror

func StatusText(errorCode uint32) string {
	switch errorCode {

	// Client-side
	case ErrorBindData:
		return "Failed to bind data"
	case ErrorValidData:
		return "Failed to valid data"

		// Server-side
	case ErrorConnect:
		return "Failed to connect database"
	case ErrorRetrieveData:
		return "Failed to retrieve data"
	case ErrorSaveData:
		return "Failed to save data"
	}
	return "Unknown Error"
}
