// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GuideNFTMetaData contains all meta data concerning the GuideNFT contract.
var GuideNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NFTBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"NFTListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokenInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMinted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"deListNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"listNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GuideNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use GuideNFTMetaData.ABI instead.
var GuideNFTABI = GuideNFTMetaData.ABI

// GuideNFT is an auto generated Go binding around an Ethereum contract.
type GuideNFT struct {
	GuideNFTCaller     // Read-only binding to the contract
	GuideNFTTransactor // Write-only binding to the contract
	GuideNFTFilterer   // Log filterer for contract events
}

// GuideNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuideNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuideNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuideNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuideNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuideNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuideNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuideNFTSession struct {
	Contract     *GuideNFT         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GuideNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuideNFTCallerSession struct {
	Contract *GuideNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GuideNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuideNFTTransactorSession struct {
	Contract     *GuideNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GuideNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuideNFTRaw struct {
	Contract *GuideNFT // Generic contract binding to access the raw methods on
}

// GuideNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuideNFTCallerRaw struct {
	Contract *GuideNFTCaller // Generic read-only contract binding to access the raw methods on
}

// GuideNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuideNFTTransactorRaw struct {
	Contract *GuideNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuideNFT creates a new instance of GuideNFT, bound to a specific deployed contract.
func NewGuideNFT(address common.Address, backend bind.ContractBackend) (*GuideNFT, error) {
	contract, err := bindGuideNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuideNFT{GuideNFTCaller: GuideNFTCaller{contract: contract}, GuideNFTTransactor: GuideNFTTransactor{contract: contract}, GuideNFTFilterer: GuideNFTFilterer{contract: contract}}, nil
}

// NewGuideNFTCaller creates a new read-only instance of GuideNFT, bound to a specific deployed contract.
func NewGuideNFTCaller(address common.Address, caller bind.ContractCaller) (*GuideNFTCaller, error) {
	contract, err := bindGuideNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuideNFTCaller{contract: contract}, nil
}

// NewGuideNFTTransactor creates a new write-only instance of GuideNFT, bound to a specific deployed contract.
func NewGuideNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*GuideNFTTransactor, error) {
	contract, err := bindGuideNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuideNFTTransactor{contract: contract}, nil
}

// NewGuideNFTFilterer creates a new log filterer instance of GuideNFT, bound to a specific deployed contract.
func NewGuideNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*GuideNFTFilterer, error) {
	contract, err := bindGuideNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuideNFTFilterer{contract: contract}, nil
}

// bindGuideNFT binds a generic wrapper to an already deployed contract.
func bindGuideNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GuideNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuideNFT *GuideNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuideNFT.Contract.GuideNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuideNFT *GuideNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuideNFT.Contract.GuideNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuideNFT *GuideNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuideNFT.Contract.GuideNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuideNFT *GuideNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuideNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuideNFT *GuideNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuideNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuideNFT *GuideNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuideNFT.Contract.contract.Transact(opts, method, params...)
}

