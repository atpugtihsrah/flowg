// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

type BaseProps struct {
	Head          templ.Component
	Notifications []string
}

func Base(props BaseProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html><head><title>FlowG</title><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><link rel=\"stylesheet\" href=\"/static/css/google-font-material-icons.css\"><link rel=\"stylesheet\" href=\"/static/css/materialize.min.css\"><link rel=\"stylesheet\" href=\"/static/css/utilities.css\"><script type=\"application/javascript\" src=\"/static/js/materialize.patched.js\"></script><script type=\"application/javascript\" src=\"/static/js/htmx.min.js\"></script><script type=\"application/javascript\">\n        document.addEventListener('htmx:load', (event) => {\n          M.AutoInit(event.detail.elt, {\n            Dropdown: {\n              constrainWidth: false,\n            }\n          })\n        })\n      </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Notifications != nil {
			for _, notification := range props.Notifications {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script type=\"application/javascript\" data-message=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(notification)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/layouts/base.templ`, Line: 38, Col: 38}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\n            (() => {\n              const S = document.currentScript\n\n              document.addEventListener('DOMContentLoaded', () => {\n                setTimeout(() => {\n                  M.toast({\n                    html: S.dataset.message,\n                    completeCallback: () => S.remove(),\n                  })\n                }, 150)\n              })\n            })()\n          </script>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		if props.Head != nil {
			templ_7745c5c3_Err = props.Head.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"blue-grey lighten-4 flex flex-col h-maxvh overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"modal\" class=\"modal\"><div class=\"modal-content\" id=\"modal-content\"><h5>hello</h5></div></div><script type=\"application/javascript\">\n        document.body.addEventListener('htmx-custom-modal-open', () => {\n          const modalElement  = document.getElementById('modal')\n          const modalInstance = M.Modal.getInstance(modalElement)\n          modalInstance.open()\n        })\n\n        document.body.addEventListener('htmx-custom-modal-close', (event) => {\n          const modalElement  = document.getElementById('modal')\n          const modalInstance = M.Modal.getInstance(modalElement)\n          modalInstance.close()\n\n          if (event.detail.after === 'reload') {\n            setTimeout(\n              () => { window.location.reload() },\n              150,\n            )\n          }\n        })\n\n        document.body.addEventListener('htmx-custom-toast', (event) => {\n          for (const msg of event.detail.messages) {\n            M.toast({ html: msg })\n          }\n        })\n      </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
