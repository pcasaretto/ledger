package ledger

import (
	"context"
	"testing"

	"github.com/numary/ledger/pkg/core"
	"github.com/numary/ledger/pkg/core/monetary"
	"github.com/numary/ledger/pkg/storage"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestVolumeAggregator(t *testing.T) {
	withContainer(fx.Invoke(func(lc fx.Lifecycle, storageDriver storage.Driver) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				name := uuid.New()

				store, _, err := storageDriver.GetStore(context.Background(), name, true)
				if err != nil {
					return err
				}

				_, err = store.Initialize(context.Background())
				if err != nil {
					return err
				}

				firstTxLog := core.NewTransactionLog(nil, core.Transaction{
					ID: 0,
					TransactionData: core.TransactionData{
						Postings: []core.Posting{
							{
								Source:      "bob",
								Destination: "zozo",
								Amount:      monetary.NewInt(100),
								Asset:       "USD",
							},
						},
					},
				})
				secondTxLog := core.NewTransactionLog(&firstTxLog, core.Transaction{
					ID: 1,
					TransactionData: core.TransactionData{
						Postings: []core.Posting{
							{
								Source:      "zozo",
								Destination: "alice",
								Amount:      monetary.NewInt(100),
								Asset:       "USD",
							},
						},
					},
				})
				require.NoError(t, store.AppendLog(context.Background(), firstTxLog, secondTxLog))

				volumeAggregator := newVolumeAggregator(store)
				firstTx := volumeAggregator.nextTx()
				require.NoError(t, firstTx.transfer(context.Background(), "bob", "alice", "USD", monetary.NewInt(100)))
				require.NoError(t, firstTx.transfer(context.Background(), "bob", "zoro", "USD", monetary.NewInt(50)))

				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(250),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(200),
						},
					},
					"zoro": {
						"USD": {
							Input: monetary.NewInt(50),
						},
					},
				}, firstTx.postCommitVolumes())
				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(100),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(100),
						},
					},
					"zoro": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(0),
						},
					},
				}, firstTx.preCommitVolumes())

				secondTx := volumeAggregator.nextTx()
				require.NoError(t, secondTx.transfer(context.Background(), "alice", "fred", "USD", monetary.NewInt(50)))
				require.NoError(t, secondTx.transfer(context.Background(), "bob", "fred", "USD", monetary.NewInt(25)))
				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(275),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input:  monetary.NewInt(200),
							Output: monetary.NewInt(50),
						},
					},
					"fred": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(75),
						},
					},
				}, secondTx.postCommitVolumes())
				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(250),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(200),
						},
					},
					"fred": core.AssetsVolumes{
						"USD": {},
					},
				}, secondTx.preCommitVolumes())

				aggregatedPostVolumes := volumeAggregator.aggregatedPostCommitVolumes()
				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(275),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input:  monetary.NewInt(200),
							Output: monetary.NewInt(50),
						},
					},
					"fred": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(75),
						},
					},
					"zoro": core.AssetsVolumes{
						"USD": {
							Input:  monetary.NewInt(50),
							Output: monetary.NewInt(0),
						},
					},
				}, aggregatedPostVolumes)

				aggregatedPreVolumes := volumeAggregator.aggregatedPreCommitVolumes()
				require.Equal(t, core.AccountsAssetsVolumes{
					"bob": core.AssetsVolumes{
						"USD": {
							Output: monetary.NewInt(100),
						},
					},
					"alice": core.AssetsVolumes{
						"USD": {
							Input: monetary.NewInt(100),
						},
					},
					"fred": core.AssetsVolumes{
						"USD": {},
					},
					"zoro": core.AssetsVolumes{
						"USD": {},
					},
				}, aggregatedPreVolumes)

				return nil
			},
		})
	}))
}
