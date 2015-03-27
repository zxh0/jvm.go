package java7.sunmisc;

import java.lang.reflect.Field;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import sun.misc.Unsafe;
import static org.junit.Assert.*;

public class UnsafeObjectTest {
    
    private static final Unsafe unsafe;
    private static final long shortArrBaseOffset;
    private static final long shortArrIndexScale;
    private static final long intArrBaseOffset;
    private static final long intArrIndexScale;
    private static final long longArrBaseOffset;
    private static final long longArrIndexScale;
    private static final long floatArrBaseOffset;
    private static final long floatArrIndexScale;
    private static final long doubleArrBaseOffset;
    private static final long doubleArrIndexScale;
    private static final long objectArrBaseOffset;
    private static final long objectArrIndexScale;
    private static final long sOffset;
    private static final long iOffset;
    private static final long jOffset;
    private static final long fOffset;
    private static final long dOffset;
    
    private boolean z;
    private byte b;
    private char c;
    private short s;
    private int i;
    private long j;
    private float f;
    private double d;
    
    static {
        unsafe = UnsafeGetter.getUnsafe();
        shortArrBaseOffset = unsafe.arrayBaseOffset(short[].class);
        shortArrIndexScale = unsafe.arrayIndexScale(short[].class);
        intArrBaseOffset = unsafe.arrayBaseOffset(int[].class);
        intArrIndexScale = unsafe.arrayIndexScale(int[].class);
        longArrBaseOffset = unsafe.arrayBaseOffset(long[].class);
        longArrIndexScale = unsafe.arrayIndexScale(long[].class);
        floatArrBaseOffset = unsafe.arrayBaseOffset(float[].class);
        floatArrIndexScale = unsafe.arrayIndexScale(float[].class);
        doubleArrBaseOffset = unsafe.arrayBaseOffset(double[].class);
        doubleArrIndexScale = unsafe.arrayIndexScale(double[].class);
        objectArrBaseOffset = unsafe.arrayBaseOffset(Object[].class);
        objectArrIndexScale = unsafe.arrayIndexScale(Object[].class);
        try {
            sOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("s"));
            iOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("i"));
            jOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("j"));
            fOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("f"));
            dOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("d"));
        } catch (NoSuchFieldException e) {
            throw new RuntimeException(e);
        }
    }
    
    @Test
    public void booleanArray() {
        boolean[] arr = {false, true, false};
        
        long booleanArrBaseOffset = unsafe.arrayBaseOffset(boolean[].class);
        long booleanArrIndexScale = unsafe.arrayIndexScale(boolean[].class);
        long index1 = booleanArrBaseOffset + booleanArrIndexScale;
        
        assertEquals(true, unsafe.getBoolean(arr, index1));
        unsafe.putBoolean(arr, index1, false);
        assertEquals(false, unsafe.getBoolean(arr, index1));
    }
    
    @Test
    public void booleanField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long zOffset = objectFieldOffset("z");
        
        assertEquals(false, unsafe.getBoolean(obj, zOffset));
        unsafe.putBoolean(obj, zOffset, true);
        assertEquals(true, unsafe.getBoolean(obj, zOffset));
    }
    
    @Test
    public void byteArray() {
        byte[] arr = {1, 3, 8};
        
        long byteArrBaseOffset = unsafe.arrayBaseOffset(byte[].class);
        long byteArrIndexScale = unsafe.arrayIndexScale(byte[].class);
        long index1 = byteArrBaseOffset + byteArrIndexScale;
        
        assertEquals((byte)3, unsafe.getByte(arr, index1));
        unsafe.putByte(arr, index1, (byte)120);
        assertEquals((byte)120, unsafe.getByte(arr, index1));
    }
    
    @Test
    public void byteField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long bOffset = objectFieldOffset("b");
        
        assertEquals((byte)0, unsafe.getByte(obj, bOffset));
        unsafe.putByte(obj, bOffset, (byte)17);
        assertEquals((byte)17, unsafe.getByte(obj, bOffset));
    }
    
    @Test
    public void charArray() {
        char[] arr = {'x', 'y', 'z'};
        
        long charArrBaseOffset = unsafe.arrayBaseOffset(char[].class);
        long charArrIndexScale = unsafe.arrayIndexScale(char[].class);
        long index1 = charArrBaseOffset + charArrIndexScale;
        
        assertEquals('y', unsafe.getChar(arr, index1));
        unsafe.putChar(arr, index1, 'a');
        assertEquals('a', unsafe.getChar(arr, index1));
    }
    
    @Test
    public void charField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long bOffset = objectFieldOffset("c");
        
        assertEquals('\0', unsafe.getChar(obj, bOffset));
        unsafe.putChar(obj, bOffset, 'x');
        assertEquals('x', unsafe.getChar(obj, bOffset));
    }
    
    private static long objectFieldOffset(String fieldName) {
        try {
            Field f = UnsafeObjectTest.class.getDeclaredField(fieldName);
            return unsafe.objectFieldOffset(f);
        } catch (NoSuchFieldException e) {
            throw new RuntimeException(e);
        }
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(UnsafeObjectTest.class);
    }
    
}
