import jvmgo.exception.JvmExTest;
import jvmgo.file.FileIoTest;
import jvmgo.reflection.ArrayClassTest;
import jvmgo.reflection.PrimitiveClassTest;
import jvmgo.cl.ClassLoaderTest;
import jvmgo.reflection.MethodTest;
import jvmgo.thread.MainThreadTest;
import jvmgo.StringTest;
import jvmgo.UnitTestRunner;
import jvmgo.cl.GetClassLoaderTest;
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
            JvmExTest.class,
            MainThreadTest.class,
            MethodTest.class,
            PrimitiveClassTest.class,
            StringTest.class,
            UnitTests.class,
        });
    }
    
}
