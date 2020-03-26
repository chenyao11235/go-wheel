package paras

const (
    errCodeNegetivePage int = iota + 1
    errCodeNegetivePageSize
    errCodeNegetivePageToken
)

var errCodeToField = map[int]string{
    errCodeNegetivePage:      "page",
    errCodeNegetivePageSize:  "page_size",
    errCodeNegetivePageToken: "page_token",
}

var errCodeToMessage = map[int]string{
    errCodeNegetivePage:      "page参数必须是正数",
    errCodeNegetivePageSize:  "page_size参数必须是正数",
    errCodeNegetivePageToken: "page_token是无效的",
}

type PaginationError interface {
    error
    Field() string
}

type paginationError struct {
    code int
}

// 返回具体的错误内容（字符串）
func (e *paginationError) Error() string {
    return errCodeToMessage[e.code]
}

// 返回是哪个参数出现了错误
func (e *paginationError) Field() string {
    return errCodeToField[e.code]
}

// 构造由page参数引发的错误
func NewNegetivePageError() *paginationError {
    return &paginationError{
        code: errCodeNegetivePage,
    }
}

// 构造由page_size参数引发的错误
func NewNegetivePageSizeError() *paginationError {
    return &paginationError{
        code: errCodeNegetivePageSize,
    }
}

// 构造由page_token参数引发的错误
func NewNegetivePageTokenError() *paginationError {
    return &paginationError{
        code: errCodeNegetivePageToken,
    }
}
