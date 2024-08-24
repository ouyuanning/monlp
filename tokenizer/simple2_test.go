package tokenizer

import (
	"bytes"
	"io"
	"os"
	"path"
	"runtime"
	"testing"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func readFile(t *testing.T, filename string) []byte {
	_, fn, _, _ := runtime.Caller(0)
	dir := path.Dir(fn)
	fpath := path.Join(dir, "..", "data", filename)
	file, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Failed to open file %s: %v", filename, err)
		return nil
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Errorf("Failed to read file %s: %v", filename, err)
		return nil
	}
	return data
}

func TestShakespear(tt *testing.T) {
	data := readFile(tt, "t8.shakespeare.txt")
	if data == nil {
		tt.Errorf("Failed to read shakespeare.txt")
		return
	}

	var mk1000, mk1100 Token
	mk1000.TokenBytes[0] = 12
	copy(mk1000.TokenBytes[1:], []byte("_MARKER_1000"))
	mk1100.TokenBytes[0] = 12
	copy(mk1100.TokenBytes[1:], []byte("_MARKER_1100"))

	var mk1000Pos, mk1100Pos, maxToken int64

	tknz := NewSimpleTokenizer(data)
	for t := range tknz.Tokenize() {
		if t.TokenBytes == mk1000.TokenBytes {
			mk1000Pos = t.TokenPos
			tt.Log("SomeText after mk1000:", string(data[mk1000Pos:mk1000Pos+50]))
		} else if t.TokenBytes == mk1100.TokenBytes {
			mk1100Pos = t.TokenPos
		}
		maxToken = t.TokenPos
	}

	tt.Log("mk1000Pos:", mk1000Pos)
	tt.Log("mk1100Pos:", mk1100Pos)
	tt.Log("MaxToken:", maxToken)
}

func TestHLM(tt *testing.T) {
	gbkdata := readFile(tt, "红楼梦.txt")
	if gbkdata == nil {
		tt.Errorf("Failed to read hlm.txt")
		return
	}

	reader := bytes.NewReader(gbkdata)
	tr := transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
	data, err := io.ReadAll(tr)
	if err != nil {
		tt.Errorf("Failed to decode hlm.txt: %v", err)
		return
	}

	var mk1000, mk1100 Token
	mk1000.TokenBytes[0] = 12
	copy(mk1000.TokenBytes[1:], []byte("_MARKER_1000"))
	mk1100.TokenBytes[0] = 12
	copy(mk1100.TokenBytes[1:], []byte("_MARKER_1100"))

	var mk1000Pos, mk1100Pos, maxToken int64

	tknz := NewSimpleTokenizer(data)
	for t := range tknz.Tokenize() {
		if t.TokenBytes == mk1000.TokenBytes {
			mk1000Pos = t.TokenPos
			tt.Log("SomeText after mk1000:", string(data[mk1000Pos:mk1000Pos+50]))
		} else if t.TokenBytes == mk1100.TokenBytes {
			mk1100Pos = t.TokenPos
		}
		maxToken = t.TokenPos
	}

	tt.Log("mk1000Pos:", mk1000Pos)
	tt.Log("mk1100Pos:", mk1100Pos)
	tt.Log("MaxToken:", maxToken)
}