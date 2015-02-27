package network;

import java.net.URL;

public class UrlTest {
    
    public static void main(String[] args) throws Exception {
        URL url = new URL("http://cn.bing.com");
        url.openStream();
        System.out.println("OK!");
    }
    
}
