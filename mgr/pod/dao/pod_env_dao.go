package dao

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"github.com/brunowang/gframe/gfcache"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type PodEnvDao struct {
	db    *sqlx.DB
	cache gfcache.Cache
}

func NewPodEnvDao(db *sqlx.DB) *PodEnvDao {
	return &PodEnvDao{db: db}
}

func (d *PodEnvDao) WithCache(cache gfcache.Cache) *PodEnvDao {
	d.cache = cache
	return d
}

func (d *PodEnvDao) InsertPodEnv(ctx context.Context, m PodEnv) (int64, error) {
	return d.InsertPodEnvTx(ctx, m, nil)
}

func (d *PodEnvDao) InsertPodEnvTx(ctx context.Context, m PodEnv, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	places := make([]string, len(cols))
	sql := fmt.Sprintf("insert into pod_env(%s) values(?%s)",
		strings.Join(cols, ","), strings.Join(places, ",?"))
	var db sqlx.ExecerContext = d.db
	if tx != nil {
		db = tx
	}
	res, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return 0, DuplicateKey
		}
		return 0, err
	}
	return res.LastInsertId()
}

func (d *PodEnvDao) TakePodEnv0(ctx context.Context, id int64) (*PodEnv, error) {
	one := &PodEnv{}
	key := fmt.Sprintf("pod_env:%v", id)
	if d.cache != nil {
		if err := d.cache.GetCache(key, one); err == nil {
			return one, nil
		} else if err == RecordNotFound {
			return nil, err
		}
	}
	var err error
	one, err = d.SelectPodEnv0(ctx, id)
	if err == RecordNotFound && d.cache != nil {
		_ = d.cache.SetCache(key, nil)
	} else if err != nil {
		return nil, err
	}
	if d.cache != nil {
		_ = d.cache.SetCache(key, one)
	}
	return one, nil
}

func (d *PodEnvDao) SelectPodEnv0(ctx context.Context, id int64) (*PodEnv, error) {
	var one PodEnv
	sql := "select * from pod_env where id=? limit 1"
	if err := d.db.GetContext(ctx, &one, sql, id); err != nil {
		if err == stdsql.ErrNoRows {
			return nil, RecordNotFound
		}
		return nil, err
	}
	return &one, nil
}

func (d *PodEnvDao) UpdatePodEnv0(ctx context.Context, m PodEnv, id int64) (int64, error) {
	return d.UpdatePodEnv0Tx(ctx, m, id, nil)
}

func (d *PodEnvDao) UpdatePodEnv0Tx(ctx context.Context, m PodEnv, id int64, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	sql := fmt.Sprintf("update pod_env set %s=? where id=?",
		strings.Join(cols, "=?,"))
	args = append(args, id)
	var db sqlx.ExecerContext = d.db
	if tx != nil {
		db = tx
	}
	res, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (d *PodEnvDao) DeletePodEnv0(ctx context.Context, id int64) (int64, error) {
	return d.DeletePodEnv0Tx(ctx, id, nil)
}

func (d *PodEnvDao) DeletePodEnv0Tx(ctx context.Context, id int64, tx *sqlx.Tx) (int64, error) {
	sql := "delete from pod_env where id=?"
	var db sqlx.ExecerContext = d.db
	if tx != nil {
		db = tx
	}
	res, err := db.ExecContext(ctx, sql, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
