package framework


type IResponse interface {
	Json(obj interface{}) IResponse

	Jsonp(obj interface{}) IResponse

	Xml(obj interface{}) IResponse

	Html(obj interface{}) IResponse

	Text(obj interface{}) IResponse

	Redirect(obj interface{}) IResponse

	SetHeader(key string,val string) IResponse

	SetCookie(key string,val string,maxAge int,path,domain string,secure,httpOnly bool) IResponse

	SetStatus(code int) IResponse

	SetOkStatus() IResponse
}


