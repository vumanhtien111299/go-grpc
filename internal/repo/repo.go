package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"msh-git.sohatv.vn/ovc-signal/signal-ovc-schema/ent"
	"msh-git.sohatv.vn/ovc-signal/signal-ovc-schema/ent/migrate"

	"msh-git.sohatv.vn/ovc-signal/ovc-signal-full/pkg/log"
)

// ProviderRepoSet is repository providers.
var ProviderRepoSet = wire.NewSet(NewRepo)
var _ IRepo = (*Repo)(nil)

// Repo is repository struct.
type Repo struct {
	*ent.Client
	ctx context.Context
}

// NewRepo creates new repository.
func NewRepo(ctx context.Context) IRepo {
	driver := viper.GetString("DB_DRIVER")
	url := viper.GetString("DB_URL")

	log.Info("Connect to database",
		zap.String("driver", driver),
		zap.String("url", url),
	)

	// Open the database connection.
	drv, err := sql.Open(driver, url)
	if err != nil {
		log.Fatal("sql.Open()", zap.Error(err))
	}

	// Create an ent.Client.
	client := ent.NewClient(ent.Driver(drv))
	opts := []schema.MigrateOption{
		// migrate.WithGlobalUniqueID(true),
		migrate.WithForeignKeys(false), // Disable foreign keys.
	}

	if viper.GetBool("DEBUG_ENABLE") {
		opts = append(opts,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
			migrate.WithFixture(true),
		)

		client = client.Debug()
	}

	if viper.GetBool("DB_MIGRATE") {
		log.Info("Migrating...")

		// Run migration.
		if err = client.Schema.Create(ctx, opts...); err != nil {
			defer func() {
				_ = client.Close()
			}()
			log.Fatal("failed creating schema resources", zap.Error(err))
		}
	}

	return &Repo{
		ctx:    ctx,
		Client: client,
	}
}

// Close closes the repository.
func (r *Repo) Close() error {
	if err := r.Client.Close(); err != nil {
		log.Error("Failed to disconnect from mysql", zap.Error(err))
		return err
	}

	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
