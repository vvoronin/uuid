package savers

/****************
 * Date: 30/05/16
 * Time: 6:46 PM
 ***************/

import (
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
	"testing"
	"time"
)

const (
	saveDuration = 3
)

func SetupFileSystemStateSaver() *FileSystemSaver {
	saver := new(FileSystemSaver)
	saver.Report = true
	saver.Duration = saveDuration * time.Second
	saver.Timestamp = uuid.Now()
	return saver
}

// Tests that the schedule is run on the timeDuration
func TestFileSystemSaver_SaveSchedule(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	count := 0

	past := time.Now()
	for i := 0; i < 5; i++ {
		if uuid.Now() > saver.Timestamp {
			time.Sleep(saver.Duration)
			count++
		}
		store := &uuid.Store{uuid.Now(), 3, []byte{0xff, 0xaa, 0x11}}
		saver.Save(store)
	}
	d := time.Since(past)

	assert.Equal(t, int(d / saver.Duration), count, "Should be as many saves as second increments")
}

func TestFileSystemSaver_Read(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	err, _ := saver.Read()

	assert.Nil(t, err)
}

func TestFileSystemSaver_Save(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	store := &uuid.Store{Timestamp: 1, Sequence: 2, Node: []byte{0xff, 0xaa, 0x33, 0x44, 0x55, 0x66}}
	saver.Save(store)
}

func TestFileSystemSaver_SaveAndRead(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	store := &uuid.Store{Timestamp: 1, Sequence: 2, Node: []byte{0xff, 0xaa, 0x33, 0x44, 0x55, 0x66}}
	saver.Save(store)

	_, saved := saver.Read()

	assert.Equal(t, store.Timestamp, saved.Timestamp)
	assert.Equal(t, store.Sequence, saved.Sequence)
	assert.Equal(t, store.Node, saved.Node)
}
