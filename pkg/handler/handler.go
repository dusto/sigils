package handler

import (
	"github.com/dusto/sigils/pkg/respository"
)

type Handler struct {
	configDB respository.DBTX
}
