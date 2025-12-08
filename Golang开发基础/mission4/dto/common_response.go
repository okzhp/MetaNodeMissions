package dto

const (
  Success  = 0
  LogicErr = 10000
)

type Response struct {
  Code int    `json:"code"`
  Msg  string `json:"msg"`
  Data any    `json:"data"`
}

func ErrMsg(errMsg string) Response {
  return Response{
    Code: LogicErr,
    Msg:  errMsg,
  }
}

func SuccessMsg() Response {
  return Response{
    Code: Success,
    Msg:  "success",
  }
}

func SuccessData(data any) Response {
  return Response{
    Code: Success,
    Msg:  "success",
    Data: data,
  }
}
