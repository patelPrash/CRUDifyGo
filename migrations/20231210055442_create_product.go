package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20231210055442 struct {
}

func (k K20231210055442) Up(d *datastore.DataStore, logger log.Logger) error {
	if _, err := d.DB().Exec(createProductTable); err != nil {
		return err
	}

	return nil
}

func (k K20231210055442) Down(d *datastore.DataStore, logger log.Logger) error {
	if _, err := d.DB().Exec(dropProductTable); err != nil {
		return err
	}

	return nil
}
