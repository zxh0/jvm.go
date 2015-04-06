package java6.sunmisc;

import libs.junit.UnitTestRunner;
import org.junit.Test;
import sun.misc.Unsafe;
import static org.junit.Assert.*;

public class UnsafeMemoryTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(UnsafeMemoryTest.class);
    }
    
    @Test
    public void test() {
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
        
        unsafe.putFloat(address, 3.14f);
        assertEquals(3.14f, unsafe.getFloat(address), 0.01);
        
        unsafe.putDouble(address, 2.71828);
        assertEquals(2.71828, unsafe.getDouble(address), 0.01);
        
        long newAddress = unsafe.reallocateMemory(address, 100);
        unsafe.freeMemory(newAddress);
    }
    
}
