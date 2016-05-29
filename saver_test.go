package uuid

/****************
 * Date: 21/06/15
 * Time: 6:46 PM
 ***************/

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

const (

	saveDuration  =  3
	generateIds = 1900000
)


func SetupFileSystemStateSaver() *FileSystemSaver {
	saver := new(FileSystemSaver)
	saver.Report = true
	saver.Duration = saveDuration * time.Second
	saver.Timestamp = Now()
	SetupSaver(saver)
	return saver
}

// Tests that the schedule is run on the timeDuration
func TestUUID_State_saveSchedule(t *testing.T) {
	saver := SetupFileSystemStateSaver()

	count := 0

	past := time.Now()
	for i := 0; i < generateIds; i++ {
		if  Now() > saver.Timestamp {
			time.Sleep(3 * time.Second)
			count++
		}
		NewV1()
	}
	d := time.Since(past)

	assert.Equal(t, int(d / saver.Duration), count, "Should be as many saves as second increments")
}
