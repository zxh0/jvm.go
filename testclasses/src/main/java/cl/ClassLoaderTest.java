package cl;

import java.net.URL;

public class ClassLoaderTest {
    
    public static void main(String[] args) {
        ClassLoader cl = ClassLoaderTest.class.getClassLoader();
        URL x = cl.getResource("org/eclipse/jetty/http/mime.properties");
        System.out.println(x);
    }
    
}
