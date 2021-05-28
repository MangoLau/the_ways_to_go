package main

import (
	"encoding/gob"
	"io"
	"log"
	"os"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu sync.RWMutex
	save chan record
}

type record struct {
	Key, URL string
}

const saveQueueLength = 1000
func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log .Println("Error loading data in URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _,present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			s.save <- record{key, url}
			return key
		}
	}
}

// load goto 程序启动时，磁盘上存储的数据必须读取到 URLStore 的 map 中。
func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err)
	return err
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		// taking a record from the channel and encoding it
		r := <-s.save
		if err := e.Encode(r); err != nil {
			log.Println("URLStore:", err)
		}
	}
}