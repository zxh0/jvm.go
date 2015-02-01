
import java.lang.reflect.Field;
import sun.misc.Unsafe;

public class UnsafeTest {
    
    public static void main(String[] args) throws Exception {
        //Unsafe unsafe = Unsafe.getUnsafe();
        Field f = Unsafe.class.getDeclaredField("theUnsafe");
        f.setAccessible(true);
        Unsafe unsafe = (Unsafe) f.get(null);

        System.out.println("OK!");
    }
    
}
