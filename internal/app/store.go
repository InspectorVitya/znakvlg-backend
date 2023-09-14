package app

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
)

// todo
func (app *App) createStore(ctx context.Context, platesUse model.PlatesUse) ([]model.Store, error) {
	store := make([]model.Store, 0, len(platesUse))
	for i := range platesUse {
		store = append(store, model.Store{
			PlateTypeID: platesUse[i],
			Used:        true,
		})
	}
	return store, nil
}

//func storeFlatPlateFix(flat1, flat2, flat3, flat4 bool, store model.PlatesUse) {
//
//	for i := range store {
//		if ((store[i].PlateTypeID == 4) && store[i].Used) && flat1 {
//			store[i].Used = false
//		}
//		if ((store[i].PlateTypeID == 6) && store[i].Used) && flat2 {
//			store[i].Used = false
//		}
//		if ((store[i].PlateTypeID == 13) && store[i].Used) && flat3 {
//			store[i].Used = false
//		}
//		if ((store[i].PlateTypeID == 2) && store[i].Used) && flat4 {
//			store[i].Used = false
//		}
//	}
//}
