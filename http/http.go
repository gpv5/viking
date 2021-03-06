package http

import (
	"github.com/valyala/fasthttp"
)

type Args = fasthttp.Args
type BalancingClient = fasthttp.BalancingClient
type Client = fasthttp.Client
type ConnState = fasthttp.ConnState
type Cookie = fasthttp.Cookie
type CookieSameSite = fasthttp.CookieSameSite
type DialFunc = fasthttp.DialFunc
type ErrBodyStreamWritePanic = fasthttp.ErrBodyStreamWritePanic
type ErrBrokenChunk = fasthttp.ErrBrokenChunk
type ErrSmallBuffer = fasthttp.ErrSmallBuffer
type FS = fasthttp.FS
type HijackHandler = fasthttp.HijackHandler
type HostClient = fasthttp.HostClient
type LBClient = fasthttp.LBClient
type Logger = fasthttp.Logger
type PathRewriteFunc = fasthttp.PathRewriteFunc
type PipelineClient = fasthttp.PipelineClient
type Request = fasthttp.Request
type RequestConfig = fasthttp.RequestConfig
type RequestCtx = fasthttp.RequestCtx
type RequestHandler = fasthttp.RequestHandler
type RequestHeader = fasthttp.RequestHeader
type Resolver = fasthttp.Resolver
type Response = fasthttp.Response
type ResponseHeader = fasthttp.ResponseHeader
type ServeHandler = fasthttp.ServeHandler
type Server = fasthttp.Server
type StreamWriter = fasthttp.StreamWriter
type TCPDialer = fasthttp.TCPDialer
type URI = fasthttp.URI




var AcquireTimer = fasthttp.AcquireTimer
var AppendDeflateBytes = fasthttp.AppendDeflateBytes
var AppendDeflateBytesLevel = fasthttp.AppendDeflateBytesLevel
var AppendGunzipBytes = fasthttp.AppendGunzipBytes
var AppendGzipBytes = fasthttp.AppendGzipBytes
var AppendGzipBytesLevel = fasthttp.AppendGzipBytesLevel
var AppendHTMLEscape = fasthttp.AppendHTMLEscape
var AppendHTMLEscapeBytes = fasthttp.AppendHTMLEscapeBytes
var AppendHTTPDate = fasthttp.AppendHTTPDate
var AppendIPv4 = fasthttp.AppendIPv4
var AppendInflateBytes = fasthttp.AppendInflateBytes
var AppendNormalizedHeaderKey = fasthttp.AppendNormalizedHeaderKey
var AppendNormalizedHeaderKeyBytes = fasthttp.AppendNormalizedHeaderKeyBytes
var AppendQuotedArg = fasthttp.AppendQuotedArg
var AppendUint = fasthttp.AppendUint
var AppendUnquotedArg = fasthttp.AppendUnquotedArg
var CoarseTimeNow = fasthttp.CoarseTimeNow
var Dial = fasthttp.Dial
var DialDualStack = fasthttp.DialDualStack
var DialDualStackTimeout = fasthttp.DialDualStackTimeout
var DialTimeout = fasthttp.DialTimeout
var Do = fasthttp.Do
var DoDeadline = fasthttp.DoDeadline
var DoTimeout = fasthttp.DoTimeout
var FileLastModified = fasthttp.FileLastModified
var Get = fasthttp.Get
var GetDeadline = fasthttp.GetDeadline
var GetTimeout = fasthttp.GetTimeout
var ListenAndServe = fasthttp.ListenAndServe
var ListenAndServeTLS = fasthttp.ListenAndServeTLS
var ListenAndServeTLSEmbed = fasthttp.ListenAndServeTLSEmbed
var ListenAndServeUNIX = fasthttp.ListenAndServeUNIX
var NewStreamReader = fasthttp.NewStreamReader
var ParseByteRange = fasthttp.ParseByteRange
var ParseHTTPDate = fasthttp.ParseHTTPDate
var ParseIPv4 = fasthttp.ParseIPv4
var ParseUfloat = fasthttp.ParseUfloat
var ParseUint = fasthttp.ParseUint
var Post = fasthttp.Post
var ReleaseArgs = fasthttp.ReleaseArgs
var ReleaseCookie = fasthttp.ReleaseCookie
var ReleaseRequest = fasthttp.ReleaseRequest
var ReleaseResponse = fasthttp.ReleaseResponse
var ReleaseTimer = fasthttp.ReleaseTimer
var ReleaseURI = fasthttp.ReleaseURI
var SaveMultipartFile = fasthttp.SaveMultipartFile
var Serve = fasthttp.Serve
var ServeConn = fasthttp.ServeConn
var ServeFile = fasthttp.ServeFile
var ServeFileBytes = fasthttp.ServeFileBytes
var ServeFileBytesUncompressed = fasthttp.ServeFileBytesUncompressed
var ServeFileUncompressed = fasthttp.ServeFileUncompressed
var ServeTLS = fasthttp.ServeTLS
var ServeTLSEmbed = fasthttp.ServeTLSEmbed
var StatusCodeIsRedirect = fasthttp.StatusCodeIsRedirect
var StatusMessage = fasthttp.StatusMessage
var WriteDeflate = fasthttp.WriteDeflate
var WriteDeflateLevel = fasthttp.WriteDeflateLevel
var WriteGunzip = fasthttp.WriteGunzip
var WriteGzip = fasthttp.WriteGzip
var WriteGzipLevel = fasthttp.WriteGzipLevel
var WriteInflate = fasthttp.WriteInflate
var WriteMultipartForm = fasthttp.WriteMultipartForm