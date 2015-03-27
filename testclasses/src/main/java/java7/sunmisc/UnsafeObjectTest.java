package java7.sunmisc;

import java.lang.reflect.Field;
import libs.junit.UnitTestRunner;
import org.junit.BeforeClass;
import org.junit.Test;
import sun.misc.Unsafe;

public class UnsafeObjectTest {
    
    private static final Unsafe unsafe;
    private static final int booleanArrBaseOffset;
    private static final int  booleanArrIndexScale;
    private static final int  byteArrBaseOffset;
    private static final int  byteArrIndexScale;
    private static final int  charArrBaseOffset;
    private static final int  charArrIndexScale;
    private static final int  shortArrBaseOffset;
    private static final int  shortArrIndexScale;
    private static final int  intArrBaseOffset;
    private static final int  intArrIndexScale;
    private static final int  longArrBaseOffset;
    private static final int  longArrIndexScale;
    private static final int  floatArrBaseOffset;
    private static final int  floatArrIndexScale;
    private static final int  doubleArrBaseOffset;
    private static final int  doubleArrIndexScale;
    private static final int  objectArrBaseOffset;
    private static final int  objectArrIndexScale;
    
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
    }
    
    @Test
    public void test() {
        
    }
    
    public static void main(String[] args) {
        //UnitTestRunner.run(UnsafeObjectTest.class);
    }
    
}
