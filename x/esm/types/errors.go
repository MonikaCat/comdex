package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/esm module sentinel errors
var (
	ErrInvalidAsset                = sdkerrors.Register(ModuleName, 501, "invalid asset")
	ErrorDuplicateESMTriggerParams = sdkerrors.Register(ModuleName, 502, "Duplicate ESM Trigger Params for AppID")
	ErrAppDataNotFound             = sdkerrors.Register(ModuleName, 503, "App Data Not Found")
	ErrBadOfferCoinType            = sdkerrors.Register(ModuleName, 504, "Bad Offer Coin Type")
	ErrESMTriggerParamsNotFound    = sdkerrors.Register(ModuleName, 505, "ESM Trigger Params Not Found")
	ErrAmtExceedsTargetValue       = sdkerrors.Register(ModuleName, 506, "Amount Exceeds Target Value")
	ErrDepositForAppReached        = sdkerrors.Register(ModuleName, 507, "Deposit For AppID Reached")
	ErrESMAlreadyExecuted          = sdkerrors.Register(ModuleName, 508, "ESM Already Executed")
	ErrCircuitBreakerEnabled       = sdkerrors.Register(ModuleName, 509, "Circuit breaker is triggered")
	ErrorUnauthorized              = sdkerrors.Register(ModuleName, 510, "Unauthorized")
	ErrorAppDoesNotExists          = sdkerrors.Register(ModuleName, 511, "App Does Not Exists")
	ErrAppIDDoesNotExists          = sdkerrors.Register(ModuleName, 512, "App Id Does Not exist")
	ErrCoolOffPeriodPassed         = sdkerrors.Register(ModuleName, 513, "Cool off period has passed")
	ErrCurrentDepositNotReached    = sdkerrors.Register(ModuleName, 514, "Current Deposit Not Reached for App")
	ErrMarketDataNotFound          = sdkerrors.Register(ModuleName, 515, "MarketData not found for App")
	ErrCoolOffPeriodRemains        = sdkerrors.Register(ModuleName, 516, "Cool off period remaining")
	ErrorInvalidAmount             = sdkerrors.Register(ModuleName, 517, "invalid amount")
	ErrorInvalidFrom               = sdkerrors.Register(ModuleName, 518, "invalid from")
	ErrESMParamsNotFound           = sdkerrors.Register(ModuleName, 519, "ESM Params Not Found")
	ErrDepositForAppNotFound       = sdkerrors.Register(ModuleName, 520, "Deposit For AppID not found")
	ErrPriceNotFound               = sdkerrors.Register(ModuleName, 521, "Price not found")
)
