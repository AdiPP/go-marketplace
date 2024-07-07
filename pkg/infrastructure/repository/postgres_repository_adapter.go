package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AdiPP/go-marketplace/pkg/domain/entity"
	"github.com/AdiPP/go-marketplace/pkg/infrastructure/config"
	"github.com/AdiPP/go-marketplace/pkg/infrastructure/migration"
	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/jackc/pgx/stdlib"
)

type PostgresRepositoryAdapter struct {
	db *goqu.Database
}

func NewPostgresRepositoryAdapter(cfg config.Database) *PostgresRepositoryAdapter {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabaseSchema,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	log.Println("Postgres database connected")

	goqu.SetDefaultPrepared(true)

	dialect := goqu.Dialect("postgres")

	repo := &PostgresRepositoryAdapter{
		db: dialect.DB(db),
	}

	migration.NewPostgresMigration(cfg).Migrate()

	return repo
}

func (d *PostgresRepositoryAdapter) Save(orderEntity *entity.OrderEntity) (*entity.OrderEntity, error) {
	db := d.db
	order := newOrderEntity(orderEntity)
	items := newOrderItemEntities(orderEntity)

	ds := db.
		Insert("orders").
		Rows(order).
		Returning(
			"id",
			"status",
			"total_price",
			"paid_value",
		)

	_, err := ds.Executor().ScanStruct(&order)
	if err != nil {
		return nil, err
	}

	items = items.setOrderId(order.Id)

	ds = db.
		Insert("order_items").
		Rows(items).
		Returning(
			"id",
			"order_id",
			"product_id",
			"product_name",
			"product_price",
			"quantity",
		)

	err = ds.Executor().ScanStructs(&items)
	if err != nil {
		return nil, err
	}

	orderEntity = order.toEntity()
	for _, v := range items.toEntity() {
		orderEntity.AddItem(v)
	}

	return orderEntity, nil
}
