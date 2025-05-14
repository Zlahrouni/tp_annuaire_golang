package main

import (
	"testing"
)

func TestAddDir(t *testing.T) {
	dir := map[string]info{
		"ziad": {"075268618"},
	}

	addDir(dir, "sabrina", "0612345678")
	if _, exists := dir["sabrina"]; !exists {
		t.Errorf("Le contact sabrina' n'a pas été ajouté")
	}
	if dir["sabrina"].phone != "0612345678" {
		t.Errorf("sabrina's phone number is incorrect")
	}
}

func TestDeleteDir(t *testing.T) {
	dir := map[string]info{
		"ziad":    {"075268618"},
		"sabrina": {"0685268848"},
	}

	deleteDir(dir, "ziad")

	if _, exists := dir["ziad"]; exists {
		t.Errorf("ziad is not deleted")
	}
}

func TestUpdateDir(t *testing.T) {
	dir := map[string]info{
		"ziad": {"075268618"},
	}

	updateDir(dir, "ziad", "0612345678")
	if dir["ziad"].phone != "0612345678" {
		t.Errorf("phone numbre is not updated")
	}
}

func TestSearchDirExisting(t *testing.T) {
	dir := map[string]info{
		"ziad": {"075268618"},
	}
	searchDir(dir, "ZIAD")
}

func TestSearchDirNonExisting(t *testing.T) {
	dir := map[string]info{
		"ziad": {"075268618"},
	}
	searchDir(dir, "sabrina")
}

func TestAddExistingContact(t *testing.T) {
	dir := map[string]info{
		"ziad": {"075268618"},
	}
	addDir(dir, "ziad", "0612345678")
	if dir["ziad"].phone != "075268618" {
		t.Errorf("ziad already exist in the directory")
	}
}
