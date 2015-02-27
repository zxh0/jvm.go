package sunmisc;

import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeTest {
    
    public int x;
    public long y;
    public String z;
    
    public static void main(String[] args) throws Exception {
        //Unsafe unsafe = Unsafe.getUnsafe();
        Field f = Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        Unsafe unsafe = (Unsafe) f.get(null);

        //memory(unsafe);
        //array(unsafe);
        //objArr(unsafe);
        casInt(unsafe);
        casLong(unsafe);
        casObj(unsafe);
        
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
    
    private static void casInt(Unsafe unsafe) throws Exception {
        int[] arr = {1, 3, 7};
        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
        unsafe.compareAndSwapInt(arr, arrayBaseOffset, 1, 8);
        if (arr[0] != 8) {
            System.out.println("casInt(arr) failed!");
        }
        
        UnsafeTest obj = new UnsafeTest();
        long xOffset = unsafe.objectFieldOffset(UnsafeTest.class.getField("x"));
        unsafe.compareAndSwapInt(obj, xOffset, 0, 7);
        if (obj.x != 7) {
            System.out.println("casInt(obj) failed!");
        }
    }
    
    private static void casLong(Unsafe unsafe) throws Exception {
        long[] arr = {1, 3, 7};
        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
        unsafe.compareAndSwapLong(arr, arrayBaseOffset, 1, 8);
        if (arr[0] != 8) {
            System.out.println("casLong(arr) failed!");
        }
        
        UnsafeTest obj = new UnsafeTest();
        long yOffset = unsafe.objectFieldOffset(UnsafeTest.class.getField("y"));
        unsafe.compareAndSwapLong(obj, yOffset, 0, 7);
        if (obj.y != 7) {
            System.out.println("casLong(obj) failed!");
        }
    }
    
    private static void casObj(Unsafe unsafe) throws Exception {
        String one = "1";
        String two = "2";
//        
//        Object[] arr = {one, two};
//        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
//        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
//        unsafe.compareAndSwapObject(arr, arrayBaseOffset, one, two);
//        if (arr[0] != two) {
//            System.out.println("casObj(arr) failed!");
//        }
        
        UnsafeTest obj = new UnsafeTest();
        long zOffset = unsafe.objectFieldOffset(UnsafeTest.class.getField("z"));
        unsafe.compareAndSwapObject(obj, zOffset, null, "two");
        if (obj.z != "two") {
            System.out.println("casObj(obj) failed!" + obj.z);
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
