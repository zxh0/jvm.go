package java7.network;

import java.net.ServerSocket;
import java.net.Socket;

/**
 * Created with IntelliJ IDEA.
 * User: beyond
 * Email: beyondblog@outlook.com
 * Date: 15/3/21
 * Time: 下午9:24
 * Description:
 * Socket
 */
public class SocketTest {

    public static void main(String[] args) throws Exception {
        ServerSocket serverSocket = new ServerSocket(5457);
        Socket socket = serverSocket.accept();
//        String s = "Hi i'm beyond\r\n";
//        socket.getOutputStream().write(s.getBytes());
//        socket.close();
//        System.out.println("success!");
        //telnet 127.0.0.1 5457
    }
}
