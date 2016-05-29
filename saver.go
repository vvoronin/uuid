package uuid

/****************
 * Date: 21/06/15
 * Time: 5:48 PM
 ***************/

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)


// This implements the Saver interface for UUIDs
type FileSystemSaver struct {
	cache        *os.File

	// Whether to log each save
	Report   bool

	// The amount of time between each save call
	time.Duration

	// The next time to save
	Timestamp
}

func (o *FileSystemSaver) Save(pStore *Store) {

	if pStore.Timestamp >= o.Timestamp {
		err := o.open()
		defer o.cache.Close()
		if err == nil {
			// do the save
			err = o.encode(pStore)
			if (err == nil) {
				if o.Report {
					log.Printf("UUID Saved State Storage: %s", pStore)
				}
			}
		}
		if err != nil {
			log.Println("uuid.State.save:", err)
		}
		o.Timestamp = pStore.Add(o.Duration)
	}
}

func (o *FileSystemSaver) Read() (err error, store Store) {
	gob.Register(Store{})

	err = o.open()
	defer o.cache.Close()

	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("'%s' created\n", "uuid.FileSystemSaver")
			o.cache, err = os.Create(os.TempDir() + "/state.unique")
			if err != nil {
				log.Println("uuid.FileSystemSaver.Init: SaveState error:", err)
				return
			}
		} else {
			log.Println("uuid.FileSystemSaver.Init: SaveState error:", err)
			return
		}
	}
	return o.decode()
}

func (o *FileSystemSaver) reset() {
	o.cache.Seek(0, 0)
}

func (o *FileSystemSaver) open() (err error) {
	o.cache, err = os.OpenFile(os.TempDir()+"/state.unique", os.O_RDWR, os.ModeExclusive)
	return
}

func (o *FileSystemSaver) encode(pStore *Store) error {
	// ensure reader state is ready for use
	o.reset()
	enc := gob.NewEncoder(o.cache)
	err := enc.Encode(&pStore)
	if err != nil {
		log.Println("uuid.FileSystemSaver.encode error:", err)
	}
	return err
}

func (o *FileSystemSaver) decode() (err error, store Store) {
	// ensure reader state is ready for use
	// o.reset()
	dec := gob.NewDecoder(o.cache)
	store = Store{}
	err = dec.Decode(&store)
	if err != nil {
		log.Println("uuid.FileSystemSaver.decode error:", err)
		return
	}
	return
}
