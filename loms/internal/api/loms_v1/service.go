package loms_v1

import (
	"route256/loms/internal/domain"
	desc "route256/loms/pkg/loms_v1"
)

type Implementation struct {
	desc.UnimplementedLomsServiceServer
	lomsService domain.Service
}

func NewLomsV1(lomsService domain.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedLomsServiceServer{},
		lomsService,
	}
}
