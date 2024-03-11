package main

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	var _uuid string
	for {
		uuidwithhyphen := uuid.New()
		_uuid = uuidwithhyphen.String()[:8]
		exists, err := CheckShortURLExists(_uuid)
		if err != nil {
			panic(err)
		}
		if !exists {
			break
		}
	}
	return _uuid, nil
}
