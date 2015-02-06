import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeTest {
    
    public static void main(String[] args) throws Exception {
        //Unsafe unsafe = Unsafe.getUnsafe();
        Field f = Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        Unsafe unsafe = (Unsafe) f.get(null);

        array(unsafe);
        memory(unsafe);
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
    
    private static void memory(Unsafe unsafe) {
        System.out.println("memory");
        long address = unsafe.allocateMemory(8);
        unsafe.putLong(address, 0x0102030405060708L);
        System.out.println(unsafe.getByte(address));
        System.out.println(unsafe.getByte(address + 1));
        
        //unsafe.putLong(address + 16, 0);
        //unsafe.freeMemory(address + 2);
    }
    
}
