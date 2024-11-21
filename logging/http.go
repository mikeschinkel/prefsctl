package logging

import (
	"net/http"
	"sort"
	"strings"
)

func HTTPRequestLogArgs(r *http.Request) []any {
	_ = r.ParseForm()
	args := []any{
		HostnameLogArg, r.Host,
		RequestProtocolLogArg, r.Proto,
		RequestMethodLogArg, r.Method,
		RequestURLLogArg, r.RequestURI,
		ContentTypeLengthLogArg, r.ContentLength,
		RequestorAddressLogArg, r.RemoteAddr,
	}
	args = append(args, HTTPHeaderLogArgs(r.Header)...)
	args = append(args)
	if r.TLS != nil {
		args = append(args, RequestTLSLogArg, "set")
		args = append(args, TLSLogArgs(r.TLS)...)
	}
	if r.TransferEncoding != nil {
		args = append(args, TransferEncodingLogArg, strings.Join(r.TransferEncoding, ","))
	}
	if r.Pattern != "" {
		args = append(args, RequestPatternLogArg, r.Pattern)
	}
	form := r.PostForm.Encode()
	if form != "" {
		args = append(args, RequestFormLogArg, form)
	}
	if r.Response != nil {
		args = append(args, RequestRedirectLogArg, r.Response)
	}
	return args
}

func HTTPHeaderLogArgs(hh http.Header) (args []any) {
	headers := make([]string, len(hh))
	args = make([]any, 0, len(hh)*2)
	args = append(args, HeadersCountLogArg, len(hh))
	index := 0
	for h := range hh {
		headers[index] = h
	}
	sort.Strings(headers)
	for _, h := range headers {
		v := strings.Join(hh[h], ",")
		if v == "" {
			continue
		}
		args = append(args, strings.ToLower(strings.Replace(h, "-", "_", -1)))
		args = append(args, v)
	}
	return args
}

func HTTPResponseWriterLogArgs(w http.ResponseWriter) []any {
	return HTTPHeaderLogArgs(w.Header())
}

func HTTPMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Debug(packageSlog, "Receiving request", func() []any {
			return HTTPRequestLogArgs(r)
		})
		// Call the next handler
		next.ServeHTTP(w, r)
		Debug(packageSlog, "Returning response", func() []any {
			return HTTPResponseWriterLogArgs(w)
		})
	}
}
