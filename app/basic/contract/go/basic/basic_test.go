package basic_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ardanlabs/smartcontract/app/basic/contract/go/basic"
)

func TestBasic(t *testing.T) {

	ctx := context.Background()

	const newAccounts = 1
	const autoComit = true
	var accountBalance = big.NewInt(100)

	backend, err := ethereum.CreateSimulatedBackend(newAccounts, autoComit, accountBalance)
	if err != nil {
		t.Fatalf("unable to create simulated backend: %s", err)
	}
	defer backend.Close()

	clt, err := ethereum.NewClient(backend, backend.PrivateKeys[0])
	if err != nil {
		t.Fatalf("unable to create ethereum api: %s", err)
	}

	// ==============================================================

	const gasLimit = 1600000
	valueGwei := big.NewFloat(0.0)
	gasPrice := currency.GWei2Wei(big.NewFloat(39.576))
	transOpts, err := clt.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
	if err != nil {
		t.Fatalf("unable to create opts for deploy: %s", err)
	}

	contractID, tx, _, err := basic.DeployBasic(transOpts, clt.Backend)
	if err != nil {
		t.Fatalf("unable to deploy basic: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for deploy: %s", err)
	}

	testBasic, err := basic.NewBasic(contractID, clt.Backend)
	if err != nil {
		t.Fatalf("error creating basic: %s", err)
	}

	// ==========================================================================

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	ver, err := testBasic.Version(callOpts)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	if ver != "1.1" {
		t.Fatalf("should be able to get the correct version, got %s exp %s", ver, "1.1")
	}

	// ===========================================================================

	transOpts, err = clt.NewTransactOpts(ctx, gasLimit, gasPrice, valueGwei)
	if err != nil {
		t.Fatalf("unable to create transaction opts for setItems: %s", err)
	}

	key := "bill"
	value := big.NewInt(1_00_000)

	tx, err = testBasic.SetItem(transOpts, key, value)
	if err != nil {
		t.Fatalf("should be able to set item: %s", err)
	}

	if _, err := clt.WaitMined(ctx, tx); err != nil {
		t.Fatalf("waiting for setitem: %s", err)
	}

	// ============================================================================

	callOpts, err = clt.NewCallOpts(ctx)
	if err != nil {
		t.Fatalf("unable to create call opts: %s", err)
	}

	item, err := testBasic.Items(callOpts, key)
	if err != nil {
		t.Fatalf("should be able to retrieve item: %s", err)
	}

	if item.Cmp(value) != 0 {
		t.Fatalf("wrong value, got %s exp %s", item, value)
	}
}
