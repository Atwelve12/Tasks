// server.go
package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", "127.0.0.1:8888")/*监听本地端口*/
    if err != nil {
        panic(err)
    }
    defer listener.Close()
    fmt.Println("正在监听 127.0.0.1:8888")

    conn, err := listener.Accept()/*等待客户端连接*/
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    fmt.Printf("Client connected: %s\n", conn.RemoteAddr())/*打印客户端ip*/
    reader := bufio.NewReader(conn)
     /*这行代码在 Go 语言中用于创建一个带缓冲的读取器，以便从网络连接 conn 中高效地读取数据。*/
    for {
        message, err := reader.ReadString('\n')
        /*
        reader 是之前通过 bufio.NewReader(conn) 创建的 *bufio.Reader 对象，它包装了 TCP 连接 conn。
        ReadString 是 bufio.Reader 的一个方法，它的作用是从缓冲区中读取数据，直到遇到指定的分隔符为止，然后将包括分隔符在内的所有已读取数据作为一个字符串返回。*/

        if err != nil {
            fmt.Println("Client disconnected")
            return
        }
        fmt.Printf("Received: %s", message)
        conn.Write([]byte("Echo: " + message))
        /*Go 语言中，Write 方法需要的是 []byte 类型（字节切片），而不是字符串。因此需要用 []byte() 将字符串转换为字节切片。*/
    }
}