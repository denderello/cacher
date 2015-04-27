package db

import (
	"encoding/csv"
	"os"

	"github.com/juju/errgo"
)

type SlowFileDatabase struct {
	file          *os.File
	state         map[string]string
	csvReader     *csv.Reader
	csvWriter     *csv.Writer
	syncThreshold uint64
	writeCounter  uint64
}

func NewSlowFileDatabase(filename string, syncThreshold uint64) (*SlowFileDatabase, error) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &SlowFileDatabase{
		file:          f,
		csvReader:     csv.NewReader(f),
		csvWriter:     csv.NewWriter(f),
		syncThreshold: syncThreshold,
		writeCounter:  0,
	}, nil
}

func (sfb *SlowFileDatabase) Open() error {
	rs, err := sfb.csvReader.ReadAll()
	if err != nil {
		return errgo.Mask(err)
	}

	sfb.state = make(map[string]string)
	for _, r := range rs {
		sfb.state[r[0]] = r[1]
	}

	return nil
}

func (sfb *SlowFileDatabase) Close() {
	sfb.file.Close()
}

func (sfb *SlowFileDatabase) Get(key string) (string, error) {
	if v, ok := sfb.state[key]; ok {
		return v, nil
	}
	return "", errgo.Newf("Could not find value for key '%s'", key)
}

func (sfb *SlowFileDatabase) Set(key, value string) error {
	sfb.state[key] = value

	if sfb.writeCounter >= sfb.syncThreshold {
		sfb.writeCounter = 0
		return errgo.Mask(sfb.synchronizeToDisk())
	}

	sfb.writeCounter++
	return nil
}

func (sfb *SlowFileDatabase) synchronizeToDisk() error {
	var rs [][]string
	for k, v := range sfb.state {
		rs = append(rs, []string{k, v})
	}

	if err := sfb.file.Truncate(0); err != nil {
		return errgo.Mask(err)
	}

	if _, err := sfb.file.Seek(0, 0); err != nil {
		return errgo.Mask(err)
	}

	return errgo.Mask(sfb.csvWriter.WriteAll(rs))
}
