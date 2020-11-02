package keeper

import (
	"encoding/hex"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/validator/types"
	cautil "gitlab.bianjie.ai/irita-pro/iritamod/utils/ca"
)

// keeper of the validator store
type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey

	paramstore paramtypes.Subspace
	hooks      staking.StakingHooks
}

func NewKeeper(cdc codec.Marshaler, storeKey sdk.StoreKey, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramstore: ps,
	}
}

// SetHooks sets the validator hooks
func (k *Keeper) SetHooks(sh staking.StakingHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}

	k.hooks = sh

	return k
}

// Create create a new validator
func (k Keeper) Create(ctx sdk.Context, msg types.MsgCreateValidator) (tmbytes.HexBytes, error) {
	if k.HasValidatorName(ctx, msg.Name) {
		return nil, types.ErrValidatorNameExists
	}

	cert, err := k.VerifyCert(ctx, msg.Certificate)
	if err != nil {
		return nil, err
	}

	pk, err := cautil.GetPubkeyFromCert(cert)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
	}

	if _, found := k.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk)); found {
		return nil, types.ErrValidatorPubkeyExists
	}

	operator, _ := hex.DecodeString(msg.Operator)
	id := tmbytes.HexBytes(tmhash.Sum(msg.GetSignBytes()))

	validator := types.NewValidator(
		id,
		msg.Name,
		msg.Description,
		pk,
		msg.Certificate,
		msg.Power,
		operator,
	)

	k.SetValidator(ctx, validator)
	k.SetValidatorConsAddrIndex(ctx, id, validator.GetConsAddr())
	k.EnqueueValidatorsUpdate(ctx, validator, msg.Power)

	k.hooks.AfterValidatorCreated(ctx, validator.GetOperator())
	k.hooks.AfterValidatorBonded(ctx, validator.GetConsAddr(), validator.GetOperator())
	return id, nil
}

// UpdateValidator updates an existing validator record
func (k Keeper) Update(ctx sdk.Context, msg types.MsgUpdateValidator) error {
	if k.HasValidatorName(ctx, msg.Name) {
		return types.ErrValidatorNameExists
	}

	id, _ := hex.DecodeString(msg.Id)
	validator, found := k.GetValidator(ctx, id)
	if !found {
		return types.ErrUnknownValidator
	}

	if len(msg.Certificate) > 0 && msg.Certificate != validator.Certificate {
		cert, err := k.VerifyCert(ctx, msg.Certificate)
		if err != nil {
			return err
		}
		pk, err := cautil.GetPubkeyFromCert(cert)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
		}
		pkStr := sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pk)

		// delete pubkey related index
		k.DeleteValidatorConsAddrIndex(ctx, validator.GetConsAddr())
		// delete from tendermint validator set
		k.EnqueueValidatorsUpdate(ctx, validator, 0)

		validator.Pubkey = pkStr
		validator.Certificate = msg.Certificate
		k.SetValidatorConsAddrIndex(ctx, id, validator.GetConsAddr())
		k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
	}
	if msg.Power > 0 {
		validator.Power = msg.Power
		// override it if already exists
		k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
	}
	if len(msg.Description) > 0 && msg.Description != types.DoNotModifyDesc {
		validator.Description = msg.Description
	}
	if len(msg.Name) > 0 && msg.Name != types.DoNotModifyDesc {
		//delete old name
		store := ctx.KVStore(k.storeKey)
		store.Delete(types.GetValidatorNameKey(validator.Name))

		validator.Name = msg.Name
	}
	k.SetValidator(ctx, validator)
	return nil
}

// Remove deletes an existing validator record
func (k Keeper) Remove(ctx sdk.Context, msg types.MsgRemoveValidator) error {
	id, _ := hex.DecodeString(msg.Id)
	validator, found := k.GetValidator(ctx, id)
	if !found {
		return types.ErrUnknownValidator
	}

	k.DeleteValidator(ctx, validator)
	k.DeleteValidatorConsAddrIndex(ctx, validator.GetConsAddr())
	// delete from tendermint validator set
	k.EnqueueValidatorsUpdate(ctx, validator, 0)

	k.hooks.AfterValidatorRemoved(ctx, validator.GetConsAddr(), validator.GetOperator())
	return nil
}

// SetValidator sets the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	id, _ := hex.DecodeString(validator.Id)
	// set validator by id
	bz := k.cdc.MustMarshalBinaryBare(&validator)
	store.Set(types.GetValidatorIDKey(id), bz)

	bz = k.cdc.MustMarshalBinaryBare(&gogotypes.BytesValue{Value: id})
	store.Set(types.GetValidatorNameKey(validator.Name), bz)
}

// GetValidator returns validator with id
func (k Keeper) GetValidator(ctx sdk.Context, id tmbytes.HexBytes) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorIDKey(id))
	if value == nil {
		return validator, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &validator)
	return validator, true
}

// HasValidatorName returns true or false with name
func (k Keeper) HasValidatorName(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorNameKey(name))
}

// DeleteValidator deletes the validator with id
func (k Keeper) DeleteValidator(ctx sdk.Context, validator types.Validator) {
	id, _ := hex.DecodeString(validator.Id)
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorIDKey(id))
	store.Delete(types.GetValidatorNameKey(validator.Name))
}

