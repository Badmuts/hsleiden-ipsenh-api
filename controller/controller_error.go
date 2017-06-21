package controller

type ControllerError struct {
	Status           string `json:"status"`
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	MoreInfo         string `json:"moreInfo"`
}

func NewControllerError(status string, developerMessage string, userMessage string, moreInfo string) *ControllerError {
	ctrl := &ControllerError{
		Status:           status,
		DeveloperMessage: developerMessage,
		UserMessage:      userMessage,
		MoreInfo:         moreInfo,
	}
	return ctrl
}
