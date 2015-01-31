package reflection;

import java.lang.reflect.Field;

public class PrimitiveClass {
    
    public static void main(String[] args) {
        Class<?> c = int.class;
        System.out.println(c.getName());
        for (Field f : c.getFields()) {
            System.out.println(f.getName());
        }
    }
    
}
