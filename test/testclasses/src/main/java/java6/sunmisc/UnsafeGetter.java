package java6.sunmisc;

import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeGetter {
    
    public static Unsafe getUnsafe() {
        //Unsafe unsafe = Unsafe.getUnsafe();
        try {
            Field f = Unsafe.class.getDeclaredField("theUnsafe");
            f.setAccessible(true);
            Unsafe unsafe = (Unsafe) f.get(null);
            return unsafe;
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
    
}
