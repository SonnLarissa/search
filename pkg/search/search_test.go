package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {
	ch := All(context.Background(), "марсоход", []string{"test.txt"})
	results, ok := <-ch
	if !ok {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}

func TestAll_singleOne(t *testing.T) {
	ch := All(context.Background(), "singleOne", []string{"test.txt"})
	results, ok := <-ch
	if !ok {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}

func TestAll_notFound(t *testing.T) {
	ch := All(context.Background(), "notFound", []string{"test.txt"})
	results, ok := <-ch
	if ok!=false {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}

func TestAny_positiv(t *testing.T) {
	ch := Any(context.Background(), "singleOne", []string{"test.txt"})
	results, ok := <-ch
	if ok!=false {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}

func TestAny_negative(t *testing.T) {
	ch := Any(context.Background(), "fake", []string{"test.txt"})
	results, ok := <-ch
	if ok!=false {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}

func TestAny_multiSearch(t *testing.T) {
	ch := Any(context.Background(), "марсоход", []string{"test.txt"})
	results, ok := <-ch
	if ok!=false {
		t.Errorf("error: %v", ok)
	}
	log.Println("result: ", results)
}
