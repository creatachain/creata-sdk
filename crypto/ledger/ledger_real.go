// +build cgo,ledger,!test_ledger_mock

package ledger

import ledger "github.com/creatachain/ledger-creata-go"

// If ledger support (build tag) has been enabled, which implies a CGO dependency,
// set the discoverLedger function which is responsible for loading the Ledger
// device at runtime or returning an error.
func init() {
	discoverLedger = func() (SECP256K1, error) {
		device, err := ledger.FindLedgerCreataUserApp()
		if err != nil {
			return nil, err
		}

		return device, nil
	}
}
