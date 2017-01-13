package azure

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/arm/storage"
	"github.com/NeowayLabs/klb/tests/lib/retrier"
)

type StorageAccount struct {
	client   storage.AccountsClient
	ctx      context.Context
	logger   *log.Logger
	retrier  *retrier.Retrier
	resgroup string
}

func NewStorageAccount(
	ctx context.Context,
	t *testing.T,
	s *Session,
	logger *log.Logger,
	resgroup string,
) *StorageAccount {
	as := &StorageAccount{
		client:   storage.NewAccountsClient(s.SubscriptionID),
		ctx:      ctx,
		resgroup: resgroup,
		logger:   logger,
		retrier:  retrier.New(ctx, t, logger),
	}
	as.client.Authorizer = s.token
	return as
}

// AssertExists checks if availability sets exists in the resource group.
// Fail tests otherwise.
func (s *StorageAccount) AssertExists(t *testing.T, name string) {
	s.retrier.Run(getIDStorageAccount("AssertExists", name), func() error {
		_, err := s.client.GetProperties(s.resgroup, name)
		return err
	})
}

// AssertDeleted checks if resource was correctly deleted.
func (s *StorageAccount) AssertDeleted(t *testing.T, name string) {
	s.retrier.Run(getIDStorageAccount("AssertDeleted", name), func() error {
		_, err := s.client.GetProperties(s.resgroup, name)
		if err == nil {
			return fmt.Errorf("resource %s should not exist", name)
		}
		return nil
	})
}

// Delete the availability set
func (s *StorageAccount) Delete(t *testing.T, name string) {
	s.retrier.Run(getIDStorageAccount("Delete", name), func() error {
		_, err := s.client.Delete(s.resgroup, os.Getenv("STORAGE_ACCOUNT_NAME"))
		return err
	})
}

func getIDStorageAccount(method string, name string) string {
	return fmt.Sprintf("StorageAccount.%s:%s", method, name)
}