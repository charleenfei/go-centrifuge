// +build integration

package did

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/identity/ethid"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/transactions"
	"github.com/stretchr/testify/assert"
)

var cfg config.Configuration
var ctx = map[string]interface{}{}

func TestMain(m *testing.M) {
	var bootstappers = []bootstrap.TestBootstrapper{
		&testlogging.TestLoggingBootstrapper{},
		&config.Bootstrapper{},
		&leveldb.Bootstrapper{},
		transactions.Bootstrapper{},
		&queue.Bootstrapper{},
		ethereum.Bootstrapper{},
		&ethid.Bootstrapper{},
		&configstore.Bootstrapper{},
		&Bootstrapper{},
		&queue.Starter{},
	}

	bootstrap.RunTestBootstrappers(bootstappers, ctx)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	result := m.Run()
	bootstrap.RunTestTeardown(bootstappers)
	os.Exit(result)
}

func TestCreateIdentity_successful(t *testing.T) {
	service := ctx[BootstrappedDIDService].(Service)
	did, _, err := service.CreateIdentity(context.Background())
	assert.Nil(t, err, "create identity should be successful")

	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)

	time.Sleep(2000 * time.Millisecond)

	contractCode, err := client.GetEthClient().CodeAt(context.Background(), did.toAddress(), nil)
	assert.Nil(t, err, "should be successful to get the contract code")

	assert.Equal(t, true, len(contractCode) > 3000, "current contract code should be arround 3378 bytes")

	fmt.Println(did)

}