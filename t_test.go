package sequebuf

import (
	"encoding/json"
	"testing"
	"time"
)

type User struct {
	Name  string
	Age   uint8
	Score float64
	Admin bool
	Roles []string
	Level map[string]uint8
}

func TestMarshal(t *testing.T) {

	user := User{
		Name:  "zhangsan",
		Age:   18,
		Score: 25.5,
		Admin: true,
		Roles: []string{
			"Admin",
			"User",
		},
		Level: map[string]uint8{
			"Admin": 1,
			"User":  2,
		},
	}

	data, err := Marshal(user)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(data)
	var u2 User
	err = Unmarshal(data, &u2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(u2)

}

func TestPerformance(t *testing.T) {
	user := User{
		Name:  "zhangsan",
		Age:   18,
		Score: 25.5,
		Admin: true,
		Roles: []string{
			"Admin",
			"User",
		},
		Level: map[string]uint8{
			"Admin": 1,
			"User":  2,
		},
	}

	const size = 1000000

	start := time.Now()
	for i := 0; i < size; i++ {
		_, _ = Marshal(user)
	}
	t.Log("binary encode time consuming", time.Since(start))

	data, _ := Marshal(user)
	start = time.Now()
	for i := 0; i < size; i++ {
		var user User
		_ = Unmarshal(data, &user)
	}

	t.Log("binary decode time consuming", time.Since(start))

	start = time.Now()

	for i := 0; i < size; i++ {
		_, _ = json.Marshal(user)
	}
	t.Log("json encode time consuming", time.Since(start))

	data, _ = json.Marshal(user)
	start = time.Now()
	for i := 0; i < size; i++ {
		var user User
		_ = json.Unmarshal(data, &user)
	}
	t.Log("json decode time consuming", time.Since(start))
}
