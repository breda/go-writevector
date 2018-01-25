package writevector

import (
	"os"
	"fmt"
)

type writeEntry struct {
	path string
	data string
	append bool
}

type WriteVector struct {
	// Data entries
	writes []writeEntry
	
	// wether to sync or not
	sync bool
}

func New(sync bool) *WriteVector {
	return &WriteVector{
		sync: sync,
	}
}

func (wv *WriteVector) Add(path, data string, append_mode bool) error {
	we := writeEntry{
		path: path,
		data: data,
		append: append_mode,
	}

	wv.writes = append(wv.writes, we)
	return nil
}

func (wv *WriteVector) Write() error {
	writes_len := len(wv.writes)

	for i := 0; i < writes_len; i++ {
		err := do_write(wv.writes[i], wv.sync) 
		if err != nil {
			return err
		}
	}
	return nil
}

func do_write(we writeEntry, sync bool) error {
	flag := os.O_WRONLY | os.O_CREATE
	
	if we.append {
		flag = flag | os.O_APPEND
	}
	
	if sync {
		flag = flag | os.O_SYNC
	}
	
	file, err := os.OpenFile(we.path, flag, 0664)
	if err != nil {
		return err
	}

	written, err := file.WriteString(we.data);
	if  err != nil { 
		return err
	}

	if written != len(we.data) {
		return fmt.Errorf("could not write data to file (data=%s, file path=%s)", we.data, we.path)	
	}
		
	return nil	
}

func main() {
	wv := New(false)
	
	wv.Add("tests/testfile1.txt", "Hola amigo 1", false)
	wv.Add("tests/testfile2.txt", "Hola amigo 2", false)

	wv.Write()
}
