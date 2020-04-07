import (
	"database/sql"

	"github.com/pkg/errors"
)

// 异常定义
type E struct {
}

func NewE() *E {
	return &E{}
}

// 记录堆栈
func (e E) Stack(err error) error {
	if b, ee := e.isNull(err); b {
		return ee
	} else {
		return errors.WithStack(ee)
	}
}

// 堆栈附带信息
func (e E) StackMsg(err error, msg string) error {
	if b, ee := e.isNull(err); b {
		return err
	} else {
		return errors.Wrap(ee, msg)
	}
}

// 堆栈附带信息 format
func (e E) StackMsgF(err error, format string, args ...interface{}) error {
	if b, ee := e.isNull(err); b {
		return err
	} else {
		return errors.Wrapf(ee, format, args...)
	}
}

// 判断是否是空的错误，则不处理
func (e E) isNull(err error) (bool, error) {
	switch err {
	case sql.ErrNoRows, sql.ErrConnDone, sql.ErrTxDone:
		return true, err
	}
	return false, err
}
