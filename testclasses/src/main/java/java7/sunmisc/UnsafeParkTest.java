package java7.sunmisc;

import sun.misc.Unsafe;

import java.util.Calendar;


/**
 * Created by beyond on 15-3-27.
 */
public class UnsafeParkTest {
    public static void main(String[] args) throws Exception {
        Unsafe unsafe = UnsafeGetter.getUnsafe();
        System.out.println(Calendar.getInstance().getTime());
        unsafe.park(false, 1000000000);
        long time = System.currentTimeMillis() + 1000;
        unsafe.park(true, time);
        System.out.println(Calendar.getInstance().getTime());
    }
}
