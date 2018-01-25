package writevector

import (
	"testing"
	"os"
	"io/ioutil"
)

var table = []writeEntry {
	{"test1", "test1", false},
	{"test2", "test2", false},
	{"test3", "test3", false},
	{"test4", "test4", false},
}

func TestAddWriteEntry(t *testing.T) {
	wv := New(true)

	if wv.sync != true {
		t.Errorf("did not create appropriate WriteVector")
	}

	for i:= 0; i<len(table); i++ {
		wv.Add(table[i].path, table[i].data, table[i].append)
	}

	if len(wv.writes) != len(table) {
		t.Errorf("does not add write entries correctly")
	}

	if wv.writes[0].data != "test1" || wv.writes[0].data != "test1" || wv.writes[0].append != false {
		t.Errorf("does not add in a correct way")		
	}
}

func TestWritesCorrectly(t *testing.T) {
	wv := New(false)

	// Add data
	for i:= 0; i<len(table); i++ {
		wv.Add(table[i].path, table[i].data, table[i].append)
	}

	wv.Write()

	// Check data
	for i:= 0; i<len(table); i++ {
		// Check existance
		if _, err := os.Stat(table[i].path); os.IsNotExist(err) {
			t.Errorf("file did not get created when written to")
		}
	
		// Check written data
		data, err := ioutil.ReadFile(table[i].path)
		if err != nil {
			t.Errorf("error reading file: %v", err)	
		}
		
		if string(data) != table[i].data {
			t.Errorf("did not write specified data to file")
		}
	}
}


