import jvmgo.StringTest;
import jvmgo.exception.JvmExTest;
import jvmgo.file.FileIoTest;
import org.junit.Test;
import jvmgo.reflection.ArrayClassTest;
import jvmgo.reflection.PrimitiveClassTest;
import jvmgo.UnitTestRunner;
import static org.junit.Assert.*;
import jvmgo.reflection.ClassLoaderTest;
import jvmgo.reflection.MethodTest;
import jvmgo.thread.MainThreadTest;

public class UnitTests {
    
    @Test
    public void test() {
        assertEquals(2, 1 + 1);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            UnitTests.class,
            PrimitiveClassTest.class,
            ArrayClassTest.class,
            StringTest.class,
            ClassLoaderTest.class,
            JvmExTest.class,
            FileIoTest.class,
            MethodTest.class,
            MainThreadTest.class,
        });
    }
    
}
