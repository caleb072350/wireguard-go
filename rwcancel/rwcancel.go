package rwcancel

import "os"

type RWCancel struct {
	fd            int
	closingReader *os.File
	closingWriter *os.File
}
