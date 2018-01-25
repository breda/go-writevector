package writevector

import (
	"os"
	"io"
)

type WriteVector struct {
	// Data entries
	writers []io.Writer
	
	// wether to sync or not
	sync bool
}

func New(sync bool) *WriteVector {
	return &WriteVector{
		sync: sync,
	}
}

func (wv *WriteVector) AddWriter(writer io.Writer) {
	wv.writers = append(wv.writers, writer)
}

func (wv *WriteVector) Add(path string, append_mode bool) error {
	writer, err := create_writer(wv, path, append_mode)
	if err != nil {
		return err
	}

	wv.writers = append(wv.writers, writer)
	return nil
}

func (wv *WriteVector) WriteString(data string) (int, error) {
	written, err := wv.Write([]byte(data))
	if err != nil {
		return 0, nil
	}
	
	return written, nil
}

func (wv *WriteVector) Write(data []byte) (int, error) {
	written := 0
	for i := 0; i < len(wv.writers); i++ {
		n, err := wv.writers[i].Write(data)
		if err != nil {
			return 0, nil
		}

		written += n
	}

	return written, nil
}

func create_writer(wv *WriteVector, filepath string, append bool) (io.Writer, error) {
	flag := os.O_WRONLY | os.O_CREATE
	
	if append {
		flag = flag | os.O_APPEND
	}
	
	if wv.sync {
		flag = flag | os.O_SYNC
	}
	
	file, err := os.OpenFile(filepath, flag, 0664)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func main() {
	wv := New(false)
	
	wv.Add("tests/testfile1.txt", false)
	wv.Add("tests/testfile2.txt", false)

	wv.WriteString("Holaaaa")
}
