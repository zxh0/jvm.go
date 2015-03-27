package java7.sunmisc;

import java.lang.reflect.Field;
import libs.junit.UnitTestRunner;
import org.junit.BeforeClass;
import org.junit.Test;
import sun.misc.Unsafe;
import static org.junit.Assert.*;

public class UnsafeObjectTest {
    
    private static final Unsafe unsafe;
    private static final long booleanArrBaseOffset;
    private static final long booleanArrIndexScale;
    private static final long byteArrBaseOffset;
    private static final long byteArrIndexScale;
    private static final long charArrBaseOffset;
    private static final long charArrIndexScale;
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
    private static final long zOffset;
    private static final long bOffset;
    private static final long cOffset;
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
        booleanArrBaseOffset = unsafe.arrayBaseOffset(boolean[].class);
        booleanArrIndexScale = unsafe.arrayIndexScale(boolean[].class);
        byteArrBaseOffset = unsafe.arrayBaseOffset(byte[].class);
        byteArrIndexScale = unsafe.arrayIndexScale(byte[].class);
        charArrBaseOffset = unsafe.arrayBaseOffset(char[].class);
        charArrIndexScale = unsafe.arrayIndexScale(char[].class);
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
            zOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("z"));
            bOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("b"));
            cOffset = unsafe.objectFieldOffset(UnsafeObjectTest.class.getDeclaredField("c"));
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
    public void getPutBooleanArray() {
        boolean[] arr = {false, true, false};
        long index1 = booleanArrBaseOffset + booleanArrIndexScale;
        assertEquals(true, unsafe.getBoolean(arr, index1));
        unsafe.putBoolean(arr, index1, false);
        assertEquals(false, unsafe.getBoolean(arr, index1));
    }
    
    @Test
    public void getPutBooleanField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
//        unsafe.objectFieldOffset(null)
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(UnsafeObjectTest.class);
    }
    
}
