package java6.sunmisc;

import sun.misc.Unsafe;

public class UnsafeTest {
    
    public int x;
    public long y;
    public String z;
    
    public static void main(String[] args) throws Exception {
        Unsafe unsafe = UnsafeGetter.getUnsafe();

        casInt(unsafe);
        casLong(unsafe);
        casObj(unsafe);
        
        System.out.println("OK!");
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
        
        Object[] arr = {one, two};
        long arrayBaseOffset = unsafe.arrayBaseOffset(arr.getClass());
        long arrayIndexScale = unsafe.arrayIndexScale(arr.getClass());
        unsafe.compareAndSwapObject(arr, arrayBaseOffset, one, two);
        if (arr[0] != two) {
            System.out.println("casObj(arr) failed!");
        }
        
        UnsafeTest obj = new UnsafeTest();
        long zOffset = unsafe.objectFieldOffset(UnsafeTest.class.getField("z"));
        unsafe.compareAndSwapObject(obj, zOffset, null, "two");
        if (obj.z != "two") {
            System.out.println("casObj(obj) failed!" + obj.z);
        }
    }
    
}
