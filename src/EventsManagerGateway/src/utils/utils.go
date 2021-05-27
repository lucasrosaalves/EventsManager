package utils

import uuid "github.com/nu7hatch/gouuid"

func GenerateGuid() string {
	u, err := uuid.NewV4()

	if err != nil {
		return ""
	}

	return u.String()
}
