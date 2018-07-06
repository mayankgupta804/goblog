package dbclient

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/go_microservices/goblog/accountservice/model"
)

// IBoltClient interface
type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
}

// BoltClient real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer bc.closeDBConnection()
}

func (bc *BoltClient) closeDBConnection() {
	err := bc.boltDB.Close()
	if err != nil {
		fmt.Println("DB connection cannot be closed\n")
		log.Fatal(err)
	}
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initializeBucket() error {
}

func (bc *BoltClient) seedAccounts() {

}
