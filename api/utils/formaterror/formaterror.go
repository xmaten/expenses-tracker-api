package formaterror

import "strings"

var errorMessages = make(map[string]string)

var err error

func FormatError(errString string) map[string]string {
	if strings.Contains(errString, "username") {
		errorMessages["Taken_username"] = "Username already taken"
	}

	if strings.Contains(errString, "email") {
		errorMessages["Taken_email"] = "Email already taken"
	}

	if strings.Contains(errString, "hashedPassword") {
		errorMessages["Incorrect_password"] = "Incorrect password"
	}

	if strings.Contains(errString, "record not found") {
		errorMessages["No_record"] = "No record found"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect details"
		return errorMessages
	}

	return nil
}