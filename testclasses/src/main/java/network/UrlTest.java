package network;

import java.net.MalformedURLException;
import java.net.URL;

public class UrlTest {
    
    public static void main(String[] args) throws MalformedURLException {
        URL url = new URL("http://cn.bing.com");
        System.out.println("OK!");
    }
    
}
