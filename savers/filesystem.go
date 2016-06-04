package savers

/****************
 * Date: 30/05/16
 * Time: 5:48 PM
 ***************/

import (
	"encoding/gob"
	"github.com/twinj/uuid"
	"log"
	"os"
	"path"
	"time"
)

var _ uuid.Saver = &FileSystemSaver{}

// This implements the Saver interface for UUIDs
type FileSystemSaver struct {
	// A file to save the state to
	// Used gob format on uuid.State entity
	file *os.File

	// Preferred location for the store
	Path string

	// Whether to log each save
	Report bool

	// The amount of time between each save call
	time.Duration

	// The next time to save
	uuid.Timestamp
}

func (o *FileSystemSaver) Save(pStore *uuid.Store) {

	if pStore.Timestamp >= o.Timestamp {
		err := o.open()
		defer o.file.Close()
		if err == nil {
			// do the save
			err = o.encode(pStore)
			if err == nil {
				if o.Report {
					log.Printf("UUID Saved State Storage: %s", pStore)
				}
			}
		}
		if err != nil {
			log.Println("uuid.FileSystemSaver.Save:", err)
		}
		o.Timestamp = pStore.Add(o.Duration)
	}
}

func (o *FileSystemSaver) Read() (err error, store uuid.Store) {
	gob.Register(uuid.Store{})

	err = o.open()
	defer o.file.Close()

	if err != nil {
		if os.IsNotExist(err) {
			dir, file := path.Split(o.Path)
			if dir == "" || dir == "/" {
				dir = os.TempDir()
			}
			o.Path = path.Join(dir, file)

			err = os.MkdirAll(dir, os.ModeDir|0755)
			if err != nil {
				goto error
			}

			o.file, err = os.Create(o.Path)
			if err != nil {
				goto error
			}

			log.Println("uuid.FileSystemSaver created", o.Path)

			// If new encode blank store
			o.encode(&uuid.Store{})
		} else {
			goto error
		}
	}
	return o.decode()

error:
	log.Println("uuid.FileSystemSaver.Read: error will autogenerate", err)
	return
}

func (o *FileSystemSaver) reset() {
	o.file.Seek(0, 0)
}

func (o *FileSystemSaver) open() (err error) {
	o.file, err = os.OpenFile(o.Path, os.O_RDWR, os.ModeExclusive)
	return
}

func (o *FileSystemSaver) encode(pStore *uuid.Store) error {
	// ensure reader state is ready for use
	enc := gob.NewEncoder(o.file)
	err := enc.Encode(&pStore)
	if err != nil {
		log.Println("uuid.FileSystemSaver.encode error:", err)
	}
	return err
}

func (o *FileSystemSaver) decode() (err error, store uuid.Store) {
	// ensure reader state is ready for use
	o.reset()
	dec := gob.NewDecoder(o.file)
	store = uuid.Store{}
	err = dec.Decode(&store)
	if err != nil {
		log.Println("uuid.FileSystemSaver.decode error:", err)
		return
	}
	return
}
