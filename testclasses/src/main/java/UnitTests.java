import java7.exception.JvmExTest;
import java7.file.FileIoTest;
import java7.reflection.ArrayClassTest;
import java7.reflection.PrimitiveClassTest;
import java7.cl.ClassLoaderTest;
import java7.reflection.MethodTest;
import java7.thread.MainThreadTest;
import java7.StringTest;
import java7.UnitTestRunner;
import java7.cl.GetClassLoaderTest;
import java8.InterfaceMethodTest;
import org.junit.Test;
import static org.junit.Assert.*;

public class UnitTests {
    
    @Test
    public void test() {
        assertEquals(2, 1 + 1);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            ArrayClassTest.class,
            ClassLoaderTest.class,
            FileIoTest.class,
            GetClassLoaderTest.class,
            InterfaceMethodTest.class,
            JvmExTest.class,
            MainThreadTest.class,
            MethodTest.class,
            PrimitiveClassTest.class,
            StringTest.class,
            UnitTests.class,
        });
    }
    
}
