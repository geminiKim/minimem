package constant

const STRING = "STRING"
const HASH = "HASH"
const LIST = "LIST"

const NOT_SUPPORTED_COMMAND = "Not Supported Command"

const GET = "GET"
const SET = "SET"

const OK = "OK"
const EMPTY = ""

const KEY = "key"
const VALUE = "value"
const FIELD = "field"
const INDEX = "index"
const COUNT = "count"

const URL_STRING_SET = "/string/{key}"
const URL_STRING_GET = "/string/{key}"

const URL_HASH_SET = "/hash/{key}/{field}"
const URL_HASH_GET = "/hash/{key}/{field}"

const LEFT_PUSH = "LEFT_PUSH"
const LEFT_PEEK = "LEFT_PEEK"
const LEFT_POP = "LEFT_POP"
const RIGHT_PUSH = "RIGHT_PUSH"
const RIGHT_PEEK = "RIGHT_PEEK"
const RIGHT_POP = "RIGHT_POP"
const BY_RANGE = "BY_RANGE"

const URL_LIST_LEFT_PUSH = "/list/{key}/leftPush" 
const URL_LIST_LEFT_PEEK = "/list/{key}/leftPeek" 
const URL_LIST_LEFT_POP = "/list/{key}/leftPop" 
const URL_LIST_RIGHT_PUSH = "/list/{key}/rightPush" 
const URL_LIST_RIGHT_PEEK = "/list/{key}/rightPeek" 
const URL_LIST_RIGHT_POP = "/list/{key}/rightPop"
const URL_LIST_BY_RANGE = "/list/{key}/{index}/{count}"