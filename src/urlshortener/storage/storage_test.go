package storage

import (
	"testing"
	"fmt"
	"time"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestIdToSlug(t *testing.T) {
	slug, _ := IdToSlug(125)
	assertEqual(t, slug, "cb", "Wrong slug")

	slug, _ = IdToSlug(19158)
	assertEqual(t, slug, "e9a", "Wrong slug")
}

func TestIdToSlug_InvalidInput(t *testing.T) {
	slug, err := IdToSlug(0)
	assertEqual(t, slug, "", "slug should be empty string on error")
	assertEqual(t, err.Error(), "id must be positive integer", "Wrong error message")
}

func TestSlugToId(t *testing.T) {
	id, _ := SlugToId("cb")
	assertEqual(t, id, 125, "Wrong id")

	id, _ = SlugToId("e9a")
	assertEqual(t, id, 19158, "Wrong id")
}

func TestSlugToId_InvalidInput(t *testing.T) {
	id, err := SlugToId("男裝")
	assertEqual(t, id, 0, "id should be 0 on error")
	assertEqual(t, err.Error(), "Invalid character found in slug", "Wrong error message")
}

func TestStoreAndLoadUrl(t *testing.T) {
	go StoreUrl(125, "www.google.com")

	// wait for 1 second, to make sure our URL was stored so we can immediately retrieve it
	time.Sleep(time.Second)

	ch := make(chan string)
	go LoadUrl(125, ch)

	url := <-ch
	assertEqual(t, url, "www.google.com", "Wrong URL")
}
