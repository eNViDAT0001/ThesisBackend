package io

import "github.com/eNViDAT0001/Thesis/Backend/external/paging"

type ListNotifyInput struct {
	UserID uint
	Paging *paging.ParamsInput
}
