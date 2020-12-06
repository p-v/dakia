package app

import (
	"errors"
	bolt "go.etcd.io/bbolt"
	"os/user"
)

// StoreCommand stores the provided command
func StoreCommand(args []string, command string) error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	db, err := bolt.Open(usr.HomeDir+"/.dakia", 0600, nil)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		var prevBucket *bolt.Bucket
		for index, bucketName := range args {
			if prevBucket != nil {
				bucket, err := prevBucket.CreateBucketIfNotExists([]byte(bucketName))
				prevBucket = bucket
				if err != nil {
					return err
				}
			} else {
				bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
				prevBucket = bucket
				if err != nil {
					return err
				}
			}

			if index == len(args)-1 {
				err := prevBucket.Put([]byte("command"), []byte(command))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// GetCommand gets the command to execute
func GetCommand(args []string) (command []byte, err error) {
	db, err := bolt.Open("/Users/pverma/.dakia", 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		var prevBucket *bolt.Bucket
		for index, bucketName := range args {
			if prevBucket != nil {
				bucket := prevBucket.Bucket([]byte(bucketName))
				prevBucket = bucket
				if bucket == nil {
					return errors.New("Cannot find command")
				}
			} else {
				bucket := tx.Bucket([]byte(bucketName))
				prevBucket = bucket
				if bucket == nil {
					return errors.New("Cannot find command")
				}
			}

			if index == len(args)-1 {
				command = prevBucket.Get([]byte("command"))
			}
		}

		if command == nil {
			return errors.New("Cannot find command")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return command, err
}
