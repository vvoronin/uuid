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
	uuid.SetupSaver(saver)
	return saver
}

// Tests that the schedule is run on the timeDuration
func TestUUID_State_saveSchedule(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	count := 0

	past := time.Now()
	for i := 0; i < 5; i++ {
		if uuid.Now() > saver.Timestamp {
			time.Sleep(saver.Duration)
			count++
		}
		uuid.NewV1()
	}
	d := time.Since(past)

	assert.Equal(t, int(d/saver.Duration), count, "Should be as many saves as second increments")
}
