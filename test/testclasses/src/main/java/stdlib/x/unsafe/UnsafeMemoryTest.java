package stdlib.x.unsafe;

import sun.misc.Unsafe;

import static helper.MyAssert.*;

public class UnsafeMemoryTest implements Runnable {

    public static void main(String[] args) {
        new UnsafeMemoryTest().run();
    }

    @Override
    public void run() {
        Unsafe unsafe = UnsafeGetter.getUnsafe();
        final long address = unsafe.allocateMemory(8);

        unsafe.putAddress(address, address);
        assertEquals(address, unsafe.getAddress(address));

        unsafe.putByte(address, (byte)7);
        assertEquals((byte)7, unsafe.getByte(address));

        unsafe.putByte(address, (byte)-7);
        assertEquals((byte)-7, unsafe.getByte(address));

        unsafe.putShort(address, (short)500);
        assertEquals((short)500, unsafe.getShort(address));

        unsafe.putShort(address, (short)-500);
        assertEquals((short)-500, unsafe.getShort(address));

        unsafe.putChar(address, 'c');
        assertEquals('c', unsafe.getChar(address));

        unsafe.putInt(address, 65536);
        assertEquals(65536, unsafe.getInt(address));

        unsafe.putInt(address, -65536);
        assertEquals(-65536, unsafe.getInt(address));

        unsafe.putLong(address, 79L);
        assertEquals(79L, unsafe.getLong(address));

        unsafe.putLong(address, -79L);
        assertEquals(-79L, unsafe.getLong(address));

        unsafe.putFloat(address, 1.5f);
        assertEquals(1.5f, unsafe.getFloat(address));

        unsafe.putDouble(address, 1.5);
        assertEquals(1.5, unsafe.getDouble(address));

        long newAddress = unsafe.reallocateMemory(address, 100);
        unsafe.freeMemory(newAddress);
    }

}
