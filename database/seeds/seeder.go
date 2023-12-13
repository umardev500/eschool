package seeds

import "github.com/umardev500/eschool/config/postgres"

type Seeder struct {
	trx *postgres.Trx
}

func NewSeeder(trx *postgres.Trx) *Seeder {
	return &Seeder{
		trx: trx,
	}
}
