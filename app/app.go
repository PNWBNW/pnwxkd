package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
    "github.com/cosmos/cosmos-sdk/simapp"
    "github.com/cosmos/cosmos-sdk/store"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/x/auth"
    authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
    "github.com/cosmos/cosmos-sdk/x/bank"
    banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
    "github.com/cosmos/cosmos-sdk/x/genutil"
    "github.com/cosmos/cosmos-sdk/x/params"

    // Custom PNW modules
    pnwzkcredential "github.com/pnwxkd/x/zkcredential"
    pnwnamingregistry "github.com/pnwxkd/x/namingregistry"
    pnwpayrollaudit "github.com/pnwxkd/x/payrollaudit"
    pnwdao "github.com/pnwxkd/x/dao"
)

var (
    ModuleBasics = module.NewBasicManager(
        auth.AppModuleBasic{},
        bank.AppModuleBasic{},
        genutil.AppModuleBasic{},
        params.AppModuleBasic{},
        pnwzkcredential.AppModuleBasic{},
        pnwnamingregistry.AppModuleBasic{},
        pnwpayrollaudit.AppModuleBasic{},
        pnwdao.AppModuleBasic{},
    )
)

func NewPnwxkdApp(logger log.Logger, db dbm.DB, traceStore io.Writer) *baseapp.BaseApp {
    encoding := simapp.MakeEncodingConfig(ModuleBasics)
    bApp := baseapp.NewBaseApp("pnwxkd", logger, db, encoding.TxConfig.TxDecoder())
    bApp.SetCommitMultiStoreTracer(traceStore)
    bApp.SetInterfaceRegistry(encoding.InterfaceRegistry)

    keys := sdk.NewKVStoreKeys(
        authtypes.StoreKey,
        banktypes.StoreKey,
        pnwzkcredential.StoreKey,
        pnwnamingregistry.StoreKey,
        pnwpayrollaudit.StoreKey,
        pnwdao.StoreKey,
    )

    // Mount stores
    bApp.MountKVStores(keys)

    // Construct app module manager
    appModules := module.NewManager(
        auth.NewAppModule(encoding.Codec, authtypes.NewAccountKeeper(...)),
        bank.NewAppModule(...),
        genutil.NewAppModule(...),
        pnwzkcredential.NewAppModule(...),
        pnwnamingregistry.NewAppModule(...),
        pnwpayrollaudit.NewAppModule(...),
        pnwdao.NewAppModule(...),
    )

    appModules.RegisterInvariants(bApp)
    appModules.RegisterRoutes(bApp.Router(), bApp.QueryRouter(), encoding.Amino)
    appModules.RegisterServices(bApp)

    bApp.SetInitChainer(...)     // for genesis logic
    bApp.SetBeginBlocker(...)    // custom begin block logic
    bApp.SetEndBlocker(...)      // custom end block logic

    return bApp
}
