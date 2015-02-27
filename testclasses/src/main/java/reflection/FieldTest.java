package reflection;

import java.lang.reflect.Field;

public class FieldTest {
    
    static int x;
    
    public static void main(String[] args) throws Exception {
        Field f = String.class.getDeclaredField("value");
        f.setAccessible(true);
        f.get("123");
        
        Field fx = FieldTest.class.getDeclaredField("x");
        fx.setAccessible(true);
        fx.get(null);
    }
    
}
