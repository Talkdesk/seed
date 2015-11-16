package main

import "testing"

func TestFindNamesInclusive(t *testing.T) {
	db_collections := []string{"accounts", "users", "entries", "customers", "presets"}
	provided_collections := []string{"accounts", "customers"}

	result, err := FindNames(db_collections, provided_collections)

	if err != nil {
		t.Fail()
	}

	if len(result) != 2 {
		t.Fail()
	}
}

func TestFindNamesExclusive(t *testing.T) {
	db_collections := []string{"accounts", "users", "entries", "customers", "presets"}
	provided_collections := []string{"accounts", "executions"}

	result, err := FindNames(db_collections, provided_collections)

	if err != nil {
		t.Fail()
	}

	if len(result) != 1 {
		t.Fail()
	}
}

func TestFindNamesEmptyProvided(t *testing.T) {
	db_collections := []string{"accounts", "users", "entries", "customers", "presets"}
	provided_collections := []string{}

	result, err := FindNames(db_collections, provided_collections)

	if err != nil {
		t.Fail()
	}

	if len(result) != 5 {
		t.Fail()
	}
}
