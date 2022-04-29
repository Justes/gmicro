package main

import(
    //"fmt"
    "math"
    "net"
    "net/rpc"
    "net/http"
    "micro/param"
)

type MathUtil struct {
}

func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
    *resp = math.Pi * req * req
    return nil
}

func (mu *MathUtil) Add(param param.AddParam, resp *float32) error {
    *resp = param.Args1 + param.Args2
    return nil
}

func main() {

    mathUtil := new(MathUtil)
    err := rpc.Register(mathUtil)

    if err != nil {
        panic(err.Error())
    }

    rpc.HandleHTTP()

    listen, err := net.Listen("tcp", ":8081")

    if err != nil {
        panic(err.Error())
    }

    http.Serve(listen, nil)

}
