package main

import (
	"github.com/n4mlz/sns-backend/internal/repository"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/repository/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	db, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}

	g.UseDB(db)
	all := g.GenerateAllTable()

	g.ApplyBasic(all...)

	g.Execute()
}
