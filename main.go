package main

import (
	"fmt"
	"github.com/pkg/errors"

	gcHorizonclient "github.com/stellar/go/clients/horizonclient"
	gKeypair "github.com/stellar/go/keypair"
	gNet "github.com/stellar/go/network"
	gTxnbuild "github.com/stellar/go/txnbuild"
)

func main() {
	kp, _ := gKeypair.Parse("SAQTXJ5C3UMS5ZIKQ6J4NHVDLDKOHQFEQZLW4NITWDPZVIOBLZUK7J2Z")
	client := gcHorizonclient.DefaultPublicNetClient
	ar := gcHorizonclient.AccountRequest{AccountID: kp.Address()}
	//
	sourceAccount, err := client.AccountDetail(ar)
	if err != nil {
		fmt.Println(errors.Wrap(err, "gcHorizonclient.AccountDetail"))
	}

	aAsset := gTxnbuild.CreditAsset{"AQUA","GBNZILSTVQZ4R7IKQDGHYGY2QXL5QOFJYQMXPKWRRM5PAV7Y4M67AQUA"}
	bAsset := gTxnbuild.CreditAsset{"USDC","GA5ZSEJYB37JRC5AVCIA5MOP4RHTM335X2KGX3IHOJAPP5RE34K4KZVN"}

	var assetList []gTxnbuild.Asset
	assetList = append(assetList, aAsset)
	assetList = append(assetList, bAsset)

	op := gTxnbuild.PathPaymentStrictSend{
		SendAsset:   gTxnbuild.NativeAsset{},
		SendAmount:  "0.1",
		Destination: kp.Address(),
		DestAsset:   gTxnbuild.NativeAsset{},
		DestMin:     "0.1001",
		Path:        assetList,
	}

	tx, err := gTxnbuild.NewTransaction(
		gTxnbuild.TransactionParams{
			SourceAccount:        &sourceAccount,
			IncrementSequenceNum: true,
			Operations:           []gTxnbuild.Operation{&op},
			BaseFee:              200,//gTxnbuild.MinBaseFee,
			Timebounds:           gTxnbuild.NewInfiniteTimeout(),
		},
	)

	if err != nil {
		fmt.Println(errors.Wrap(err, "gTxnbuild.NewTransaction"))
	}

	tx, err = tx.Sign(gNet.PublicNetworkPassphrase, kp.(*gKeypair.Full))
	if err != nil {
		fmt.Println(errors.Wrap(err, "tx.Sign"))
	}

	resp, err := client.SubmitTransaction(tx)
	fmt.Println("resp")
	fmt.Println(resp)
	fmt.Println("err")
	fmt.Println(err)
	if err != nil {
		fmt.Println(errors.Wrap(err, "gcHorizonclient.client.SubmitTransaction"))
	}
}
