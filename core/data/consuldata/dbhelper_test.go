package consuldata

import (
	"testing"
)

func TestFindAll(t *testing.T) {
	object := FindAll()
	t.Log(len(object))
}
