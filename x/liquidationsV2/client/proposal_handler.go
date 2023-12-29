package client

import (
	"github.com/MonikaCat/comdex/v5/x/liquidationsV2/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var LiquidationsV2Handler = []govclient.ProposalHandler{
	govclient.NewProposalHandler(cli.NewCmdSubmitWhitelistingLiquidationProposal),
}
