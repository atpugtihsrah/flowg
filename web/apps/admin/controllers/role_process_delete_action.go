package controllers

import (
	"log/slog"

	"encoding/json"
	"strconv"

	"net/http"

	"link-society.com/flowg/internal/data/auth"
)

func ProcessRoleDeleteAction(
	roleSys *auth.RoleSystem,
	userSys *auth.UserSystem,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		permissions := auth.Permissions{}
		notifications := []string{}

		user := auth.GetContextUser(r.Context())
		scopes, err := userSys.ListUserScopes(user.Name)
		if err != nil {
			slog.ErrorContext(
				r.Context(),
				"error listing user scopes",
				"channel", "web",
				"error", err.Error(),
			)

			notifications = append(notifications, "&#10060; Could not fetch user permissions")
		} else {
			permissions = auth.PermissionsFromScopes(scopes)
		}

		if !permissions.CanEditACLs {
			w.Header().Add("HX-Reswap", "none")

			notifications = append(notifications, "&#10060; You do not have permission to delete roles")
		} else {
			roleName := r.PathValue("name")

			err := roleSys.DeleteRole(roleName)
			if err != nil {
				slog.ErrorContext(
					r.Context(),
					"error deleting role",
					"channel", "web",
					"error", err.Error(),
				)

				w.Header().Add("HX-Reswap", "none")

				notifications = append(notifications, "&#10060; Could not delete role")
			} else {
				w.Header().Add("HX-Reswap", "delete")
				w.Header().Add("HX-Retarget", "tr[data-role="+strconv.Quote(roleName)+"]")

				notifications = append(notifications, "&#9989; Role deleted")
			}
		}

		trigger := map[string]interface{}{
			"htmx-custom-toast": map[string]interface{}{
				"messages": notifications,
			},
		}

		triggerData, err := json.Marshal(trigger)
		if err != nil {
			slog.ErrorContext(
				r.Context(),
				"error marshalling trigger",
				"channel", "web",
				"error", err.Error(),
			)
		} else {
			w.Header().Add("HX-Trigger", string(triggerData))
		}

		w.WriteHeader(http.StatusOK)
	}
}