// TokenInfos is a free data retrieval call binding the contract method 0x8676192b.
//
// Solidity: function _tokenInfos(uint256 ) view returns(uint256 price, address owner, bool isListed, bool isMinted)
func (_GuideNFT *GuideNFTCaller) TokenInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price    *big.Int
	Owner    common.Address
	IsListed bool
	IsMinted bool
}, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "_tokenInfos", arg0)

	outstruct := new(struct {
		Price    *big.Int
		Owner    common.Address
		IsListed bool
		IsMinted bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsListed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IsMinted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// TokenInfos is a free data retrieval call binding the contract method 0x8676192b.
//
// Solidity: function _tokenInfos(uint256 ) view returns(uint256 price, address owner, bool isListed, bool isMinted)
func (_GuideNFT *GuideNFTSession) TokenInfos(arg0 *big.Int) (struct {
	Price    *big.Int
	Owner    common.Address
	IsListed bool
	IsMinted bool
}, error) {
	return _GuideNFT.Contract.TokenInfos(&_GuideNFT.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0x8676192b.
//
// Solidity: function _tokenInfos(uint256 ) view returns(uint256 price, address owner, bool isListed, bool isMinted)
func (_GuideNFT *GuideNFTCallerSession) TokenInfos(arg0 *big.Int) (struct {
	Price    *big.Int
	Owner    common.Address
	IsListed bool
	IsMinted bool
}, error) {
	return _GuideNFT.Contract.TokenInfos(&_GuideNFT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GuideNFT *GuideNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GuideNFT *GuideNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GuideNFT.Contract.BalanceOf(&_GuideNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_GuideNFT *GuideNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _GuideNFT.Contract.BalanceOf(&_GuideNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GuideNFT.Contract.GetApproved(&_GuideNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _GuideNFT.Contract.GetApproved(&_GuideNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GuideNFT *GuideNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GuideNFT *GuideNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GuideNFT.Contract.IsApprovedForAll(&_GuideNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_GuideNFT *GuideNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _GuideNFT.Contract.IsApprovedForAll(&_GuideNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GuideNFT *GuideNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GuideNFT *GuideNFTSession) Name() (string, error) {
	return _GuideNFT.Contract.Name(&_GuideNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GuideNFT *GuideNFTCallerSession) Name() (string, error) {
	return _GuideNFT.Contract.Name(&_GuideNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GuideNFT.Contract.OwnerOf(&_GuideNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_GuideNFT *GuideNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _GuideNFT.Contract.OwnerOf(&_GuideNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GuideNFT *GuideNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GuideNFT *GuideNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GuideNFT.Contract.SupportsInterface(&_GuideNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GuideNFT *GuideNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GuideNFT.Contract.SupportsInterface(&_GuideNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GuideNFT *GuideNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GuideNFT *GuideNFTSession) Symbol() (string, error) {
	return _GuideNFT.Contract.Symbol(&_GuideNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GuideNFT *GuideNFTCallerSession) Symbol() (string, error) {
	return _GuideNFT.Contract.Symbol(&_GuideNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GuideNFT *GuideNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _GuideNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GuideNFT *GuideNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GuideNFT.Contract.TokenURI(&_GuideNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_GuideNFT *GuideNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _GuideNFT.Contract.TokenURI(&_GuideNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.Approve(&_GuideNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.Approve(&_GuideNFT.TransactOpts, to, tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 _tokenId) payable returns()
func (_GuideNFT *GuideNFTTransactor) BuyNFT(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "buyNFT", _tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 _tokenId) payable returns()
func (_GuideNFT *GuideNFTSession) BuyNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.BuyNFT(&_GuideNFT.TransactOpts, _tokenId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 _tokenId) payable returns()
func (_GuideNFT *GuideNFTTransactorSession) BuyNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.BuyNFT(&_GuideNFT.TransactOpts, _tokenId)
}

// DeListNFT is a paid mutator transaction binding the contract method 0x9f606dc2.
//
// Solidity: function deListNFT(uint256 _tokenId) returns()
func (_GuideNFT *GuideNFTTransactor) DeListNFT(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "deListNFT", _tokenId)
}

// DeListNFT is a paid mutator transaction binding the contract method 0x9f606dc2.
//
// Solidity: function deListNFT(uint256 _tokenId) returns()
func (_GuideNFT *GuideNFTSession) DeListNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.DeListNFT(&_GuideNFT.TransactOpts, _tokenId)
}

// DeListNFT is a paid mutator transaction binding the contract method 0x9f606dc2.
//
// Solidity: function deListNFT(uint256 _tokenId) returns()
func (_GuideNFT *GuideNFTTransactorSession) DeListNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.DeListNFT(&_GuideNFT.TransactOpts, _tokenId)
}

// ListNFT is a paid mutator transaction binding the contract method 0x94383f14.
//
// Solidity: function listNFT(uint256 _tokenId, uint256 _price) returns(uint256)
func (_GuideNFT *GuideNFTTransactor) ListNFT(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "listNFT", _tokenId, _price)
}

// ListNFT is a paid mutator transaction binding the contract method 0x94383f14.
//
// Solidity: function listNFT(uint256 _tokenId, uint256 _price) returns(uint256)
func (_GuideNFT *GuideNFTSession) ListNFT(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.ListNFT(&_GuideNFT.TransactOpts, _tokenId, _price)
}

// ListNFT is a paid mutator transaction binding the contract method 0x94383f14.
//
// Solidity: function listNFT(uint256 _tokenId, uint256 _price) returns(uint256)
func (_GuideNFT *GuideNFTTransactorSession) ListNFT(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.ListNFT(&_GuideNFT.TransactOpts, _tokenId, _price)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.SafeTransferFrom(&_GuideNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.SafeTransferFrom(&_GuideNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GuideNFT *GuideNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GuideNFT *GuideNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GuideNFT.Contract.SafeTransferFrom0(&_GuideNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_GuideNFT *GuideNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GuideNFT.Contract.SafeTransferFrom0(&_GuideNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GuideNFT *GuideNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GuideNFT *GuideNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GuideNFT.Contract.SetApprovalForAll(&_GuideNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GuideNFT *GuideNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GuideNFT.Contract.SetApprovalForAll(&_GuideNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.TransferFrom(&_GuideNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_GuideNFT *GuideNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _GuideNFT.Contract.TransferFrom(&_GuideNFT.TransactOpts, from, to, tokenId)
}

// GuideNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GuideNFT contract.
type GuideNFTApprovalIterator struct {
	Event *GuideNFTApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTApproval represents a Approval event raised by the GuideNFT contract.
type GuideNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*GuideNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTApprovalIterator{contract: _GuideNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GuideNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTApproval)
				if err := _GuideNFT.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) ParseApproval(log types.Log) (*GuideNFTApproval, error) {
	event := new(GuideNFTApproval)
	if err := _GuideNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuideNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GuideNFT contract.
type GuideNFTApprovalForAllIterator struct {
	Event *GuideNFTApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTApprovalForAll represents a ApprovalForAll event raised by the GuideNFT contract.
type GuideNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GuideNFT *GuideNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*GuideNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTApprovalForAllIterator{contract: _GuideNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GuideNFT *GuideNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GuideNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTApprovalForAll)
				if err := _GuideNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_GuideNFT *GuideNFTFilterer) ParseApprovalForAll(log types.Log) (*GuideNFTApprovalForAll, error) {
	event := new(GuideNFTApprovalForAll)
	if err := _GuideNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuideNFTNFTBoughtIterator is returned from FilterNFTBought and is used to iterate over the raw logs and unpacked data for NFTBought events raised by the GuideNFT contract.
type GuideNFTNFTBoughtIterator struct {
	Event *GuideNFTNFTBought // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTNFTBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTNFTBought)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTNFTBought)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTNFTBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTNFTBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTNFTBought represents a NFTBought event raised by the GuideNFT contract.
type GuideNFTNFTBought struct {
	Buyer     common.Address
	Seller    common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTBought is a free log retrieval operation binding the contract event 0x7a0a7460f8075212b07f19a9ab30b15d4d17c0984bfc2670bcd0430a46913d87.
//
// Solidity: event NFTBought(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 price, uint256 timestamp)
func (_GuideNFT *GuideNFTFilterer) FilterNFTBought(opts *bind.FilterOpts, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (*GuideNFTNFTBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "NFTBought", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTNFTBoughtIterator{contract: _GuideNFT.contract, event: "NFTBought", logs: logs, sub: sub}, nil
}

// WatchNFTBought is a free log subscription operation binding the contract event 0x7a0a7460f8075212b07f19a9ab30b15d4d17c0984bfc2670bcd0430a46913d87.
//
// Solidity: event NFTBought(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 price, uint256 timestamp)
func (_GuideNFT *GuideNFTFilterer) WatchNFTBought(opts *bind.WatchOpts, sink chan<- *GuideNFTNFTBought, buyer []common.Address, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "NFTBought", buyerRule, sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTNFTBought)
				if err := _GuideNFT.contract.UnpackLog(event, "NFTBought", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTBought is a log parse operation binding the contract event 0x7a0a7460f8075212b07f19a9ab30b15d4d17c0984bfc2670bcd0430a46913d87.
//
// Solidity: event NFTBought(address indexed buyer, address indexed seller, uint256 indexed tokenId, uint256 price, uint256 timestamp)
func (_GuideNFT *GuideNFTFilterer) ParseNFTBought(log types.Log) (*GuideNFTNFTBought, error) {
	event := new(GuideNFTNFTBought)
	if err := _GuideNFT.contract.UnpackLog(event, "NFTBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuideNFTNFTDelistedIterator is returned from FilterNFTDelisted and is used to iterate over the raw logs and unpacked data for NFTDelisted events raised by the GuideNFT contract.
type GuideNFTNFTDelistedIterator struct {
	Event *GuideNFTNFTDelisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTNFTDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTNFTDelisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTNFTDelisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTNFTDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTNFTDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTNFTDelisted represents a NFTDelisted event raised by the GuideNFT contract.
type GuideNFTNFTDelisted struct {
	Operator common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNFTDelisted is a free log retrieval operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed operator, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) FilterNFTDelisted(opts *bind.FilterOpts, operator []common.Address, tokenId []*big.Int) (*GuideNFTNFTDelistedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "NFTDelisted", operatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTNFTDelistedIterator{contract: _GuideNFT.contract, event: "NFTDelisted", logs: logs, sub: sub}, nil
}

// WatchNFTDelisted is a free log subscription operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed operator, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) WatchNFTDelisted(opts *bind.WatchOpts, sink chan<- *GuideNFTNFTDelisted, operator []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "NFTDelisted", operatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTNFTDelisted)
				if err := _GuideNFT.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTDelisted is a log parse operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed operator, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) ParseNFTDelisted(log types.Log) (*GuideNFTNFTDelisted, error) {
	event := new(GuideNFTNFTDelisted)
	if err := _GuideNFT.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuideNFTNFTListedIterator is returned from FilterNFTListed and is used to iterate over the raw logs and unpacked data for NFTListed events raised by the GuideNFT contract.
type GuideNFTNFTListedIterator struct {
	Event *GuideNFTNFTListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTNFTListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTNFTListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTNFTListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTNFTListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTNFTListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTNFTListed represents a NFTListed event raised by the GuideNFT contract.
type GuideNFTNFTListed struct {
	Recipient common.Address
	TokenId   *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTListed is a free log retrieval operation binding the contract event 0xbbf44e68296791df5535457da734bf1017c4b233a740af6d7a89b9f1af4ecb42.
//
// Solidity: event NFTListed(address indexed recipient, uint256 indexed tokenId, uint256 price)
func (_GuideNFT *GuideNFTFilterer) FilterNFTListed(opts *bind.FilterOpts, recipient []common.Address, tokenId []*big.Int) (*GuideNFTNFTListedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "NFTListed", recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTNFTListedIterator{contract: _GuideNFT.contract, event: "NFTListed", logs: logs, sub: sub}, nil
}

// WatchNFTListed is a free log subscription operation binding the contract event 0xbbf44e68296791df5535457da734bf1017c4b233a740af6d7a89b9f1af4ecb42.
//
// Solidity: event NFTListed(address indexed recipient, uint256 indexed tokenId, uint256 price)
func (_GuideNFT *GuideNFTFilterer) WatchNFTListed(opts *bind.WatchOpts, sink chan<- *GuideNFTNFTListed, recipient []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "NFTListed", recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTNFTListed)
				if err := _GuideNFT.contract.UnpackLog(event, "NFTListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNFTListed is a log parse operation binding the contract event 0xbbf44e68296791df5535457da734bf1017c4b233a740af6d7a89b9f1af4ecb42.
//
// Solidity: event NFTListed(address indexed recipient, uint256 indexed tokenId, uint256 price)
func (_GuideNFT *GuideNFTFilterer) ParseNFTListed(log types.Log) (*GuideNFTNFTListed, error) {
	event := new(GuideNFTNFTListed)
	if err := _GuideNFT.contract.UnpackLog(event, "NFTListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuideNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GuideNFT contract.
type GuideNFTTransferIterator struct {
	Event *GuideNFTTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuideNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuideNFTTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuideNFTTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuideNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuideNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuideNFTTransfer represents a Transfer event raised by the GuideNFT contract.
type GuideNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*GuideNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &GuideNFTTransferIterator{contract: _GuideNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GuideNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _GuideNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuideNFTTransfer)
				if err := _GuideNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_GuideNFT *GuideNFTFilterer) ParseTransfer(log types.Log) (*GuideNFTTransfer, error) {
	event := new(GuideNFTTransfer)
	if err := _GuideNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
