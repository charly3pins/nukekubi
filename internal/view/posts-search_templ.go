// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/charly3pins/nukekubi/internal/model"

func PostSearch(posts []model.Post) templ.Component {
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
		for _, post := range posts {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article class=\"flex flex-col gap-y-3 p-6 mt-6 mx-2 md:mx-0 rounded-lg shadow-md bg-white dark:bg-gray-700\"><h2 class=\"text-4xl font-semibold text-slate-800 dark:text-slate-200\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 templ.SafeURL = templ.URL("/blog/" + post.Slug)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(post.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/posts-search.templ`, Line: 11, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></h2><div class=\"my-4 text-large text-slate-600 dark:text-slate-300\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(post.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/posts-search.templ`, Line: 13, Col: 85}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><ul class=\"flex flex-row flex-wrap text-slate-500 dark:text-slate-300\"><li><a href=\"/categories/themes/\" class=\"text-sm mr-2 px-2 py-1 rounded border border-emerald-800 bg-emerald-800 text-slate-50\">themes</a></li><li><a href=\"/categories/syntax/\" class=\"text-sm mr-2 px-2 py-1 rounded border border-emerald-800 bg-emerald-800 text-slate-50\">syntax</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, tag := range post.Tags {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 templ.SafeURL = templ.URL("/tags/" + tag)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex flex-row text-sm mr-2 py-1\"><i class=\"h-5 w-5 flex-none\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-hash\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentcolor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M5 9h14\"></path> <path d=\"M5 15h14\"></path> <path d=\"M11 4 7 20\"></path> <path d=\"M17 4l-4 16\"></path></svg></i> <span class=\"ml-0\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(tag)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/posts-search.templ`, Line: 48, Col: 31}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul><div class=\"flex flex-col gap-y-1 md:flex-row md:gap-y-0 md:gap-x-4 text-slate-500 dark:text-slate-300\"><div class=\"flex flex-row text-base gap-x-1\"><i class=\"h-6 w-6 flex-none\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-calendar\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentcolor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M4 7a2 2 0 012-2h12a2 2 0 012 2v12a2 2 0 01-2 2H6a2 2 0 01-2-2V7z\"></path> <path d=\"M16 3v4\"></path> <path d=\"M8 3v4\"></path> <path d=\"M4 11h16\"></path> <path d=\"M11 15h1\"></path> <path d=\"M12 15v3\"></path></svg></i> <time datetime=\"2019-03-11T00:00:00+00:00\">2019-03-11</time></div><div class=\"flex flex-row text-base gap-x-1\"><i class=\"h-6 w-6 flex-none\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"icon icon-tabler icon-tabler-hourglass-high\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke=\"currentcolor\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path stroke=\"none\" d=\"M0 0h24v24H0z\" fill=\"none\"></path> <path d=\"M6.5 7h11\"></path> <path d=\"M6 20v-2a6 6 0 1112 0v2a1 1 0 01-1 1H7a1 1 0 01-1-1z\"></path> <path d=\"M6 4v2a6 6 0 1012 0V4a1 1 0 00-1-1H7A1 1 0 006 4z\"></path></svg></i> <span>3 minutes to read</span></div></div></article>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
