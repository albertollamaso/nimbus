package nimbus

import "testing"

func TestAddKV(t *testing.T) {
	db, err := InitDB("db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	AddKV(db, "key1", "myvalue")
}

func TestListKV(t *testing.T) {
	db, err := InitDB("db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ListKV(db)
}

func TestReadKV(t *testing.T) {
	db, err := InitDB("db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	value, _ := ReadKV(db, "key1")
	if value != "myvalue" {
		t.Fatal("value is not equal to myvalue")
	}
}

func TestRemoveKV(t *testing.T) {
	db, err := InitDB("db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = RemoveKV(db, "key1")
	if err != nil {
		t.Log("key1 was not found in database")
		t.Fatal(err)

	}
}
