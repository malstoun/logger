package logger

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"os"
)

var (
	Trace     *log.Logger
	Info      *log.Logger
	Warning   *log.Logger
	Error     *log.Logger
	infoFile  *os.File
	errorFile *os.File
)

func New(params map[string]string) {
	if params["env"] == "dev" {
		Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

		return
	}

	infoLog := params["info_file_path"]
	errorLog := params["error_file_path"]

	infoFile, err := os.OpenFile(infoLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	errorFile, err := os.OpenFile(errorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		infoFile.Close()
		panic(err)
	}

	Info = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(errorFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func FM(err error, args ...interface{}) (string, string) {
	uuid := uuid.NewV4().String()

	return uuid, fmt.Sprintf("%s\n%s\n%s\n", uuid, err, fmt.Sprint(args))
}

func Finish() {
	infoFile.Close()
	errorFile.Close()
}
