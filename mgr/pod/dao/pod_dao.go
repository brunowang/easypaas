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

type PodDao struct {
	db    *sqlx.DB
	cache gfcache.Cache
}

func NewPodDao(db *sqlx.DB) *PodDao {
	return &PodDao{db: db}
}

func (d *PodDao) WithCache(cache gfcache.Cache) *PodDao {
	d.cache = cache
	return d
}

func (d *PodDao) InsertPod(ctx context.Context, m Pod) (int64, error) {
	return d.InsertPodTx(ctx, m, nil)
}

func (d *PodDao) InsertPodTx(ctx context.Context, m Pod, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	places := make([]string, len(cols))
	sql := fmt.Sprintf("insert into pod(%s) values(?%s)",
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

func (d *PodDao) TakePod0(ctx context.Context, id int64) (*Pod, error) {
	one := &Pod{}
	key := fmt.Sprintf("pod:%v", id)
	if d.cache != nil {
		if err := d.cache.GetCache(key, one); err == nil {
			return one, nil
		} else if err == RecordNotFound {
			return nil, err
		}
	}
	var err error
	one, err = d.SelectPod0(ctx, id)
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

func (d *PodDao) SelectPod0(ctx context.Context, id int64) (*Pod, error) {
	var one Pod
	sql := "select * from pod where id=? limit 1"
	if err := d.db.GetContext(ctx, &one, sql, id); err != nil {
		if err == stdsql.ErrNoRows {
			return nil, RecordNotFound
		}
		return nil, err
	}
	return &one, nil
}

func (d *PodDao) TakePod1(ctx context.Context, podName string) (*Pod, error) {
	one := &Pod{}
	key := fmt.Sprintf("pod:%v", podName)
	if d.cache != nil {
		if err := d.cache.GetCache(key, one); err == nil {
			return one, nil
		} else if err == RecordNotFound {
			return nil, err
		}
	}
	var err error
	one, err = d.SelectPod1(ctx, podName)
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

func (d *PodDao) SelectPod1(ctx context.Context, podName string) (*Pod, error) {
	var one Pod
	sql := "select * from pod where pod_name=? limit 1"
	if err := d.db.GetContext(ctx, &one, sql, podName); err != nil {
		if err == stdsql.ErrNoRows {
			return nil, RecordNotFound
		}
		return nil, err
	}
	return &one, nil
}

func (d *PodDao) UpdatePod0(ctx context.Context, m Pod, id int64) (int64, error) {
	return d.UpdatePod0Tx(ctx, m, id, nil)
}

func (d *PodDao) UpdatePod0Tx(ctx context.Context, m Pod, id int64, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	sql := fmt.Sprintf("update pod set %s=? where id=?",
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

func (d *PodDao) UpdatePod1(ctx context.Context, m Pod, podName string) (int64, error) {
	return d.UpdatePod1Tx(ctx, m, podName, nil)
}

func (d *PodDao) UpdatePod1Tx(ctx context.Context, m Pod, podName string, tx *sqlx.Tx) (int64, error) {
	cols, args := genSqlCols(m)
	sql := fmt.Sprintf("update pod set %s=? where pod_name=?",
		strings.Join(cols, "=?,"))
	args = append(args, podName)
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

func (d *PodDao) DeletePod0(ctx context.Context, id int64) (int64, error) {
	return d.DeletePod0Tx(ctx, id, nil)
}

func (d *PodDao) DeletePod0Tx(ctx context.Context, id int64, tx *sqlx.Tx) (int64, error) {
	sql := "delete from pod where id=?"
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

func (d *PodDao) DeletePod1(ctx context.Context, podName string) (int64, error) {
	return d.DeletePod1Tx(ctx, podName, nil)
}

func (d *PodDao) DeletePod1Tx(ctx context.Context, podName string, tx *sqlx.Tx) (int64, error) {
	sql := "delete from pod where pod_name=?"
	var db sqlx.ExecerContext = d.db
	if tx != nil {
		db = tx
	}
	res, err := db.ExecContext(ctx, sql, podName)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
