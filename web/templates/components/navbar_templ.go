// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"link-society.com/flowg/internal/data/auth"
	"link-society.com/flowg/internal/webutils"
)

type navbarItemProps struct {
	Icon   string
	Label  string
	Link   string
	Active bool
}

func navbarItem(props navbarItemProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if props.Active {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"active\"><a><i class=\"left material-icons\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.Icon)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 19, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 20, Col: 20}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL(props.Link)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><i class=\"left material-icons\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(props.Icon)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 26, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.Label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 27, Col: 20}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

type NavbarProps struct {
	CurrentNav string
}

func Navbar(props NavbarProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"nav-dropdown-user\" class=\"dropdown-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = navbarItem(navbarItemProps{
			Icon:   "account_circle",
			Label:  "Account",
			Link:   "/web/account",
			Active: props.CurrentNav == "account",
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if webutils.Permissions(ctx).CanViewACLs {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "dashboard",
				Label:  "Admin",
				Link:   "/web/admin",
				Active: props.CurrentNav == "admin",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"divider\"></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = navbarItem(navbarItemProps{
			Icon:   "exit_to_app",
			Label:  "Logout",
			Link:   "/auth/logout",
			Active: false,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><ul id=\"nav-dropdown-settings\" class=\"dropdown-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if webutils.Permissions(ctx).CanViewTransformers {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "filter_alt",
				Label:  "Transformers",
				Link:   "/web/transformers",
				Active: props.CurrentNav == "transformers",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewPipelines {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "settings",
				Label:  "Pipelines",
				Link:   "/web/pipelines",
				Active: props.CurrentNav == "pipelines",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewAlerts {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "notifications_active",
				Label:  "Alerts",
				Link:   "/web/alerts",
				Active: props.CurrentNav == "alerts",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewStreams {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "sd_storage",
				Label:  "Storage",
				Link:   "/web/storage",
				Active: props.CurrentNav == "storage",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><nav class=\"blue darken-3\" style=\"z-index: 900;\"><div class=\"nav-wrapper flex flex-row items-center\"><ul><li><a class=\"text-4xl font-semibold\" href=\"/\">FlowG</a></li></ul><ul class=\"flex-grow hide-on-small-only\"><li><a href=\"https://github.com/link-society/flowg\" target=\"_blank\"><i class=\"left material-icons\">code</i> GitHub</a></li><li><a href=\"/api/docs\" target=\"_blank\"><i class=\"left material-icons\">cloud</i> API Docs</a></li></ul><ul class=\"hide-on-small-only\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if webutils.Permissions(ctx).CanViewStreams {
			templ_7745c5c3_Err = navbarItem(navbarItemProps{
				Icon:   "storage",
				Label:  "Streams",
				Link:   "/web/streams",
				Active: props.CurrentNav == "streams",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a class=\"dropdown-trigger\" href=\"#\" data-target=\"nav-dropdown-settings\"><i class=\"left material-icons\">settings</i> Settings <i class=\"right material-icons\">arrow_drop_down</i></a></li><li><a class=\"dropdown-trigger\" href=\"#\" data-target=\"nav-dropdown-user\"><i class=\"left material-icons\">account_circle</i> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(auth.GetContextUser(ctx).Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 137, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <i class=\"right material-icons\">arrow_drop_down</i></a></li></ul></div></nav><div class=\"w-full hide-on-med-and-up\"><div class=\"collection\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if webutils.Permissions(ctx).CanViewStreams {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/streams\"><i class=\"left material-icons\">storage</i> Streams</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewTransformers {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/transformers\"><i class=\"left material-icons\">filter_alt</i> Transformers</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewPipelines {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/pipelines\"><i class=\"left material-icons\">settings</i> Pipelines</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewAlerts {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/alerts\"><i class=\"left material-icons\">notifications_active</i> Alerts</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if webutils.Permissions(ctx).CanViewStreams {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/storage\"><i class=\"left material-icons\">sd_storage</i> Storage</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"collection\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if webutils.Permissions(ctx).CanViewACLs {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/admin\"><i class=\"left material-icons\">dashboard</i> Admin</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"collection-item\" href=\"/web/account\"><i class=\"left material-icons\">account_circle</i> Account</a> <a class=\"collection-item\" href=\"/auth/logout\"><i class=\"left material-icons\">exit_to_app</i> Logout</a></div><div class=\"collection\"><a class=\"collection-item\" href=\"https://github.com/link-society/flowg\" target=\"_blank\"><i class=\"left material-icons\">code</i> GitHub</a> <a class=\"collection-item\" href=\"/api/docs\" target=\"_blank\"><i class=\"left material-icons\">cloud</i> API Docs</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
