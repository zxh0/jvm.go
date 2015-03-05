package reflection;

import java.util.Arrays;

public class ArrayClassTest {
    
    public static void main(String[] args) {
        Class<?> c = new boolean[0].getClass();
        System.out.println(c.getName());
        System.out.println("superclass:" + c.getSuperclass());
        System.out.println(Arrays.toString(c.getInterfaces()));
        System.out.println("fields:" + c.getFields().length);
        System.out.println("methods:" + c.getMethods().length);
        System.out.println("declaredMethods:" + c.getDeclaredMethods().length);
    }
    
}
