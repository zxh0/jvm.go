package jetty;

import java.util.ResourceBundle;

public class ResourceBundleTest {
    
    public static void main(String[] args) {
        ResourceBundle b = ResourceBundle.getBundle("org/eclipse/jetty/http/mime");
        System.out.println(b);
    }
    
}
