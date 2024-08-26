package auth

const (
	SCOPE_READ_PIPELINES     Scope = "read_pipelines"
	SCOPE_WRITE_PIPELINES    Scope = "write_pipelines"
	SCOPE_READ_TRANSFORMERS  Scope = "read_transformers"
	SCOPE_WRITE_TRANSFORMERS Scope = "write_transformers"
	SCOPE_READ_STREAMS       Scope = "read_streams"
	SCOPE_WRITE_STREAMS      Scope = "write_streams"
	SCOPE_CREATE_USERS       Scope = "create_users"
	SCOPE_SEND_LOGS          Scope = "send_logs"
)

type Permissions struct {
	CanViewPipelines bool
	CanEditPipelines bool

	CanViewTransformers bool
	CanEditTransformers bool

	CanViewStreams bool
	CanEditStreams bool

	CanCreateUsers bool
	CanSendLogs    bool
}

func ParseScope(s string) (Scope, error) {
	switch s {
	case "read_pipelines":
		return SCOPE_READ_PIPELINES, nil
	case "write_pipelines":
		return SCOPE_WRITE_PIPELINES, nil
	case "read_transformers":
		return SCOPE_READ_TRANSFORMERS, nil
	case "write_transformers":
		return SCOPE_WRITE_TRANSFORMERS, nil
	case "read_streams":
		return SCOPE_READ_STREAMS, nil
	case "write_streams":
		return SCOPE_WRITE_STREAMS, nil
	case "create_users":
		return SCOPE_CREATE_USERS, nil
	case "send_logs":
		return SCOPE_SEND_LOGS, nil
	default:
		return "", &InvalidScopeError{Scope: s}
	}
}

func PermissionsFromScopes(scopes []Scope) Permissions {
	permissions := Permissions{}
	for _, scope := range scopes {
		switch scope {
		case SCOPE_READ_PIPELINES:
			permissions.CanViewPipelines = true
		case SCOPE_WRITE_PIPELINES:
			permissions.CanEditPipelines = true
		case SCOPE_READ_TRANSFORMERS:
			permissions.CanViewTransformers = true
		case SCOPE_WRITE_TRANSFORMERS:
			permissions.CanEditTransformers = true
		case SCOPE_READ_STREAMS:
			permissions.CanViewStreams = true
		case SCOPE_WRITE_STREAMS:
			permissions.CanEditStreams = true
		case SCOPE_CREATE_USERS:
			permissions.CanCreateUsers = true
		case SCOPE_SEND_LOGS:
			permissions.CanSendLogs = true
		}
	}
	return permissions
}
