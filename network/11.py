# client.py
import socket

def main():
    # 创建 TCP socket
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    # 连接 Go 服务器
    client.connect(('localhost', 8888))
    print("Connected to Go server. Type messages (type 'exit' to quit)")

    try:
        while True:
            msg = input()
            if msg.lower() == 'exit':
                break
            # 发送消息，必须加上换行符以匹配 Go 服务器的 ReadString('\n')
            client.sendall((msg + '\n').encode('utf-8'))
            # 接收服务器回复
            response = client.recv(1024).decode('utf-8')
            print(f"Server says: {response}", end='')
    finally:
        client.close()

if __name__ == '__main__':
    main()