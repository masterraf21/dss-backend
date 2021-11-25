package models

func ValidateMenuBody(body MenuBody) bool {
	return !((body.Name == "") || (body.Calorie == 0))
}
