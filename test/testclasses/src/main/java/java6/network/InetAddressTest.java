package java6.network;

import junit.framework.Assert;

import java.net.InetAddress;

/**
 * Created with IntelliJ IDEA.
 * User: beyond
 * Email: beyondblog@outlook.com
 * Date: 15/4/10
 * Time: 上午8:32
 * Description:
 */
public class InetAddressTest {

    public static void main(String[] args) throws Exception {
        //InetAddress test
        InetAddress inetAddress = InetAddress.getByName("127.0.0.1");
        System.out.println(inetAddress.getHostName());
        
    }
}
