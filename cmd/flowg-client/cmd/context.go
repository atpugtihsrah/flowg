package cmd

type ContextKey string

const (
	ApiClient     ContextKey = "api_client"
	MgmtApiClient ContextKey = "mgmt_api_client"
)
