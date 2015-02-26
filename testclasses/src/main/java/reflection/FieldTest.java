package reflection;

import java.lang.reflect.Field;

public class FieldTest {
    
    public static void main(String[] args) throws Exception {
        Field f = String.class.getDeclaredField("value");
        f.setAccessible(true);
        f.get("123");
    }
    
}
