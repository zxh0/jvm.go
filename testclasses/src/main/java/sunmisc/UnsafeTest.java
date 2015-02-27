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
        //objArr(unsafe);
        //cmpInt(unsafe);
        memory(unsafe);
        System.out.println("OK!");
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
    
    private static void objArr(Unsafe unsafe) {
        String[] arr = {"one", "two"};
        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
        System.out.println(unsafe.getObject(arr, arrayBaseOffset));
        System.out.println(unsafe.getObject(arr, arrayBaseOffset + arrayIndexScale));
    }
    
    private static void cmpInt(Unsafe unsafe) {
        int[] arr = {1, 3, 7};
        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
        unsafe.compareAndSwapInt(arr, arrayBaseOffset, 1, 8);
        if (arr[0] != 8) {
            System.out.println("cmpInt() failed!");
        }
    }
    
    private static void memory(Unsafe unsafe) {
        final long address = unsafe.allocateMemory(8);
        
        unsafe.putAddress(address, address);
        if (unsafe.getAddress(address) != address) {
            System.out.println("getAddress() failed!");
        }
        
        unsafe.putByte(address, (byte)7);
        if (unsafe.getByte(address) != 7) {
            System.out.println("getByte() failed!");
        }
        unsafe.putByte(address, (byte)-7);
        if (unsafe.getByte(address) != -7) {
            System.out.println("getByte() failed!");
        }
        
        unsafe.putShort(address, (short)500);
        if (unsafe.getShort(address) != 500) {
            System.out.println("getShort() failed!");
        }
        unsafe.putShort(address, (short)-500);
        if (unsafe.getShort(address) != -500) {
            System.out.println("getShort() failed!");
        }
        
        unsafe.putChar(address, 'c');
        if (unsafe.getChar(address) != 'c') {
            System.out.println("getChar() failed!");
        }
        
        unsafe.putInt(address, 65536);
        if (unsafe.getInt(address) != 65536) {
            System.out.println("getInt() failed!");
        }
        unsafe.putInt(address, -65536);
        if (unsafe.getInt(address) != -65536) {
            System.out.println("getInt() failed!");
        }
        
        unsafe.putLong(address, 79L);
        if (unsafe.getLong(address) != 79L) {
            System.out.println("getLong() failed!");
        }
        unsafe.putLong(address, -79L);
        if (unsafe.getLong(address) != -79L) {
            System.out.println("getLong() failed!");
        }
        
        unsafe.putFloat(address, 3.14f);
        if (unsafe.getFloat(address) != 3.14f) {
            System.out.println("getFloat() failed!");
        }
        
        unsafe.putDouble(address, 2.71828);
        if (unsafe.getDouble(address) != 2.71828) {
            System.out.println("getDouble() failed");
        }
        
        long newAddress = unsafe.reallocateMemory(address, 100);
        unsafe.freeMemory(newAddress);
        System.out.println("memory testing ok!");
    }
    
}
