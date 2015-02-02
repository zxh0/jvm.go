package reflection;

public class PrimitiveClass {
    
    public static void main(String[] args) {
        Class<?> c = int.class;
        System.out.println(c.getName());
        System.out.println("superclass:" + c.getSuperclass());
        System.out.println("fields:" + c.getFields().length);
        System.out.println("methods:" + c.getMethods().length);
    }
    
}
