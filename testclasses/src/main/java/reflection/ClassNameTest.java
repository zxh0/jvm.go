package reflection;

public class ClassNameTest {
    
    public static void main(String[] args) {
        // normal classes
        System.out.println(Object.class.getName());
        System.out.println(ClassNameTest.class.getName());        
        // primivive types
        System.out.println(void.class.getName());
        System.out.println(boolean.class.getName());
        System.out.println(byte.class.getName());
        System.out.println(short.class.getName());
        System.out.println(char.class.getName());
        System.out.println(int.class.getName());
        System.out.println(long.class.getName());
        System.out.println(float.class.getName());
        System.out.println(double.class.getName());
        // array types
        System.out.println(new boolean[0].getClass().getName());
        System.out.println(new byte[0].getClass().getName());
        System.out.println(new short[0].getClass().getName());
        System.out.println(new char[0].getClass().getName());
        System.out.println(new int[0].getClass().getName());
        System.out.println(new long[0].getClass().getName());
        System.out.println(new float[0].getClass().getName());
        System.out.println(new double[0].getClass().getName());
        System.out.println(new Object[0].getClass().getName());
        System.out.println(args.getClass().getName());
        //System.out.println(new int[0][0].getClass().getName());
        //System.out.println(new Object[0][0].getClass().getName());
    }
    
}
