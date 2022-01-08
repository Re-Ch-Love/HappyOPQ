import socket

from multiprocessing import Process


def handle_client(client):
    """
    处理客户端请求
    """
    request_data = client.recv(1024)
    print("request data:", request_data)
    # 构造响应数据
    response_start_line = "HTTP/1.1 200 OK\r\n"
    response_headers = "Server: My server\r\n"
    response_body = "<h1>Python HTTP Test</h1>"
    response = response_start_line + response_headers + "\r\n" + response_body

    # 向客户端返回响应数据
    client.send(bytes(response, "utf-8"))

    # 关闭客户端连接
    client.close()


if __name__ == "__main__":
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind(("127.0.0.1", 8081))
    server_socket.listen(128)

    while True:
        client_socket, client_address = server_socket.accept()
        handle_client_process = Process(target=handle_client, args=(client_socket,))
        handle_client_process.start()
        client_socket.close()
