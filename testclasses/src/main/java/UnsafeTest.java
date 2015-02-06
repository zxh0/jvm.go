import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeTest {
    
    public static void main(String[] args) throws Exception {
        //Unsafe unsafe = Unsafe.getUnsafe();
        Field f = Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        Unsafe unsafe = (Unsafe) f.get(null);

        array(unsafe);
    }
    
    private static void array(Unsafe unsafe) {
        System.out.println("arrayBaseOffset");
        System.out.println(unsafe.arrayBaseOffset(new int[0].getClass()));
        System.out.println(unsafe.arrayBaseOffset(new long[0].getClass()));
        System.out.println(unsafe.arrayBaseOffset(new Object[0].getClass()));
        System.out.println(unsafe.arrayBaseOffset(new Class<?>[0].getClass()));
        System.out.println("arrayIndexScale");
        System.out.println(unsafe.arrayIndexScale(new int[0].getClass()));
        System.out.println(unsafe.arrayIndexScale(new long[0].getClass()));
        System.out.println(unsafe.arrayIndexScale(new Object[0].getClass()));
        System.out.println(unsafe.arrayIndexScale(new Class<?>[0].getClass()));
    }
    
}
