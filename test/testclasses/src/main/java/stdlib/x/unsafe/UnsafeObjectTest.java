package stdlib.x.unsafe;

import java.lang.reflect.Field;
import sun.misc.Unsafe;

import static helper.MyAssert.*;

public class UnsafeObjectTest implements Runnable {

    private static final Unsafe unsafe = UnsafeGetter.getUnsafe();

    private boolean z;
    private byte b;
    private char c;
    private short s;
    private int i;
    private long j;
    private float f;
    private double d;
    private String str;

    public static void main(String[] args) {
        new UnsafeObjectTest().run();
    }

    @Override
    public void run() {
        testBooleanArray();
        testBooleanField();
        testByteArray();
        testByteField();
        testCharArray();
        testCharField();
        testShortArray();
        testShortField();
        testIntArray();
        testIntField();
        testLongArray();
        testLongField();
        testFloatArray();
        testFloatField();
        testDoubleArray();
        testDoubleField();
        testObjectArray();
        testObjectField();
    }

    private static void testBooleanArray() {
        boolean[] arr = {false, true, false};

        long booleanArrBaseOffset = unsafe.arrayBaseOffset(boolean[].class);
        long booleanArrIndexScale = unsafe.arrayIndexScale(boolean[].class);
        long index1 = booleanArrBaseOffset + booleanArrIndexScale;

        assertEquals(true, unsafe.getBoolean(arr, index1));
        unsafe.putBoolean(arr, index1, false);
        assertEquals(false, unsafe.getBoolean(arr, index1));
    }

    private static void testBooleanField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long zOffset = objectFieldOffset("z");

