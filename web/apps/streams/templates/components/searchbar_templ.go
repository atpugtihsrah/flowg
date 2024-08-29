// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "time"

type SearchbarProps struct {
	From        time.Time
	To          time.Time
	Filter      string
	AutoRefresh string
}

func Searchbar(props SearchbarProps) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"form_stream\" hx-get=\"\" hx-target=\"#stream_viewer\" hx-swap=\"outerHTML\" class=\"flex flex-row items-center gap-2 px-3 py-1 z-depth-1\"><input type=\"hidden\" id=\"data_stream_timeoffset\" name=\"timeoffset\" value=\"\"><div class=\"flex-grow\"><label for=\"data_stream_filter\">Filter:</label> <input id=\"data_stream_filter\" name=\"filter\" type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.Filter)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `apps/streams/templates/components/searchbar.templ`, Line: 33, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" placeholder=\"field = &#34;value&#34;\"></div><div><label for=\"data_stream_from\">From:</label> <input id=\"data_stream_from\" name=\"from\" type=\"datetime-local\" step=\"1\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.From.Format("2006-01-02T15:04:05"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `apps/streams/templates/components/searchbar.templ`, Line: 44, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div><label for=\"data_stream_to\">To:</label> <input id=\"data_stream_to\" name=\"to\" type=\"datetime-local\" step=\"1\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(props.To.Format("2006-01-02T15:04:05"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `apps/streams/templates/components/searchbar.templ`, Line: 54, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div><label for=\"data_stream_autorefresh\">Auto Refresh:</label> <select id=\"data_stream_autorefresh\" name=\"autorefresh\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch props.AutoRefresh {
		case "0":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"0\" selected>No Auto Refresh</option> <option value=\"5\">Every 5s</option> <option value=\"10\">Every 10s</option> <option value=\"30\">Every 30s</option> <option value=\"60\">Every 1m</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "5":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"0\">No Auto Refresh</option> <option value=\"5\" selected>Every 5s</option> <option value=\"10\">Every 10s</option> <option value=\"30\">Every 30s</option> <option value=\"60\">Every 1m</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "10":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"0\">No Auto Refresh</option> <option value=\"5\">Every 5s</option> <option value=\"10\" selected>Every 10s</option> <option value=\"30\">Every 30s</option> <option value=\"60\">Every 1m</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "30":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"0\">No Auto Refresh</option> <option value=\"5\">Every 5s</option> <option value=\"10\">Every 10s</option> <option value=\"30\" selected>Every 30s</option> <option value=\"60\">Every 1m</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "60":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"0\">No Auto Refresh</option> <option value=\"5\">Every 5s</option> <option value=\"10\">Every 10s</option> <option value=\"30\">Every 30s</option> <option value=\"60\" selected>Every 1m</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><button type=\"submit\" class=\"btn waves-effect waves-light ml-5\"><i class=\"material-icons right\">search</i> Run Query</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
