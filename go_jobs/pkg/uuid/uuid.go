package uuid

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	uid := uuid.Must(uuid.NewV4(),nil)
	return uid.String()
}