        assertEquals(false, unsafe.getBoolean(obj, zOffset));
        unsafe.putBoolean(obj, zOffset, true);
        assertEquals(true, unsafe.getBoolean(obj, zOffset));
    }

    private static void testByteArray() {
        byte[] arr = {1, 3, 8, 2};

        long byteArrBaseOffset = unsafe.arrayBaseOffset(byte[].class);
        long byteArrIndexScale = unsafe.arrayIndexScale(byte[].class);
        long index1 = byteArrBaseOffset + byteArrIndexScale;

        assertEquals(0x0803, unsafe.getShort(arr, index1));

        assertEquals((byte)3, unsafe.getByte(arr, index1));
        unsafe.putByte(arr, index1, (byte)120);
        assertEquals((byte)120, unsafe.getByte(arr, index1));
    }

    private static void testByteField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long bOffset = objectFieldOffset("b");

        assertEquals((byte)0, unsafe.getByte(obj, bOffset));
        unsafe.putByte(obj, bOffset, (byte)17);
        assertEquals((byte)17, unsafe.getByte(obj, bOffset));
    }

    private static void testCharArray() {
        char[] arr = {'x', 'y', 'z'};

        long charArrBaseOffset = unsafe.arrayBaseOffset(char[].class);
        long charArrIndexScale = unsafe.arrayIndexScale(char[].class);
        long index1 = charArrBaseOffset + charArrIndexScale;

        assertEquals('y', unsafe.getChar(arr, index1));
        unsafe.putChar(arr, index1, 'a');
        assertEquals('a', unsafe.getChar(arr, index1));
    }

    private static void testCharField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long cOffset = objectFieldOffset("c");

        assertEquals('\0', unsafe.getChar(obj, cOffset));
        unsafe.putChar(obj, cOffset, 'x');
        assertEquals('x', unsafe.getChar(obj, cOffset));
    }

    private static void testShortArray() {
        short[] arr = {3, 4, 5};

        long shortArrBaseOffset = unsafe.arrayBaseOffset(short[].class);
        long shortArrIndexScale = unsafe.arrayIndexScale(short[].class);
        long index1 = shortArrBaseOffset + shortArrIndexScale;

        assertEquals(4, unsafe.getShort(arr, index1));
        unsafe.putShort(arr, index1, (short)12345);
        assertEquals((short)12345, unsafe.getShort(arr, index1));
    }

    private static void testShortField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long sOffset = objectFieldOffset("s");

        assertEquals((short)0, unsafe.getShort(obj, sOffset));
        unsafe.putShort(obj, sOffset, (short)12345);
        assertEquals((short)12345, unsafe.getShort(obj, sOffset));
    }

    private static void testIntArray() {
        int[] arr = {3, 4, 5};

        long intArrBaseOffset = unsafe.arrayBaseOffset(int[].class);
        long intArrIndexScale = unsafe.arrayIndexScale(int[].class);
        long index1 = intArrBaseOffset + intArrIndexScale;

        assertEquals(4, unsafe.getInt(arr, index1));
        unsafe.putInt(arr, index1, 12345);
        assertEquals(12345, unsafe.getInt(arr, index1));
    }

    private static void testIntField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long iOffset = objectFieldOffset("i");

        assertEquals(0, unsafe.getInt(obj, iOffset));
        unsafe.putInt(obj, iOffset, 12345);
        assertEquals(12345, unsafe.getInt(obj, iOffset));
    }

    private static void testLongArray() {
        long[] arr = {3, 4, 5};

        long longArrBaseOffset = unsafe.arrayBaseOffset(long[].class);
        long longArrIndexScale = unsafe.arrayIndexScale(long[].class);
        long index1 = longArrBaseOffset + longArrIndexScale;

        assertEquals(4L, unsafe.getLong(arr, index1));
        unsafe.putLong(arr, index1, 12345L);
        assertEquals(12345L, unsafe.getLong(arr, index1));
    }

    private static void testLongField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long jOffset = objectFieldOffset("j");

        assertEquals(0L, unsafe.getLong(obj, jOffset));
        unsafe.putLong(obj, jOffset, 12345L);
        assertEquals(12345L, unsafe.getLong(obj, jOffset));
    }

    private static void testFloatArray() {
        float[] arr = {1.4f, 1.5f, 0f};

        long floatArrBaseOffset = unsafe.arrayBaseOffset(float[].class);
        long floatArrIndexScale = unsafe.arrayIndexScale(float[].class);
        long index1 = floatArrBaseOffset + floatArrIndexScale;

        assertEquals(1.5f, unsafe.getFloat(arr, index1));
        unsafe.putFloat(arr, index1, 2.5f);
        assertEquals(2.5f, unsafe.getFloat(arr, index1));
    }

    private static void testFloatField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long fOffset = objectFieldOffset("f");

        assertEquals(0f, unsafe.getFloat(obj, fOffset));
        unsafe.putFloat(obj, fOffset, 1.5f);
        assertEquals(1.5f, unsafe.getFloat(obj, fOffset));
    }

    private static void testDoubleArray() {
        double[] arr = {1.4, 1.5, 0d};

        long doubleArrBaseOffset = unsafe.arrayBaseOffset(double[].class);
        long doubleArrIndexScale = unsafe.arrayIndexScale(double[].class);
        long index1 = doubleArrBaseOffset + doubleArrIndexScale;

        assertEquals(1.5, unsafe.getDouble(arr, index1));
        unsafe.putDouble(arr, index1, 2.71828);
        assertEquals(2.71828, unsafe.getDouble(arr, index1));
    }

    private static void testDoubleField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long dOffset = objectFieldOffset("d");

        assertEquals(0, unsafe.getDouble(obj, dOffset));
        unsafe.putDouble(obj, dOffset, 1.5);
        assertEquals(1.5, unsafe.getDouble(obj, dOffset));
    }

    private static void testObjectArray() {
        String[] arr = {"a", "b", "c"};

        long objectArrBaseOffset = unsafe.arrayBaseOffset(String[].class);
        long objectArrIndexScale = unsafe.arrayIndexScale(String[].class);
        long index1 = objectArrBaseOffset + objectArrIndexScale;

        assertEquals("b", unsafe.getObject(arr, index1));
        unsafe.putObject(arr, index1, "hello");
        assertEquals("hello", unsafe.getObject(arr, index1));
    }

    private static void testObjectField() {
        UnsafeObjectTest obj = new UnsafeObjectTest();
        long strOffset = objectFieldOffset("str");

        assertNull(unsafe.getObject(obj, strOffset));
        unsafe.putObject(obj, strOffset, "world");
        assertEquals("world", unsafe.getObject(obj, strOffset));
    }

    private static long objectFieldOffset(String fieldName) {
        try {
            Field f = UnsafeObjectTest.class.getDeclaredField(fieldName);
            return unsafe.objectFieldOffset(f);
        } catch (NoSuchFieldException e) {
            throw new RuntimeException(e);
        }
    }

}
