package constanta

import "time"

const DefaultDateFormat = `2006-01-02 15:04:05`

const RequestPOST = "POST"
const RequestGET = "GET"
const RequestPUT = "PUT"
const RequestDELETE = "DELETE"

const TokenHeaderNameConstanta = "Authorization"
const AccessHeaderNameConstanta = "access"
const ClientIDHeaderNameConstanta = "client-id"

const ParamSearchID = "search-id"

const Time8Hour = 8 * time.Hour

const DescIncorrectFormat = "incorrect format"
const DescLoginFailed = "login failed"
const DescRedisDeleteDataFailed = "failed to delete data in redis by key"
const DescUnauthorized = "access forbidden"
const DescUpdateFailed = "failed to update"
const DescEmptyField = "field is empty"
const DescDataNotFound = "data not found"
const DescActivationFailed = "failed to activate your account"
const DescInvalidPassword = "password invalid"
