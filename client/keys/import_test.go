package keys

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/creatachain/creata-sdk/client"
	"github.com/creatachain/creata-sdk/client/flags"
	"github.com/creatachain/creata-sdk/crypto/keyring"
	"github.com/creatachain/creata-sdk/testutil"
	sdk "github.com/creatachain/creata-sdk/types"
)

func Test_runImportCmd(t *testing.T) {
	cmd := ImportKeyCommand()
	cmd.Flags().AddFlagSet(Commands("home").PersistentFlags())
	mockIn := testutil.ApplyMockIODiscardOutErr(cmd)

	// Now add a temporary keybase
	kbHome := t.TempDir()
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, kbHome, mockIn)

	clientCtx := client.Context{}.WithKeyring(kb)
	ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

	require.NoError(t, err)
	t.Cleanup(func() {
		kb.Delete("keyname1") // nolint:errcheck
	})

	keyfile := filepath.Join(kbHome, "key.asc")
	armoredKey := `-----BEGIN AUGUSTEUM PRIVATE KEY-----
salt: A790BB721D1C094260EA84F5E5B72289
kdf: bcrypt

HbP+c6JmeJy9JXe2rbbF1QtCX1gLqGcDQPBXiCtFvP7/8wTZtVOPj8vREzhZ9ElO
3P7YnrzPQThG0Q+ZnRSbl9MAS8uFAM4mqm5r/Ys=
=f3l4
-----END AUGUSTEUM PRIVATE KEY-----
`
	require.NoError(t, ioutil.WriteFile(keyfile, []byte(armoredKey), 0644))

	mockIn.Reset("123456789\n")
	cmd.SetArgs([]string{
		"keyname1", keyfile,
		fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
	})
	require.NoError(t, cmd.ExecuteContext(ctx))
}
