import java7.StringTest;
import java7.cl.ClassLoaderTest;
import java7.cl.GetClassLoaderTest;
import java7.exception.ClassLibExTest;
import java7.exception.InstructionExTest;
import java7.file.FileIoTest;
import java7.reflection.ArrayClassTest;
import java7.reflection.MethodTest;
import java7.reflection.PrimitiveClassTest;
import java7.thread.MainThreadTest;
import java7.wrapper.DoubleTest;
import java8.InterfaceMethodTest;
import libs.junit.UnitTestRunner;
import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class UnitTests {
    
    @Test
    public void test() {
        assertEquals(2, 1 + 1);
    }
    
    public static void main(String[] args) {
        UnitTestRunner.run(new Class<?>[] {
            ArrayClassTest.class,
            ClassLibExTest.class,
            ClassLoaderTest.class,
            DoubleTest.class,
            FileIoTest.class,
            GetClassLoaderTest.class,
            InstructionExTest.class,
            InterfaceMethodTest.class,
            MainThreadTest.class,
            MethodTest.class,
            PrimitiveClassTest.class,
            StringTest.class,
            UnitTests.class,
        });
    }
    
}
