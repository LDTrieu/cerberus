package exception

import (
	"errors"
	"fmt"
	"time"
)

/*
EXAMPLE

	func main() {
		var err = exception.New(exception.DB_EXCEPTION, "DBException at %v", "test")
		PrintTypeError(err)
		err = exception.Wrap(exception.DB_EXCEPTION, errors.New("DBException at"), "DBException at %v", "test")
		PrintTypeError(err)
	}

	func PrintTypeError(err error) {
		switch err.(type) {
		case *exception.DbException:
			log.Println("DbException erorr::: ", err.Error())
		case *exception.SysException:
			log.Println("SystemException erorr::: ", err)
		default:
			log.Println("Exception error::: ")
		}
	}
*/
const (
	DB_EXCEPTION  = "DbException"
	SYS_EXCEPTION = "SysException"
	DATA_VALIDATE = "DataValidate"
)

type DataValidate struct {
	Type     string
	CreateAt time.Time
	Message  string
	Err      error
}

func (e *DataValidate) Error() string {
	return e.Err.Error()
}

type DbException struct {
	Type     string
	CreateAt time.Time
	Message  string
	Err      error
}

func (e *DbException) Error() string {
	return e.Err.Error()
}

type SysException struct {
	Type     string
	CreateAt time.Time
	Message  string
	Err      error
}

func (e *SysException) Error() string {
	return e.Err.Error()
}

type Exception struct {
	Type     string
	CreateAt time.Time
	Message  string
	Err      error
}

func (e *Exception) Error() string {
	return e.Err.Error()
}

// Wrap error
func Wrap(_type string, err error, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	switch _type {
	case DB_EXCEPTION:
		return &DbException{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	case SYS_EXCEPTION:
		return &SysException{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	default:
		return &Exception{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	}

}

// New create a new error
func New(_type string, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	err := errors.New(message)
	switch _type {
	case DB_EXCEPTION:
		return &DbException{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	case SYS_EXCEPTION:
		return &SysException{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	default:
		return &Exception{
			CreateAt: time.Now(),
			Type:     _type,
			Err:      err,
			Message:  message,
		}
	}
}
