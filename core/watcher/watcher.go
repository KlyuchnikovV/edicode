package watcher

import (
	"context"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

type FileOperations interface {
	OpenFile(string) error
	CloseFile(string) error
	ReloadFile(string) error
}

type Watcher struct {
	context.Context
	*fsnotify.Watcher
	FileOperations
}

func New(ctx context.Context, operations FileOperations, paths ...string) (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, path := range paths {
		if err := w.Add(path); err != nil {
			return nil, err
		}
	}

	var watcher = &Watcher{
		Context: ctx,
		Watcher: w,
	}

	go func() {
		if err := watcher.Start(); err != nil {
			panic(err)
		}
	}()

	return watcher, nil
}

func (watcher *Watcher) Start() error {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return fmt.Errorf("events closed")
			}

			switch event.Op {
			case fsnotify.Create:
				if err := watcher.OpenFile(event.Name); err != nil {
					return err
				}
			case fsnotify.Write:
				if err := watcher.ReloadFile(event.Name); err != nil {
					return err
				}
			case fsnotify.Remove:
				if err := watcher.CloseFile(event.Name); err != nil {
					return err
				}
			case fsnotify.Rename:
				log.Printf("%s %s\n", event.Name, event.Op)
				// TODO: nothing to do now
			case fsnotify.Chmod:
				log.Printf("%s %s\n", event.Name, event.Op)
				// TODO: nothing to do now
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return fmt.Errorf("errors closed")
			}
			log.Println("error:", err)
		case <-watcher.Done():
			return nil
		}
	}
}

func (watcher *Watcher) Stop() error {
	return watcher.Close()
}

// func (watcher *Watcher) FileNames() ([]string, error) {

// }
