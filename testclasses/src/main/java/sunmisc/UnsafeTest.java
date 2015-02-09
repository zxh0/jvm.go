package sunmisc;

import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeTest {
    
    public static void main(String[] args) throws Exception {
        //Unsafe unsafe = Unsafe.getUnsafe();
        Field f = Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        Unsafe unsafe = (Unsafe) f.get(null);

        //array(unsafe);
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
        final long address = unsafe.allocateMemory(8);
        
        unsafe.putByte(address, (byte)7);
        if (unsafe.getByte(address) != 7) {
            System.out.println("getByte() failed!");
        }
        
        unsafe.putShort(address, (short)500);
        if (unsafe.getShort(address) != 500) {
            System.out.println("getShort() failed!");
        }
        
//        unsafe.putChar(address, 'c');       address+=2;
//        unsafe.putInt(address, 29);         address+=4;
//        unsafe.putLong(address, 79L);       address+=8;
//        unsafe.putFloat(address, 3.14f);    address+=4;
//        unsafe.putDouble(address, 2.71828); address+=8;
        
        
        
//        address+=2;
//        if (unsafe.getChar(address) != 'c') {
//            System.out.println("getChar() failed!");
//        }
//        address+=2;
//        if (unsafe.getInt(address) != 29) {
//            System.out.println("getInt() failed!");
//        }
//        address+=4;
//        if (unsafe.getLong(address) != 79L) {
//            System.out.println("getLong() failed!");
//        }
//        address+=8;
//        if (unsafe.getFloat(address) != 3.14f) {
//            System.out.println("getFloat() failed!");
//        }
//        address+=4;
//        if (unsafe.getDouble(address) != 2.71828) {
//            System.out.println("getDouble() failed");
//        }
        
        unsafe.freeMemory(address);
        System.out.println("memory testing ok!");
    }
    
}
