// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Just-A-NoobieDev/go-htmx-templ-todo/models"

func RenderTodo(todo *models.SingleTodo, tasks map[string]interface{}) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-[518px] h-[257px] bg-[#EEA83599] p-8 rounded-3xl show-modal cursor-pointer relative\" id=\"s\"><div class=\"flex items-center justify-between mb-6\"><button class=\"cursor-pointer\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/todo/" + todo.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dialog\"><h1 class=\"text-[28px] font-medium\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(todo.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\renderTodo.templ`, Line: 10, Col: 59}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1></button> <button hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/dlt/" + todo.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"this\"><svg width=\"40\" height=\"40\" viewBox=\"0 0 40 40\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M20 35C28.2843 35 35 28.2843 35 20C35 11.7157 28.2843 5 20 5C11.7157 5 5 11.7157 5 20C5 28.2843 11.7157 35 20 35Z\" stroke=\"#40393E\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path d=\"M20.0167 20V20.0167H20V20H20.0167ZM27.5167 20V20.0167H27.5V20H27.5167ZM12.5167 20V20.0167H12.5V20H12.5167Z\" stroke=\"#40393E\" stroke-width=\"2.25\" stroke-linejoin=\"round\"></path></svg></button></div><button class=\"cursor-pointer\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/todo/" + todo.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#dialog\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for key, task := range tasks {
			if key == todo.Id {
				for _, t := range task.(models.TaskList) {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full flex items-center gap-4 rounded-md px-5 \">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if !t.Status {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"18\" height=\"18\" viewBox=\"0 0 18 18\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><circle cx=\"9\" cy=\"9\" r=\"8\" stroke=\"#40393E\" stroke-width=\"2\"></circle></svg>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"30\" height=\"30\" viewBox=\"0 0 30 30\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 26.25C16.4774 26.25 17.9403 25.959 19.3052 25.3936C20.6701 24.8283 21.9103 23.9996 22.955 22.955C23.9996 21.9103 24.8283 20.6701 25.3936 19.3052C25.959 17.9403 26.25 16.4774 26.25 15C26.25 13.5226 25.959 12.0597 25.3936 10.6948C24.8283 9.3299 23.9996 8.08971 22.955 7.04505C21.9103 6.00039 20.6701 5.17172 19.3052 4.60636C17.9403 4.04099 16.4774 3.75 15 3.75C12.0163 3.75 9.15483 4.93526 7.04505 7.04505C4.93526 9.15483 3.75 12.0163 3.75 15C3.75 17.9837 4.93526 20.8452 7.04505 22.955C9.15483 25.0647 12.0163 26.25 15 26.25ZM14.71 19.55L20.96 12.05L19.04 10.45L13.665 16.8988L10.8837 14.1163L9.11625 15.8837L12.8663 19.6337L13.8338 20.6013L14.71 19.55Z\" fill=\"#5B801A\"></path></svg>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#40393E] text-[26px]\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var3 string
					templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(t.Description)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\renderTodo.templ`, Line: 35, Col: 80}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
