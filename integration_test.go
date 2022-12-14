package matata

import (
	"context"
	"os"
	"testing"
)

func TestPong(t *testing.T) {
	c := NewClient(os.Getenv("HAKUNA_TENNANT"), os.Getenv("HAKUNA_TOKEN"))

	res, err := c.Ping(context.TODO())

	if err != nil {
		t.Error(err)
	}

	if res.Pong.IsZero() {
		t.Error("No Pong came back")
	}

}

func TestGetAbsenceTypes(t *testing.T) {

	c := NewClient(os.Getenv("HAKUNA_TENNANT"), os.Getenv("HAKUNA_TOKEN"))

	res, err := c.GetAbsenceTypes(context.TODO())
	if res == nil || err != nil {
		t.Error(err)
	}
}

func TestGetAbsences(t *testing.T) {

	c := NewClient(os.Getenv("HAKUNA_TENNANT"), os.Getenv("HAKUNA_TOKEN"))

	res, err := c.GetAbsences(context.TODO(), nil)
	if res == nil || err != nil {
		t.Error(err)
	}

}
