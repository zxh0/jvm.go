package reflection;

public class GetClass {
    
    public static void main(String[] args) {
        System.out.println(new GetClass().getClass().getName());
        // primivive types
        System.out.println(boolean.class.getName());
        System.out.println(byte.class.getName());
        System.out.println(short.class.getName());
        System.out.println(char.class.getName());
        System.out.println(int.class.getName());
        System.out.println(long.class.getName());
        System.out.println(float.class.getName());
        System.out.println(double.class.getName());
        // array types
        System.out.println(new boolean[0].getClass());
        System.out.println(new byte[0].getClass());
        System.out.println(new short[0].getClass());
        System.out.println(new char[0].getClass());
        System.out.println(new int[0].getClass());
        System.out.println(new long[0].getClass());
        System.out.println(new float[0].getClass());
        System.out.println(new double[0].getClass());
        System.out.println(args.getClass());
    }
    
}
