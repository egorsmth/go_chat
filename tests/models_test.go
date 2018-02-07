package tests

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/egorsmth/go_chat/models"
	"github.com/egorsmth/go_chat/shared"
)

func setup() error {
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

	err = genPairChatroom()
	if err != nil {
		log.Println("gen user to pair", err)
		return err
	}

	err = genMessages()
	if err != nil {
		log.Println("gen rmessages", err)
		return err
	}

	return nil
}

func teardown() {
	err := clearDb()
	if err != nil {
		log.Println("teardown", err)
	}
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		teardown()
		log.Fatal("err while set up", err)
	}
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

var expectedRoom = models.ChatRoom{1, nil, 1, "pair"}

var ID = 1
var msg = "first message"
var status = "unread"
var date = time.Date(2018, 02, 07, 15, 33, 00, 00, time.UTC)

var expectedRoomIds = []int{1}

var expectedUnreaded = 1

func TestChatRoom(t *testing.T) {
	user := models.User{}
	id := 1
	user.ID = &id
	rooms, roomIds, unreaded, err := models.GetChatRooms(&user)
	if err != nil {
		t.Error("error", err)
	}

	if len((*rooms)) != 1 {
		t.Error("expected", 1, "got", len((*rooms)))
	}

	if (*rooms)[0].ID != expectedRoom.ID {
		t.Error("expected", expectedRoom, "got", *rooms)
	}

	if (*rooms)[0].LastMessageID != expectedRoom.LastMessageID {
		t.Error("expected", expectedRoom, "got", *rooms)
	}

	if (*rooms)[0].Status != expectedRoom.Status {
		t.Error("expected", expectedRoom, "got", *rooms)
	}

	message := (*rooms)[0].LastMessage
	if *(*message).ID != ID {
		t.Error("expected", *(*message).ID, "got", ID)
	}
	if *(*message).Message != msg {
		t.Error("expected", *(*message).Message, "got", msg)
	}
	if *(*message).Status != status {
		t.Error("expected", *(*message).Status, "got", status)
	}

	if *(*message).Date != date {
		t.Error("expected", *(*message).Date, "got", date)
	}

	if len((*roomIds)) != len(expectedRoomIds) {
		t.Error("expected", roomIds, "got", expectedRoomIds)
	}

	if (*roomIds)[0] != expectedRoomIds[0] {
		t.Error("expected", roomIds, "got", expectedRoomIds)
	}

	if unreaded != expectedUnreaded {
		t.Error("expected", expectedUnreaded, "got", unreaded)
	}
}
