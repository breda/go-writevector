package writevector

import (
	"testing"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

func TestAddWriter(t *testing.T) {
	wv := New(true)

	if wv.sync != true {
		t.Errorf("did not create appropriate WriteVector")
	}

	for i := 0; i<3; i++ {
		wv.Add(fmt.Sprintf("testfiles/testFile%d", i), false)
	}

	wv.AddWriter(os.Stdout)
	
	if len(wv.writers) != 4 {
		t.Errorf("does not add writer correctly")
	}

	exec.Command("rm", "-f", "testfiles/testFile*").Run()
}

func TestWrite(t *testing.T) {
	exec.Command("rm", "-f", "testfiles/testfile").Run()
	wv := New(false)
	wv.Add("testfiles/testfile", true)
	wv.Add("testfiles/testfile2", false)

	wv.WriteString("Hola")

	data, err := ioutil.ReadFile("testfiles/testfile")
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}

	if len(data) != len("Hola") {
		t.Error("did not write sufficient ammout of data to file")
	}
}



