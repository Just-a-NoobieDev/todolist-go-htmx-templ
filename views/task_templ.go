// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Just-A-NoobieDev/go-htmx-templ-todo/models"

func Task(task *models.SingleTask) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center gap-4 w-full\" hx-target=\"this\" hx-swap=\"outerHTML\"><div class=\"bg-[#B5AFA499] h-[75px] w-full flex items-center gap-4 rounded-md px-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !task.Status {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"18\" height=\"18\" viewBox=\"0 0 18 18\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"cursor-pointer\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/edit-status?taskId=" + task.Taskid))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><circle cx=\"9\" cy=\"9\" r=\"8\" stroke=\"#40393E\" stroke-width=\"2\"></circle></svg><p class=\"text-[#40393E] text-[26px]\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(task.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\task.templ`, Line: 11, Col: 65}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"30\" height=\"30\" viewBox=\"0 0 30 30\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"cursor-pointer\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/edit-status?taskId=" + task.Taskid))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 26.25C16.4774 26.25 17.9403 25.959 19.3052 25.3936C20.6701 24.8283 21.9103 23.9996 22.955 22.955C23.9996 21.9103 24.8283 20.6701 25.3936 19.3052C25.959 17.9403 26.25 16.4774 26.25 15C26.25 13.5226 25.959 12.0597 25.3936 10.6948C24.8283 9.3299 23.9996 8.08971 22.955 7.04505C21.9103 6.00039 20.6701 5.17172 19.3052 4.60636C17.9403 4.04099 16.4774 3.75 15 3.75C12.0163 3.75 9.15483 4.93526 7.04505 7.04505C4.93526 9.15483 3.75 12.0163 3.75 15C3.75 17.9837 4.93526 20.8452 7.04505 22.955C9.15483 25.0647 12.0163 26.25 15 26.25ZM14.71 19.55L20.96 12.05L19.04 10.45L13.665 16.8988L10.8837 14.1163L9.11625 15.8837L12.8663 19.6337L13.8338 20.6013L14.71 19.55Z\" fill=\"#5B801A\"></path></svg><p class=\"text-[#40393E80] line-through text-[26px]\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(task.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views\task.templ`, Line: 16, Col: 80}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><button class=\"flex items-center justify-center w-[85px] h-[75px] bg-[#B5AFA499] rounded-md\" hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("/edit?taskId=" + task.Taskid))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><svg width=\"30\" height=\"30\" viewBox=\"0 0 30 30\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M7.01875 25.0001C6.44375 25.0001 5.96375 24.8076 5.57875 24.4226C5.19292 24.0367 5 23.5563 5 22.9813V7.01881C5 6.44381 5.19292 5.96381 5.57875 5.57881C5.96375 5.19298 6.44375 5.00006 7.01875 5.00006H17.5025L16.2525 6.25006H7.01875C6.82708 6.25006 6.65083 6.33006 6.49 6.49006C6.33 6.65089 6.25 6.82714 6.25 7.01881V22.9813C6.25 23.173 6.33 23.3492 6.49 23.5101C6.65083 23.6701 6.82708 23.7501 7.01875 23.7501H22.9813C23.1729 23.7501 23.3492 23.6701 23.51 23.5101C23.67 23.3492 23.75 23.173 23.75 22.9813V13.6201L25 12.3701V22.9813C25 23.5563 24.8075 24.0363 24.4225 24.4213C24.0367 24.8071 23.5562 25.0001 22.9813 25.0001H7.01875ZM12.5 17.5001V14.2313L23.68 3.05006C23.8183 2.91256 23.9633 2.81756 24.115 2.76506C24.2675 2.71173 24.4279 2.68506 24.5963 2.68506C24.7529 2.68506 24.9062 2.71173 25.0562 2.76506C25.2062 2.81756 25.3421 2.90506 25.4638 3.02756L26.8538 4.37506C26.9871 4.51256 27.0879 4.66423 27.1562 4.83006C27.2254 4.99673 27.26 5.16381 27.26 5.33131C27.26 5.49881 27.235 5.65797 27.185 5.80881C27.1342 5.95881 27.0396 6.10256 26.9012 6.24006L15.65 17.5001H12.5ZM13.75 16.2501H15.115L23.4475 7.91881L22.765 7.23506L22.0025 6.50506L13.75 14.7576V16.2501Z\" fill=\"#40393E\"></path></svg></button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
