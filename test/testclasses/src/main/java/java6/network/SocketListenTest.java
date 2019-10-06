package java6.network;

import libs.junit.UnitTestRunner;
import org.junit.Test;

import java.io.InputStreamReader;
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
public class SocketListenTest {

    public static void main(String[] args) throws Exception {
        try {
            ServerSocket serverSocket = new ServerSocket(5457);
            System.out.println("Accept:");
            Socket socket = serverSocket.accept();
            //read from client
            //char chars[] = new char[64];
            byte[] buffer = new byte[1024];
            int len;
            StringBuilder sb = new StringBuilder();
            //Reader reader = new InputStreamReader(socket.getInputStream());
            socket.setSoTimeout(2000);
            while ((len = socket.getInputStream().read(buffer)) != -1) {
                sb.append(new String(buffer, 0, len));
            }
            System.out.println("from client: " + sb);
            //write to client
            String s = "Hi i'm beyond\r\n";
            socket.getOutputStream().write(s.getBytes());
            socket.close();
        } catch (Exception e) {
            System.out.println(e.toString());
        }
        //telnet 127.0.0.1 5457
    }
}
