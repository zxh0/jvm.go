package stdlib.net;

import java.io.InputStream;
import java.net.URL;

public class UrlTest {

    public static void main(String[] args) throws Exception {
        URL url = new URL("http://cn.bing.com");
        try (InputStream is = url.openStream()) {
            byte[] bytes = new byte[is.available()];
            int n = is.read(bytes);
            System.out.println(new String(bytes));
        }
        System.out.println("OK!");
    }

}
