//// This file is safe to edit. Once it exists it will not be overwritten
//
package restapi
//
//import (
//	"crypto/tls"
//	"dcn_sdk/api/handlers"
//	"io"
//	"net/http"
//
//	"github.com/go-openapi/errors"
//	"github.com/go-openapi/runtime"
//	"github.com/go-openapi/runtime/middleware"
//
//	"dcn_sdk/restapi/operations"
//	"dcn_sdk/restapi/operations/account_controller"
//	"dcn_sdk/restapi/operations/account_mapping_controller"
//	"dcn_sdk/restapi/operations/account_transactions_controller"
//	"dcn_sdk/restapi/operations/block_controller"
//	"dcn_sdk/restapi/operations/block_transactions_controller"
//	"dcn_sdk/restapi/operations/configuration_controller"
//	"dcn_sdk/restapi/operations/contract_controller"
//	"dcn_sdk/restapi/operations/contract_events_controller"
//	"dcn_sdk/restapi/operations/contract_transactions_controller"
//	"dcn_sdk/restapi/operations/dashboard_tokens_totals_controller"
//	"dcn_sdk/restapi/operations/dashboard_totals_controller"
//	"dcn_sdk/restapi/operations/dashboard_transactions_totals_controller"
//	"dcn_sdk/restapi/operations/erc_20_transfers_controller"
//	"dcn_sdk/restapi/operations/erc_721_transfers_controller"
//	"dcn_sdk/restapi/operations/erc_777_transfers_controller"
//	"dcn_sdk/restapi/operations/events_controller"
//	"dcn_sdk/restapi/operations/gas_oracle_controller"
//	"dcn_sdk/restapi/operations/index_controller"
//	"dcn_sdk/restapi/operations/metadata_controller"
//	"dcn_sdk/restapi/operations/node_controller"
//	"dcn_sdk/restapi/operations/search_controller"
//	"dcn_sdk/restapi/operations/token_controller"
//	"dcn_sdk/restapi/operations/token_events_controller"
//	"dcn_sdk/restapi/operations/token_transactions_controller"
//	"dcn_sdk/restapi/operations/transaction_controller"
//	"dcn_sdk/restapi/operations/transaction_events_controller"
//)
//
////go:generate swagger generate server --target ../../dcn_sdk --name DcnSdk --spec ../swagger.yaml --principal interface{} --exclude-main
//
//func configureFlags(api *operations.DcnSdkAPI) {
//	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
//}
//
//func configureAPI(api *operations.DcnSdkAPI) http.Handler {
//	// configure the api here
//	api.ServeError = errors.ServeError
//
//	// Set your custom logger if needed. Default one is log.Printf
//	// Expected interface func(string, ...interface{})
//	//
//	// Example:
//	// api.Logger = log.Printf
//
//	api.UseSwaggerUI()
//	// To continue using redoc as your UI, uncomment the following line
//	// api.UseRedoc()
//
//	api.JSONConsumer = runtime.JSONConsumer()
//	api.MultipartformConsumer = runtime.DiscardConsumer
//
//	api.EmptyProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
//		return errors.NotImplemented(" producer has not yet been implemented")
//	})
//
//	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
//	// metadata_controller.UploadMetadataUsingPOSTMaxParseMemory = 32 << 20
//
//	if api.AccountMappingControllerAddAccountUsingPOSTHandler == nil {
//		api.AccountMappingControllerAddAccountUsingPOSTHandler = account_mapping_controller.AddAccountUsingPOSTHandlerFunc(func(params account_mapping_controller.AddAccountUsingPOSTParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_mapping_controller.AddAccountUsingPOST has not yet been implemented")
//		})
//	}
//	if api.AccountMappingControllerDeleteAccountUsingDELETEHandler == nil {
//		api.AccountMappingControllerDeleteAccountUsingDELETEHandler = account_mapping_controller.DeleteAccountUsingDELETEHandlerFunc(func(params account_mapping_controller.DeleteAccountUsingDELETEParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_mapping_controller.DeleteAccountUsingDELETE has not yet been implemented")
//		})
//	}
//	if api.MetadataControllerDeleteMetadataUsingDELETEHandler == nil {
//		api.MetadataControllerDeleteMetadataUsingDELETEHandler = metadata_controller.DeleteMetadataUsingDELETEHandlerFunc(func(params metadata_controller.DeleteMetadataUsingDELETEParams) middleware.Responder {
//			return middleware.NotImplemented("operation metadata_controller.DeleteMetadataUsingDELETE has not yet been implemented")
//		})
//	}
//	if api.NodeControllerDetailsUsingGETHandler == nil {
//		api.NodeControllerDetailsUsingGETHandler = node_controller.DetailsUsingGETHandlerFunc(func(params node_controller.DetailsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation node_controller.DetailsUsingGET has not yet been implemented")
//		})
//	}
//	if api.AccountMappingControllerEditAccountUsingPUTHandler == nil {
//		api.AccountMappingControllerEditAccountUsingPUTHandler = account_mapping_controller.EditAccountUsingPUTHandlerFunc(func(params account_mapping_controller.EditAccountUsingPUTParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_mapping_controller.EditAccountUsingPUT has not yet been implemented")
//		})
//	}
//	if api.IndexControllerErrorPageUsingGETHandler == nil {
//		api.IndexControllerErrorPageUsingGETHandler = index_controller.ErrorPageUsingGETHandlerFunc(func(params index_controller.ErrorPageUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation index_controller.ErrorPageUsingGET has not yet been implemented")
//		})
//	}
//	if api.AccountTransactionsControllerFindAccountTransactionsUsingGETHandler == nil {
//		api.AccountTransactionsControllerFindAccountTransactionsUsingGETHandler = account_transactions_controller.FindAccountTransactionsUsingGETHandlerFunc(func(params account_transactions_controller.FindAccountTransactionsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_transactions_controller.FindAccountTransactionsUsingGET has not yet been implemented")
//		})
//	}
//	if api.AccountMappingControllerFindAccountsUsingGETHandler == nil {
//		api.AccountMappingControllerFindAccountsUsingGETHandler = account_mapping_controller.FindAccountsUsingGETHandlerFunc(func(params account_mapping_controller.FindAccountsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_mapping_controller.FindAccountsUsingGET has not yet been implemented")
//		})
//	}
//	if api.BlockControllerFindBlocksUsingGETHandler == nil {
//		api.BlockControllerFindBlocksUsingGETHandler = block_controller.FindBlocksUsingGETHandlerFunc(func(params block_controller.FindBlocksUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation block_controller.FindBlocksUsingGET has not yet been implemented")
//		})
//	}
//	if api.ContractTransactionsControllerFindContractTransactionsUsingGETHandler == nil {
//		api.ContractTransactionsControllerFindContractTransactionsUsingGETHandler = contract_transactions_controller.FindContractTransactionsUsingGETHandlerFunc(func(params contract_transactions_controller.FindContractTransactionsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation contract_transactions_controller.FindContractTransactionsUsingGET has not yet been implemented")
//		})
//	}
//	if api.TokenTransactionsControllerFindContractTransactionsUsingGET1Handler == nil {
//		api.TokenTransactionsControllerFindContractTransactionsUsingGET1Handler = token_transactions_controller.FindContractTransactionsUsingGET1HandlerFunc(func(params token_transactions_controller.FindContractTransactionsUsingGET1Params) middleware.Responder {
//			return middleware.NotImplemented("operation token_transactions_controller.FindContractTransactionsUsingGET1 has not yet been implemented")
//		})
//	}
//	if api.MetadataControllerFindContractsBySwarmHashUsingGETHandler == nil {
//		api.MetadataControllerFindContractsBySwarmHashUsingGETHandler = metadata_controller.FindContractsBySwarmHashUsingGETHandlerFunc(func(params metadata_controller.FindContractsBySwarmHashUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation metadata_controller.FindContractsBySwarmHashUsingGET has not yet been implemented")
//		})
//	}
//	if api.ContractControllerFindContractsUsingGETHandler == nil {
//		api.ContractControllerFindContractsUsingGETHandler = contract_controller.FindContractsUsingGETHandlerFunc(func(params contract_controller.FindContractsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation contract_controller.FindContractsUsingGET has not yet been implemented")
//		})
//	}
//	if api.Erc20TransfersControllerFindERC20TransfersUsingGETHandler == nil {
//		api.Erc20TransfersControllerFindERC20TransfersUsingGETHandler = erc_20_transfers_controller.FindERC20TransfersUsingGETHandlerFunc(func(params erc_20_transfers_controller.FindERC20TransfersUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation erc_20_transfers_controller.FindERC20TransfersUsingGET has not yet been implemented")
//		})
//	}
//	if api.Erc721TransfersControllerFindERC721TransfersUsingGETHandler == nil {
//		api.Erc721TransfersControllerFindERC721TransfersUsingGETHandler = erc_721_transfers_controller.FindERC721TransfersUsingGETHandlerFunc(func(params erc_721_transfers_controller.FindERC721TransfersUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation erc_721_transfers_controller.FindERC721TransfersUsingGET has not yet been implemented")
//		})
//	}
//	if api.Erc777TransfersControllerFindERC777TransfersUsingGETHandler == nil {
//		api.Erc777TransfersControllerFindERC777TransfersUsingGETHandler = erc_777_transfers_controller.FindERC777TransfersUsingGETHandlerFunc(func(params erc_777_transfers_controller.FindERC777TransfersUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation erc_777_transfers_controller.FindERC777TransfersUsingGET has not yet been implemented")
//		})
//	}
//	if api.MetadataControllerFindMetadataUsingGETHandler == nil {
//		api.MetadataControllerFindMetadataUsingGETHandler = metadata_controller.FindMetadataUsingGETHandlerFunc(func(params metadata_controller.FindMetadataUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation metadata_controller.FindMetadataUsingGET has not yet been implemented")
//		})
//	}
//	if api.TokenControllerFindTokensUsingGETHandler == nil {
//		api.TokenControllerFindTokensUsingGETHandler = token_controller.FindTokensUsingGETHandlerFunc(func(params token_controller.FindTokensUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation token_controller.FindTokensUsingGET has not yet been implemented")
//		})
//	}
//	if api.TransactionControllerFindTransactionsUsingGETHandler == nil {
//		api.TransactionControllerFindTransactionsUsingGETHandler = transaction_controller.FindTransactionsUsingGETHandlerFunc(func(params transaction_controller.FindTransactionsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation transaction_controller.FindTransactionsUsingGET has not yet been implemented")
//		})
//	}
//	if api.ConfigurationControllerFrontendConfigurationUsingGETHandler == nil {
//		api.ConfigurationControllerFrontendConfigurationUsingGETHandler = configuration_controller.FrontendConfigurationUsingGETHandlerFunc(func(params configuration_controller.FrontendConfigurationUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation configuration_controller.FrontendConfigurationUsingGET has not yet been implemented")
//		})
//	}
//	if api.GasOracleControllerGasPriceInfoUsingGETHandler == nil {
//		api.GasOracleControllerGasPriceInfoUsingGETHandler = gas_oracle_controller.GasPriceInfoUsingGETHandlerFunc(func(params gas_oracle_controller.GasPriceInfoUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation gas_oracle_controller.GasPriceInfoUsingGET has not yet been implemented")
//		})
//	}
//	if api.AccountControllerGetAccountUsingGETHandler == nil {
//		api.AccountControllerGetAccountUsingGETHandler = account_controller.GetAccountUsingGETHandlerFunc(func(params account_controller.GetAccountUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation account_controller.GetAccountUsingGET has not yet been implemented")
//		})
//	}
//	if api.AccountMappingControllerGetAccountUsingGET1Handler == nil {
//		api.AccountMappingControllerGetAccountUsingGET1Handler = account_mapping_controller.GetAccountUsingGET1HandlerFunc(func(params account_mapping_controller.GetAccountUsingGET1Params) middleware.Responder {
//			return middleware.NotImplemented("operation account_mapping_controller.GetAccountUsingGET1 has not yet been implemented")
//		})
//	}
//	if api.BlockTransactionsControllerGetBlockTransactionsUsingGETHandler == nil {
//		api.BlockTransactionsControllerGetBlockTransactionsUsingGETHandler = block_transactions_controller.GetBlockTransactionsUsingGETHandlerFunc(func(params block_transactions_controller.GetBlockTransactionsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation block_transactions_controller.GetBlockTransactionsUsingGET has not yet been implemented")
//		})
//	}
//	if api.BlockControllerGetBlockUsingGETHandler == nil {
//		//api.BlockControllerGetBlockUsingGETHandler = block_controller.GetBlockUsingGETHandlerFunc(
//		//	func(params block_controller.GetBlockUsingGETParams) middleware.Responder {
//		//		return middleware.NotImplemented("operation block_controller.GetBlockUsingGET has not yet been implemented")
//		//	},
//		//)
//		api.BlockControllerGetBlockUsingGETHandler = handlers.NewApiGetBlockUsing()
//	}
//	if api.ContractEventsControllerGetContractEventsUsingGETHandler == nil {
//		api.ContractEventsControllerGetContractEventsUsingGETHandler = contract_events_controller.GetContractEventsUsingGETHandlerFunc(func(params contract_events_controller.GetContractEventsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation contract_events_controller.GetContractEventsUsingGET has not yet been implemented")
//		})
//	}
//	if api.ContractControllerGetContractUsingGETHandler == nil {
//		api.ContractControllerGetContractUsingGETHandler = contract_controller.GetContractUsingGETHandlerFunc(func(params contract_controller.GetContractUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation contract_controller.GetContractUsingGET has not yet been implemented")
//		})
//	}
//	if api.DashboardTotalsControllerGetDashboardTotalsUsingGETHandler == nil {
//		api.DashboardTotalsControllerGetDashboardTotalsUsingGETHandler = dashboard_totals_controller.GetDashboardTotalsUsingGETHandlerFunc(func(params dashboard_totals_controller.GetDashboardTotalsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation dashboard_totals_controller.GetDashboardTotalsUsingGET has not yet been implemented")
//		})
//	}
//	if api.EventsControllerGetEventsUsingGETHandler == nil {
//		api.EventsControllerGetEventsUsingGETHandler = events_controller.GetEventsUsingGETHandlerFunc(func(params events_controller.GetEventsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation events_controller.GetEventsUsingGET has not yet been implemented")
//		})
//	}
//	if api.MetadataControllerGetMetadataUsingGETHandler == nil {
//		api.MetadataControllerGetMetadataUsingGETHandler = metadata_controller.GetMetadataUsingGETHandlerFunc(func(params metadata_controller.GetMetadataUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation metadata_controller.GetMetadataUsingGET has not yet been implemented")
//		})
//	}
//	if api.TokenEventsControllerGetTokenEventsUsingGETHandler == nil {
//		api.TokenEventsControllerGetTokenEventsUsingGETHandler = token_events_controller.GetTokenEventsUsingGETHandlerFunc(func(params token_events_controller.GetTokenEventsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation token_events_controller.GetTokenEventsUsingGET has not yet been implemented")
//		})
//	}
//	if api.DashboardTokensTotalsControllerGetTokenTotalsUsingGETHandler == nil {
//		api.DashboardTokensTotalsControllerGetTokenTotalsUsingGETHandler = dashboard_tokens_totals_controller.GetTokenTotalsUsingGETHandlerFunc(func(params dashboard_tokens_totals_controller.GetTokenTotalsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation dashboard_tokens_totals_controller.GetTokenTotalsUsingGET has not yet been implemented")
//		})
//	}
//	if api.TokenControllerGetTokenUsingGETHandler == nil {
//		api.TokenControllerGetTokenUsingGETHandler = token_controller.GetTokenUsingGETHandlerFunc(func(params token_controller.GetTokenUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation token_controller.GetTokenUsingGET has not yet been implemented")
//		})
//	}
//	if api.DashboardTransactionsTotalsControllerGetTotalsUsingGETHandler == nil {
//		api.DashboardTransactionsTotalsControllerGetTotalsUsingGETHandler = dashboard_transactions_totals_controller.GetTotalsUsingGETHandlerFunc(func(params dashboard_transactions_totals_controller.GetTotalsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation dashboard_transactions_totals_controller.GetTotalsUsingGET has not yet been implemented")
//		})
//	}
//	if api.DashboardTransactionsTotalsControllerGetTotalsUsingGET1Handler == nil {
//		api.DashboardTransactionsTotalsControllerGetTotalsUsingGET1Handler = dashboard_transactions_totals_controller.GetTotalsUsingGET1HandlerFunc(func(params dashboard_transactions_totals_controller.GetTotalsUsingGET1Params) middleware.Responder {
//			return middleware.NotImplemented("operation dashboard_transactions_totals_controller.GetTotalsUsingGET1 has not yet been implemented")
//		})
//	}
//	if api.TransactionEventsControllerGetTransactionEventsUsingGETHandler == nil {
//		api.TransactionEventsControllerGetTransactionEventsUsingGETHandler = transaction_events_controller.GetTransactionEventsUsingGETHandlerFunc(func(params transaction_events_controller.GetTransactionEventsUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation transaction_events_controller.GetTransactionEventsUsingGET has not yet been implemented")
//		})
//	}
//	if api.TransactionControllerGetTransactionUsingGETHandler == nil {
//		api.TransactionControllerGetTransactionUsingGETHandler = transaction_controller.GetTransactionUsingGETHandlerFunc(func(params transaction_controller.GetTransactionUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation transaction_controller.GetTransactionUsingGET has not yet been implemented")
//		})
//	}
//	//if api.IndexControllerHomePageUsingGETHandler == nil {
//	//	api.IndexControllerHomePageUsingGETHandler = index_controller.HomePageUsingGETHandlerFunc(func(params index_controller.HomePageUsingGETParams) middleware.Responder {
//	//		return middleware.NotImplemented("operation index_controller.HomePageUsingGET has not yet been implemented")
//	//	})
//	//}
//	//if api.IndexControllerIndexUsingGETHandler == nil {
//	//	api.IndexControllerIndexUsingGETHandler = index_controller.IndexUsingGETHandlerFunc(func(params index_controller.IndexUsingGETParams) middleware.Responder {
//	//		return middleware.NotImplemented("operation index_controller.IndexUsingGET has not yet been implemented")
//	//	})
//	//}
//	if api.SearchControllerSearchUsingGETHandler == nil {
//		api.SearchControllerSearchUsingGETHandler = search_controller.SearchUsingGETHandlerFunc(func(params search_controller.SearchUsingGETParams) middleware.Responder {
//			return middleware.NotImplemented("operation search_controller.SearchUsingGET has not yet been implemented")
//		})
//	}
//	if api.SearchControllerSearchUsingGET1Handler == nil {
//		api.SearchControllerSearchUsingGET1Handler = search_controller.SearchUsingGET1HandlerFunc(func(params search_controller.SearchUsingGET1Params) middleware.Responder {
//			return middleware.NotImplemented("operation search_controller.SearchUsingGET1 has not yet been implemented")
//		})
//	}
//	if api.MetadataControllerUploadMetadataUsingPOSTHandler == nil {
//		api.MetadataControllerUploadMetadataUsingPOSTHandler = metadata_controller.UploadMetadataUsingPOSTHandlerFunc(func(params metadata_controller.UploadMetadataUsingPOSTParams) middleware.Responder {
//			return middleware.NotImplemented("operation metadata_controller.UploadMetadataUsingPOST has not yet been implemented")
//		})
//	}
//
//	api.PreServerShutdown = func() {}
//
//	api.ServerShutdown = func() {}
//
//	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
//}
//
//// The TLS configuration before HTTPS server starts.
//func configureTLS(tlsConfig *tls.Config) {
//	// Make all necessary changes to the TLS configuration here.
//}
//
//// As soon as server is initialized but not run yet, this function will be called.
//// If you need to modify a config, store server instance to stop it individually later, this is the place.
//// This function can be called multiple times, depending on the number of serving schemes.
//// scheme value will be set accordingly: "http", "https" or "unix".
//func configureServer(s *http.Server, scheme, addr string) {
//}
//
//// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
//// The middleware executes after routing but before authentication, binding and validation.
//func setupMiddlewares(handler http.Handler) http.Handler {
//	return handler
//}
//
//// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
//// So this is a good place to plug in a panic handling middleware, logging and metrics.
//func setupGlobalMiddleware(handler http.Handler) http.Handler {
//	return handler
//}
