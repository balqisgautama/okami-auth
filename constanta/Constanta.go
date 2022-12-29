package constanta

import "time"

const DefaultDateFormat = `2006-01-02 15:04:05`

const RequestPOST = "POST"
const RequestGET = "GET"
const RequestPUT = "PUT"
const RequestDELETE = "DELETE"

const HeaderKeyContentType = "Content-Type"
const HeaderValueContentTypeHTML = "text/html"

const TokenHeaderName = "Authorization"
const AccessHeaderName = "access"
const ClientIDHeaderName = "client-id"
const SecretTokenHeaderName = "secret-token"

const ParamSearchID = "search-id"

const VarToken = "token"
const VarCode = "code"

const ResponseTypeRedirect = "redirect"

const UserPending = 1 // user belum melakukan aktivasi
const UserActive = 2
const UserDeleted = 3

const Time8Hour = 8 * time.Hour
const Time3Minute = 3 * time.Minute
const Time30Minute = 30 * time.Minute

const DescIncorrectFormat = "incorrect format"
const DescLoginFailed = "login failed"
const DescRedisDeleteDataFailed = "failed to delete data in redis by key"
const DescUnauthorized = "access forbidden"
const DescUpdateFailed = "failed to update"
const DescEmptyField = "field is empty"
const DescDataNotFound = "data not found"
const DescActivationFailed = "failed to activate your account"
const DescInvalidPassword = "password invalid"

const EmailProject = "info.okami.project@gmail.com"
const EmailAppPassword = "dfdaohklehxnsccl"
const EmailHostGmail = "smtp.gmail.com"
const EmailHostGmailWithPort = "smtp.gmail.com:587"
const EmailSubjectTokenLogin = "Token For Login"
const EmailSubjectActivationAccount = "Activation Account"

const PathAssetResponseHTMLButton = "D:\\okami\\playground\\auth\\public\\template-response-button.html"
const PathAssetResponseHTMLWithoutButton = "D:\\okami\\playground\\auth\\public\\template-response-without-button.html"
