// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func form() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"w-full mt-4 flex flex-col gap-3\" id=\"todoForm\" action=\"/add\" method=\"POST\" hx-post=\"/add\" hx-target=\"#todo-list\"><div class=\"w-full flex flex-col gap-2\"><label for=\"title\" class=\"text-[22px] font-medium tracing-[1.2px]\">Board Title</label> <input type=\"text\" name=\"title\" class=\"bg-[#B5AFA433] w-full h-[51px] rounded-lg\"></div><div class=\"w-full flex flex-col gap-2 mt-2\"><label for=\"task1\" class=\"text-[22px] font-medium tracing-[1.2px]\">Task 1</label> <input type=\"text\" name=\"task1\" class=\"bg-[#B5AFA433] w-full h-[51px] rounded-lg\"></div><div class=\"w-full flex flex-col gap-2\"><label for=\"task2\" class=\"text-[22px] font-medium tracing-[1.2px]\">Task 2</label> <input type=\"text\" name=\"task2\" class=\"bg-[#B5AFA433] w-full h-[51px] rounded-lg\"></div><div class=\"w-full flex flex-col gap-2\"><label for=\"task3\" class=\"text-[22px] font-medium tracing-[1.2px]\">Task 3</label> <input type=\"text\" name=\"task3\" class=\"bg-[#B5AFA433] w-full h-[51px] rounded-lg\"></div><button type=\"submit\" class=\"w-full bg-[#EEA83599] h-[51px] rounded-lg mt-5 text-[22px] font-semibold text-[#1E1E1E]\" hx-on:click=\"\">Save</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
