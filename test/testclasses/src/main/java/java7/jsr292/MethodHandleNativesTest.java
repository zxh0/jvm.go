package java7.jsr292;

import helper.ReflectionHelper;
import libs.junit.UnitTestRunner;
import org.junit.Test;
import static org.junit.Assert.*;

public class MethodHandleNativesTest {
    
    public static void main(String[] args) {
        UnitTestRunner.run(MethodHandleNativesTest.class);
    }
    
    @Test
    public void getConstant() throws Exception {
        Class<?> mhnClass = Class.forName("java.lang.invoke.MethodHandleNatives");
        
        for (int i = 0; i < 100; i++) {
            Object x =  ReflectionHelper.call(mhnClass, "getConstant", i);
            if (i == 4) {
                assertEquals(1, x);
            } else {
                assertEquals(0, x);
            }
        }
    }
    
}
