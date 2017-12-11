package main

import (
	"bytes"
	"encoding/json"
	"github.com/penutty/momentservice/moment"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const (
	tMomentID = 1

	tUser  = "testuser0"
	tUser1 = "testuser1"
	tUser2 = "testuser2"
	tUser3 = "testuser3"

	tMessage1 = "testMessage1"
	tMessage2 = "testMessage2"

	tLat  = 1.00
	tLong = 1.00
)

var (
	defaultRecipients = []recipient{
		recipient{tUser1},
		recipient{tUser2},
		recipient{tUser3},
	}
	defaultMedia = []medium{
		medium{tMessage1, moment.DNE},
		medium{tMessage2, moment.DNE},
	}
)

type recipient struct {
	UserID string
}

type medium struct {
	Message string
	Mtype   uint8
}

func TestMain(m *testing.M) {
	// Create http.Client or http.Transport if necessary

	os.Exit(m.Run())
}

type reqFull struct {
	reqSemi
	Recipients []recipient
}

func NewReqFull(lat float32, long float32

type reqSemi struct {
	Latitude   float32
	Longitude  float32
	UserID     string
	Public     bool
	Hidden     bool
	CreateDate time.Time
	Media      []medium
}

type reqLoc struct {
	Latitude  float32
	Longitude float32
}

type reqLocUser struct {
	reqLoc
	userID string
}

func Test_postPrivateMoment(t *testing.T) {
	type test struct {
		b        reqFull
		expected error
	}
	tests := []test{
		test{body{tLat, tLong, tUser, false, false, time.Now().UTC(), defaultMedia, defaultRecipients}, nil},
	}

	for _, v := range tests {
		j, err := json.Marshal(v.b)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, MomentEndpoint, bytes.NewReader(j))

		a := MockApp()
		err = a.postPrivateMoment(req)
		assert.Exactly(t, v.expected, err)
	}
}

func Test_postPublicMoment(t *testing.T) {
	type test struct {
		b        reqSemi
		expected error
	}
	tests := []test{
		test{body{tLat, tLong, tUser, false, false, time.Now().UTC(), defaultMedia}, nil},
	}

	for _, v := range tests {
		j, err := json.Marshal(v.b)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, MomentEndpoint, bytes.NewReader(j))

		a := MockApp()
		err = a.postPublicMoment(req)
		assert.Exactly(t, v.expected, err)
	}
}

func Test_getHiddenMoment(t *testing.T) {
	type test struct {
		req      reqLoc
		expected error
	}
	tests := []test{
		test{reqbody{tLat, tLong}, nil},
	}

	for _, v := range tests {
		reqJson, err := json.Marshal(v.req)
		assert.Nil(t, err)
		req := httptest.NewRequest(http.MethodGet, MomentEndpoint, bytes.NewReader(reqJson))

		rec := httptest.NewRecorder()

		a := MockApp()
		err = a.getHiddenMoment(rec, req)
		assert.Exactly(t, v.expected, err)
	}
}

func Test_getLostMoment(t *testing.T) {
	type test struct {
		req      reqLocUser
		expected error
	}
	tests := []test{
		test{reqbody{tLat, tLong, tUser}, nil},
	}
	for _, v := range tests {
		reqJson, err := json.Marshal(v.req)
		assert.Nil(t, err)
		req := httptest.NewRequest(http.MethodGet, MomentEndpoint, bytes.NewReader(reqJson))

		rec := httptest.NewRecorder()

		a := MockApp()
		err = a.getLostMoment(rec, req)
		assert.Exactly(t, v.expected, err)
	}
}

func Test_getSharedMomentbyLocation(t *testing.T) {
	type reqbody struct {
		Latitude float32
		Longitude
	}
}

func MockApp() *app {
	c := new(MockClient)
	c.c = new(moment.MomentClient)

	a := new(app)
	a.c = c
	return a
}

type MockClient struct {
	c moment.Client
}

func (mc *MockClient) Err() error {
	return mc.c.Err()
}

func (mc *MockClient) FindPublic(db moment.DbRunner, f *moment.FindsRow) (int64, error) {
	return 1, nil
}

func (mc *MockClient) FindPrivate(db moment.DbRunner, f *moment.FindsRow) error {
	return nil
}

func (mc *MockClient) Share(db moment.DbRunnerTrans, s *moment.SharesRow, ms []*moment.RecipientsRow) error {
	return nil
}

func (mc *MockClient) CreatePublic(db moment.DbRunnerTrans, m *moment.MomentsRow, ms []*moment.MediaRow) error {
	return nil
}

func (mc *MockClient) CreatePrivate(db moment.DbRunnerTrans, m *moment.MomentsRow, ms []*moment.MediaRow, fs []*moment.FindsRow) error {
	return nil
}

func (mc *MockClient) NewMomentsRow(l *moment.Location, userID string, public bool, hidden bool, createDate *time.Time) *moment.MomentsRow {
	return mc.c.NewMomentsRow(l, userID, public, hidden, createDate)
}

func (mc *MockClient) NewLocation(lat float32, long float32) *moment.Location {
	return mc.c.NewLocation(lat, long)
}

func (mc *MockClient) NewMediaRow(momentID int64, userID string, mtype uint8, dir string) *moment.MediaRow {
	return mc.c.NewMediaRow(momentID, userID, mtype, dir)
}

func (mc *MockClient) NewFindsRow(momentID int64, userID string, found bool, findDate *time.Time) *moment.FindsRow {
	return mc.c.NewFindsRow(momentID, userID, found, findDate)
}

func (mc *MockClient) NewSharesRow(sharesID int64, momentID int64, userID string) *moment.SharesRow {
	return mc.c.NewSharesRow(sharesID, momentID, userID)
}

func (mc *MockClient) NewRecipientsRow(sharesID int64, all bool, recipientID string) *moment.RecipientsRow {
	return mc.c.NewRecipientsRow(sharesID, all, recipientID)
}

func (mc *MockClient) LocationShared(db moment.DbRunner, l *moment.Location, me string) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) LocationPublic(db moment.DbRunner, l *moment.Location) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) LocationHidden(db moment.DbRunner, l *moment.Location) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) LocationLost(db moment.DbRunner, l *moment.Location, me string) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) UserShared(db moment.DbRunner, you string, me string) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) UserLeft(db moment.DbRunner, me string) ([]*moment.Moment, error) {
	return nil, nil
}

func (mc *MockClient) UserFound(db moment.DbRunner, me string) ([]*moment.Moment, error) {
	return nil, nil
}
