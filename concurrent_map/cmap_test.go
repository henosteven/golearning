package concurrent_map

import (
	"testing"
	"fmt"
)


func TestCmap_Set(t *testing.T) {
	cmap := NewCmap()
	cmap.Set("name", "jinjing")
}

func TestCmap_Get(t *testing.T) {
	cmap := NewCmap()
	cmap.Set("name", "jinjing")
	val, ok := cmap.Get("name").(string)
	if !ok {
		fmt.Printf("error-type")
	}

	if val != "jinjing" {
		fmt.Println("error-val")
	}
}
