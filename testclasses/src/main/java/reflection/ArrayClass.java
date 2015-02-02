package reflection;

public class ArrayClass {
 
    public static void main(String[] args) {
        Class<?> c = new boolean[0].getClass();
        System.out.println(c.getName());
//        System.out.println("superclass:" + c.getSuperclass());
//        System.out.println("fields:" + c.getFields().length);
//        System.out.println("methods:" + c.getMethods().length);
//        System.out.println("declaredMethods:" + c.getDeclaredMethods().length);
    }
    
}
