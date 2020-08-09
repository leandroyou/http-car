package repository

import "errors"

var ErrNoCar = errors.New("noCar: no cars found in hotStorage")

type IRepository interface {
	ICarRepository
}