// SetValidatorConsAddrIndex sets the validator index by pubkey
func (k Keeper) SetValidatorConsAddrIndex(ctx sdk.Context, id tmbytes.HexBytes, addr sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.BytesValue{Value: id})
	store.Set(types.GetValidatorConsAddrKey(addr), bz)
}

// DeleteValidatorConsAddrIndex deletes the validator index with pubkey
func (k Keeper) DeleteValidatorConsAddrIndex(ctx sdk.Context, addr sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorConsAddrKey(addr))
}

// GetValidatorByConsAddr returns validator with pubkey
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorConsAddrKey(addr))
	if value == nil {
		return validator, false
	}

	var id gogotypes.BytesValue
	k.cdc.MustUnmarshalBinaryBare(value, &id)

	return k.GetValidator(ctx, id.Value)
}

// EnqueueValidatorsUpdate enqueue to the validators update queue
func (k Keeper) EnqueueValidatorsUpdate(ctx sdk.Context, validator types.Validator, power int64) {
	// do not update this validator if already jailed
	if validator.Jailed {
		return
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetValidatorUpdateQueueKey(validator.Pubkey), bz)
}

// DequeueValidatorsUpdate dequeue from the validators update queue
func (k Keeper) DequeueValidatorsUpdate(ctx sdk.Context, pubkey string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorUpdateQueueKey(pubkey))
}

// IterateUpdateValidators iterates through the validators update queue
func (k Keeper) IterateUpdateValidators(ctx sdk.Context, fn func(index int64, pubkey string, power int64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsUpdateQueueKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var power gogotypes.Int64Value
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &power)
		stop := fn(i, string(iterator.Key()[1:]), power.Value)

		if stop {
			break
		}
		i++
	}
}

// IterateValidators iterates through the validator set and perform the provided function
func (k Keeper) IterateValidators(ctx sdk.Context, fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		stop := fn(i, validator)

		if stop {
			break
		}
		i++
	}
}

// GetAllValidators gets the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		validators = append(validators, validator)
	}

	return validators
}

// ValidatorByID return the validator imformation by id
func (k Keeper) ValidatorByID(ctx sdk.Context, id tmbytes.HexBytes) staking.ValidatorI {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorIDKey(id))

	var validator types.Validator
	k.cdc.MustUnmarshalBinaryBare(bz, &validator)
	return validator
}

// Validator return the validator imformation by valAddr
func (k Keeper) Validator(ctx sdk.Context, valAddr sdk.ValAddress) staking.ValidatorI {
	return k.ValidatorByConsAddr(ctx, sdk.ConsAddress(valAddr))
}

// ValidatorByConsAddr return the validator imformation by consAddr
func (k Keeper) ValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) staking.ValidatorI {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return nil
	}
	return validator
}

// Slash not implement
func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, i int64, i2 int64, dec sdk.Dec) {}

// Jail disable the validator
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	if validator.Jailed {
		panic(fmt.Sprintf("cannot jail already jailed validator, validator: %v\n", validator))
	}

	k.EnqueueValidatorsUpdate(ctx, validator, 0)
	validator.Jailed = true
	k.SetValidator(ctx, validator)
}

// Unjail enable the validator
func (k Keeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}
	if !validator.Jailed {
		panic(fmt.Sprintf("cannot unjail already unjailed validator, validator: %v\n", validator))
	}
	validator.Jailed = false
	k.SetValidator(ctx, validator)
	k.EnqueueValidatorsUpdate(ctx, validator, validator.Power)
}

func (k Keeper) Delegation(context sdk.Context, accAddr sdk.AccAddress, consAddr sdk.ValAddress) staking.DelegationI {
	return staking.Delegation{}
}

func (k Keeper) MaxValidators(context sdk.Context) uint32 {
	return 0
}

// get the group of the bonded validators
func (k Keeper) GetLastValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		if !validator.Jailed {
			validators = append(validators, validator)
		}
	}

	return validators
}

func (k *Keeper) IterateBondedValidatorsByPower(ctx sdk.Context, fn func(index int64, validator staking.ValidatorI) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		var validator types.Validator
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &validator)
		if !validator.Jailed {
			stop := fn(i, validator)
			if stop {
				break
			}
			i++
		}
	}
}

func (k *Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	total := sdk.NewInt(0)
	k.IterateValidators(ctx,
		func(index int64, validator staking.ValidatorI) bool {
			if !validator.IsJailed() {
				total = total.Sub(validator.GetTokens())
			}
			return false
		},
	)
	return total
}

func (k *Keeper) IterateDelegations(
	ctx sdk.Context, delegator sdk.AccAddress,
	fn func(index int64, delegation staking.DelegationI) (stop bool),
) {
}

func (k *Keeper) VerifyCert(ctx sdk.Context, certStr string) (cert cautil.Cert, err error) {
	rootCertStr, _ := k.GetRootCert(ctx)
	rootCert, err := cautil.ReadCertificateFromMem([]byte(rootCertStr))
	if err != nil {
		return cert, sdkerrors.Wrap(types.ErrInvalidRootCert, err.Error())
	}

	cert, err = cautil.ReadCertificateFromMem([]byte(certStr))
	if err != nil {
		return cert, sdkerrors.Wrap(types.ErrInvalidCert, err.Error())
	}

	if err = cert.VerifyCertFromRoot(rootCert); err != nil {
		return cert, sdkerrors.Wrapf(types.ErrInvalidCert, "cannot be verified by root certificate, err: %s", err.Error())
	}

	return cert, nil
}
