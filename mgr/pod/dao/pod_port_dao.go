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

type PodPortDao struct {
	db    *sqlx.DB
	cache gfcache.Cache
}

func NewPodPortDao(db *sqlx.DB) *PodPortDao {
	return &PodPortDao{db: db}
}

func (d *PodPortDao) WithCache(cache gfcache.Cache) *PodPortDao {
	d.cache = cache
	return d
}

func (d *PodPortDao) InsertPodPort(ctx context.Context, m PodPort) (int64, error) {
	return d.InsertPodPortTx(ctx, m, nil)
}

func (d *PodPortDao) InsertPodPortTx(ctx context.Context, m PodPort, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	places := make([]string, len(cols))
	sql := fmt.Sprintf("insert into pod_port(%s) values(?%s)",
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

func (d *PodPortDao) TakePodPort0(ctx context.Context, id int64) (*PodPort, error) {
	one := &PodPort{}
	key := fmt.Sprintf("pod_port:%v", id)
	if d.cache != nil {
		if err := d.cache.GetCache(key, one); err == nil {
			return one, nil
		} else if err == RecordNotFound {
			return nil, err
		}
	}
	var err error
	one, err = d.SelectPodPort0(ctx, id)
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

func (d *PodPortDao) SelectPodPort0(ctx context.Context, id int64) (*PodPort, error) {
	var one PodPort
	sql := "select * from pod_port where id=? limit 1"
	if err := d.db.GetContext(ctx, &one, sql, id); err != nil {
		if err == stdsql.ErrNoRows {
			return nil, RecordNotFound
		}
		return nil, err
	}
	return &one, nil
}

func (d *PodPortDao) UpdatePodPort0(ctx context.Context, m PodPort, id int64) (int64, error) {
	return d.UpdatePodPort0Tx(ctx, m, id, nil)
}

func (d *PodPortDao) UpdatePodPort0Tx(ctx context.Context, m PodPort, id int64, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	sql := fmt.Sprintf("update pod_port set %s=? where id=?",
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

func (d *PodPortDao) DeletePodPort0(ctx context.Context, id int64) (int64, error) {
	return d.DeletePodPort0Tx(ctx, id, nil)
}

func (d *PodPortDao) DeletePodPort0Tx(ctx context.Context, id int64, tx *sqlx.Tx) (int64, error) {
	sql := "delete from pod_port where id=?"
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
