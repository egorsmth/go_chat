package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/egorsmth/go_chat/models"
	"github.com/egorsmth/go_chat/shared"
)

func setup() {
	err := shared.Init("user=root password=root dbname=test_social_net sslmode=disable")
	if err != nil {
		log.Println("Something wrong with db, check django settings", err)
		return err
	}
	err = genUser()
	if err != nil {
		log.Println("gen user", err)
		return err
	}
	err = genChatRooms()
	if err != nil {
		log.Println("gen chatrooms", err)
		return err
	}
	return nil
}

func teardown() {
	clearDb()
}

func TestMain(m *testing.M) {
	err := setup()
	if err {
		teardown()
		log.Fatal("err while set up", err)
	}
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func TestChatRoom(*testing.T) {
	user := models.User{}
	id := 1
	user.ID = &id
	c, i, err := models.GetChatRooms(&user)
	fmt.Println(c, i, err)
}
